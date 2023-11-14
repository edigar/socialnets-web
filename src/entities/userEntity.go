package entities

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/edigar/socialnets-web/src/config"
	"github.com/edigar/socialnets-web/src/requests"
	"net/http"
	"time"
)

type User struct {
	Id        uint64    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Nick      string    `json:"nick"`
	CreatedAt time.Time `json:"createdAt"`
	Followers []User    `json:"followers"`
	Following []User    `json:"following"`
	Posts     []Post    `json:"posts"`
}

func GetUser(userId uint64, r *http.Request) (User, error) {
	userChannel := make(chan User)
	followersChannel := make(chan []User)
	followingChannel := make(chan []User)
	postsChannel := make(chan []Post)

	go GetUserData(userChannel, userId, r)
	go GetFollowers(followersChannel, userId, r)
	go GetFollowing(followingChannel, userId, r)
	go GetPosts(postsChannel, userId, r)

	var (
		user      User
		followers []User
		following []User
		posts     []Post
	)

	for i := 0; i < 4; i++ {
		select {
		case userResponse := <-userChannel:
			if userResponse.Id == 0 {
				return User{}, errors.New("getting user error")
			}

			user = userResponse

		case followersResponse := <-followersChannel:
			if followersResponse == nil {
				return User{}, errors.New("getting user followers error")
			}

			followers = followersResponse

		case followingResponse := <-followingChannel:
			if followingResponse == nil {
				return User{}, errors.New("getting user following error")
			}

			following = followingResponse

		case postsResponse := <-postsChannel:
			if postsResponse == nil {
				return User{}, errors.New("getting user posts error")
			}

			posts = postsResponse
		}
	}

	user.Followers = followers
	user.Following = following
	user.Posts = posts

	return user, nil
}

func GetUserData(channel chan<- User, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/user/%d", config.ApiUrl, userId)
	response, err := requests.Authenticated(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- User{}
		return
	}
	defer response.Body.Close()

	var user User
	if err = json.NewDecoder(response.Body).Decode(&user); err != nil {
		channel <- User{}
		return
	}

	channel <- user
}

func GetFollowers(channel chan<- []User, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/user/%d/followers", config.ApiUrl, userId)
	response, err := requests.Authenticated(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var followers []User
	if err = json.NewDecoder(response.Body).Decode(&followers); err != nil {
		channel <- nil
		return
	}

	if followers == nil {
		channel <- make([]User, 0)
		return
	}

	channel <- followers
}

func GetFollowing(channel chan<- []User, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/user/%d/following", config.ApiUrl, userId)
	response, err := requests.Authenticated(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var following []User
	if err = json.NewDecoder(response.Body).Decode(&following); err != nil {
		channel <- nil
		return
	}

	if following == nil {
		channel <- make([]User, 0)
		return
	}

	channel <- following
}

func GetPosts(channel chan<- []Post, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/user/%d/posts", config.ApiUrl, userId)
	response, err := requests.Authenticated(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var posts []Post
	if err = json.NewDecoder(response.Body).Decode(&posts); err != nil {
		channel <- nil
		return
	}

	if posts == nil {
		channel <- make([]Post, 0)
		return
	}

	channel <- posts
}
