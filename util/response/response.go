package response

type Response struct {
	Msg string `json:"msg"`
}

type ResponseSuccess struct {
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}