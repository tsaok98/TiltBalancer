package main

import (
	"bufio"
	"os"
	"fmt"
	"log"
	"Untiltable/riotapi"
	"regexp"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter your Summoner Name:")
	reader.Scan()
	sumName := reader.Text()
	b, _ := regexp.MatchString("[0-9\\p{L}_.]+", sumName)
	if !b {
		log.Fatal("Invalid Summoner Name.")
	}
	matchList := riotapi.MatchListCaller(sumName)
	fmt.Println(matchList)
}