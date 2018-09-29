package main

import (
	"bufio"
	"os"
	"fmt"
	"log"
	"Untiltable/riotapi"
	"regexp"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your Summoner Name:")
	sumName, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Invalid Summoner Name.")
	}
	sumName = strings.Replace(sumName, "\n", "", -1)
	b, err := regexp.MatchString("[0-9\\p{L}_.]+", sumName)
	if !b {
		log.Fatal("Invalid Summoner Name.")
	}
	matchList := riotapi.MatchListCaller(sumName)
	fmt.Println(matchList)
}