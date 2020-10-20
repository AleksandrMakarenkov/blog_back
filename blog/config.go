package blog

import (
	"database/sql"
	"fmt"
	"gopkg.in/reform.v1"
)

type Config struct {
	Secret string
	DSN    string
	DB     *sql.DB
	Reform *reform.DB
}

func NewConfig(
	secret string,
	dsn string,
	db *sql.DB,
	reform *reform.DB,
) (*Config, error) {
	if secret == "" {
		return nil, fmt.Errorf("secret string is empty, please export it or set as argument")
	}
	if dsn == "" {
		return nil, fmt.Errorf("dsn empty, please export it or set as argument")
	}
	return &Config{
		Secret: secret,
		DSN:    dsn,
		DB:     db,
		Reform: reform,
	}, nil
}
