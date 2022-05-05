package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
)

type RepoList struct {
	TotalCount int    `json:"total_count"`
	Items      []Repo `json:"items"`
}

type Repo struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	PushedAt    time.Time `json:"pushed_at"`
	CloneUrl    string    `json:"clone_url"`
}

// TODO
// Make sortable
// Allow user to create username config json file
// Add pretty formatting
// Add docs and general info to readme

func main() {
	query := flag.String("query", "", "A query string to filter results.")
	sortBy := flag.String(
		"sort",
		"",
		`A string indicating how to sort the results. 
		One of: interactions, reactions, author-date, comitter-date, updated`,
	)
	language := flag.String("lang", "", "Filter results by language")
	flag.Parse()

	if *sortBy != "" {
		sortString := "sort:" + *sortBy
		sortBy = &sortString
	}

	if *language != "" {
		languageString := "language:" + *language
		language = &languageString
	}

	var username string = "jewarner57"
	responseData := fetchRepoData(username, *query, *sortBy, *language)

	var repoList RepoList
	if err := json.Unmarshal([]byte(responseData), &repoList); err != nil {
		log.Fatal(err)
	}

	printRepoList(repoList)
}

func printRepoList(repoList RepoList) {
	for _, repo := range repoList.Items {

		green := color.New(color.FgGreen)
		boldGreen := green.Add(color.Bold)

		boldGreen.Println(repo.Name)
		fmt.Println(repo.Description)
		fmt.Println(repo.CloneUrl)
		fmt.Println("------------------")
	}
}

func fetchRepoData(username string, q string, sortBy string, language string) []byte {
	// +language:assembly&sort=stars&order=desc
	response, err := http.Get(
		"https://api.github.com/search/repositories?q=" + q + "%20user:" + username + "%20" + sortBy + "%20" + language,
	)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return responseData
}
