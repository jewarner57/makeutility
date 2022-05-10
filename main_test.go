package main

import (
	"testing"
)

type Args struct {
	Q        string
	SortBy   string
	Language string
}

/*
	Name        string    `json:"name"`
	GithubUrl   string    `json:"html_url"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	PushedAt    time.Time `json:"pushed_at"`
	CloneUrl    string    `json:"clone_url"`
*/

func TestTableFetchRepoData(t *testing.T) {
	var tests = []struct {
		input    Args
		expected []string
	}{
		{Args{Q: "SvgToPointConverter"}, []string{"SvgToPointConverter"}},
		{Args{Q: "prosperousuniversepriceapi"}, []string{"prosperousuniversepriceapi"}},
		{Args{Q: "this-is-a-package-name-that-doesnt-exist"}, []string{}},
		{Args{Language: "%20language:hack"}, []string{"MLH-Orientation-Hackathon"}},
		{Args{SortBy: "%20sort:updated", Language: "%20language:java"}, []string{"Energy-Comparison-Java-Applet", "SvgToPointConverter"}},
	}

	for _, test := range tests {
		output := fetchRepoData("jewarner57", test.input.Q, test.input.SortBy, test.input.Language)

		// Check result length is expected length
		if len(output) != len(test.expected) {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}

		// Check names match
		for index, repo := range output {
			// fmt.Println(repo.Name, test.expected[index])
			if repo.Name != test.expected[index] {
				t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
			}
		}
	}
}
