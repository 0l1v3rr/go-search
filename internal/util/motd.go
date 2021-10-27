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

	fmt.Println()
	fmt.Print(string(colorCyan), " [>] GitHub:      ")
	fmt.Println(string(colorReset), "github.com/0l1v3rr/go-search")
	fmt.Print(string(colorGreen), " [>] Version:     ")
	fmt.Println(string(colorReset), "go-search v1.0.1")
	fmt.Print(string(colorLightGreen), " [$] show options:")
	fmt.Println(string(colorReset), "for more information.")
	fmt.Print(string(colorYellow), " [$] search:      ")
	fmt.Println(string(colorReset), "runs the search")
	fmt.Print(string(colorOrange), " [$] exit:        ")
	fmt.Println(string(colorReset), "exits the app")
	fmt.Println()
}

func logo() {
	//Doh
	colorRed := "\033[31;1m"
	fmt.Println("")
	fmt.Println(string(colorRed), "        GGGGGGGGGGGGG   SSSSSSSSSSSSSSS ")
	fmt.Println(string(colorRed), "     GGG::::::::::::G SS:::::::::::::::S")
	fmt.Println(string(colorRed), "   GG:::::::::::::::GS:::::SSSSSS::::::S")
	fmt.Println(string(colorRed), "  G:::::GGGGGGGG::::GS:::::S     SSSSSSS")
	fmt.Println(string(colorRed), " G:::::G       GGGGGGS:::::S            ")
	fmt.Println(string(colorRed), "G:::::G              S:::::S            ")
	fmt.Println(string(colorRed), "G:::::G               S::::SSSS         ")
	fmt.Println(string(colorRed), "G:::::G    GGGGGGGGGG  SS::::::SSSSS    ")
	fmt.Println(string(colorRed), "G:::::G    G::::::::G    SSS::::::::SS  ")
	fmt.Println(string(colorRed), "G:::::G    GGGGG::::G       SSSSSS::::S ")
	fmt.Println(string(colorRed), "G:::::G        G::::G            S:::::S")
	fmt.Println(string(colorRed), " G:::::G       G::::G            S:::::S")
	fmt.Println(string(colorRed), "  G:::::GGGGGGGG::::GSSSSSSS     S:::::S")
	fmt.Println(string(colorRed), "   GG:::::::::::::::GS::::::SSSSSS:::::S")
	fmt.Println(string(colorRed), "     GGG::::::GGG:::GS:::::::::::::::SS ")
	fmt.Println(string(colorRed), "        GGGGGG   GGGG SSSSSSSSSSSSSSS   ")
}
