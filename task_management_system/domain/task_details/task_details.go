package task_details

import (
	"time"
)

type TaskDetails struct {
	Id          *string    `db:"id"`
	UserId      *string    `db:"user_id"`
	Status      *string    `db:"status"`
	Title       *string    `db:"title"`
	Description *string    `db:"description"`
	DueDate     *time.Time `db:"due_date"`
	IsDeleted   *bool      `db:"is_deleted"`
	CreatedAt   *time.Time `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
}

type TaskDetailsColumn string

const (
	ColID          TaskDetailsColumn = "id"
	ColUserId      TaskDetailsColumn = "user_id"
	ColStatus      TaskDetailsColumn = "status"
	ColTitle       TaskDetailsColumn = "title"
	ColDescription TaskDetailsColumn = "description"
	ColDueDate     TaskDetailsColumn = "due_date"
	ColIsDeleted   TaskDetailsColumn = "is_deleted"
	ColCreatedAt   TaskDetailsColumn = "created_at"
	ColUpdatedAt   TaskDetailsColumn = "updated_at"
)
