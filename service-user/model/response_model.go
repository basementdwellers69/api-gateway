package model

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type LoginResponse struct {
	Code        int         `json:"code"`
	Status      string      `json:"status"`
	AccessToken string      `json:"access_token"`
	Data        interface{} `json:"data"`
}
