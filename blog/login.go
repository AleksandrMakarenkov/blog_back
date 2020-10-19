package blog

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/sessions"
	_ "github.com/jackc/pgx/stdlib"
	_ "gopkg.in/reform.v1"
	"net/http"
	"vue_back/blog/repository"
	"vue_back/blog/service/password"
	"vue_back/blog/service/session"
)

type loginAttempt struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Authenticator struct {
	store sessions.Store
	config *Config
	userRepo *repository.UserRepository
	comparator *password.Comparator
	sessionSaver *session.Saver
}

func NewAuthenticator(
	store sessions.Store,
	config *Config,
	userRepo *repository.UserRepository,
	comparator *password.Comparator,
	sessionSaver *session.Saver,
) *Authenticator {
	return &Authenticator{
		store: store,
		config: config,
		userRepo: userRepo,
		comparator: comparator,
		sessionSaver: sessionSaver,
	}
}


func (a *Authenticator) LoginHandler(writer http.ResponseWriter, req *http.Request) {
	// parse json
	var la loginAttempt
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&la)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	// find user
	user, err := a.userRepo.FindByEmail(la.Email)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	// compare password
	isEqual, err := a.comparator.Compare(la.Password, user.Password)
	if err != nil || isEqual != true {
		fmt.Println(err)
		writer.WriteHeader(http.StatusUnauthorized)
		_, err = writer.Write([]byte("bad creds"))
		return
	}
	// save session
	err = a.sessionSaver.Save(user, req, writer)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
}
