package blog

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"io/ioutil"
	"net/http"
)

const SecretName = "BLOG_SECRET"

type Blog struct {
	router *mux.Router
	store  *sessions.Store
	db *sql.DB
}

func NewBlog(
	store sessions.Store,
	authenticator *Authenticator,
	db *sql.DB,
) *Blog {
	r := mux.NewRouter()
	r.HandleFunc("/", Index)
	r.HandleFunc("/login", authenticator.LoginHandler).Methods("POST")
	//r.HandleFunc("/article", Article).Methods("POST")

	return &Blog{
		router: r,
		store:  &store,
	}
}

func (blog *Blog) Run () {
	var err error
	err = http.ListenAndServe(":9090", blog.router)
	if err != nil {
		fmt.Println(err)
	}
}

func (blog *Blog) CloseDB() error {
	return blog.db.Close()
}

func Index(writer http.ResponseWriter, req *http.Request) {
	var b , _ = ioutil.ReadFile("./static/index.html")
	_, _ = writer.Write(b)
}

//func Article(writer http.ResponseWriter, req *http.Request) {
	//session, _ := store.Get(req, CookieName)
	//fmt.Println(session.Values["id"])
//}
