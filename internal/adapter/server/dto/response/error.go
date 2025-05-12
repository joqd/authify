package response

type ErrorResponse struct {
	BaseResponse `example:"false"`
	ErrorCode    int    `json:"error_code"`
	Description  string `json:"description"`
}
