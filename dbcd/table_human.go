package dbcd

// Human 是表“Human”的模型
type Human struct {
	Username string `form:"name"`
	Password string `form:"pass"`
}
