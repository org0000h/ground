package service

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if _, ok := fmt.Fprint(w, "Welcome!\n"); ok != nil {
		// todo
	}
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if _, ok := fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name")); ok != nil {
		// err handle
	}
}

func GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	if _, ok := fmt.Fprintf(w, "you are get user %s", uid); ok != nil {
		// err handle
	}
}

func ModifyUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	if _, ok := fmt.Fprintf(w, "you are modify user %s", uid); ok != nil {

	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uid := ps.ByName("uid")
	if _, ok := fmt.Fprintf(w, "you are delete user %s", uid); ok != nil {
		// err handle
	}

}

func AddUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// uid := r.FormValue("uid")
	uid := ps.ByName("uid")
	if _, ok := fmt.Fprintf(w, "you are add user %s", uid); ok != nil {
		// err handle
	}
	// fmt.Fprintf(w, "you are add user %s", r)
}
