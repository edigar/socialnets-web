package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/edigar/socialnets-web/src/config"
	"github.com/edigar/socialnets-web/src/cookies"
	"github.com/edigar/socialnets-web/src/dtos"
	"github.com/edigar/socialnets-web/src/responses"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	login, err := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/login", config.ApiUrl)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(login))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleErrorStatusCode(w, response)
		return
	}

	var authenticationDTO dtos.AuthenticationDTO
	if err = json.NewDecoder(response.Body).Decode(&authenticationDTO); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Error: err.Error()})
		return
	}

	if err = cookies.Save(w, authenticationDTO); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Error: err.Error()})
		return
	}

	responses.JSON(w, http.StatusOK, nil)
}
