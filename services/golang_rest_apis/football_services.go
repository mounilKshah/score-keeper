package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var request_url = "https://v3.football.api-sports.io"
var host_url = "v3.football.api-sports.io"
var api_key = ""
var GET = "GET"

type paramBody struct {
	request_param string
	param_value   string
}

// Fetch API key and links from file and use them
func readAPISecretsFile() {
	file, err := os.Open("API_SECRET.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	content, err := ioutil.ReadFile("API_SECRET.txt")
	if err != nil {
		fmt.Println(err)
	}
	str := string(content)
	key_value := strings.Split(str, "|")
	api_key = key_value[1]
}

func handleRESTRequests() {
	readAPISecretsFile()
	router := gin.Default()
	router.GET("/country-data", getAllCountries)
	router.GET("/season-data", getAllSeasons)
	router.GET("/league-data", getAllLeagues)
	router.GET("/standings/:league_id/:season_id", getStandings)
	router.GET("/player-details/:player_id/:season_id", getPlayerInfo)
	router.GET("/squad-data/:team_id", getSquadInfo)
	router.GET("/team-data/:team_id", getTeamInfo)
	router.GET("/team-statistics/:season_id/:team_id/:league_id", getTeamStats)
	router.GET("/prediction/:fixture_id", getPredictions)

	router.Run("localhost:8888")
}

func main() {
	handleRESTRequests()
}

func getPredictions(c *gin.Context) {
	/*
		curl --request GET \
			--url 'https://v3.football.api-sports.io/predictions?fixture=198772' \
			--header 'x-rapidapi-host: v3.football.api-sports.io' \
			--header 'x-rapidapi-key: XxXxXxXxXxXxXxXxXxXxXxXx'
	*/
	url := request_url + "/" + "predictions"
	req, err := http.NewRequest(GET, url, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	var param_bodyList []paramBody

	fixture := c.Param("fixture_id")

	param_bodyList = append(param_bodyList, paramBody{"fixture", fixture})

	processRequest(c, req, param_bodyList)
	// processReqWithoutParamBody(c, req)

}
func getTeamStats(c *gin.Context) {
	/*
		curl --request GET \
			--url 'https://v3.football.api-sports.io/teams/statistics?season=2019&team=33&league=39' \
			--header 'x-rapidapi-host: v3.football.api-sports.io' \
			--header 'x-rapidapi-key: XxXxXxXxXxXxXxXxXxXxXxXx'
	*/
	url := request_url + "/" + "teams" + "/" + "statistics"
	req, err := http.NewRequest(GET, url, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	var param_bodyList []paramBody

	season := c.Param("season_id")
	team := c.Param("team_id")
	league := c.Param("league_id")

	param_bodyList = append(param_bodyList, paramBody{"season", season})
	param_bodyList = append(param_bodyList, paramBody{"team", team})
	param_bodyList = append(param_bodyList, paramBody{"league", league})

	processRequest(c, req, param_bodyList)

}

func getTeamInfo(c *gin.Context) {
	/*
		curl --request GET \
			--url 'https://v3.football.api-sports.io/teams?id=33' \
			--header 'x-rapidapi-host: v3.football.api-sports.io' \
			--header 'x-rapidapi-key: XxXxXxXxXxXxXxXxXxXxXxXx'
	*/
	url := request_url + "/" + "teams"
	req, err := http.NewRequest(GET, url, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	var param_bodyList []paramBody

	id := c.Param("team_id")

	param_bodyList = append(param_bodyList, paramBody{"id", id})

	processRequest(c, req, param_bodyList)

}

func getSquadInfo(c *gin.Context) {
	/*
		curl --request GET \
			--url 'https://v3.football.api-sports.io/players?id=276&season=2019' \
			--header 'x-rapidapi-host: v3.football.api-sports.io' \
			--header 'x-rapidapi-key: XxXxXxXxXxXxXxXxXxXxXxXx'
	*/
	url := request_url + "/" + "players" + "/" + "squads"
	req, err := http.NewRequest(GET, url, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	var param_bodyList []paramBody

	team := c.Param("team_id")

	param_bodyList = append(param_bodyList, paramBody{"team", team})

	processRequest(c, req, param_bodyList)

}

func getPlayerInfo(c *gin.Context) {
	/*
		curl --request GET \
			--url 'https://v3.football.api-sports.io/players?id=276&season=2019' \
			--header 'x-rapidapi-host: v3.football.api-sports.io' \
			--header 'x-rapidapi-key: XxXxXxXxXxXxXxXxXxXxXxXx'
	*/
	url := request_url + "/" + "players"
	req, err := http.NewRequest(GET, url, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	var param_bodyList []paramBody

	id := c.Param("player_id")
	season := c.Param("season_id")

	param_bodyList = append(param_bodyList, paramBody{"id", id})
	param_bodyList = append(param_bodyList, paramBody{"season", season})

	processRequest(c, req, param_bodyList)

}

func getStandings(c *gin.Context) {
	/*
		curl --request GET \
			--url 'https://v3.football.api-sports.io/standings?league=39&season=2019' \
			--header 'x-rapidapi-host: v3.football.api-sports.io' \
			--header 'x-rapidapi-key: XxXxXxXxXxXxXxXxXxXxXxXx'

	*/
	url := request_url + "/" + "standings"
	req, err := http.NewRequest(GET, url, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	var param_bodyList []paramBody

	league := c.Param("league_id")
	season := c.Param("season_id")

	param_bodyList = append(param_bodyList, paramBody{"league", league})
	param_bodyList = append(param_bodyList, paramBody{"season", season})

	temp_var := c.Params
	fmt.Println("Params in context: ")
	fmt.Println(temp_var)

	processRequest(c, req, param_bodyList)

}

func getAllLeagues(c *gin.Context) {
	/*
		curl --request GET \
		--url https://v3.football.api-sports.io/leagues \
		--header 'x-rapidapi-host: v3.football.api-sports.io' \
		--header 'x-rapidapi-key: XxXxXxXxXxXxXxXxXxXxXxXx'
	*/
	url := request_url + "/" + "leagues"
	req, err := http.NewRequest(GET, url, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	var param_bodyList []paramBody
	processRequest(c, req, param_bodyList)
}

func getAllSeasons(c *gin.Context) {
	/*
		curl --request GET \
		--url https://v3.football.api-sports.io/leagues/seasons \
		--header 'x-rapidapi-host: v3.football.api-sports.io' \
		--header 'x-rapidapi-key: XxXxXxXxXxXxXxXxXxXxXxXx'
	*/

	url := request_url + "/" + "leagues" + "/" + "seasons"
	req, err := http.NewRequest(GET, url, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	var param_bodyList []paramBody
	processRequest(c, req, param_bodyList)
}

func getAllCountries(c *gin.Context) {
	/*
		curl --request GET \
		--url https://v3.football.api-sports.io/countries \
		--header 'x-rapidapi-host: v3.football.api-sports.io' \
		--header 'x-rapidapi-key: XxXxXxXxXxXxXxXxXxXxXxXx'

	*/
	url := request_url + "/" + "countries"
	req, err := http.NewRequest(GET, url, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	var param_bodyList []paramBody
	processRequest(c, req, param_bodyList)
}

func processRequest(c *gin.Context, get_request *http.Request, param_list []paramBody) {
	request_query := get_request.URL.Query()
	if len(param_list) != 0 {
		for _, v := range param_list {
			fmt.Println(v.request_param + ": " + v.param_value)
			request_query.Add(v.request_param, v.param_value)

		}
	}
	get_request.URL.RawQuery = request_query.Encode()
	get_request.Header.Add("x-rapidapi-key", api_key)
	get_request.Header.Add("x-rapidapi-host", host_url)
	client := &http.Client{}
	res, err := client.Do(get_request)
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

	var json_map map[string]any

	err = json.Unmarshal(body, &json_map)
	if err != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
		return
	}

	c.IndentedJSON(http.StatusOK, json_map)
}

// This uses gin.Params instead of gin.Param
// Currently function is not used as the earlier version works fine
func processReqWithoutParamBody(c *gin.Context, get_request *http.Request) {
	request_query := get_request.URL.Query()

	if len(c.Params) != 0 {
		for _, v := range c.Params {
			fmt.Println(v.Key)
			fmt.Println(v.Value)
			request_query.Add(v.Key, v.Value)
		}
	}

	get_request.URL.RawQuery = request_query.Encode()
	get_request.Header.Add("x-rapidapi-key", api_key)
	get_request.Header.Add("x-rapidapi-host", host_url)
	client := &http.Client{}
	res, err := client.Do(get_request)
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

	var json_map map[string]any

	err = json.Unmarshal(body, &json_map)
	if err != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
		return
	}

	c.IndentedJSON(http.StatusOK, json_map)
}
