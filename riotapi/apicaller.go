package riotapi

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"github.com/joho/godotenv"
	"os"
)

type SummonerDto struct {
	ProfileIconID int    `json:"profileIconId"`
	Name          string `json:"name"`
	SummonerLevel int    `json:"summonerLevel"`
	RevisionDate  int64  `json:"revisionDate"`
	ID            int    `json:"id"`
	AccountID     int    `json:"accountId"`
}

type MatchListDto struct {
	Matches []struct {
		Lane       string `json:"lane"`
		GameID     int64  `json:"gameId"`
		Champion   int    `json:"champion"`
		PlatformID string `json:"platformId"`
		Timestamp  int64  `json:"timestamp"`
		Queue      int    `json:"queue"`
		Role       string `json:"role"`
		Season     int    `json:"season"`
	} `json:"matches"`
	TotalGames int `json:"totalGames"`
	StartIndex int `json:"startIndex"`
	EndIndex   int `json:"endIndex"`
}

// For out of testing use, when not using a temp dev key.

func MatchListCaller(sumName string) MatchListDto {
	//Gets SummonerID from Summoner Name
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("RIOT_API_KEY")
	url := fmt.Sprintf("https://na1.api.riotgames.com/lol/summoner/v3/summoners/by-name/%s", sumName)
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("X-Riot-Token", apiKey)

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		log.Fatalf("Bad HTTP Request with error: %s.", err)
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading json body: %s", err)
	}
	var summdto SummonerDto
	//fmt.Println(string(data))

	if err := json.Unmarshal(data, &summdto); err != nil {
		log.Fatalf("Error unmarshaling json: %s", err)
	}

	//Gets match history (25 games) from ID.
	res, err = http.Get(fmt.Sprintf("https://na1.api.riotgames.com/lol/match/v3/matchlists/by-account/%d?api_key=%s&endIndex=5", summdto.AccountID, apiKey))
	if err != nil {
		log.Fatalf("Bad HTTP Request with error %s.", err)
	}
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading json body: %s", err)
	}
	//fmt.Println(string(data))
	var matchList MatchListDto
	if err := json.Unmarshal(data, &matchList); err != nil {
		log.Fatalf("Error unmarshaling json: %s", err)
	}
	return matchList
}


