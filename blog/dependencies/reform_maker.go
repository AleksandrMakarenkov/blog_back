package dependencies

import (
	"database/sql"
	"fmt"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
	"log"
	"os"
	"vue_back/blog"
)

func MakeReform() (*reform.DB, error) {
	config, err := blog.NewConfig(os.Getenv(blog.SecretName), os.Getenv("DB_DSN"), nil, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	db, err := sql.Open("pgx", config.DSN)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	logger := log.New(os.Stderr, "SQL: ", log.Flags())
	return reform.NewDB(db, postgresql.Dialect, reform.NewPrintfLogger(logger.Printf)), nil
}
