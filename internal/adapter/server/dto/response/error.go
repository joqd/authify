package response

type ErrorResponse struct {
	BaseResponse
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}
