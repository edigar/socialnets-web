package main

import (
	"fmt"
	"github.com/edigar/socialnets-web/src/config"
	"github.com/edigar/socialnets-web/src/cookies"
	"github.com/edigar/socialnets-web/src/router"
	"github.com/edigar/socialnets-web/src/template"
	"log"
	"net/http"
)

//func init() {
//	hashKey := hex.EncodeToString(securecookie.GenerateRandomKey(32))
//	fmt.Println(hashKey)
//
//	blockKey := hex.EncodeToString(securecookie.GenerateRandomKey(32))
//	fmt.Println(blockKey)
//}

func main() {
	config.Load()
	cookies.Setup()
	template.LoadTemplates()
	r := router.Generate()

	fmt.Printf("Running SocialNetsWEB on port %d...\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
