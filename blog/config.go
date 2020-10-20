package blog

import (
	"database/sql"
	"gopkg.in/reform.v1"
)

type Config struct {
	Secret string
	Db     *sql.DB
	Reform *reform.DB
}

func NewConfig(secret string, db *sql.DB, reform *reform.DB) *Config {
	return &Config{
		Secret: secret,
		Db:     db,
		Reform: reform,
	}
}
