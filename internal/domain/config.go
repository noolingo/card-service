package domain

import "time"

type Config struct {
	AppName    string           `yaml:"appname" env-default:"test"`
	GrpcServer GrpcServer       `yaml:"grpcserver" env-prefix:"CARD_SERVICE_"`
	Mysql      Mysql            `yaml:"mysql" env-prefix:"CARD_SERVICE_"`
	Api        YandexDictionary `yaml:"yandexapi" env-prefix:"YANDEX_API_"`
}

type GrpcServer struct {
	Host string `yaml:"host" env-default:"0.0.0.0"`
	Port string `yaml:"port" env-default:"9002"`
}

type Mysql struct {
	DSN             string        `yaml:"dsn" env:"MYSQL_DSN"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime" env-default:"5m"`
	MaxOpenConns    int           `yaml:"max_open_conns" env-default:"10"`
	MaxIdleConns    int           `yaml:"max_idle_conns" env-default:"10"`
}

type YandexDictionary struct {
	Api string `yaml:"yandex-api" env:"YANDEX_API"`
}
