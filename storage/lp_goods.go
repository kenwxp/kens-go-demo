package storage

import (
	"context"
	"database/sql"
	"fmt"
	"kens/demo/storage/types"
)

const goodsSchema = `
 	CREATE TABLE IF NOT EXISTS lp_goods  (
     	goods_id			serial primary key,
     	goods_name			text,
     	goods_type			text,
     	category			text,
     	content				text,
     	images				text,
		price				text,
		membership_price 	text,
		unit				text,
		stock				integer,
		min_buy_num			integer,
		use_property		integer,
		property_groups		text,
		spec_groups			text,
		is_sell				text,
		sort				integer,
		create_time			text,
		update_time			text,
		is_valid			integer
    );
    comment on column lp_goods.goods_id		    is '商品id';
    comment on column lp_goods.goods_name		is '商品名';
    comment on column lp_goods.goods_type		is '商品类别 1-普通 2-积分' ;
    comment on column lp_goods.category			is '商品类别';
    comment on column lp_goods.content			is '商品描述';
    comment on column lp_goods.images			is '商品图片 数组';
    comment on column lp_goods.price			is '售价 单位分';
    comment on column lp_goods.membership_price is '会员价 单位分';
    comment on column lp_goods.unit			    is '单位';
    comment on column lp_goods.stock			is '库存';
    comment on column lp_goods.min_buy_num		is '最小起售数';
    comment on column lp_goods.use_property	    is '是否使用规格';
    comment on column lp_goods.property_groups	is '规格数组';
    comment on column lp_goods.spec_groups		is '参数数组';
    comment on column lp_goods.is_sell			is '上架标志 0-下架 1-上架';
    comment on column lp_goods.sort			    is '排序';
    comment on column lp_goods.create_time		is '创建时间';
    comment on column lp_goods.update_time		is '更新时间';
    comment on column lp_goods.is_valid		    is '启用标志 0-启用 1-禁用';
`
const selectGoodsListSQL = "" +
	" select " +
	"  goods_id		," +
	"  goods_name	," +
	"  goods_type	," +
	"  category		," +
	"  content		," +
	"  images		," +
	"  price		," +
	"  membership_price ," +
	"  unit			," +
	"  stock		," +
	"  min_buy_num	," +
	"  use_property	," +
	"  property_groups," +
	"  spec_groups	," +
	"  is_sell		," +
	"  sort			," +
	"  create_time	," +
	"  update_time	," +
	"  is_valid	" +
	" from lp_goods" +
	"  where is_valid = 0"

type goodsStatements struct {
	selectGoodsListStmts *sql.Stmt
}

func (s *goodsStatements) execSchema(db *sql.DB) error {
	_, err := db.Exec(goodsSchema)
	return err
}

func (s *goodsStatements) prepare(db *sql.DB) (err error) {
	if s.selectGoodsListStmts, err = db.Prepare(selectGoodsListSQL); err != nil {
		return
	}
	return
}

func (s *goodsStatements) selectGoodsList(ctx context.Context, txn *sql.Tx) ([]types.Goods, error) {
	list := make([]types.Goods, 0)
	row, err := TxStmt(txn, s.selectGoodsListStmts).QueryContext(ctx)
	defer row.Close()
	if err != nil {
		fmt.Print("selectGoodsList error:", err)
		return nil, err
	}
	for row.Next() {
		var item types.Goods
		err := row.Scan(
			&item.GoodsId,
			&item.GoodsName,
			&item.GoodsType,
			&item.Category,
			&item.Content,
			&item.Images,
			&item.Price,
			&item.MembershipPrice,
			&item.Unit,
			&item.Stock,
			&item.MinBuyNum,
			&item.UseProperty,
			&item.PropertyGroups,
			&item.SpecGroups,
			&item.IsSell,
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
