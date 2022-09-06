package storage

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"kens/demo/storage/types"
	"kens/demo/util"
	"kens/demo/util/environment"
)

const (
	//host = "192.168.2.239"
	host     = "127.0.0.1"
	port     = 5501
	user     = "backend"
	password = "liupai666"
	dbname   = "liupai"
)

// Database represents an account database
type Database struct {
	Db              *sql.DB
	user            userStatements
	role            roleStatements
	client          clientStatements
	customer        customerStatements
	region          regionStatements
	shippingAddress shippingAddressStatements
	goods           goodsStatements
	goodsCategory   goodsCategoryStatements
	goodsProperty   goodsPropertyStatements
	order           orderStatements
	orderPayment    orderPaymentStatements
}

func (d *Database) WithTransaction(fn func(txn *sql.Tx) error) (err error) {
	return util.WithTransaction(d.Db, fn)
}

func TxStmt(transaction *sql.Tx, statement *sql.Stmt) *sql.Stmt {
	return util.TxStmt(transaction, statement)
}

// NewDatabase creates a new accounts and profiles database
func NewDatabase() (*Database, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	if environment.ZeeEnv == environment.EnvDev {
		psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			"127.0.0.1", port, user, password, dbname)
	}
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	d := &Database{
		Db: db,
	}
	// Create tables before executing migrations so we don't fail if the table is missing,
	// and THEN prepare statements so we don't fail due to referencing new columns
	//=============== execSchema =========================
	if err = d.user.execSchema(db); err != nil {
		fmt.Print("user execSchema:", err)
		return nil, err
	}
	if err = d.role.execSchema(db); err != nil {
		fmt.Print("role execSchema:", err)
		return nil, err
	}
	if err = d.client.execSchema(db); err != nil {
		fmt.Print("client execSchema:", err)
		return nil, err
	}
	if err = d.customer.execSchema(db); err != nil {
		fmt.Print("customer execSchema:", err)
		return nil, err
	}

	if err = d.region.execSchema(db); err != nil {
		fmt.Print("region execSchema:", err)
		return nil, err
	}
	if err = d.shippingAddress.execSchema(db); err != nil {
		fmt.Print("shippingAddress execSchema:", err)
		return nil, err
	}
	if err = d.goods.execSchema(db); err != nil {
		fmt.Print("goods execSchema:", err)
		return nil, err
	}
	if err = d.goodsCategory.execSchema(db); err != nil {
		fmt.Print("goodsCategory execSchema:", err)
		return nil, err
	}
	if err = d.goodsProperty.execSchema(db); err != nil {
		fmt.Print("goodsProperty execSchema:", err)
		return nil, err
	}
	if err = d.order.execSchema(db); err != nil {
		fmt.Print("order execSchema:", err)
		return nil, err
	}
	if err = d.orderPayment.execSchema(db); err != nil {
		fmt.Print("orderPayment execSchema:", err)
		return nil, err
	}

	//=============== prepare  =========================
	if err = d.user.prepare(db); err != nil {
		fmt.Print("user prepare:", err)
		return nil, err
	}
	if err = d.role.prepare(db); err != nil {
		fmt.Print("role prepare:", err)
		return nil, err
	}
	if err = d.client.prepare(db); err != nil {
		fmt.Print("client prepare:", err)
		return nil, err
	}
	if err = d.customer.prepare(db); err != nil {
		fmt.Print("customer prepare:", err)
		return nil, err
	}

	if err = d.region.prepare(db); err != nil {
		fmt.Print("region prepare:", err)
		return nil, err
	}
	if err = d.shippingAddress.prepare(db); err != nil {
		fmt.Print("shippingAddress prepare:", err)
		return nil, err
	}
	if err = d.goods.prepare(db); err != nil {
		fmt.Print("goods prepare:", err)
		return nil, err
	}
	if err = d.goodsCategory.prepare(db); err != nil {
		fmt.Print("goodsCategory prepare:", err)
		return nil, err
	}
	if err = d.goodsProperty.prepare(db); err != nil {
		fmt.Print("goodsProperty prepare:", err)
		return nil, err
	}
	if err = d.order.prepare(db); err != nil {
		fmt.Print("order prepare:", err)
		return nil, err
	}
	if err = d.orderPayment.prepare(db); err != nil {
		fmt.Print("orderPayment prepare:", err)
		return nil, err
	}
	return d, nil
}

func (d *Database) SelectUserListWithCondition(ctx context.Context, txn *sql.Tx, input types.User) ([]types.User, error) {
	return d.user.selectUserListWithCondition(ctx, txn, input)
}

func (d *Database) SelectUserByToken(ctx context.Context, txn *sql.Tx, token string) (*types.User, error) {
	return d.user.selectUserByToken(ctx, txn, token)
}

func (d *Database) UpdateUserToken(ctx context.Context, txn *sql.Tx, accNo string, token string) error {
	return d.user.updateUserToken(ctx, txn, accNo, token)
}
func (d *Database) UpdateUserLogon(ctx context.Context, txn *sql.Tx, accNo string, onWork string) (err error) {
	return d.user.updateUserLogon(ctx, txn, accNo, onWork)
}
func (d *Database) SelectUserByAccNo(ctx context.Context, txn *sql.Tx, accNo string) (*types.User, error) {
	return d.user.selectUserByAccNo(ctx, txn, accNo)
}
func (d *Database) SelectUserByPhone(ctx context.Context, txn *sql.Tx, phone string) (*types.User, error) {
	return d.user.selectUserByPhone(ctx, txn, phone)
}
func (d *Database) UpdateUserPassword(ctx context.Context, txn *sql.Tx, accNo string, salt string, password string) error {
	return d.user.updateUserPassword(ctx, txn, accNo, salt, password)
}

func (d *Database) InsertUser(ctx context.Context, txn *sql.Tx, user *types.User) error {
	return d.user.insertUser(ctx, txn, user)
}
func (d *Database) DeleteUser(ctx context.Context, txn *sql.Tx, accNo string) error {
	return d.user.deleteUser(ctx, txn, accNo)
}
func (d *Database) EditUser(ctx context.Context, txn *sql.Tx, user *types.User) error {
	return d.user.editUser(ctx, txn, user)
}
func (d *Database) SelectMaxAccNoByRoleId(ctx context.Context, txn *sql.Tx, roleId int64) (int64, error) {
	return d.user.selectMaxAccNoByRoleId(ctx, txn, roleId)
}

func (d *Database) SelectRoleList(ctx context.Context, txn *sql.Tx) ([]types.Role, error) {
	return d.role.selectRoleList(ctx, txn)
}
func (d *Database) GetRoleList(ctx context.Context, txn *sql.Tx) ([]types.Role, error) {
	return d.role.getRoleList(ctx, txn)
}
func (d *Database) InsertRole(ctx context.Context, txn *sql.Tx, req *types.Role) (int64, error) {
	return d.role.insertRole(ctx, txn, req)
}
func (d *Database) DeleteRole(ctx context.Context, txn *sql.Tx, roleId int64) error {
	return d.role.deleteRole(ctx, txn, roleId)
}
func (d *Database) EditRole(ctx context.Context, txn *sql.Tx, req *types.Role) error {
	return d.role.editRole(ctx, txn, req)
}

func (d *Database) SelectClientListWithCondition(ctx context.Context, txn *sql.Tx, clientName string, phone string) ([]types.Client, error) {
	return d.client.selectClientListWithCondition(ctx, txn, clientName, phone)
}

func (d *Database) SelectClientByToken(ctx context.Context, txn *sql.Tx, token string) (*types.Client, error) {
	return d.client.selectClientByToken(ctx, txn, token)
}

func (d *Database) UpdateClientToken(ctx context.Context, txn *sql.Tx, clientId int64, token string) error {
	return d.client.updateClientToken(ctx, txn, clientId, token)
}
func (d *Database) SelectClientByPhone(ctx context.Context, txn *sql.Tx, phone string) (*types.Client, error) {
	return d.client.selectClientByPhone(ctx, txn, phone)
}
func (d *Database) UpdateClientPassword(ctx context.Context, txn *sql.Tx, clientId int64, salt string, password string) error {
	return d.client.updateClientPassword(ctx, txn, clientId, salt, password)
}
func (d *Database) InsertClient(ctx context.Context, txn *sql.Tx, client *types.Client) (int64, error) {
	return d.client.insertClient(ctx, txn, client)
}
func (d *Database) DeleteClientByClientId(ctx context.Context, txn *sql.Tx, clientId int64) error {
	return d.client.deleteClientByClientId(ctx, txn, clientId)
}
func (d *Database) EditClient(ctx context.Context, txn *sql.Tx, clientId int64, clientName string, phone string) error {
	return d.client.editClient(ctx, txn, clientId, clientName, phone)
}
func (d *Database) SelectClientByClientId(ctx context.Context, txn *sql.Tx, clientId int64) (*types.Client, error) {
	return d.client.selectClientByClientId(ctx, txn, clientId)
}

func (d *Database) SelectCustomerByToken(ctx context.Context, txn *sql.Tx, token string) (*types.Customer, error) {
	return d.customer.selectCustomerByToken(ctx, txn, token)
}

func (d *Database) UpdateCustomerToken(ctx context.Context, txn *sql.Tx, customerId int64, token string, sessionKey string) error {
	return d.customer.updateCustomerToken(ctx, txn, customerId, token, sessionKey)
}

func (d *Database) SelectCustomerByOpenId(ctx context.Context, txn *sql.Tx, openId string) (*types.Customer, error) {
	return d.customer.selectCustomerByOpenId(ctx, txn, openId)
}

func (d *Database) InsertCustomer(ctx context.Context, txn *sql.Tx, customer *types.Customer) (int64, error) {
	return d.customer.insertCustomer(ctx, txn, customer)
}
func (d *Database) SelectRegionList(ctx context.Context, txn *sql.Tx) ([]types.Region, error) {
	return d.region.selectRegionList(ctx, txn)
}
