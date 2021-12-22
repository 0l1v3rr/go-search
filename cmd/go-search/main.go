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
	showHttp    bool   = true
	intitle     string = "-"
	inurl       string = "-"
	intext      string = "-"
	related     string = "-"
	save        bool   = false
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
			if intitle != "-" {
				keywords += fmt.Sprintf(" intitle:%s", intitle)
			}
			if inurl != "-" {
				keywords += fmt.Sprintf(" inurl:%s", inurl)
			}
			if intext != "-" {
				keywords += fmt.Sprintf(" intext:%s", intext)
			}
			if related != "-" {
				keywords += fmt.Sprintf(" related:%s", related)
			}

			if strings.HasPrefix(input, "search -ef") || strings.HasPrefix(input, "search -fe") {
				search(keywords, true)
				continue
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
			u.ShowOptions(searchTerm, pages, resultCount, site, filetype, showHttp, intitle, inurl, intext, related, save)
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
		} else if strings.HasPrefix(input, "set intitle") {
			if len(args) < 3 {
				printError("Please provide valid arguments!")
				continue
			}
			intitle = args[2]
			fmt.Printf("intitle => %v\n", intitle)
		} else if strings.HasPrefix(input, "set inurl") {
			if len(args) < 3 {
				printError("Please provide valid arguments!")
				continue
			}
			inurl = args[2]
			fmt.Printf("inurl => %v\n", inurl)
		} else if strings.HasPrefix(input, "set intext") {
			if len(args) < 3 {
				printError("Please provide valid arguments!")
				continue
			}
			intext = args[2]
			fmt.Printf("intext => %v\n", intext)
		} else if strings.HasPrefix(input, "set related") {
			if len(args) < 3 {
				printError("Please provide valid arguments!")
				continue
			}
			related = args[2]
			fmt.Printf("related => %v\n", related)
		} else if strings.HasPrefix(input, "set http") {
			if len(args) < 3 {
				printError("Please provide valid arguments!")
				continue
			}
			if args[2] != "true" && args[2] != "false" {
				printError("Please provide valid arguments!")
				continue
			}
			if args[2] == "true" {
				showHttp = true
			} else {
				showHttp = false
			}
			fmt.Printf("show http => %v\n", args[2])
		} else if strings.HasPrefix(input, "set save") {
			if len(args) < 3 {
				printError("Please provide valid arguments!")
				continue
			}
			if args[2] != "true" && args[2] != "false" {
				printError("Please provide valid arguments!")
				continue
			}
			if args[2] == "true" {
				save = true
			} else {
				save = false
			}
			fmt.Printf("save to txt => %v\n", args[2])
		} else if strings.HasPrefix(input, "clear") {
			u.Clear()
			continue
		} else if strings.HasPrefix(input, "reset") {
			reset()
		} else if input == "" || input == " " || input == "  " {
			continue
		} else {
			printError("Unknown command.")
		}

	}
}

func reset() {
	searchTerm = "-"
	pages = 1
	resultCount = 20
	site = "-"
	filetype = "-"
	showHttp = true
	intitle = "-"
	inurl = "-"
	intext = "-"
	related = "-"
	save = false
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

func printSuccess(s string) {
	colorGreen := "\033[0m"
	colorRed := "\033[32;1m"

	fmt.Print(string(colorRed), "[^]")
	fmt.Println(string(colorGreen), s)
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

	if save {
		f, err := os.Create("urls.txt")
		if err != nil {
			printError("An unknown error occurred.")
			return
		}
		defer f.Close()
		for _, res := range results {
			if save {
				f.WriteString(fmt.Sprintf("%v. %s\n", res.Rank, res.Url))
			}
		}
	}

	var links []string
	color := "\033[34;1m"
	counter := 0

	fmt.Println()
	for _, res := range results {
		if !showHttp {
			if strings.HasPrefix(res.Url, "http://") {
				continue
			}
		}
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

		if save && counter == resultCount {
			printSuccess("The results have been saved to the urls.txt file!")
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
