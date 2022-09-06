package storage

import (
	"context"
	"database/sql"
	"fmt"
	"kens/demo/storage/types"
	"kens/demo/util"
)

//CREATE SEQUENCE IF NOT EXISTS user_seq START 1;
//goland:noinspection SqlNoDataSourceInspection
const customerSchema = `
 	CREATE TABLE IF NOT EXISTS lp_customer  (
     	customer_id		serial primary key,
		nickname		text,
		avatar			text,
		phone			text,
		gender			text,
 		birth			text,
		token			text,
		open_id			text,
		session_key		text,
		is_member		INTEGER,
		join_time		text,
		join_end_time	text,
		create_time		text,
		update_time		text,
		is_black		INTEGER	                                
    );
	comment on column lp_customer.customer_id		is  '自增	顾客id	pk ';
	comment on column lp_customer.nickname		is  '用户昵称';	
	comment on column lp_customer.avatar		is  '用户头像';	
	comment on column lp_customer.phone			is  '手机号';	
	comment on column lp_customer.gender			is  '性别 0-男 1-女';	
	comment on column lp_customer.birth			is  '生日 yyyy-mm-dd';	
	comment on column lp_customer.token			is  'token';	
	comment on column lp_customer.open_id			is  '微信open_id';
	comment on column lp_customer.session_key		is  '微信session_key';
	comment on column lp_customer.is_member		is  '是否会员 0-否 1-是';
	comment on column lp_customer.join_time		is '加入会员时间   （10位数字时间戳）';
	comment on column lp_customer.join_end_time	is '退出会员时间   （10位数字时间戳）';
	comment on column lp_customer.create_time		is '创建时间   （10位数字时间戳）';
	comment on column lp_customer.update_time		is '更新时间   （10位数字时间戳）';
	comment on column lp_customer.is_black		is  '是否黑名单 0-否 1-是';	
`
const selectCustomerByTokenSQL = "" +
	" select " +
	"  customer_id	 ," +
	"  nickname	 	 ," +
	"  avatar	 	 ," +
	"  phone		 ," +
	"  gender		 ," +
	"  birth		 ," +
	"  token		 ," +
	"  open_id		 ," +
	"  session_key	 ," +
	"  is_member	 ," +
	"  join_time	 ," +
	"  join_end_time ," +
	"  create_time	 ," +
	"  update_time	 ," +
	"  is_black	 " +
	" from lp_customer" +
	"  where token like $1"

const updateCustomerTokenSQL = "" +
	" UPDATE " +
	"  lp_customer " +
	" SET " +
	"  token = $2, " +
	"  session_key = $3," +
	"  update_time = $4 " +
	" WHERE " +
	"  customer_id = $1"

const selectCustomerByOpenIdSQL = "" +
	" select " +
	"  customer_id		  ," +
	"  nickname		  	  ," +
	"  avatar		  	  ," +
	"  phone			  ," +
	"  gender			  ," +
	"  birth			  ," +
	"  token			  ," +
	"  open_id			  ," +
	"  session_key		  ," +
	"  is_member		  ," +
	"  join_time		  ," +
	"  join_end_time	  ," +
	"  create_time		  ," +
	"  update_time		  ," +
	"  is_black	   		" +
	" from lp_customer" +
	"  where open_id = $1"

const insertCustomerSQL = "" +
	"INSERT INTO lp_customer " +
	" ( " +
	"  nickname		 	  ," +
	"  avatar		 	  ," +
	"  phone			  ," +
	"  gender			  ," +
	"  birth			  ," +
	"  token			  ," +
	"  open_id			  ," +
	"  session_key		  ," +
	"  is_member		  ," +
	"  join_time		  ," +
	"  join_end_time	  ," +
	"  create_time		  ," +
	"  update_time		  ," +
	"  is_black	   		  " +
	") " +
	" VALUES " +
	"($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14) RETURNING customer_id"

type customerStatements struct {
	selectCustomerByTokenStmts  *sql.Stmt
	updateCustomerTokenStmts    *sql.Stmt
	selectCustomerByOpenIdStmts *sql.Stmt
	insertCustomerStmts         *sql.Stmt
}

func (s *customerStatements) execSchema(db *sql.DB) error {
	_, err := db.Exec(customerSchema)
	return err
}

func (s *customerStatements) prepare(db *sql.DB) (err error) {
	if s.selectCustomerByTokenStmts, err = db.Prepare(selectCustomerByTokenSQL); err != nil {
		return
	}
	if s.updateCustomerTokenStmts, err = db.Prepare(updateCustomerTokenSQL); err != nil {
		return
	}
	if s.selectCustomerByOpenIdStmts, err = db.Prepare(selectCustomerByOpenIdSQL); err != nil {
		return
	}
	if s.insertCustomerStmts, err = db.Prepare(insertCustomerSQL); err != nil {
		return
	}
	return
}

func (s *customerStatements) selectCustomerByToken(ctx context.Context, txn *sql.Tx, token string) (*types.Customer, error) {
	row := TxStmt(txn, s.selectCustomerByTokenStmts).QueryRowContext(ctx, "%"+token+"%")
	item := &types.Customer{}
	if err := row.Scan(
		&item.CustomerId,
		&item.Nickname,
		&item.Avatar,
		&item.Phone,
		&item.Gender,
		&item.Birth,
		&item.Token,
		&item.OpenId,
		&item.SessionKey,
		&item.IsMember,
		&item.JoinTime,
		&item.JoinEndTime,
		&item.CreateTime,
		&item.UpdateTime,
		&item.IsBlack,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return item, nil
}

func (s *customerStatements) updateCustomerToken(ctx context.Context, txn *sql.Tx, customerId int64, token string, sessionKey string) (err error) {
	updateTs := util.TimeNowUnixStr()
	stmt := TxStmt(txn, s.updateCustomerTokenStmts)
	_, err = stmt.ExecContext(ctx, customerId, token, sessionKey, updateTs)
	if err != nil {
		return err
	}
	return
}

func (s *customerStatements) selectCustomerByOpenId(ctx context.Context, txn *sql.Tx, openId string) (*types.Customer, error) {
	row := TxStmt(txn, s.selectCustomerByOpenIdStmts).QueryRowContext(ctx, openId)
	item := &types.Customer{}
	if err := row.Scan(
		&item.CustomerId,
		&item.Nickname,
		&item.Avatar,
		&item.Phone,
		&item.Gender,
		&item.Birth,
		&item.Token,
		&item.OpenId,
		&item.SessionKey,
		&item.IsMember,
		&item.JoinTime,
		&item.JoinEndTime,
		&item.CreateTime,
		&item.UpdateTime,
		&item.IsBlack,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return item, nil
}

func (s *customerStatements) insertCustomer(ctx context.Context, txn *sql.Tx, customer *types.Customer) (customerId int64, err error) {
	stmt := TxStmt(txn, s.insertCustomerStmts)
	err = stmt.QueryRowContext(ctx,
		customer.Nickname,
		customer.Avatar,
		customer.Phone,
		customer.Gender,
		customer.Birth,
		customer.Token,
		customer.OpenId,
		customer.SessionKey,
		customer.IsMember,
		customer.JoinTime,
		customer.JoinEndTime,
		customer.CreateTime,
		customer.UpdateTime,
		customer.IsBlack,
	).Scan(&customerId)
	if err != nil {
		fmt.Print("insertCustomer error:", err)
	}
	return
}
