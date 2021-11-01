package util

import "fmt"

func Motd() {
	logo()

	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	colorLightGreen := "\033[32;1m"
	colorCyan := "\033[36m"
	colorLightCyan := "\033[36;1m"
	colorYellow := "\033[33;1m"
	colorOrange := "\033[33m"

	fmt.Println()
	fmt.Print(string(colorLightCyan), " [>] GitHub:      ")
	fmt.Println(string(colorReset), "github.com/0l1v3rr/go-search")
	fmt.Print(string(colorCyan), " [>] Version:     ")
	fmt.Println(string(colorReset), "go-search v1.1.0")
	fmt.Print(string(colorGreen), " [@] README.md:   ")
	fmt.Println(string(colorReset), "to know more about the app.")
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
	fmt.Println(string(colorRed), "         GGGGGGGGGGGGG    SSSSSSSSSSSSSSS ")
	fmt.Println(string(colorRed), "      GGG::::::::::::G  SS:::::::::::::::S")
	fmt.Println(string(colorRed), "    GG:::::::::::::::G S:::::SSSSSS::::::S")
	fmt.Println(string(colorRed), "   G:::::GGGGGGGG::::G S:::::S     SSSSSSS")
	fmt.Println(string(colorRed), "  G:::::G       GGGGGG S:::::S            ")
	fmt.Println(string(colorRed), " G:::::G               S:::::S            ")
	fmt.Println(string(colorRed), " G:::::G                S::::SSSS         ")
	fmt.Println(string(colorRed), " G:::::G    GGGGGGGGGG   SS::::::SSSSS    ")
	fmt.Println(string(colorRed), " G:::::G    G::::::::G     SSS::::::::SS  ")
	fmt.Println(string(colorRed), " G:::::G    GGGGG::::G        SSSSSS::::S ")
	fmt.Println(string(colorRed), " G:::::G        G::::G             S:::::S")
	fmt.Println(string(colorRed), "  G:::::G       G::::G             S:::::S")
	fmt.Println(string(colorRed), "   G:::::GGGGGGGG::::G SSSSSSS     S:::::S")
	fmt.Println(string(colorRed), "    GG:::::::::::::::G S::::::SSSSSS:::::S")
	fmt.Println(string(colorRed), "      GGG::::::GGG:::G S:::::::::::::::SS ")
	fmt.Println(string(colorRed), "         GGGGGG   GGGG  SSSSSSSSSSSSSSS   ")
}
