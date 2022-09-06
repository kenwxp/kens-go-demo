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
const clientSchema = `
 	CREATE TABLE IF NOT EXISTS lp_client  (
     	client_id		serial primary key,
		client_name		text,
		phone			text,
		password		text,
		salt			text,
		token			text,
		create_time		text,
		update_time		text,
		logon_time		text,
		is_valid		INTEGER,
		unique(phone)	                                 
    );

	comment on column lp_client.client_id		is '自增	客户id';
	comment on column lp_client.client_name	is '客户名';
	comment on column lp_client.phone			is '手机';
	comment on column lp_client.password		is '登录密码';
	comment on column lp_client.salt			is '盐';
	comment on column lp_client.token			is 'token';
	comment on column lp_client.create_time	is '创建时间   （10位数字时间戳）';
	comment on column lp_client.update_time	is '更新时间   （10位数字时间戳）';
	comment on column lp_client.logon_time		is '登录时间   （10位数字时间戳）';
	comment on column lp_client.is_valid		is '启用标志   (0-启用 1-禁用)';
`
const selectClientByTokenSQL = "" +
	" select " +
	"  client_id   ," +
	"  client_name ," +
	"  phone," +
	"  password," +
	"  salt," +
	"  token," +
	"  create_time," +
	"  update_time," +
	"  logon_time," +
	"  is_valid " +
	" from lp_client" +
	"  where token like $1" +
	"    and is_valid = 0"

const updateClientTokenSQL = "" +
	"UPDATE lp_client SET token = $2, logon_time = $3 WHERE client_id = $1"

const selectClientByPhoneSQL = "" +
	" select " +
	"  client_id   ," +
	"  client_name ," +
	"  phone," +
	"  password," +
	"  salt," +
	"  token," +
	"  create_time," +
	"  update_time," +
	"  logon_time," +
	"  is_valid " +
	" from lp_client" +
	"  where phone = $1" +
	"    and is_valid = 0"

const updateClientPasswordSQL = "" +
	"UPDATE lp_client SET salt = $2, password = $3, update_time=$4 WHERE client_id = $1"

const insertClientSQL = "" +
	" INSERT INTO lp_client (" +
	"  client_name ," +
	"  phone," +
	"  password," +
	"  salt," +
	"  token," +
	"  create_time," +
	"  update_time," +
	"  logon_time," +
	"  is_valid )" +
	" VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING client_id"

const deleteClientByClientIdSQL = "" +
	" UPDATE" +
	"   lp_client" +
	" SET is_valid = 1," +
	"   update_time = $2" +
	"  where client_id = $1 "

const editClientSQL = "" +
	" UPDATE" +
	" lp_client SET client_name = $2, phone = $3, update_time = $4" +
	" where client_id = $1"

const selectClientListWithConditionSQL = "" +
	" select " +
	"  client_id   ," +
	"  client_name ," +
	"  phone," +
	"  password," +
	"  salt," +
	"  token," +
	"  create_time," +
	"  update_time," +
	"  logon_time," +
	"  is_valid " +
	" from lp_client" +
	"  where " +
	"   ('%%'= $1 or client_name like $1 )" +
	"  and ( '' = $2 or phone = $2)" +
	"  and is_valid = 0" +
	" order by create_time desc"

const selectClientByClientIdSQL = "" +
	" select " +
	"  client_id   ," +
	"  client_name ," +
	"  phone," +
	"  password," +
	"  salt," +
	"  token," +
	"  create_time," +
	"  update_time," +
	"  logon_time," +
	"  is_valid " +
	" from lp_client" +
	"  where client_id = $1" +
	"    and is_valid = 0"

type clientStatements struct {
	selectClientByTokenStmts           *sql.Stmt
	updateClientTokenStmts             *sql.Stmt
	selectClientByPhoneStmts           *sql.Stmt
	updateClientPasswordStmts          *sql.Stmt
	insertClientStmts                  *sql.Stmt
	deleteClientByClientIdStmts        *sql.Stmt
	editClientStmts                    *sql.Stmt
	selectClientListWithConditionStmts *sql.Stmt
	selectClientByClientIdStmts        *sql.Stmt
}

func (s *clientStatements) execSchema(db *sql.DB) error {
	_, err := db.Exec(clientSchema)
	return err
}

func (s *clientStatements) prepare(db *sql.DB) (err error) {
	if s.selectClientByTokenStmts, err = db.Prepare(selectClientByTokenSQL); err != nil {
		return
	}
	if s.updateClientTokenStmts, err = db.Prepare(updateClientTokenSQL); err != nil {
		return
	}
	if s.selectClientByPhoneStmts, err = db.Prepare(selectClientByPhoneSQL); err != nil {
		return
	}
	if s.updateClientPasswordStmts, err = db.Prepare(updateClientPasswordSQL); err != nil {
		return
	}
	if s.insertClientStmts, err = db.Prepare(insertClientSQL); err != nil {
		return
	}
	if s.deleteClientByClientIdStmts, err = db.Prepare(deleteClientByClientIdSQL); err != nil {
		return
	}
	if s.editClientStmts, err = db.Prepare(editClientSQL); err != nil {
		return
	}
	if s.selectClientListWithConditionStmts, err = db.Prepare(selectClientListWithConditionSQL); err != nil {
		return
	}
	if s.selectClientByClientIdStmts, err = db.Prepare(selectClientByClientIdSQL); err != nil {
		return
	}
	return
}

func (s *clientStatements) selectClientListWithCondition(ctx context.Context, txn *sql.Tx, clientName string, phone string) ([]types.Client, error) {
	list := make([]types.Client, 0)
	row, err := TxStmt(txn, s.selectClientListWithConditionStmts).QueryContext(ctx, "%"+clientName+"%", phone)
	defer row.Close()
	if err != nil {
		fmt.Print("selectClientListWithCondition error:", err)
		return nil, err
	}
	for row.Next() {
		var item types.Client
		err := row.Scan(
			&item.ClientId,
			&item.ClientName,
			&item.ClientPhone,
			&item.Password,
			&item.Salt,
			&item.Token,
			&item.CreateTime,
			&item.UpdateTime,
			&item.LogonTime,
			&item.IsValid,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}

func (s *clientStatements) selectClientByToken(ctx context.Context, txn *sql.Tx, token string) (*types.Client, error) {
	row := TxStmt(txn, s.selectClientByTokenStmts).QueryRowContext(ctx, "%"+token+"%")
	item := &types.Client{}
	if err := row.Scan(
		&item.ClientId,
		&item.ClientName,
		&item.ClientPhone,
		&item.Password,
		&item.Salt,
		&item.Token,
		&item.CreateTime,
		&item.UpdateTime,
		&item.LogonTime,
		&item.IsValid,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return item, nil
}

func (s *clientStatements) updateClientToken(ctx context.Context, txn *sql.Tx, clientId int64, token string) (err error) {
	stmt := TxStmt(txn, s.updateClientTokenStmts)
	_, err = stmt.ExecContext(ctx, clientId, token, util.TimeNowUnixStr())
	if err != nil {
		return err
	}
	return
}

func (s *clientStatements) selectClientByPhone(ctx context.Context, txn *sql.Tx, phone string) (*types.Client, error) {
	row := TxStmt(txn, s.selectClientByPhoneStmts).QueryRowContext(ctx, phone)
	item := &types.Client{}
	if err := row.Scan(
		&item.ClientId,
		&item.ClientName,
		&item.ClientPhone,
		&item.Password,
		&item.Salt,
		&item.Token,
		&item.CreateTime,
		&item.UpdateTime,
		&item.LogonTime,
		&item.IsValid,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return item, nil
}

func (s *clientStatements) updateClientPassword(ctx context.Context, txn *sql.Tx, clientId int64, salt string, password string) (err error) {
	updateTs := util.TimeNowUnixStr()
	stmt := TxStmt(txn, s.updateClientPasswordStmts)
	_, err = stmt.ExecContext(ctx, clientId, salt, password, updateTs)
	if err != nil {
		return err
	}
	return
}

func (s *clientStatements) insertClient(ctx context.Context, txn *sql.Tx, client *types.Client) (clientId int64, err error) {
	stmt := TxStmt(txn, s.insertClientStmts)
	err = stmt.QueryRowContext(ctx,
		client.ClientName,
		client.ClientPhone,
		client.Password,
		client.Salt,
		client.Token,
		client.CreateTime,
		client.UpdateTime,
		client.LogonTime,
		client.IsValid,
	).Scan(&clientId)
	if err != nil {
		fmt.Print("insertClient error:", err)
	}
	return
}

func (s *clientStatements) deleteClientByClientId(ctx context.Context, txn *sql.Tx, clientId int64) error {
	_, err := TxStmt(txn, s.deleteClientByClientIdStmts).ExecContext(ctx, clientId, util.TimeNowUnixStr())
	if err != nil {
		return err
	}
	return nil
}

func (s *clientStatements) editClient(ctx context.Context, txn *sql.Tx, clientId int64, clientName string, phone string) (err error) {
	stmt := TxStmt(txn, s.editClientStmts)
	_, err = stmt.ExecContext(ctx, clientId, clientName, phone, util.TimeNowUnixStr())
	if err != nil {
		return err
	}
	return
}

func (s *clientStatements) selectClientByClientId(ctx context.Context, txn *sql.Tx, clientId int64) (*types.Client, error) {
	row := TxStmt(txn, s.selectClientByClientIdStmts).QueryRowContext(ctx, clientId)
	item := &types.Client{}
	if err := row.Scan(
		&item.ClientId,
		&item.ClientName,
		&item.ClientPhone,
		&item.Password,
		&item.Salt,
		&item.Token,
		&item.CreateTime,
		&item.UpdateTime,
		&item.LogonTime,
		&item.IsValid,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return item, nil
}
