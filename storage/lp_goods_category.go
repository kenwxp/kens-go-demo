package storage

import (
	"context"
	"database/sql"
	"fmt"
	"kens/demo/storage/types"
)

const goodsCategorySchema = `
 	CREATE TABLE IF NOT EXISTS lp_goods_category  (
     	category_id		serial primary key,
     	category_name	text,
     	icon			text,
     	sort			integer,
		create_time		text,
		update_time		text,
		is_valid		integer
    );
    comment on column lp_goods_category.category_id		is '类别id';
    comment on column lp_goods_category.category_name	is '类别名';
    comment on column lp_goods_category.icon			is '图表';
    comment on column lp_goods_category.sort			is '排序';
    comment on column lp_goods_category.create_time		is '创建时间';
    comment on column lp_goods_category.update_time		is '更新时间';
    comment on column lp_goods_category.is_valid		is '是否启用 0-启用 1-禁用';
`

const selectGoodsCategoryListSQL = "" +
	" select " +
	"  category_id	 	," +
	"  category_name	," +
	"  icon			  	," +
	"  sort			  	," +
	"  create_time		," +
	"  update_time		," +
	"  is_valid	  " +
	" from lp_goods_category" +
	"  where is_valid = 0"

type goodsCategoryStatements struct {
	selectGoodsCategoryListStmts *sql.Stmt
}

func (s *goodsCategoryStatements) execSchema(db *sql.DB) error {
	_, err := db.Exec(goodsCategorySchema)
	return err
}

func (s *goodsCategoryStatements) prepare(db *sql.DB) (err error) {
	if s.selectGoodsCategoryListStmts, err = db.Prepare(selectGoodsCategoryListSQL); err != nil {
		return
	}
	return
}

func (s *goodsCategoryStatements) selectGoodsCategoryList(ctx context.Context, txn *sql.Tx) ([]types.GoodsCategory, error) {
	list := make([]types.GoodsCategory, 0)
	row, err := TxStmt(txn, s.selectGoodsCategoryListStmts).QueryContext(ctx)
	defer row.Close()
	if err != nil {
		fmt.Print("selectGoodsCategoryList error:", err)
		return nil, err
	}
	for row.Next() {
		var item types.GoodsCategory
		err := row.Scan(
			&item.CategoryId,
			&item.CategoryName,
			&item.Icon,
			&item.Sort,
			&item.CreateTime,
			&item.UpdateTime,
			&item.IsValid,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}
