package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"time"
)

type Repo struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	PushedAt    time.Time `json:"pushed_at"`
	CloneUrl    string    `json:"clone_url"`
}

func main() {
	var username string = "jewarner57"
	responseData := fetchRepoData(username)

	repoList := make([]Repo, 0)
	if err := json.Unmarshal([]byte(responseData), &repoList); err != nil {
		log.Fatal(err)
	}

	sort.Slice(repoList, func(i, j int) bool {
		return repoList[i].UpdatedAt.After(repoList[j].UpdatedAt)
	})

	fmt.Println(string(repoList[0].CloneUrl))
}

func fetchRepoData(username string) []byte {
	info, err := os.Stat("./cache.json")
	hoursSinceCacheUpdated := time.Since(info.ModTime()).Hours()
	if errors.Is(err, os.ErrNotExist) || hoursSinceCacheUpdated > 0.5 {
		updateCacheData(username)
	}

	jsonFile, err := os.Open("cache.json")
	if err != nil {
		log.Fatal(err)
	}

	cacheData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	return cacheData
}

func updateCacheData(username string) {
	response, err := http.Get("https://api.github.com/users/" + username + "/repos")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile("cache.json", []byte(responseData), 0644)
}
