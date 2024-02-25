package model

type UserBodyReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Code        int    `json:"Code"`
	Status      string `json:"Status"`
	AccessToken string `json:"access_token"`
	Data        struct {
		ID       string `json:"id"`
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"Data"`
}
