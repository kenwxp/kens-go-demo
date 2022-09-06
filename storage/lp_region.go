package storage

import (
	"context"
	"database/sql"
	"fmt"
	"kens/demo/storage/types"
)

//CREATE SEQUENCE IF NOT EXISTS region_seq START 1;
//goland:noinspection SqlNoDataSourceInspection
const regionSchema = `
 	CREATE TABLE IF NOT EXISTS lp_region  (
        region_id			text primary key,
        region_name			text,
        region_short_name	text,
        region_code			text,
        region_parent_id	text,
        region_level		text
    );              
	comment on column lp_region.region_id				is '地区主键编号	pk';
	comment on column lp_region.region_name				is '地区名称	';
	comment on column lp_region.region_short_name		is '地区缩写	';
	comment on column lp_region.region_code				is '行政地区编号	';
	comment on column lp_region.region_parent_id		is '地区父id	';
	comment on column lp_region.region_level			is '地区级别 1-省、自治区、直辖市 2-地级市、地区、自治州、盟 3-市辖区、县级市、县';
`
const selectRegionListSQL = "" +
	" select " +
	"  region_id		," +
	"  region_name		," +
	"  region_short_name," +
	"  region_code		," +
	"  region_parent_id ," +
	"  region_level" +
	" from lp_region"

type regionStatements struct {
	selectRegionListStmts *sql.Stmt
}

func (s *regionStatements) execSchema(db *sql.DB) error {
	_, err := db.Exec(regionSchema)
	return err
}

func (s *regionStatements) prepare(db *sql.DB) (err error) {
	if s.selectRegionListStmts, err = db.Prepare(selectRegionListSQL); err != nil {
		return
	}
	return
}

func (s *regionStatements) selectRegionList(ctx context.Context, txn *sql.Tx) ([]types.Region, error) {
	list := make([]types.Region, 0)
	row, err := TxStmt(txn, s.selectRegionListStmts).QueryContext(ctx)
	defer row.Close()
	if err != nil {
		fmt.Print("selectRegionList error:", err)
		return nil, err
	}
	for row.Next() {
		var item types.Region
		err := row.Scan(
			&item.RegionId,
			&item.RegionName,
			&item.RegionShortName,
			&item.RegionCode,
			&item.RegionParentId,
			&item.RegionLevel,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}
