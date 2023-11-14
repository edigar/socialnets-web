package cookies

import (
	"github.com/edigar/socialnets-web/src/config"
	"github.com/edigar/socialnets-web/src/dtos"
	"github.com/gorilla/securecookie"
	"net/http"
	"time"
)

var s *securecookie.SecureCookie

func Setup() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

func Save(w http.ResponseWriter, authenticationDTO dtos.AuthenticationDTO) error {
	data := map[string]string{
		"id":    authenticationDTO.Id,
		"token": authenticationDTO.Token,
	}

	encodeData, err := s.Encode("data", data)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "data",
		Value:    encodeData,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}

func Read(r *http.Request) (map[string]string, error) {
	cookie, err := r.Cookie("data")
	if err != nil {
		return nil, err
	}

	values := make(map[string]string)
	if err = s.Decode("data", cookie.Value, &values); err != nil {
		return nil, err
	}

	return values, nil
}

func Delete(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "data",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	})
}
