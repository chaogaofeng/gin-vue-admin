package config

type TronGrid struct {
	BaseURL string   `mapstructure:"base-url" json:"base-url" yaml:"base-url"`
	ApiKeys []string `mapstructure:"api-keys" json:"api-keys" yaml:"api-keys"`
	TRC20   string   `mapstructure:"trc20" json:"trc20" yaml:"trc20"`
}
