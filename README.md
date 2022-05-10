# BEW - Statically Typed Languages Final Project

## Github Repo Viewer CLI Utility
  * This utility provides a simple CLI to quickly view, search, and sort a user's github repositories with one quick command.

## Installation
  ```
  git clone git@github.com:jewarner57/makeutility.git
  cd makeutility
  go build ./main.go
  ```

## Usage

### Flags
  * lang string
    * Filter results by language
  * query string
    * A query string to filter results.
  * setUser string
    * Save a new default username.
  * sort string
    * A string indicating how to sort the results.
    	* One of: interactions, reactions, author-date, comitter-date, updated

## Examples
![example usage](./assets/ex1.png)
![example usage](./assets/ex2.png)