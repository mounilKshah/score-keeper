/*
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    req, err := http.NewRequest("GET", "http://api.themoviedb.org/3/tv/popular", nil)
    if err != nil {
        log.Print(err)
        os.Exit(1)
    }

    q := req.URL.Query()
    q.Add("api_key", "key_from_environment_or_flag")
    q.Add("another_thing", "foo & bar")
    req.URL.RawQuery = q.Encode()

    fmt.Println(req.URL.String())
    // Output:
    // http://api.themoviedb.org/3/tv/popular?another_thing=foo+%26+bar&api_key=key_from_environment_or_flag
}
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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
	myRouter.HandleFunc("/fixturesold/:id", getAlbumByID)
	myRouter.HandleFunc("/fixtures/:id", getId)
	router := gin.Default()
	router.GET("/albums/:id", getAlbumByIDNew)

	// myRouter.HandleFunc("/all", returnAllArticles)
	// myRouter.HandleFunc("/article/{id}", returnSingleArticle)

	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	// log.Fatal(http.ListenAndServe(":9000", myRouter))
	router.Run("localhost:9000")

}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByIDNew(c *gin.Context) {
	id := c.Param("id")

	fmt.Println(id)
	url := "https://v3.football.api-sports.io/fixtures/events"
	method := "GET"

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	req.Header.Add("x-rapidapi-key", "b8920a4d3ff30be1e631b5f5c9e676f6")
	req.Header.Add("x-rapidapi-host", "v3.football.api-sports.io")

	q := req.URL.Query()
	// q.Add("api_key", "key_from_environment_or_flag")
	q.Add("fixture", id)
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())

	client := &http.Client{}
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
	// fmt.Fprintf(w, "Timezone API hit")
	// // fmt.Fprintf(w, string(body))
	fmt.Println(string(body))
	fmt.Println("Endpoint Hit: Get fixtures by ID page")
	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	// for _, a := range albums {
	//     if a.ID == id {
	//         c.IndentedJSON(http.StatusOK, a)
	//         return
	//     }
	// }
	// c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func getId(w http.ResponseWriter, r *http.Request) {
	url := "https://v3.football.api-sports.io/fixtures/events"
	method := "GET"

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	req.Header.Add("x-rapidapi-key", "b8920a4d3ff30be1e631b5f5c9e676f6")
	req.Header.Add("x-rapidapi-host", "v3.football.api-sports.io")

	q := req.URL.Query()
	// q.Add("api_key", "key_from_environment_or_flag")
	q.Add("fixtures", "foo & bar")
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())

	// Output:
	// http://api.themoviedb.org/3/tv/popular?another_thing=foo+%26+bar&api_key=key_from_environment_or_flag
}

func getAlbumByID(w http.ResponseWriter, r *http.Request) {
	// id := c.Param("id")
	fmt.Println(r)
	fmt.Println(r.Body)
	fmt.Println(r.URL)
	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	// for _, a := range albums {
	// 	if a.ID == id {
	// 		c.IndentedJSON(http.StatusOK, a)
	// 		return
	// 	}
	// }
	// c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})

}

func getTimezone(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
	fmt.Println(r.URL)

	// url := "https://v3.football.api-sports.io/teams"
	url := "https://v3.football.api-sports.io/fixtures/events?fixture="
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

/*
package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbumByID)
    router.POST("/albums", postAlbums)

    router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop through the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

*/
