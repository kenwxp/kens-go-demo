package storage

import (
	"context"
	"database/sql"
	"fmt"
	"kens/demo/storage/types"
)

const goodsPropertySchema = `
 	CREATE TABLE IF NOT EXISTS lp_goods_property  (
     	id 			serial 	primary key,
		goods_id 	INTEGER,
		sub_id	 	INTEGER,
		sub_params 	text,
		stock 		INTEGER,
		price 		text,
		pic_url 	text,
		is_valid 	INTEGER
    );
    comment on column lp_goods_property.id		    		is '自增id 主键';
    comment on column lp_goods_property.goods_id 		    is '商品id';
    comment on column lp_goods_property.sub_id	 		    is '规格id';
    comment on column lp_goods_property.sub_params 		    is '规格参数';
    comment on column lp_goods_property.stock 			    is '库存';
    comment on column lp_goods_property.price 			    is '价格';
    comment on column lp_goods_property.pic_url 		    is '图片';
    comment on column lp_goods_property.is_valid 		    is '启用标志 0-启用 1-禁用';
`

const selectGoodsPropertyListSQL = "" +
	" select " +
	"  id 		," +
	"  goods_id ," +
	"  sub_id	 ," +
	"  sub_params," +
	"  stock 	," +
	"  price 	," +
	"  pic_url ," +
	"  is_valid " +
	" from lp_goods_property" +
	"  where is_valid = 0"

type goodsPropertyStatements struct {
	selectGoodsPropertyListStmts *sql.Stmt
}

func (s *goodsPropertyStatements) execSchema(db *sql.DB) error {
	_, err := db.Exec(goodsPropertySchema)
	return err
}

func (s *goodsPropertyStatements) prepare(db *sql.DB) (err error) {
	if s.selectGoodsPropertyListStmts, err = db.Prepare(selectGoodsPropertyListSQL); err != nil {
		return
	}
	return
}

func (s *goodsPropertyStatements) selectGoodsPropertyList(ctx context.Context, txn *sql.Tx) ([]types.GoodsProperty, error) {
	list := make([]types.GoodsProperty, 0)
	row, err := TxStmt(txn, s.selectGoodsPropertyListStmts).QueryContext(ctx)
	defer row.Close()
	if err != nil {
		fmt.Print("selectGoodsPropertyList error:", err)
		return nil, err
	}
	for row.Next() {
		var item types.GoodsProperty
		err := row.Scan(
			&item.Id,
			&item.GoodsId,
			&item.SubId,
			&item.SubParams,
			&item.Stock,
			&item.Price,
			&item.PicUrl,
			&item.IsValid,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}
