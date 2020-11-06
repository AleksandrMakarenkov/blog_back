package dependencies

import (
	"database/sql"
	"fmt"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
	"io/ioutil"
	"log"
	"os"
	"vue_back/blog"
)

func MakeReform() (*reform.DB, error) {
	env := os.Getenv("BLOG_ENV")
	dbName := os.Getenv("POSTGRES_DB")

	dbUser, err := ioutil.ReadFile("/run/secrets/postgres_user")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	dbPassword, err := ioutil.ReadFile("/run/secrets/postgres_password")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	dsn := fmt.Sprintf("postgresql://%s:%s@db/%s", dbUser, dbPassword, dbName)
	config, err := blog.NewConfig(os.Getenv(blog.EnvNameOfSecret), dsn, nil, nil, env)
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
