package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/edigar/socialnets-web/src/config"
	"github.com/edigar/socialnets-web/src/cookies"
	"github.com/edigar/socialnets-web/src/entities"
	"github.com/edigar/socialnets-web/src/requests"
	"github.com/edigar/socialnets-web/src/responses"
	"github.com/edigar/socialnets-web/src/template"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

func ShowLogin(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}

	template.ExecuteTemplate(w, "login.html", nil)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	template.ExecuteTemplate(w, "register-user.html", nil)
}

func ShowHome(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/post", config.ApiUrl)
	response, err := requests.Authenticated(r, http.MethodGet, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer response.Body.Close()
	if response.StatusCode >= 400 {
		responses.HandleErrorStatusCode(w, response)
		return
	}

	var posts []entities.Post
	if err = json.NewDecoder(response.Body).Decode(&posts); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Error: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	template.ExecuteTemplate(w, "home.html", struct {
		Posts  []entities.Post
		UserId uint64
	}{
		Posts:  posts,
		UserId: userId,
	})
}

func EditPosts(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	postId, err := strconv.ParseUint(param["postId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/post/%d", config.ApiUrl, postId)
	response, err := requests.Authenticated(r, http.MethodGet, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer response.Body.Close()
	if response.StatusCode >= 400 {
		responses.HandleErrorStatusCode(w, response)
		return
	}

	var post entities.Post
	if err = json.NewDecoder(response.Body).Decode(&post); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Error: err.Error()})
		return
	}

	template.ExecuteTemplate(w, "update-post.html", post)
}

func ShowUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))
	url := fmt.Sprintf("%s/user?search=%s", config.ApiUrl, nameOrNick)

	response, err := requests.Authenticated(r, http.MethodGet, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer response.Body.Close()
	if response.StatusCode >= 400 {
		responses.HandleErrorStatusCode(w, response)
		return
	}

	var users []entities.User
	if err = json.NewDecoder(response.Body).Decode(&users); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Error: err.Error()})
		return
	}

	template.ExecuteTemplate(w, "users.html", users)
}

func ShowProfile(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	userId, err := strconv.ParseUint(param["userId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	loggedUserId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	if userId == loggedUserId {
		http.Redirect(w, r, "/perfil", http.StatusFound)
		return
	}

	user, err := entities.GetUser(userId, r)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}

	template.ExecuteTemplate(w, "user.html", struct {
		User       entities.User
		LoggedUser uint64
	}{
		User:       user,
		LoggedUser: loggedUserId,
	})
}

func ShowLoggedUserProfile(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	userId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	user, err := entities.GetUser(userId, r)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}

	template.ExecuteTemplate(w, "profile.html", user)
}

func ShowEditUser(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	userId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	channel := make(chan entities.User)
	go entities.GetUserData(channel, userId, r)
	user := <-channel

	if user.Id == 0 {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: "Erro ao buscar o usuário"})
		return
	}

	template.ExecuteTemplate(w, "edit-user.html", user)
}

func ShowUpdatePassword(w http.ResponseWriter, r *http.Request) {
	template.ExecuteTemplate(w, "update-password.html", nil)
}
