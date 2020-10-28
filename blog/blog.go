package blog

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"io/ioutil"
	"net/http"
)

const EnvNameOfSecret = "BLOG_SECRET"

type Blog struct {
	router *mux.Router
	store  *sessions.Store
	db     *sql.DB
	config *Config
}

func NewBlog(
	store sessions.Store,
	authenticator *Authenticator,
	db *sql.DB,
	config *Config,
) *Blog {
	r := mux.NewRouter()
	r.HandleFunc("/", Index)
	r.HandleFunc("/login", authenticator.LoginHandler).Methods("POST")

	return &Blog{
		router: r,
		store:  &store,
		db:     db,
		config: config,
	}
}

func (blog *Blog) Run() {
	var err error
	err = http.ListenAndServe(fmt.Sprintf(":%d",blog.config.port), blog.router)
	if err != nil {
		fmt.Println(err)
	}
}

func (blog *Blog) CloseDB() error {
	return blog.db.Close()
}

func Index(writer http.ResponseWriter, req *http.Request) {
	var b, _ = ioutil.ReadFile("./static/index.html")
	_, _ = writer.Write(b)
}
