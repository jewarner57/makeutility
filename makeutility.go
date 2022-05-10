package makeutility

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"
)

type RepoList struct {
	TotalCount int    `json:"total_count"`
	Items      []Repo `json:"items"`
}

type Repo struct {
	Name        string    `json:"name"`
	GithubUrl   string    `json:"html_url"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	PushedAt    time.Time `json:"pushed_at"`
	CloneUrl    string    `json:"clone_url"`
}

type User struct {
	Username string `json:"username"`
}

// TODO
// Add table tests and benchmark tests
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
	setUsername := flag.String("setUser", "", "Save a new default username.")
	flag.Parse()

	if *sortBy != "" {
		sortString := "%20sort:" + *sortBy
		sortBy = &sortString
	}

	if *language != "" {
		languageString := "%20language:" + *language
		language = &languageString
	}

	if *setUsername != "" {
		updateUserConfig(*setUsername)
	}

	var username string = loadConfig()
	repoList := fetchRepoData(username, *query, *sortBy, *language)

	printRepoList(repoList)
}

func printRepoList(repos []Repo) {
	width, _, _ := terminal.GetSize(0)
	fmt.Println(strings.Repeat("-", width))
	for _, repo := range repos {

		boldGreen := color.New(color.FgGreen).Add(color.Bold)
		underLineCyan := color.New(color.FgBlue).Add(color.Underline)

		boldGreen.Println(repo.Name)
		if repo.Description != "" {
			fmt.Println(repo.Description)
		}
		fmt.Println()

		fmt.Print("View: ")
		underLineCyan.Println(repo.GithubUrl)
		fmt.Print("Clone: ")
		underLineCyan.Println(repo.CloneUrl)
		fmt.Println(strings.Repeat("-", width))
	}
}

func fetchRepoData(username string, q string, sortBy string, language string) []Repo {
	response, err := http.Get(
		"https://api.github.com/search/repositories?q=" + q + "%20user:" + username + sortBy + language,
	)

	// fmt.Println("https://api.github.com/search/repositories?q=" + q + "%20user:" + username + sortBy + language)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var repoList RepoList
	if err := json.Unmarshal([]byte(responseData), &repoList); err != nil {
		log.Fatal(err)
	}

	return repoList.Items
}

func loadConfig() string {
	config, err := os.Open("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	configData, err := ioutil.ReadAll(config)
	if err != nil {
		log.Fatal(err)
	}

	var user User
	if json.Unmarshal([]byte(configData), &user); err != nil {
		log.Fatal(err)
	}

	return user.Username
}

func updateUserConfig(username string) {
	newUser := User{Username: username}

	file, _ := json.MarshalIndent(newUser, "", " ")

	_ = ioutil.WriteFile("config.json", file, 0644)
}
