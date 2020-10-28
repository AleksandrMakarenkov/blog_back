package blog

import (
	"database/sql"
	"errors"
	"fmt"
	"gopkg.in/reform.v1"
)

type Config struct {
	Secret string
	DSN    string
	DB     *sql.DB
	Reform *reform.DB
	port   int
}

func NewConfig(
	secret string,
	dsn string,
	db *sql.DB,
	reform *reform.DB,
	env string,
) (*Config, error) {
	if secret == "" {
		return nil, fmt.Errorf("please provide env var BLOG_SECRET")
	}
	if dsn == "" {
		return nil, fmt.Errorf("please provide env var DB_DSN")
	}
	if env == "" {
		return nil, errors.New("please provide env var BLOG_ENV (dev, prod)")
	}
	var port int
	if env == "dev" {
		port = 9090
	} else {
		port = 80
	}
	return &Config{
		Secret: secret,
		DSN:    dsn,
		DB:     db,
		Reform: reform,
		port:   port,
	}, nil
}
