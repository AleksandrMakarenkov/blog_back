package session

import (
	"github.com/gorilla/sessions"
	"net/http"
	"vue_back/blog/model"
)

type Saver struct {
	store sessions.Store
}

const CookieName = "blog_session"

func NewSaver(store sessions.Store) *Saver {
	return &Saver{store: store}
}

func (s *Saver) Save(user *model.Account, req *http.Request, writer http.ResponseWriter) error {
	session, err := s.store.Get(req, CookieName)
	if err != nil {
		return err
	}
	session.Values["id"] = user.Id
	err = session.Save(req, writer)
	if err != nil {
		return err
	}
	return nil
}
