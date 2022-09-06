package storage

import (
	"context"
	"database/sql"
	"fmt"
	"kens/demo/storage/types"
)

const orderPaymentSchema = `
 CREATE TABLE IF NOT EXISTS lp_order_payment  (
    payment_id	 	serial primary key,
	pay_order_num    text,
	prepay_id        text default '',
	mch_id 			 text,
	open_id          text,
	customer_id		 integer,
	shop_id		 	 integer,
    pay_amount		 text,
    pay_fee			 text,
	pay_status		 text,
	create_time		 text,
	update_time		 text,
    is_valid		 integer
);
	comment on column lp_order_payment.payment_id	  is '支付id' ;
	comment on column lp_order_payment.pay_order_num  is '微信支付的订单号' ;
	comment on column lp_order_payment.prepay_id      is '微信支付预付款单号' ;
	comment on column lp_order_payment.mch_id 		  is '微信支付商户号' ;
	comment on column lp_order_payment.open_id        is '支付的用户的微信openid' ;
	comment on column lp_order_payment.customer_id	  is '支付的用户id' ;
	comment on column lp_order_payment.shop_id		  is '门店id' ;
	comment on column lp_order_payment.pay_amount	 is '支付金额 单位分' ;
	comment on column lp_order_payment.pay_fee		 is '支付手续费 单位分' ;
	comment on column lp_order_payment.pay_status	 is '支付状态 0-待支付 1-已支付 2-支付失败' ;
	comment on column lp_order_payment.create_time	 is '创建时间' ;
	comment on column lp_order_payment.update_time	 is '更新时间' ;
	comment on column lp_order_payment.is_valid	 	 is '启用标志 0-启用 1-禁用' ;
`

const selectOrderPaymentListSQL = "" +
	" select " +
	"  	payment_id	 ," +
	"   pay_order_num ," +
	"   prepay_id     ," +
	"   mch_id 		," +
	"   open_id       ," +
	"   customer_id	," +
	"   shop_id		 ," +
	"   pay_amount	," +
	"   pay_fee		," +
	"   pay_status	," +
	"  	create_time	," +
	"  	update_time	," +
	"  	is_valid	" +
	" from lp_order_payment" +
	"  where is_valid = 0"

type orderPaymentStatements struct {
	selectOrderPaymentListStmts *sql.Stmt
}

func (s *orderPaymentStatements) execSchema(db *sql.DB) error {
	_, err := db.Exec(orderPaymentSchema)
	return err
}

func (s *orderPaymentStatements) prepare(db *sql.DB) (err error) {
	if s.selectOrderPaymentListStmts, err = db.Prepare(selectOrderPaymentListSQL); err != nil {
		return
	}
	return
}

func (s *orderPaymentStatements) selectOrderPaymentList(ctx context.Context, txn *sql.Tx) ([]types.OrderPayment, error) {
	list := make([]types.OrderPayment, 0)
	row, err := TxStmt(txn, s.selectOrderPaymentListStmts).QueryContext(ctx)
	defer row.Close()
	if err != nil {
		fmt.Print("selectOrderPaymentList error:", err)
		return nil, err
	}
	for row.Next() {
		var item types.OrderPayment
		err := row.Scan(
			&item.PaymentId,
			&item.PayOrderNum,
			&item.PrepayId,
			&item.MchId,
			&item.OpenId,
			&item.CustomerId,
			&item.ShopId,
			&item.PayAmount,
			&item.PayFee,
			&item.PayStatus,
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
