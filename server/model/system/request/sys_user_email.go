package request

type EmailCaptcha struct {
	Email string `json:"email"` // 邮箱
}

type EmailLogin struct {
	Email   string `json:"email"`   // 邮箱
	Captcha string `json:"captcha"` // 验证码
}
