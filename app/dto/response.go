package dto

type ResponseList struct {
	Data  interface{} `json:"data"`
	Total interface{} `json:"total"`
}

type Response struct {
	Data interface{} `json:"data"`
}
