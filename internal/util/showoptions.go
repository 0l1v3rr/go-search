package util

import (
	"fmt"
	"strconv"
)

func ShowOptions(searchTerms string, pages int, resCount int, site string) {
	fmt.Println()
	fmt.Println(" NAME         | REQUIRED | COMMAND            | CURRNET VALUE")
	fmt.Println(" -------------|----------|--------------------|--------------")
	fmt.Println(" search terms | yes      | set terms <terms>  | " + searchTerms)
	fmt.Println(" pages        | yes      | set pages <page>   | " + strconv.Itoa(pages))
	fmt.Println(" result count | yes      | set count <count>  | " + strconv.Itoa(resCount))
	fmt.Println(" site         | no       | set site <website> | " + site)
	fmt.Println()
}
