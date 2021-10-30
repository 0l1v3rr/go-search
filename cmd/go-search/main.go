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
	site        string = "-"
	filetype    string = "-"
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
			fmt.Print(string("\033[0m"))
			break
		}

		if strings.HasPrefix(input, "search") {
			moreinfo := false
			keywords := ""
			if searchTerm != "-" {
				keywords += searchTerm
			}
			if site != "-" {
				keywords += fmt.Sprintf(" site:%s", site)
			}
			if filetype != "-" {
				keywords += fmt.Sprintf(" filetype:%s", filetype)
			}

			if strings.HasPrefix(input, "search -e") {
				moreinfo = true
			}
			if strings.HasPrefix(input, "search -f") {
				search(keywords, false)
				continue
			}

			if searchTerm == "-" {
				printError("Please provide a valid search term!")
				continue
			}
			search(keywords, moreinfo)
		} else if strings.HasPrefix(input, "show options") {
			u.ShowOptions(searchTerm, pages, resultCount, site, filetype)
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
		} else if strings.HasPrefix(input, "set count") {
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
		} else if strings.HasPrefix(input, "set site") {
			if len(args) < 3 {
				printError("Please provide valid arguments!")
				continue
			}
			if len(args[2]) < 5 {
				printError("Please provide valid arguments!")
				continue
			}
			site = args[2]
			fmt.Printf("site => %v\n", site)
		} else if strings.HasPrefix(input, "set file") {
			if len(args) < 3 {
				printError("Please provide valid arguments!")
				continue
			}
			filetype = args[2]
			fmt.Printf("filetype => %v\n", filetype)
		} else if input == "" || input == " " || input == "  " {
			continue
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
	fmt.Print(string("\033[32;1m"))
}

func printError(e string) {
	colorReset := "\033[0m"
	colorRed := "\033[31;1m"

	fmt.Print(string(colorRed), "[!]")
	fmt.Println(string(colorReset), e)
}

func search(keywords string, moreinfo bool) {
	results, err := s.GoogleSearch(keywords, pages, resultCount)
	if err != nil {
		printError("An unknown error occurred.")
		return
	}
	if len(results) == 0 {
		printError("Could not find any result. Try again with other keywords.")
		return
	}

	var links []string
	color := "\033[34;1m"
	counter := 0

	fmt.Println()
	for _, res := range results {
		if contains(links, res.Url) {
			continue
		}
		links = append(links, res.Url)
		counter++

		if strings.HasPrefix(res.Url, "https") {
			color = "\033[34;1m"
		} else {
			color = "\033[31;1m"
		}

		if moreinfo {
			fmt.Print(string(color), "[#]")
			fmt.Println(string("\033[0m"), counter)
			fmt.Print(string(color), "[@]")
			fmt.Println(string("\033[0m"), res.Url)
			fmt.Print(string(color), "[!]")
			fmt.Println(string("\033[0m"), res.Title)
			fmt.Print(string(color), "[*]")
			fmt.Println(string("\033[0m"), res.Description)
			fmt.Println()
		} else {
			fmt.Print(string(color), "[")
			if counter < 10 {
				fmt.Print("0")
			}
			fmt.Print(counter)
			fmt.Print("] ")
			fmt.Println(string("\033[0m"), res.Url)
		}

		if counter == resultCount {
			fmt.Println()
			return
		}
	}
	if counter == 0 {
		printError("Could not find any result. Try again with other keywords.")
		return
	}
	fmt.Println()
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
