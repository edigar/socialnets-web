package controllers

import (
	"github.com/edigar/socialnets-web/src/cookies"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	cookies.Delete(w)
	http.Redirect(w, r, "/login", http.StatusFound)
}
