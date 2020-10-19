package blog

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/sessions"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
	"log"
	"os"
	"vue_back/blog/repository"
	"vue_back/blog/service/password"
	"vue_back/blog/service/session"
)

func MakeBlog() (*Blog, error) {
	db, err := sql.Open("pgx", os.Getenv("DB_DSN"))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}


	logger := log.New(os.Stderr, "SQL: ", log.Flags())
	reformDb := reform.NewDB(db, postgresql.Dialect, reform.NewPrintfLogger(logger.Printf))

	config := NewConfig(os.Getenv(SecretName), db, reformDb)
	userRepo := repository.NewUserRepository(reformDb)

	store := sessions.NewCookieStore([]byte(config.Secret))
	comparator := password.NewComparator()
	sessionSaver := session.NewSaver(store)
	auth := NewAuthenticator(store, config, userRepo, comparator, sessionSaver)
	return NewBlog(store, auth, db), nil
}

