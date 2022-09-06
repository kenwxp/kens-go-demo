package storage

import (
	"context"
	"database/sql"
	"fmt"
	"kens/demo/storage/types"
)

const orderSchema = `
 CREATE TABLE IF NOT EXISTS lp_order  (
    order_id	 		serial primary key, 
	order_serial    	text,
	payment_id     		integer,
	comment_id     		integer,
	customer_id   		integer,
	shop_id   			integer,
	goods_num			text,
	goods_amount		text,
	send_fee			text,
	discount			text,
	total_amount		text,
	goods_detail    	integer,
	order_status		text,
    send_type	   		text,
    shipping_address_id	integer,
	create_time			text,
	produce_time		text,
    send_time			text,
    receive_time		text,
    finish_time			text,
    is_valid		 	integer  
);
	comment on column lp_order.order_id	 			is '订单id 主键';
	comment on column lp_order.order_serial    		is '订单序列号';
	comment on column lp_order.payment_id     		is '支付id';
	comment on column lp_order.comment_id     		is '评论id';
	comment on column lp_order.customer_id   		is '顾客id';
	comment on column lp_order.shop_id   			is '店铺id';
	comment on column lp_order.goods_num			is '商品总件数';
	comment on column lp_order.goods_amount			is '商品金额';
	comment on column lp_order.send_fee				is '配送金额';
	comment on column lp_order.discount				is '折扣金额';
	comment on column lp_order.total_amount			is '总金额';
	comment on column lp_order.goods_detail    		is '商品详情';
	comment on column lp_order.order_status			is '订单状态';
	comment on column lp_order.send_type	   		is '寄送类型 1-自取 2-外卖';
	comment on column lp_order.shipping_address_id	is '寄送地址id';
	comment on column lp_order.create_time			is '创建时间';
	comment on column lp_order.produce_time			is '制作时间';
	comment on column lp_order.send_time			is '寄送时间';
	comment on column lp_order.receive_time			is '送达时间';
	comment on column lp_order.finish_time			is '完成时间';
	comment on column lp_order.is_valid				is '启用标志 0-启用 1-禁用';
`

const selectOrderListSQL = "" +
	" select " +
	"  order_id	 	," +
	"  order_serial ," +
	"  payment_id   ," +
	"  comment_id   ," +
	"  customer_id  ," +
	"  shop_id   	," +
	"  goods_num	," +
	"  goods_amount	," +
	"  send_fee		," +
	"  discount		," +
	"  total_amount	," +
	"  goods_detail ," +
	"  order_status	," +
	"  send_type	," +
	"  shipping_address_id," +
	"  create_time	," +
	"  produce_time	," +
	"  send_time	," +
	"  receive_time	," +
	"  finish_time	," +
	"  is_valid   " +
	" from lp_order  " +
	"  where is_valid = 0"

type orderStatements struct {
	selectOrderListStmts *sql.Stmt
}

func (s *orderStatements) execSchema(db *sql.DB) error {
	_, err := db.Exec(orderSchema)
	return err
}

func (s *orderStatements) prepare(db *sql.DB) (err error) {
	if s.selectOrderListStmts, err = db.Prepare(selectOrderListSQL); err != nil {
		return
	}
	return
}

func (s *orderStatements) selectOrderList(ctx context.Context, txn *sql.Tx) ([]types.Order, error) {
	list := make([]types.Order, 0)
	row, err := TxStmt(txn, s.selectOrderListStmts).QueryContext(ctx)
	defer row.Close()
	if err != nil {
		fmt.Print("selectOrderList error:", err)
		return nil, err
	}
	for row.Next() {
		var item types.Order
		err := row.Scan(
			&item.OrderId,
			&item.OrderSerial,
			&item.PaymentId,
			&item.CommentId,
			&item.CustomerId,
			&item.ShopId,
			&item.GoodsNum,
			&item.GoodsAmount,
			&item.SendFee,
			&item.Discount,
			&item.TotalAmount,
			&item.GoodsDetail,
			&item.OrderStatus,
			&item.SendType,
			&item.ShippingAddressId,
			&item.CreateTime,
			&item.ProduceTime,
			&item.SendTime,
			&item.ReceiveTime,
			&item.FinishTime,
			&item.IsValid,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}
