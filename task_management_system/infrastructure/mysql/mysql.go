package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"reflect"
	"time"

	"task_management_system/appcontext"
	"task_management_system/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// IMysqlStore interface

//go:generate mockery --name IMysqlStore --inpackage --filename=mysql_mock.go
type IMysqlStore interface {
	PrepareStatement(ctx context.Context, stmt string) (*sql.Stmt, error)
	InsertOne(ctx context.Context, sqlStmt *sql.Stmt, args []interface{}) (int64, error)
	FindOne(ctx context.Context, sqlStmt *sql.Stmt, args []interface{}, dest interface{}) error
	UpdateOne(ctx context.Context, sqlStmt *sql.Stmt, args []interface{}) error
	DeleteOne(ctx context.Context, sqlStmt *sql.Stmt, args []interface{}) error
	Query(ctx context.Context, sqlStmt *sql.Stmt, inputs ...interface{}) (*sql.Rows, error)
	QueryRows(ctx context.Context, sqlStmt *sql.Stmt, args []interface{}) (*sql.Rows, error)
	DB() *sqlx.DB
	CloseResource(c Closeable) error
}

// MysqlDatastore struct
type MysqlDatastore struct {
	db     *sqlx.DB
	Config *config.SQLConfig
}

type BaseMysqlStore struct {
	Config    *config.SQLConfig
	Datastore IMysqlStore
}

func NewBaseMysqlStore(mysqlStore IMysqlStore, config *config.SQLConfig) *BaseMysqlStore {
	base := new(BaseMysqlStore)
	base.Datastore = mysqlStore
	base.Config = config
	return base
}

// NewDatastore initializes MysqlDatastore
func NewDatastore(cfg *config.SQLConfig) IMysqlStore {
	connectionString := getConnectionString(cfg)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.DialTimeOut)*time.Second)
	defer cancel()
	db, err := sqlx.ConnectContext(ctx, cfg.Driver, connectionString)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConnsint)

	mysqlDB := &MysqlDatastore{
		Config: cfg,
		db:     db,
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return mysqlDB
}

func getConnectionString(cfg *config.SQLConfig) string {
	if cfg.URI == "" {
		return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", cfg.Username, cfg.Password, cfg.DatabaseHost, cfg.DatabaseName)
	}
	return cfg.URI
}

func (m *MysqlDatastore) DB() *sqlx.DB {
	m.db.Ping()
	return m.db
}

// PrepareStatement internally executes sql.DB.Prepare, returns *sql.Stmt
func (ds *MysqlDatastore) PrepareStatement(ctx context.Context, stmt string) (*sql.Stmt, error) {
	return ds.db.Prepare(stmt)
}

// InsertOne inserts object to the datastore and returns last insert ID
func (ds *MysqlDatastore) InsertOne(ctx context.Context, sqlStmt *sql.Stmt, args []interface{}) (int64, error) {
	if sqlStmt == nil {
		return -1, errors.New("prepared statement is null")
	}
	defer ds.CloseResource(sqlStmt)

	sqlResult, err := sqlStmt.ExecContext(ctx, args...)
	if err != nil {
		return -1, err
	}

	return sqlResult.LastInsertId()
}

// FindOne retrieves single record and populates dest
func (ds *MysqlDatastore) FindOne(ctx context.Context, sqlStmt *sql.Stmt, args []interface{}, dest interface{}) error {
	if sqlStmt == nil {
		return errors.New("prepared statement is null")
	}
	defer ds.CloseResource(sqlStmt)

	sqlResult := sqlStmt.QueryRowContext(ctx, args...)
	if sqlResult.Err() != nil {
		return sqlResult.Err()
	}
	fields := structToSqlRow(dest, sqlResult)
	return sqlResult.Scan(fields...)
}

// UpdateOne updates single record
func (ds *MysqlDatastore) UpdateOne(ctx context.Context, sqlStmt *sql.Stmt, args []interface{}) error {
	if sqlStmt == nil {
		return errors.New("prepared statement is null")
	}
	defer ds.CloseResource(sqlStmt)

	_, err := sqlStmt.ExecContext(ctx, args...)
	return err
}

// DeleteOne deletes single record
func (ds *MysqlDatastore) DeleteOne(ctx context.Context, sqlStmt *sql.Stmt, args []interface{}) error {
	if sqlStmt == nil {
		return errors.New("prepared statement is null")
	}
	defer ds.CloseResource(sqlStmt)

	_, err := sqlStmt.ExecContext(ctx, args...)
	return err
}

func (ds *MysqlDatastore) Query(ctx context.Context, sqlStmt *sql.Stmt, inputs ...interface{}) (*sql.Rows, error) {
	if sqlStmt == nil {
		return nil, errors.New("prepared statement is null")
	}
	defer ds.CloseResource(sqlStmt)

	results, err := sqlStmt.QueryContext(ctx, inputs...)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// structToSqlRow populates u with sql.Row values
func structToSqlRow(u interface{}, row *sql.Row) []interface{} {
	val := reflect.ValueOf(u).Elem()
	v := make([]interface{}, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		v[i] = valueField.Addr().Interface()
	}
	return v
}

func (ds *MysqlDatastore) QueryRows(ctx context.Context, sqlStmt *sql.Stmt, args []interface{}) (*sql.Rows, error) {
	if sqlStmt == nil {
		return nil, errors.New("prepared statement is null")
	}
	defer ds.CloseResource(sqlStmt)

	result, err := sqlStmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type Closeable interface {
	Close() error
}

// CloseResource executes Close()
func (ds *MysqlDatastore) CloseResource(c Closeable) error {
	if c == nil {
		return nil
	}
	err := c.Close()
	if err != nil {
		return fmt.Errorf("failed to close resource: %s", err.Error())
	}
	return nil
}

// Close PreparedStatement
func (ds *MysqlDatastore) ClosePreparedStatement(stmt *sql.Stmt) error {
	if stmt == nil {
		return nil
	}
	err := stmt.Close()
	if err != nil {
		return errors.New("failed to close the prepared statement")
	}
	return nil
}

// Close Rows
func (ds *MysqlDatastore) CloseRows(rows *sql.Rows) error {
	if rows == nil {
		return nil
	}
	err := rows.Close()
	if err != nil {
		return errors.New("failed to close the rows")
	}
	return nil
}

// logMysqlError function
func logMysqlError(ctx context.Context, err error, query string, msg string, values interface{}) {
	log.Printf("mysql Error | requestId:%s | query:%s | values:%+v | msg:%s | err:%s ",
		appcontext.GetRequestContext(ctx).RequestID(),
		query,
		values,
		msg,
		err.Error())
}
