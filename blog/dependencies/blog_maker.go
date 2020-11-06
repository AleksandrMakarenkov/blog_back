package dependencies

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/sessions"
	_ "github.com/jackc/pgx/stdlib"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
	"io/ioutil"
	"log"
	"os"
	"vue_back/blog"
	"vue_back/blog/repository"
	"vue_back/blog/service/password"
	"vue_back/blog/service/session"
)

func MakeBlog() (*blog.Blog, error) {
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

	dsn := fmt.Sprintf( "postgresql://%s:%s@db/%s", string(dbUser), string(dbPassword), dbName)
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
	reformDb := reform.NewDB(db, postgresql.Dialect, reform.NewPrintfLogger(logger.Printf))

	config.DB = db
	config.Reform = reformDb

	userRepo := repository.NewUserRepository(reformDb)

	store := sessions.NewCookieStore([]byte(config.Secret))
	comparator := password.NewComparator()
	sessionSaver := session.NewSaver(store)
	auth := blog.NewAuthenticator(store, config, userRepo, comparator, sessionSaver)

	return blog.NewBlog(store, auth, db, config), nil
}
