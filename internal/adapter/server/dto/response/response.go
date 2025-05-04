package response

type BaseResponse struct {
	Ok bool `json:"ok"`
}

type Response[T any] struct {
	BaseResponse
	Result *T `json:"result"`
}
