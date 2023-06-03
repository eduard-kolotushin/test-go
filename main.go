package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/eduard-kolotushin/test-go/models"
	"github.com/eduard-kolotushin/test-go/secman"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(models.Articles)
}

func returnStaticFiles(prefix string, dir string) http.Handler {
	fmt.Println("Endpoint Hit: returnStaticFiles")
	return http.StripPrefix(prefix, http.FileServer(http.Dir(dir)))
}

func getVaultSecretsList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getVaultSecretsList")
	data, err := secman.GetListSecrets("secret/my-app/")
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(data)
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllArticles)
	myRouter.PathPrefix("/static/").Handler(returnStaticFiles("/static/", "./stats/"))
	myRouter.HandleFunc("/vault/secrets/list", getVaultSecretsList)
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	models.Articles = []models.Article{
		{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}
