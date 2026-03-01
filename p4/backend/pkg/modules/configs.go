package modules

import "time"

type PostgresqlConfig struct {
	Host        string
	Port        string
	Username    string
	Password    string
	DBName      string
	SSLMode     string
	ExecTimeout time.Duration
}
