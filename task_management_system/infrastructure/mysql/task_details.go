package mysql

import (
	"context"
	"database/sql"
	errs "errors"
	"fmt"
	"strings"

	"task_management_system/constant"
	"task_management_system/domain/task_details"
	"task_management_system/errors"
	"task_management_system/util"

	"github.com/go-sql-driver/mysql"
)

var (
	taskDetailsColumns = strings.Join([]string{
		string(task_details.ColID),
		string(task_details.ColUserId),
		string(task_details.ColStatus),
		string(task_details.ColTitle),
		string(task_details.ColDescription),
		string(task_details.ColDueDate),
		string(task_details.ColIsDeleted),
	}, ",")
	itaskDetailsInsertValues = "?, ?, ?, ?, ?, ?, ?"
)

type TaskDetailsRepo struct {
	dataStore IMysqlStore
	table     string
}

func NewTaskDetailsDatastoreApi(ds IMysqlStore) task_details.ITaskDetailsRepo {
	return &TaskDetailsRepo{
		dataStore: ds,
		table:     constant.TableTaskDetails,
	}
}

func (cr *TaskDetailsRepo) initGetTdByIdStmt(ctx context.Context) (*sql.Stmt, errors.IError, string) {
	getTdByIdQry := fmt.Sprintf("SELECT %s FROM %s WHERE %s=?", taskDetailsColumns, cr.table, task_details.ColID)
	getTdByIdStmt, err := cr.dataStore.PrepareStatement(ctx, getTdByIdQry)
	if err != nil {
		err1 := errors.New(errors.MySqlReadFailID, errors.MysqlReadFailCode, err.Error())
		logMysqlError(ctx, err1, getTdByIdQry, "initGetTdByIdPPStmt|PrepareStatement|failed", nil)
		return nil, err1, ""
	}
	return getTdByIdStmt, nil, getTdByIdQry
}

func (cr *TaskDetailsRepo) GetTaskDetailsById(ctx context.Context, id string) (*task_details.TaskDetails, errors.IError) {
	getTdByIdStmt, err1, getTdByIdQry := cr.initGetTdByIdStmt(ctx)
	if err1 != nil {
		return nil, err1
	}
	taskDetails := new(task_details.TaskDetails)
	values := []interface{}{id}
	err := cr.dataStore.FindOne(ctx, getTdByIdStmt, values, taskDetails)
	if err != nil {
		logMysqlError(ctx, err, getTdByIdQry, "GetTaskDetailsById|FindOne|failed", values)

		if err == sql.ErrNoRows {
			err1 = errors.New(errors.ResourceNotFoundGenericID, errors.DataNotFoundCode, err.Error())
			return nil, err1
		}

		err1 = errors.NewErrMySQLReadFail(err.Error())
		return nil, err1
	}
	return taskDetails, nil
}

func (cr *TaskDetailsRepo) initGetTdByUserIdStmt(ctx context.Context) (*sql.Stmt, errors.IError, string) {
	getTdByUserIdQry := fmt.Sprintf("SELECT %s FROM %s WHERE %s=?", taskDetailsColumns, cr.table, task_details.ColUserId)
	getTdByUserIdStmt, err := cr.dataStore.PrepareStatement(ctx, getTdByUserIdQry)
	if err != nil {
		err1 := errors.New(errors.MySqlReadFailID, errors.MysqlReadFailCode, err.Error())
		logMysqlError(ctx, err1, getTdByUserIdQry, "initGetTdByUserIdStmt|PrepareStatement|failed", nil)
		return nil, err1, ""
	}
	return getTdByUserIdStmt, nil, getTdByUserIdQry
}

func (cr *TaskDetailsRepo) GetTaskDetailsByUserId(ctx context.Context, userId string) ([]task_details.TaskDetails, errors.IError) {
	getTdByUserIdStmt, err1, getTdByUserIdQry := cr.initGetTdByUserIdStmt(ctx)
	if err1 != nil {
		return nil, err1
	}
	values := []interface{}{userId}
	sqlRows, err := cr.dataStore.QueryRows(ctx, getTdByUserIdStmt, values)
	if err != nil {
		logMysqlError(ctx, err, getTdByUserIdQry, "GetTaskDetailsByUserId|QueryRows|failed", values)

		if err == sql.ErrNoRows {
			err1 = errors.New(errors.ResourceNotFoundGenericID, errors.DataNotFoundCode, err.Error())
			return nil, err1
		}

		err1 = errors.NewErrMySQLReadFail(err.Error())
		return nil, err1
	}

	taskDetailsList := []task_details.TaskDetails{}
	for sqlRows.Next() {
		var taskDetailsModel task_details.TaskDetails
		err := sqlRows.Scan(
			&taskDetailsModel.Id,
			&taskDetailsModel.UserId,
			&taskDetailsModel.Status,
			&taskDetailsModel.Title,
			&taskDetailsModel.Description,
			&taskDetailsModel.DueDate,
			&taskDetailsModel.IsDeleted,
		)

		if err != nil {
			err1 := errors.New(errors.MySqlReadFailID, errors.MysqlReadFailCode, err.Error())
			logMysqlError(ctx, err1, getTdByUserIdQry, "GetTaskDetailsByUserId|Scan|failed", values)
			continue
		}

		taskDetailsList = append(taskDetailsList, taskDetailsModel)
	}

	return taskDetailsList, nil
}

func (cr *TaskDetailsRepo) initGetTdByUserIdAndStatusStmt(ctx context.Context) (*sql.Stmt, errors.IError, string) {
	getTdByUserIdAndStatusQry := fmt.Sprintf("SELECT %s FROM %s WHERE %s=? and %s=?", taskDetailsColumns, cr.table, task_details.ColUserId, task_details.ColStatus)
	getTdByUserIdAndStatusStmt, err := cr.dataStore.PrepareStatement(ctx, getTdByUserIdAndStatusQry)
	if err != nil {
		err1 := errors.New(errors.MySqlReadFailID, errors.MysqlReadFailCode, err.Error())
		logMysqlError(ctx, err1, getTdByUserIdAndStatusQry, "initGetTdByUserIdAndStatusStmt|PrepareStatement|failed", nil)
		return nil, err1, ""
	}
	return getTdByUserIdAndStatusStmt, nil, getTdByUserIdAndStatusQry
}

func (cr *TaskDetailsRepo) GetTaskDetailsByUserIdAndStatus(ctx context.Context, userId, status string) ([]task_details.TaskDetails, errors.IError) {
	getTdByUserIdAndStatusStmt, err1, getTdByUserIdAndStatusQry := cr.initGetTdByUserIdAndStatusStmt(ctx)
	if err1 != nil {
		return nil, err1
	}
	values := []interface{}{userId, status}
	sqlRows, err := cr.dataStore.QueryRows(ctx, getTdByUserIdAndStatusStmt, values)
	if err != nil {
		logMysqlError(ctx, err, getTdByUserIdAndStatusQry, "GetTaskDetailsByUserIdAndStatus|QueryRows|failed", values)

		if err == sql.ErrNoRows {
			err1 = errors.New(errors.ResourceNotFoundGenericID, errors.DataNotFoundCode, err.Error())
			return nil, err1
		}

		err1 = errors.NewErrMySQLReadFail(err.Error())
		return nil, err1
	}

	taskDetailsList := []task_details.TaskDetails{}
	for sqlRows.Next() {
		var taskDetailsModel task_details.TaskDetails
		err := sqlRows.Scan(
			&taskDetailsModel.Id,
			&taskDetailsModel.UserId,
			&taskDetailsModel.Status,
			&taskDetailsModel.Title,
			&taskDetailsModel.Description,
			&taskDetailsModel.DueDate,
			&taskDetailsModel.IsDeleted,
		)

		if err != nil {
			err1 := errors.New(errors.MySqlReadFailID, errors.MysqlReadFailCode, err.Error())
			logMysqlError(ctx, err1, getTdByUserIdAndStatusQry, "GetTaskDetailsByUserIdAndStatus|Scan|failed", values)
			continue
		}

		taskDetailsList = append(taskDetailsList, taskDetailsModel)
	}

	return taskDetailsList, nil
}

func (cr *TaskDetailsRepo) initAddTaskDetails(ctx context.Context) (*sql.Stmt, errors.IError, string) {
	insertTaskDetailsQuery := fmt.Sprintf("INSERT INTO %s (%s) values (%s)", cr.table, taskDetailsColumns, itaskDetailsInsertValues)
	insertTdStmt, err := cr.dataStore.PrepareStatement(ctx, insertTaskDetailsQuery)
	if err != nil {
		logMysqlError(ctx, err, insertTaskDetailsQuery, "initAddTaskDetails|PrepareStatement|failed", nil)
		return nil, errors.NewErrMySQLWriteFail(err.Error()), ""
	}

	return insertTdStmt, nil, insertTaskDetailsQuery
}

func (cr *TaskDetailsRepo) AddTaskDetails(ctx context.Context, data *task_details.TaskDetails) (*string, errors.IError) {
	insertTdStmt, err1, insertTaskDetailsQuery := cr.initAddTaskDetails(ctx)
	if err1 != nil {
		return nil, err1
	}

	uuid := util.GenerateUUIDPk()
	newRec := false
	status := "Pending"
	data.Id = &uuid
	data.IsDeleted = &newRec
	data.Status = &status

	insertArgs := []interface{}{
		data.Id,
		data.UserId,
		data.Status,
		data.Title,
		data.Description,
		data.DueDate,
		data.IsDeleted,
	}

	_, err := cr.dataStore.InsertOne(ctx, insertTdStmt, insertArgs)
	if err != nil {
		logMysqlError(ctx, err, insertTaskDetailsQuery, "AddTaskDetails|InsertOne|failed", insertArgs)

		var mysqlErr *mysql.MySQLError
		errCode := errors.MysqlWriteFailCode

		if errs.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			errCode = errors.DuplicateSpecificationFailedCode
		}

		err2 := errors.New(errors.MySqlWriteFailID, errCode, err.Error())
		return nil, err2
	}

	return data.Id, nil
}

func (cr *TaskDetailsRepo) initUpdateTaskDetailsStmt(ctx context.Context) (*sql.Stmt, errors.IError, string) {
	updateTaskDetailsQuery := fmt.Sprintf("UPDATE %s SET %s = ?, %s = ?, %s = ?, %s = ?, %s = ? WHERE %s = ? ;", cr.table, task_details.ColStatus, task_details.ColTitle, task_details.ColDescription, task_details.ColDueDate, task_details.ColUserId, task_details.ColID)
	updateTaskDetailsStmt, err := cr.dataStore.PrepareStatement(ctx, updateTaskDetailsQuery)
	if err != nil {
		logMysqlError(ctx, err, updateTaskDetailsQuery, "initUpdateTaskDetailsStmt|PrepareStatement|failed", nil)
		return nil, errors.NewErrMySQLWriteFail(err.Error()), ""
	}

	return updateTaskDetailsStmt, nil, updateTaskDetailsQuery
}

func (cr *TaskDetailsRepo) UpdateTaskDetails(ctx context.Context, data *task_details.TaskDetails) errors.IError {
	updateTaskDetailsStmt, err1, updateTaskDetailsQuery := cr.initUpdateTaskDetailsStmt(ctx)
	if err1 != nil {
		return err1
	}

	updateArgs := []interface{}{
		data.Status,
		data.Title,
		data.Description,
		data.DueDate,
		data.UserId,
		data.Id,
	}

	err := cr.dataStore.UpdateOne(ctx, updateTaskDetailsStmt, updateArgs)
	if err != nil {
		logMysqlError(ctx, err, updateTaskDetailsQuery, "UpdateTaskDetails|UpdateOne|failed", updateArgs)
		return errors.NewErrMySQLWriteFail(err.Error())
	}

	return nil
}

func (cr *TaskDetailsRepo) initDeleteTaskDetailsStmt(ctx context.Context) (*sql.Stmt, errors.IError, string) {

	deleteTaskDetailsQuery := fmt.Sprintf("UPDATE %s SET %s = 1 WHERE %s = ? ;", cr.table, task_details.ColIsDeleted, task_details.ColID)
	deleteTaskDetailsStmt, err := cr.dataStore.PrepareStatement(ctx, deleteTaskDetailsQuery)
	if err != nil {
		logMysqlError(ctx, err, deleteTaskDetailsQuery, "initDeleteTaskDetailsStmt|PrepareStatement|failed", nil)
		return nil, errors.NewErrMySQLWriteFail(err.Error()), ""
	}

	return deleteTaskDetailsStmt, nil, deleteTaskDetailsQuery
}

func (cr *TaskDetailsRepo) DeleteTaskDetailsById(ctx context.Context, id string) errors.IError {
	deleteTaskDetailsStmt, err1, deleteTaskDetailsQuery := cr.initDeleteTaskDetailsStmt(ctx)
	if err1 != nil {
		return err1
	}
	values := []interface{}{id}
	err := cr.dataStore.DeleteOne(ctx, deleteTaskDetailsStmt, values)
	if err != nil {
		logMysqlError(ctx, err, deleteTaskDetailsQuery, "DeleteTaskDetailsById|DeleteOne|failed", values)
		return errors.NewErrMySQLWriteFail(err.Error())
	}

	return nil
}
