package adminlogin

type MyError struct {
    Message interface{}
}

type Login struct{
	UserName string `json:"username"`
	Password string `json:"password"`
}

type ResultTokenDTO struct {
	Token string `json:"token"`
}