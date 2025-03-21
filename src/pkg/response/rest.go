package response

import "reflect"

type (
	data interface {
		string | any | []any
	}

	errors interface {
		any | []any
	}

	jsonDefaultMessage struct {
		Message     string `json:"message"`
		Code        int    `json:"code"`
		TotalRecord int    `json:"total_recode,omitempty"`
	}

	JSONMessageResponse struct {
		*jsonDefaultMessage
		Data   data   `json:"data,omitempty"`
		Errors errors `json:"errors,omitempty"`
	}
)

func NewJSONMessage() *jsonDefaultMessage {
	return &jsonDefaultMessage{}
}

func (j *jsonDefaultMessage) AddMessage(s string) *jsonDefaultMessage {
	j.Message = s
	return j
}

func (j *jsonDefaultMessage) AddCode(code int) *jsonDefaultMessage {
	j.Code = code
	return j
}

func BuildMessageWithData[T data](d T, jsonDef *jsonDefaultMessage) JSONMessageResponse {
	jsonResMsg := JSONMessageResponse{
		jsonDefaultMessage: jsonDef,
		Data:               d,
	}

	v := reflect.ValueOf(d)

	if v.Kind() == reflect.Slice {
		jsonResMsg.TotalRecord = v.Len()
	}

	return jsonResMsg
}

func BuildMessageWithErrors[T errors](errs T, jsonDef *jsonDefaultMessage) JSONMessageResponse {
	return JSONMessageResponse{
		jsonDefaultMessage: jsonDef,
		Errors:             errs,
	}
}
