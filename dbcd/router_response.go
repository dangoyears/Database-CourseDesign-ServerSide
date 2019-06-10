package dbcd

// RouterResponse 响应载体，将送回前端的信息
type RouterResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data map[string]interface{} `json:"data"` 
}
// NewRouterResponse 返回带有默认值的响应载体
func NewRouterResponse() RouterResponse {
	var response RouterResponse
	response.setCodeAndMsg(-1, "lightyears忘记设置Code和Msg了，请🔨他,,ԾㅂԾ,,")
	response.Data = make(map[string]interface{})
	return response
}

func (response *RouterResponse) setCodeAndMsg(code int, msg string) {
	response.Code, response.Msg = code, msg
}
