package main

import (
	"fmt"

	s "github.com/0l1v3rr/go-search/internal/search"
)

func main() {
	results, err := s.GoogleSearch("test", 1)
	if err != nil {
		return
	}

	for _, res := range results {
		fmt.Print(string("\033[34;1m"), "[")
		if res.Rank < 10 {
			fmt.Print("0")
		}
		fmt.Print(res.Rank)
		fmt.Print("] ")
		fmt.Println(string("\033[0m"), res.Url)
	}
}
