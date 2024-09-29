package mysql

import (
	"context"
	"database/sql"
	errs "errors"
	"fmt"
	"strings"

	"task_management_system/constant"
	"task_management_system/domain/users"
	"task_management_system/errors"
	"task_management_system/util"

	"github.com/go-sql-driver/mysql"
)

var (
	usersColumns = strings.Join([]string{
		string(users.ColID),
		string(users.ColName),
		string(users.ColEmail),
		string(users.ColIsDeleted),
		string(users.ColCreatedAt),
		string(users.ColUpdatedAt),
	}, ",")
	usersInsertValues = "?, ?, ?, ?"
)

type UserRepo struct {
	dataStore IMysqlStore
	table     string
}

func NewUsersDatastoreApi(ds IMysqlStore) users.IUsersRepo {
	return &UserRepo{
		dataStore: ds,
		table:     constant.TableUsers,
	}
}

func (cr *UserRepo) initGetUserByIdStmt(ctx context.Context) (*sql.Stmt, errors.IError, string) {
	getUserByIdQry := fmt.Sprintf("SELECT %s FROM %s WHERE %s=?", usersColumns, cr.table, users.ColID)
	getUserByIdStmt, err := cr.dataStore.PrepareStatement(ctx, getUserByIdQry)
	if err != nil {
		err1 := errors.New(errors.MySqlReadFailID, errors.MysqlReadFailCode, err.Error())
		logMysqlError(ctx, err1, getUserByIdQry, "initGetUserByIdStmt|PrepareStatement|failed", nil)
		return nil, err1, ""
	}
	return getUserByIdStmt, nil, getUserByIdQry
}

func (cr *UserRepo) GetUserById(ctx context.Context, id string) (*users.Users, errors.IError) {
	getUserByIdStmt, err1, getUserByIdQry := cr.initGetUserByIdStmt(ctx)
	if err1 != nil {
		return nil, err1
	}
	user := new(users.Users)
	values := []interface{}{id}
	err := cr.dataStore.FindOne(ctx, getUserByIdStmt, values, user)
	if err != nil {
		logMysqlError(ctx, err, getUserByIdQry, "GetUsersById|FindOne|failed", values)

		if err == sql.ErrNoRows {
			err1 = errors.New(errors.ResourceNotFoundGenericID, errors.DataNotFoundCode, err.Error())
			return nil, err1
		}

		err1 = errors.NewErrMySQLReadFail(err.Error())
		return nil, err1
	}
	return user, nil
}

func (cr *UserRepo) initGetUserByEmailStmt(ctx context.Context) (*sql.Stmt, errors.IError, string) {
	getUserByEmailQry := fmt.Sprintf("SELECT %s FROM %s WHERE %s=?", usersColumns, cr.table, users.ColEmail)
	getUserByEmailStmt, err := cr.dataStore.PrepareStatement(ctx, getUserByEmailQry)
	if err != nil {
		err1 := errors.New(errors.MySqlReadFailID, errors.MysqlReadFailCode, err.Error())
		logMysqlError(ctx, err1, getUserByEmailQry, "initGetUserByEmailStmt|PrepareStatement|failed", nil)
		return nil, err1, ""
	}
	return getUserByEmailStmt, nil, getUserByEmailQry
}

func (cr *UserRepo) GetUserByEmail(ctx context.Context, email string) (*users.Users, errors.IError) {
	getUserByEmailStmt, err1, getUserByEmailQry := cr.initGetUserByEmailStmt(ctx)
	if err1 != nil {
		return nil, err1
	}
	user := new(users.Users)
	values := []interface{}{email}
	err := cr.dataStore.FindOne(ctx, getUserByEmailStmt, values, user)
	if err != nil {
		logMysqlError(ctx, err, getUserByEmailQry, "GetUserByEmail|FindOne|failed", values)

		if err == sql.ErrNoRows {
			err1 = errors.New(errors.ResourceNotFoundGenericID, errors.DataNotFoundCode, err.Error())
			return nil, err1
		}

		err1 = errors.NewErrMySQLReadFail(err.Error())
		return nil, err1
	}

	return user, nil
}

func (cr *UserRepo) initAddUser(ctx context.Context) (*sql.Stmt, errors.IError, string) {
	insertUserQuery := fmt.Sprintf("INSERT INTO %s (%s) %s", cr.table, usersColumns, usersInsertValues)
	insertUserStmt, err := cr.dataStore.PrepareStatement(ctx, insertUserQuery)
	if err != nil {
		logMysqlError(ctx, err, insertUserQuery, "initAddUser|PrepareStatement|failed", nil)
		return nil, errors.NewErrMySQLWriteFail(err.Error()), ""
	}

	return insertUserStmt, nil, insertUserQuery
}

func (cr *UserRepo) AddUsers(ctx context.Context, data *users.Users) (string, errors.IError) {
	insertTdStmt, err1, insertTaskDetailsQuery := cr.initAddUser(ctx)
	if err1 != nil {
		return "", err1
	}

	uuid := util.GenerateUUIDPk()
	newRec := false
	data.Id = &uuid
	data.IsDeleted = &newRec

	insertArgs := []interface{}{
		data.Id,
		data.Name,
		data.EmailId,
		data.IsDeleted,
	}

	_, err := cr.dataStore.InsertOne(ctx, insertTdStmt, insertArgs)
	if err != nil {
		logMysqlError(ctx, err, insertTaskDetailsQuery, "AddUsers|InsertOne|failed", insertArgs)

		var mysqlErr *mysql.MySQLError
		errCode := errors.MysqlWriteFailCode

		if errs.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			errCode = errors.DuplicateSpecificationFailedCode
		}

		err2 := errors.New(errors.MySqlWriteFailID, errCode, err.Error())
		return "", err2
	}

	return *data.Id, nil
}

func (cr *UserRepo) initUpdateUserStmt(ctx context.Context) (*sql.Stmt, errors.IError, string) {
	updateUserQuery := fmt.Sprintf("UPDATE %s SET %s = ?, %s = ? WHERE %s = ? ;", cr.table, users.ColName, users.ColEmail, users.ColID)
	updateUserStmt, err := cr.dataStore.PrepareStatement(ctx, updateUserQuery)
	if err != nil {
		logMysqlError(ctx, err, updateUserQuery, "initUpdateUserStmt|PrepareStatement|failed", nil)
		return nil, errors.NewErrMySQLWriteFail(err.Error()), ""
	}

	return updateUserStmt, nil, updateUserQuery
}

func (cr *UserRepo) UpdateUsers(ctx context.Context, data *users.Users) errors.IError {
	updateUserStmt, err1, updateUserQuery := cr.initUpdateUserStmt(ctx)
	if err1 != nil {
		return err1
	}

	updateArgs := []interface{}{
		data.Name,
		data.EmailId,
		data.Id,
	}

	err := cr.dataStore.UpdateOne(ctx, updateUserStmt, updateArgs)
	if err != nil {
		logMysqlError(ctx, err, updateUserQuery, "UpdateUsers|UpdateOne|failed", updateArgs)
		return errors.NewErrMySQLWriteFail(err.Error())
	}

	return nil
}

func (cr *UserRepo) initDeleteUserStmt(ctx context.Context) (*sql.Stmt, errors.IError, string) {

	deleteUserQuery := fmt.Sprintf("UPDATE %s SET %s = 1 WHERE %s = ?", cr.table, users.ColIsDeleted, users.ColID)
	deleteUserStmt, err := cr.dataStore.PrepareStatement(ctx, deleteUserQuery)
	if err != nil {
		logMysqlError(ctx, err, deleteUserQuery, "initDeleteUserStmt|PrepareStatement|failed", nil)
		return nil, errors.NewErrMySQLWriteFail(err.Error()), ""
	}

	return deleteUserStmt, nil, deleteUserQuery
}

func (cr *UserRepo) DeleteUsersById(ctx context.Context, id string) errors.IError {
	deleteUserStmt, err1, deleteUserQuery := cr.initDeleteUserStmt(ctx)
	if err1 != nil {
		return err1
	}
	values := []interface{}{id}
	err := cr.dataStore.DeleteOne(ctx, deleteUserStmt, values)
	if err != nil {
		logMysqlError(ctx, err, deleteUserQuery, "DeleteUsersById|DeleteOne|failed", values)
		return errors.NewErrMySQLWriteFail(err.Error())
	}

	return nil
}
