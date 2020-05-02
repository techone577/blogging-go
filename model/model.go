package model

type UnifiedResponse struct {
	Success   bool        `json:"success"`
	Msg       string      `json:"msg"`
	ErrorCode int         `json:"errorCode"`
	Data      interface{} `json:"data"`
}

func NewUnifiedResponse(success bool, data interface{}) UnifiedResponse {
	return UnifiedResponse{
		Success: success,
		Data:    data,
	}
}
