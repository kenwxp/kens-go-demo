package storage

import (
	"context"
	"database/sql"
	"fmt"
	"kens/demo/storage/types"
)

const shippingAddressSchema = `
 	CREATE TABLE IF NOT EXISTS lp_shipping_address  (
     	address_id		serial primary key,
     	customer_id		integer,
     	accept_name		text,
     	mobile			text,
     	area_code		text,
     	street			text,
		door_number		text,
		is_default		integer,
		create_time		text,
		update_time		text,
		is_valid		integer
    );
    comment on column lp_shipping_address.address_id	is '自增 id pk';
    comment on column lp_shipping_address.customer_id	is '顾客id';
    comment on column lp_shipping_address.accept_name	is '收货人';
    comment on column lp_shipping_address.mobile		is '手机';
    comment on column lp_shipping_address.area_code	    is '地区号';
    comment on column lp_shipping_address.street		is '详细街道地址';
    comment on column lp_shipping_address.door_number	is '门牌号';
    comment on column lp_shipping_address.is_default	is '是否默认 0-否 1-是';
    comment on column lp_shipping_address.create_time	is '创建时间';
    comment on column lp_shipping_address.update_time	is '更新时间';
    comment on column lp_shipping_address.is_valid	    is '是否启用 0-启用，1-禁用';
`

const selectShippingAddressListSQL = "" +
	" select " +
	"  address_id	  ," +
	"  customer_id	  ," +
	"  accept_name	  ," +
	"  mobile		  ," +
	"  area_code	  ," +
	"  street		  ," +
	"  door_number	  ," +
	"  is_default	 , " +
	"  create_time	 , " +
	"  update_time	 , " +
	"  is_valid	  " +
	" from lp_shipping_address" +
	"  where is_valid = 0"

type shippingAddressStatements struct {
	selectShippingAddressListStmts *sql.Stmt
}

func (s *shippingAddressStatements) execSchema(db *sql.DB) error {
	_, err := db.Exec(shippingAddressSchema)
	return err
}

func (s *shippingAddressStatements) prepare(db *sql.DB) (err error) {
	if s.selectShippingAddressListStmts, err = db.Prepare(selectShippingAddressListSQL); err != nil {
		return
	}
	return
}

func (s *shippingAddressStatements) selectShippingAddressList(ctx context.Context, txn *sql.Tx) ([]types.ShippingAddress, error) {
	list := make([]types.ShippingAddress, 0)
	row, err := TxStmt(txn, s.selectShippingAddressListStmts).QueryContext(ctx)
	defer row.Close()
	if err != nil {
		fmt.Print("selectShippingAddressList error:", err)
		return nil, err
	}
	for row.Next() {
		var item types.ShippingAddress
		err := row.Scan(
			&item.AddressId,
			&item.CustomerId,
			&item.AcceptName,
			&item.Mobile,
			&item.AreaCode,
			&item.Street,
			&item.DoorNumber,
			&item.IsDefault,
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
