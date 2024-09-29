package response

type AddUsersResponse struct {
	UserId    string `json:"userId"`
	RequestId string `json:"requestId"`
}

type AssingUsersTaskResponse struct {
	RequestId string `json:"requestId"`
}
