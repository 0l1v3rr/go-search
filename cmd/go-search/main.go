package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	s "github.com/0l1v3rr/go-search/internal/search"
	u "github.com/0l1v3rr/go-search/internal/util"
)

var (
	searchTerm  string = "-"
	pages       int    = 1
	resultCount int    = 20
)

func main() {
	u.Clear()
	u.Motd()
	reader := bufio.NewReader(os.Stdin)

	for {
		scanner()
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		input = strings.ToLower(input)
		fmt.Print(string("\033[0m"))

		args := strings.Split(input, " ")

		if strings.HasPrefix(input, "ex") {
			break
		}

		if strings.HasPrefix(input, "search") {
			if searchTerm == "-" {
				printError("Please provide a valid search term!")
				continue
			}

			results, err := s.GoogleSearch(searchTerm, pages, resultCount)
			if err != nil {
				return
			}

			fmt.Println()
			for _, res := range results {
				fmt.Print(string("\033[34;1m"), "[")
				if res.Rank < 10 {
					fmt.Print("0")
				}
				fmt.Print(res.Rank)
				fmt.Print("] ")
				fmt.Println(string("\033[0m"), res.Url)
			}
			fmt.Println()
		} else if strings.HasPrefix(input, "show options") {
			u.ShowOptions(searchTerm, pages, resultCount)
		} else if strings.HasPrefix(input, "set terms") {
			if len(args) < 3 {
				printError("Please provide valid arguments!")
				continue
			}
			searchTerm = strings.Join(args[2:], " ")
			fmt.Printf("terms => %s\n", searchTerm)
		} else if strings.HasPrefix(input, "set pages") {
			if len(args) < 3 {
				printError("Please provide valid arguments!")
				continue
			}
			converted, err := strconv.Atoi(args[2])
			if err != nil {
				printError("Please provide valid arguments!")
				continue
			}
			pages = converted
			fmt.Printf("pages => %v\n", pages)
		} else if strings.HasPrefix(input, "set count ") {
			if len(args) < 3 {
				printError("Please provide valid arguments!")
				continue
			}
			converted, err := strconv.Atoi(args[2])
			if err != nil {
				printError("Please provide valid arguments!")
				continue
			}
			resultCount = converted
			fmt.Printf("count => %v\n", resultCount)
		} else {
			printError("Unknown command.")
		}

	}
}

func scanner() {
	colorReset := "\033[0m"

	fmt.Print(string(colorReset))
	fmt.Print(string("\033[4m"), "search")
	fmt.Print(string(colorReset), " > ")
	fmt.Print(string("\033[33;1m"))
}

func printError(e string) {
	colorReset := "\033[0m"
	colorRed := "\033[31;1m"

	fmt.Print(string(colorRed), "[!]")
	fmt.Println(string(colorReset), e)
}
