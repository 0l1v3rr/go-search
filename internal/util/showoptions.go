package util

import (
	"fmt"
	"strconv"
)

func ShowOptions(searchTerms string, pages int) {
	fmt.Println()
	fmt.Println(" NAME         | REQUIRED | COMMAND           | CURRNET VALUE")
	fmt.Println(" -------------|----------|-------------------|--------------")
	fmt.Println(" search terms | yes      | set terms <terms> | " + searchTerms)
	fmt.Println(" pages        | yes      | set pages <page>  | " + strconv.Itoa(pages))
	fmt.Println()
}
