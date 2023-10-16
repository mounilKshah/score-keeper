package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// getTimezoneOld()
	handleRequests()
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/timezone", getTimezone)
	// myRouter.HandleFunc("/all", returnAllArticles)
	// myRouter.HandleFunc("/article/{id}", returnSingleArticle)

	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":9000", myRouter))
}

func getTimezone(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
	fmt.Println(r.URL)

	url := "https://v3.football.api-sports.io/teams"
	// url := "https://v3.football.api-sports.io/players/seasons"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("x-rapidapi-key", "xxx")
	req.Header.Add("x-rapidapi-host", "v3.football.api-sports.io")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, "Timezone API hit")
	// // fmt.Fprintf(w, string(body))
	fmt.Println(string(body))
	fmt.Println("Endpoint Hit: Timezone page")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}	


