package response

import "time"

type GetTaskResponse struct {
	RequestId string      `json:"requestId"`
	Data      interface{} `json:"data"`
}

type GetTaskDetails struct {
	TaskId      string     `json:"taskId"`
	UserId      *string    `json:"userId"`
	Status      string     `json:"status"`
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	DueDate     *time.Time `json:"dueDate"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

type AddTaskResponse struct {
	RequestId string `json:"requestId"`
	TaskId    string `json:"taskId"`
}

type UpdateTaskResponse struct {
	RequestId string `json:"requestId"`
}

type DeleteTaskResponse struct {
	RequestId string `json:"requestId"`
}
