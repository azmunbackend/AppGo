package appresult

import "encoding/json"

var  Success = NewAppSuccess("Success!", "SS-10000", nil)


type AppSuccess struct {
	Status  bool        `json:"status"`
	Message string      `json:"message,omitempty"`
	Code    string      `json:"code,omitempty"`
	Data    interface{} `json:"data"`
}

func (s *AppSuccess) Success() string {
	return s.Message
}

func (s *AppSuccess) Marshal() []byte {
	marshal, err := json.Marshal(s)
	if err != nil {
		return nil
	}
	return marshal
}

func NewAppSuccess(message, code string, data interface{}) *AppSuccess {
	return &AppSuccess{
		Status:  true,
		Message: message,
		Code:    code,
		Data:    data,
	}
}
