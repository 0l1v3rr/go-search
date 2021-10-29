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
	fmt.Println("")
	printLogoLine("         GGGGGGGGGGGGG    SSSSSSSSSSSSSSS ")
	printLogoLine("      GGG::::::::::::G  SS:::::::::::::::S")
	printLogoLine("    GG:::::::::::::::G S:::::SSSSSS::::::S")
	printLogoLine("   G:::::GGGGGGGG::::G S:::::S     SSSSSSS")
	printLogoLine("  G:::::G       GGGGGG S:::::S            ")
	printLogoLine(" G:::::G               S:::::S            ")
	printLogoLine(" G:::::G                S::::SSSS         ")
	printLogoLine(" G:::::G    GGGGGGGGGG   SS::::::SSSSS    ")
	printLogoLine(" G:::::G    G::::::::G     SSS::::::::SS  ")
	printLogoLine(" G:::::G    GGGGG::::G        SSSSSS::::S ")
	printLogoLine(" G:::::G        G::::G             S:::::S")
	printLogoLine("  G:::::G       G::::G             S:::::S")
	printLogoLine("   G:::::GGGGGGGG::::G SSSSSSS     S:::::S")
	printLogoLine("    GG:::::::::::::::G S::::::SSSSSS:::::S")
	printLogoLine("      GGG::::::GGG:::G S:::::::::::::::SS ")
	printLogoLine("         GGGGGG   GGGG  SSSSSSSSSSSSSSS   ")
}

func printLogoLine(line string) {
	colorRed := "\033[31;1m"
	colorReset := "\033[0m"
	bold := "\033[1m"

	for _, char := range line {
		if char == ':' {
			fmt.Print(bold)
			fmt.Print(string(colorReset), string(char))
		} else {
			fmt.Print(bold)
			fmt.Print(string(colorRed), string(char))
		}
	}
	fmt.Println()
}
