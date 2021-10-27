package util

import "fmt"

func Motd() {
	logo()

	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	colorLightGreen := "\033[32;1m"
	colorCyan := "\033[36m"
	colorYellow := "\033[33;1m"
	colorOrange := "\033[33m"

	fmt.Print(string(colorCyan), "[>] GitHub:      ")
	fmt.Println(string(colorReset), "github.com/0l1v3rr/go-search")
	fmt.Print(string(colorGreen), "[>] Version:     ")
	fmt.Println(string(colorReset), "go-search v1.0.1")
	fmt.Print(string(colorLightGreen), "[$] show options:")
	fmt.Println(string(colorReset), "for more information.")
	fmt.Print(string(colorYellow), "[$] exit:        ")
	fmt.Println(string(colorReset), "exits the app")
	fmt.Print(string(colorOrange), "[$] search:      ")
	fmt.Println(string(colorReset), "runs the search")
}

func logo() {

}
