package config

type ContactUs struct {
	Telegram string `mapstructure:"telegram" json:"telegram" yaml:"telegram"`
	Email    string `mapstructure:"email" json:"email" yaml:"email"`
}
