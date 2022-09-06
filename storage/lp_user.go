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
const userSchema = `
 	CREATE TABLE IF NOT EXISTS lp_user  (
     	acc_no		text   primary key,
     	acc_name	text,
		password	text   ,
		salt		text   ,
		token		text   ,
		phone		text   ,
		email		text   ,
		role_id		INTEGER,
		on_work		text,
		create_time	text  ,
		update_time	text  ,
		logon_time	text  ,
		is_valid	INTEGER                            
    );
	comment on column lp_user.acc_no		is	'工号、登录账号';
	comment on column lp_user.acc_name		is	'员工名';
	comment on column lp_user.password		is	'密码';
	comment on column lp_user.salt			is	'盐';
	comment on column lp_user.token		is	'token';
	comment on column lp_user.phone		is	'手机号';
	comment on column lp_user.email		is	'邮箱';
	comment on column lp_user.role_id			is	'角色权限（1-审核员 2-云值守员 3-工客服人员 8-开发人员 9-超级管理员）';
	comment on column lp_user.on_work		is	'在岗状态（0-离岗 1-在岗）';
	comment on column lp_user.create_time	is	'创建时间（10位数字时间戳）';
	comment on column lp_user.update_time	is	'更新时间（10位数字时间戳）';
	comment on column lp_user.logon_time	is	'登录时间（10位数字时间戳）';
	comment on column lp_user.is_valid		is	'启用标志(0-启用 1-禁用)';
`
const selectUserListWithConditionSQL = "" +
	" select " +
	"  acc_no     ," +
	"  acc_name     ," +
	"  password   ," +
	"  salt       ," +
	"  token      ," +
	"  phone      ," +
	"  email      ," +
	"  role_id       ," +
	"  on_work    ," +
	"  create_time," +
	"  update_time," +
	"  logon_time," +
	"  is_valid " +
	" from lp_user" +
	"  where is_valid = 0" +
	"    and ('' = $1 or acc_no = $1) " +
	"    and ('%%' = $2 or acc_name like $2)" +
	"    and ( 0 = $3 or role_id = $3 )" +
	"    and ('' = $4 or on_work = $4)" +
	" order by create_time desc"

const selectUserByTokenSQL = "" +
	" select " +
	"  acc_no     ," +
	"  acc_name     ," +
	"  password   ," +
	"  salt       ," +
	"  token      ," +
	"  phone      ," +
	"  email      ," +
	"  role_id       ," +
	"  on_work    ," +
	"  create_time," +
	"  update_time," +
	"  logon_time," +
	"  is_valid " +
	" from lp_user" +
	"  where token like $1"

const updateUserTokenSQL = "" +
	"UPDATE lp_user SET token = $2, update_time = $3 WHERE acc_no = $1"

const updateUserLogonSQL = "" +
	"UPDATE lp_user SET on_work = $2, logon_time = $3 WHERE acc_no = $1"

const selectUserByAccNoSQL = "" +
	" select " +
	"  acc_no     ," +
	"  acc_name     ," +
	"  password   ," +
	"  salt       ," +
	"  token      ," +
	"  phone      ," +
	"  email      ," +
	"  role_id       ," +
	"  on_work    ," +
	"  create_time," +
	"  update_time," +
	"  logon_time," +
	"  is_valid " +
	" from lp_user" +
	"  where acc_no = $1" +
	"    and is_valid = 0"

const selectUserByPhoneSQL = "" +
	" select " +
	"  acc_no     ," +
	"  acc_name     ," +
	"  password   ," +
	"  salt       ," +
	"  token      ," +
	"  phone      ," +
	"  email      ," +
	"  role_id       ," +
	"  on_work    ," +
	"  create_time," +
	"  update_time," +
	"  logon_time," +
	"  is_valid " +
	" from lp_user" +
	"  where phone = $1"

const updateUserPasswordSQL = "" +
	"UPDATE lp_user SET salt = $2, password = $3, update_time=$4 WHERE acc_no = $1"

const selectMaxAccNoByRoleIdSQL = "" +
	"select coalesce(max(acc_no::int8),0) from lp_user where role_id = $1"

const insertUserSQL = "" +
	" INSERT INTO lp_user (" +
	"  acc_no," +
	"  acc_name," +
	"  password," +
	"  salt," +
	"  token," +
	"  phone," +
	"  email," +
	"  role_id," +
	"  on_work," +
	"  create_time," +
	"  update_time," +
	"  logon_time," +
	"  is_valid )" +
	" VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12, $13)"

const deleteUserSQL = "" +
	" UPDATE" +
	"   lp_user" +
	" SET is_valid = 1," +
	"	update_time = $2" +
	"  where acc_no = $1 "

const editUserSQL = "" +
	" UPDATE" +
	" lp_user SET acc_name = $2, phone = $3, email = $4, update_time = $5" +
	" where acc_no = $1"

type userStatements struct {
	selectUserListWithConditionStmt *sql.Stmt
	selectUserByTokenStmts          *sql.Stmt
	updateUserTokenStmts            *sql.Stmt
	updateUserLogonStmts            *sql.Stmt
	selectUserByAccNoStmts          *sql.Stmt
	selectUserByPhoneStmts          *sql.Stmt
	updateUserPasswordStmts         *sql.Stmt
	selectMaxAccNoByRoleIdStmts     *sql.Stmt
	insertUserStmts                 *sql.Stmt
	deleteUserStmts                 *sql.Stmt
	editUserStmts                   *sql.Stmt
}

func (s *userStatements) execSchema(db *sql.DB) error {
	_, err := db.Exec(userSchema)
	return err
}

func (s *userStatements) prepare(db *sql.DB) (err error) {
	if s.selectUserListWithConditionStmt, err = db.Prepare(selectUserListWithConditionSQL); err != nil {
		return
	}
	if s.selectUserByTokenStmts, err = db.Prepare(selectUserByTokenSQL); err != nil {
		return
	}
	if s.updateUserTokenStmts, err = db.Prepare(updateUserTokenSQL); err != nil {
		return
	}
	if s.updateUserLogonStmts, err = db.Prepare(updateUserLogonSQL); err != nil {
		return
	}
	if s.selectUserByAccNoStmts, err = db.Prepare(selectUserByAccNoSQL); err != nil {
		return
	}
	if s.selectUserByPhoneStmts, err = db.Prepare(selectUserByPhoneSQL); err != nil {
		return
	}
	if s.updateUserPasswordStmts, err = db.Prepare(updateUserPasswordSQL); err != nil {
		return
	}
	if s.selectMaxAccNoByRoleIdStmts, err = db.Prepare(selectMaxAccNoByRoleIdSQL); err != nil {
		return
	}
	if s.insertUserStmts, err = db.Prepare(insertUserSQL); err != nil {
		return
	}
	if s.deleteUserStmts, err = db.Prepare(deleteUserSQL); err != nil {
		return
	}
	if s.editUserStmts, err = db.Prepare(editUserSQL); err != nil {
		return
	}
	return
}

func (s *userStatements) selectUserListWithCondition(ctx context.Context, txn *sql.Tx, input types.User) ([]types.User, error) {
	list := make([]types.User, 0)
	row, err := TxStmt(txn, s.selectUserListWithConditionStmt).QueryContext(ctx,
		input.AccNo,
		"%"+input.AccName+"%",
		input.RoleId,
		input.OnWork,
	)
	defer row.Close()
	if err != nil {
		fmt.Print("selectUserListWithCondition error:", err)
		return nil, err
	}
	for row.Next() {
		var item types.User
		err := row.Scan(
			&item.AccNo,
			&item.AccName,
			&item.Password,
			&item.Salt,
			&item.Token,
			&item.Phone,
			&item.Email,
			&item.RoleId,
			&item.OnWork,
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

func (s *userStatements) selectUserByToken(ctx context.Context, txn *sql.Tx, token string) (*types.User, error) {
	row := TxStmt(txn, s.selectUserByTokenStmts).QueryRowContext(ctx, "%"+token+"%")
	item := &types.User{}
	if err := row.Scan(
		&item.AccNo,
		&item.AccName,
		&item.Password,
		&item.Salt,
		&item.Token,
		&item.Phone,
		&item.Email,
		&item.RoleId,
		&item.OnWork,
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

func (s *userStatements) updateUserToken(ctx context.Context, txn *sql.Tx, accNo string, token string) (err error) {
	stmt := TxStmt(txn, s.updateUserTokenStmts)
	_, err = stmt.ExecContext(ctx, accNo, token, util.TimeNowUnixStr())
	if err != nil {
		return err
	}
	return
}
func (s *userStatements) updateUserLogon(ctx context.Context, txn *sql.Tx, accNo string, onWork string) (err error) {
	stmt := TxStmt(txn, s.updateUserLogonStmts)
	_, err = stmt.ExecContext(ctx, accNo, onWork, util.TimeNowUnixStr())
	if err != nil {
		return err
	}
	return
}

func (s *userStatements) selectUserByAccNo(ctx context.Context, txn *sql.Tx, accNo string) (*types.User, error) {
	row := TxStmt(txn, s.selectUserByAccNoStmts).QueryRowContext(ctx, accNo)
	item := &types.User{}
	if err := row.Scan(
		&item.AccNo,
		&item.AccName,
		&item.Password,
		&item.Salt,
		&item.Token,
		&item.Phone,
		&item.Email,
		&item.RoleId,
		&item.OnWork,
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

func (s *userStatements) selectUserByPhone(ctx context.Context, txn *sql.Tx, phone string) (*types.User, error) {
	row := TxStmt(txn, s.selectUserByAccNoStmts).QueryRowContext(ctx, phone)
	item := &types.User{}
	if err := row.Scan(
		&item.AccNo,
		&item.AccName,
		&item.Password,
		&item.Salt,
		&item.Token,
		&item.Phone,
		&item.Email,
		&item.RoleId,
		&item.OnWork,
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

func (s *userStatements) updateUserPassword(ctx context.Context, txn *sql.Tx, accNo string, salt string, password string) (err error) {
	stmt := TxStmt(txn, s.updateUserPasswordStmts)
	_, err = stmt.ExecContext(ctx, accNo, salt, password, util.TimeNowUnixStr())
	if err != nil {
		return err
	}
	return
}

func (s *userStatements) selectMaxAccNoByRoleId(ctx context.Context, txn *sql.Tx, roleId int64) (maxAccNo int64, err error) {
	if err = TxStmt(txn, s.selectMaxAccNoByRoleIdStmts).QueryRowContext(ctx, roleId).Scan(&maxAccNo); err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	return
}

func (s *userStatements) insertUser(ctx context.Context, txn *sql.Tx, user *types.User) error {
	_, err := TxStmt(txn, s.insertUserStmts).ExecContext(ctx,
		user.AccNo,
		user.AccName,
		user.Password,
		user.Salt,
		user.Token,
		user.Phone,
		user.Email,
		user.RoleId,
		user.OnWork,
		user.CreateTime,
		user.UpdateTime,
		user.LogonTime,
		user.IsValid,
	)
	if err != nil {
		fmt.Print("insertUser error:", err)
		return err
	}
	return nil
}

func (s *userStatements) deleteUser(ctx context.Context, txn *sql.Tx, accNo string) error {
	_, err := TxStmt(txn, s.deleteUserStmts).ExecContext(ctx, accNo, util.TimeNowUnixStr())
	if err != nil {
		return err
	}
	return nil
}

func (s *userStatements) editUser(ctx context.Context, txn *sql.Tx, user *types.User) error {
	_, err := TxStmt(txn, s.editUserStmts).ExecContext(ctx,
		user.AccNo,
		user.AccName,
		user.Phone,
		user.Email,
		util.TimeNowUnixStr(),
	)
	if err != nil {
		fmt.Print("editUser error:", err)
		return err
	}
	return nil
}
