package exampleop

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	queryAuthRequestID = "authRequestID"
)

type login struct {
	authenticate authenticate
	router       *mux.Router
	callback     func(string) string
}

func NewLogin(authenticate authenticate, callback func(string) string) *login {
	l := &login{
		authenticate: authenticate,
		callback:     callback,
	}
	l.createRouter()
	return l
}

func (l *login) createRouter() {
	l.router = mux.NewRouter()
	l.router.Path("/username").Methods("POST").HandlerFunc(l.checkLoginHandler)
}

type authenticate interface {
	CheckUsernamePassword(username, password, id string) error
}

type loginResponse struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	CallbackURL string `json:"callback_url"`
}

func (l *login) checkLoginHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, fmt.Sprintf("cannot parse form:%s", err), http.StatusInternalServerError)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	id := r.FormValue("id")
	err = l.authenticate.CheckUsernamePassword(username, password, id)
	if err != nil {
		//renderLogin(w, id, err)
		buf, _ := json.Marshal(loginResponse{Code: http.StatusInternalServerError, Message: "login failed, " + err.Error(), CallbackURL: ""})
		_, _ = w.Write(buf)
		return
	} else {
		http.Redirect(w, r, l.callback(id), http.StatusFound)
	}

}
