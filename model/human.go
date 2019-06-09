package model

type Human struct {
	Username string `form:"name"`
	Password string `form:"pass"`
}
