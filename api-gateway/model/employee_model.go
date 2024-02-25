package model

type EmployeeBodyReq struct {
	Name string `json:"name"`
}

type EmployeeResponse struct {
	Code   int    `json:"Code"`
	Status string `json:"Status"`
	Data   struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"Data"`
}
