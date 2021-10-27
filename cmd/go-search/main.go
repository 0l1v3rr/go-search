package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	s "github.com/0l1v3rr/go-search/internal/search"
	u "github.com/0l1v3rr/go-search/internal/util"
)

var (
	searchTerm string = "-"
	pages      int    = 1
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

			results, err := s.GoogleSearch(searchTerm, pages)
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
			u.ShowOptions(searchTerm, pages)
		} else if strings.HasPrefix(input, "set terms") {
			searchTerm = strings.Join(args[2:], " ")
			fmt.Printf("terms => %s\n", searchTerm)
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
