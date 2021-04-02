package main

//User struct
type User struct {
	Id        int32     `json:"id"`
	FirstName string  `jso:"fname"`
	City      string  `json:"city"`
	Phone     int64     `json:"phone"`
	Height    float32 `json:"height"`
	Married   bool    `json:"married"`
}

//List of ids
type IDList struct {
	Id []int `json:"id"`
}

//Error Response Struct
type ResponseError struct {
	ErrorMessage string `json:"message"`
	ErrorCode    int    `json:"status_code"`
}
