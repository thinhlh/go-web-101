package dto

type BaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func NewSuccessResponse(data any) BaseResponse {
	return BaseResponse{
		Success: true,
		Message: "",
		Data:    data,
	}
}

func NewFailedResponse(err error) BaseResponse {
	return BaseResponse{
		Success: true,
		Message: err.Error(),
		Data:    nil,
	}
}
