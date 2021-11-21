package entity

import "time"

type Config struct {
	Debug bool `env:"DEBUG,default=true"`

	ListenPort   string        `env:"LISTEN_PORT,default=80"`
	ReadTimeout  time.Duration `env:"READ_TIMEOUT,default=15"`
	WriteTimeout time.Duration `env:"WRITE_TIMEOUT,default=15"`
	IdleTimeout  time.Duration `env:"IDLE_TIMEOUT,default=60"`

	Token string `env:"TELEGRAM_BOT_TOKEN"`
	Hook  string `env:"TELEGRAM_HOOK_LINK"`
}
