package apiResponse

type JsonObject map[string]interface{}

type ResponseData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Response(data interface{}, msg string, code int) interface{} {
	return &ResponseData{Msg: msg, Code: code, Data: data}
}

func ResponseOk(msg string) interface{} {
	return &ResponseData{Msg: msg, Code: StatusCodeOk, Data: JsonObject{}}
}

func ResponseFail(msg string, code int) interface{} {
	return &ResponseData{Msg: msg, Code: code, Data: JsonObject{}}
}

func ResponsesItems(items interface{}) interface{} {
	return &ResponseData{
		Msg:  StatusMsgSuccess,
		Code: StatusCodeOk,
		Data: map[string]interface{}{
			"Items": items,
		},
	}
}
