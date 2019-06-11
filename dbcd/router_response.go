package dbcd

// RouterResponse 响应载体，将送回前端的信息
type RouterResponse map[string]interface{}

// NewRouterResponse 返回带有默认值的响应载体
func NewRouterResponse() RouterResponse {
	var response = make(RouterResponse)
	response.SetCodeAndMsg(-1, "lightyears忘记设置Code和Msg了，请🔨他,,ԾㅂԾ,,")
	return response
}

// SetCodeAndMsg 设置RouterResponse的code字段和msg字段。
func (response *RouterResponse) SetCodeAndMsg(code int, msg string) {
	(*response)["code"], (*response)["msg"] = code, msg
}
