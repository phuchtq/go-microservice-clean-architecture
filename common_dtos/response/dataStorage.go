package response

type DataStorage struct {
	Data   interface{} `json:"data"`
	ErrMsg error       `json:"err_msg"`
}
