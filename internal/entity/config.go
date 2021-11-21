package entity

type Config struct {
	DebugMode     bool   `env:"DEBUG_MODE,default=true"`
	Port          string `env:"LISTEN_PORT"`
	TelegramToken string `env:"TELEGRAM_TOKEN"`
}
