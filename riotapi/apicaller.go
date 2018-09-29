package riotapi

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
)

type SummonerDto struct {
	ID            int    `json:"id"`
	AccountID     int    `json:"accountId"`
	Name          string `json:"name"`
	ProfileIconID int    `json:"profileIconId"`
	RevisionDate  int64  `json:"revisionDate"`
	SummonerLevel int    `json:"summonerLevel"`
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
	EndIndex   int `json:"endIndex"`
	StartIndex int `json:"startIndex"`
	TotalGames int `json:"totalGames"`
}

// For out of testing use, when not using a temp dev key.
//var (
//	apiKey = os.Getenv("RIOTAPIKEY")
//)

func MatchListCaller(sumName string) MatchListDto {
	//Gets SummonerID from Summoner Name
	fmt.Println(sumName)
	//"https://na1.api.riotgames.com/lol/summoner/v3/summoners/by-name/Starfirephoenix?api_key=RGAPI-c438d1a9-cfbe-4fa4-811f-98a8988c494e"
	//fmt.Sprintf("https://na1.api.riotgames.com/lol/summoner/v3/summoners/by-name/%s?api_key=%s", sumName, apiKey)
	//var apiKey = ""
	//var url = fmt.Sprintf("https://na1.api.riotgames.com/lol/summoner/v3/summoners/by-name/%s?api_key=%s", sumName, apiKey)
	request, _ := http.NewRequest("GET", fmt.Sprintf("https://na1.api.riotgames.com/lol/summoner/v3/summoners/by-name/%s", sumName), nil)
	//request.Header.Set("X-Riot-Token", apiKey)
	//request.Header.Set("Accept-Charset", "application/x-www-form-urlencoded; charset=UTF-8")
	//request.Header.Set("Origin", "https://developer.riotgames.com")
	//request.Header.Set("Accept-Language", "en-US,en;q=0.9")
	//request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")

	client := &http.Client{}
	//var url = strings.Join([]string{"https://na1.api.riotgames.com/lol/summoner/v3/summoners/by-name/", sumName, "?api_key=", apiKey}, "")
	fmt.Println(fmt.Sprintf("https://na1.api.riotgames.com/lol/summoner/v3/summoners/by-name/Starfirephoenix"))
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
	fmt.Println(string(data))

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
	fmt.Println(string(data))
	var matchList MatchListDto
	if err := json.Unmarshal(data, &matchList); err != nil {
		log.Fatalf("Error unmarshaling json: %s", err)
	}
	return matchList
}


