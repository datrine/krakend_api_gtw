package responsedtos

type LoginSuccessResponseData struct {
	AccessToken string `json:"access_token"`
	Email       string `json:"email"`
	LastName    string `json:"last_name"`
	FirstName   string `json:"first_name"`
}
type LoginSuccessResponse struct {
	Status  int
	Message string
	Data    *LoginSuccessResponseData `json:"data"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type RegisterSuccessResponse struct {
	Response
}

type GetUserByEmailSuccessResponseData struct {
	Email     string `json:"email"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
}
type GetUserByEmailSuccessResponse struct {
	Response
	Data *GetUserByEmailSuccessResponseData `json:"data"`
}
