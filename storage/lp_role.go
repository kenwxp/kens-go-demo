package storage

import (
	"context"
	"database/sql"
	"fmt"
	"kens/demo/storage/types"
	"kens/demo/util"
)

//CREATE SEQUENCE IF NOT EXISTS role_seq START 1;
//goland:noinspection SqlNoDataSourceInspection
const roleSchema = `
 	CREATE TABLE IF NOT EXISTS lp_role  (
		role_id		INTEGER primary key ,
		role_name	text  ,
		menu_items	text  ,
		create_time	text  ,
		update_time	text  ,
		is_valid	text                
    );
	comment on column lp_role.role_id		is	'角色ID';
	comment on column lp_role.role_name	is	'角色名 ';
	comment on column lp_role.menu_items	is	'菜单项 ';
	comment on column lp_role.create_time	is	'创建时间（10位数字时间戳）';
	comment on column lp_role.update_time	is	'更新时间（10位数字时间戳）';
	comment on column lp_role.is_valid		is	'启用标志(0-启用 1-禁用) ';
`
const selectRoleListSQL = "" +
	" select " +
	"  role_id	   ," +
	"  role_name   ," +
	"  menu_items  ," +
	"  create_time ," +
	"  update_time ," +
	"  is_valid    " +
	" from lp_role" +
	"  where is_valid = '0'"

const getRoleListSQL = "" +
	" select " +
	"  role_id	   ," +
	"  role_name   ," +
	"  menu_items  ," +
	"  create_time ," +
	"  update_time ," +
	"  is_valid    " +
	" from lp_role"

const insertRoleSQL = "" +
	" INSERT INTO lp_role (" +
	"  role_id," +
	"  role_name," +
	"  menu_items," +
	"  create_time," +
	"  update_time," +
	"  is_valid )" +
	" VALUES (" +
	" (case when $1 = '管理员' then 1 " +
	"       when $1 = '云值守员' then 2 " +
	"       when $1 = '工客服人员' then 3 " +
	"       when $1 = '开发人员' then 8 " +
	"       when $1 = '超级管理员' then 9 end), " +
	" $1,$2,$3,$4,$5) RETURNING role_id"

const deleteRoleSQL = "" +
	" UPDATE" +
	"   lp_role" +
	" SET is_valid = 1" +
	"  where role_id = $1 "

const editRoleSQL = "" +
	" UPDATE" +
	"   lp_role" +
	" SET role_name = $2, menu_items = $3, update_time = $4," +
	" role_id = ( " +
	"  case when $2 = '管理员' then 1 " +
	"       when $2 = '云值守员' then 2 " +
	"       when $2 = '工客服人员' then 3 " +
	"       when $2 = '开发人员' then 8 " +
	"       when $2 = '超级管理员' then 9 end) " +
	"  where role_id = $1 "

type roleStatements struct {
	selectRoleListStmts *sql.Stmt
	getRoleListStmts    *sql.Stmt
	insertRoleStmts     *sql.Stmt
	deleteRoleStmts     *sql.Stmt
	editRoleStmts       *sql.Stmt
}

func (s *roleStatements) execSchema(db *sql.DB) error {
	_, err := db.Exec(roleSchema)
	return err
}

func (s *roleStatements) prepare(db *sql.DB) (err error) {
	if s.selectRoleListStmts, err = db.Prepare(selectRoleListSQL); err != nil {
		return
	}
	if s.getRoleListStmts, err = db.Prepare(getRoleListSQL); err != nil {
		return
	}
	if s.insertRoleStmts, err = db.Prepare(insertRoleSQL); err != nil {
		return
	}
	if s.deleteRoleStmts, err = db.Prepare(deleteRoleSQL); err != nil {
		return
	}
	if s.editRoleStmts, err = db.Prepare(editRoleSQL); err != nil {
		return
	}
	return
}

func (s *roleStatements) selectRoleList(ctx context.Context, txn *sql.Tx) ([]types.Role, error) {
	list := make([]types.Role, 0)
	row, err := TxStmt(txn, s.selectRoleListStmts).QueryContext(ctx)
	defer row.Close()
	if err != nil {
		fmt.Print("selectRoleList error:", err)
		return nil, err
	}
	for row.Next() {
		var item types.Role
		err := row.Scan(
			&item.RoleId,
			&item.RoleName,
			&item.MenuItems,
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

func (s *roleStatements) getRoleList(ctx context.Context, txn *sql.Tx) ([]types.Role, error) {
	list := make([]types.Role, 0)
	row, err := TxStmt(txn, s.getRoleListStmts).QueryContext(ctx)
	defer row.Close()
	if err != nil {
		fmt.Print("getRoleList error:", err)
		return nil, err
	}
	for row.Next() {
		var item types.Role
		err := row.Scan(
			&item.RoleId,
			&item.RoleName,
			&item.MenuItems,
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

func (s *roleStatements) insertRole(ctx context.Context, txn *sql.Tx, role *types.Role) (roleId int64, err error) {
	err = TxStmt(txn, s.insertRoleStmts).QueryRowContext(ctx,
		role.RoleName,
		role.MenuItems,
		role.CreateTime,
		role.UpdateTime,
		role.IsValid,
	).Scan(&roleId)

	if err != nil {
		fmt.Print("insertRole error:", err)
	}

	return
}

func (s *roleStatements) deleteRole(ctx context.Context, txn *sql.Tx, RoleId int64) error {
	_, err := TxStmt(txn, s.deleteRoleStmts).ExecContext(ctx, RoleId)
	if err != nil {
		return err
	}
	return nil
}

func (s *roleStatements) editRole(ctx context.Context, txn *sql.Tx, role *types.Role) error {
	updateTs := util.TimeNowUnixStr()
	_, err := TxStmt(txn, s.editRoleStmts).ExecContext(ctx,
		role.RoleId,
		role.RoleName,
		role.MenuItems,
		updateTs,
	)
	if err != nil {
		fmt.Print("editRole error:", err)
	}
	return nil
}
