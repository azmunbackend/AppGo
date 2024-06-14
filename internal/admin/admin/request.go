package adminlogin

type CreateLoginRequest struct{
	UserName string `json:"username" form:"username"`
	Password string `json:"password" form:"username"`
}