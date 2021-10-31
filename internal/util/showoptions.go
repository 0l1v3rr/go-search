package util

import (
	"fmt"
	"strconv"
)

func ShowOptions(searchTerms string, pages int, resCount int, site string, filetype string, showHttp bool, intitle string, inurl string, intext string) {
	http := "false"
	if showHttp {
		http = "true"
	}
	fmt.Println()
	fmt.Println(" NAME         | REQUIRED | COMMAND            | VALUE")
	fmt.Println(" -------------|----------|--------------------|------")
	fmt.Println(" search terms | yes      | set terms <terms>  | " + searchTerms)
	fmt.Println(" pages        | yes      | set pages <page>   | " + strconv.Itoa(pages))
	fmt.Println(" result count | yes      | set count <count>  | " + strconv.Itoa(resCount))
	fmt.Println(" show http    | no       | set http <false>   | " + http)
	fmt.Println(" site         | no       | set site <website> | " + site)
	fmt.Println(" filetype     | no       | set file <type>    | " + filetype)
	fmt.Println(" intitle      | no       | set intitle <word> | " + intitle)
	fmt.Println(" inurl        | no       | set inurl <word>   | " + inurl)
	fmt.Println(" intext       | no       | set intext <word>  | " + intext)
	fmt.Println()
}
