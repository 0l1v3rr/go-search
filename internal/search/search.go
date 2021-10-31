package search

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/briandowns/spinner"

	"github.com/PuerkitoBio/goquery"
)

type SearchResult struct {
	Rank        int
	Url         string
	Title       string
	Description string
}

const domain = "https://www.google.com/search?q="
const agent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36"

func buildURL(term string, pages int, resultCount int) []string {
	urls := []string{}
	term = strings.Trim(term, " ")
	term = strings.Trim(term, "\n")
	term = strings.Replace(term, " ", "+", -1)

	for i := 0; i < pages; i++ {
		resUrl := fmt.Sprintf("%s%s&num=%d&hl=%s&start=%d&filter=0", domain, term, 30, "en", i*resultCount)
		urls = append(urls, resUrl)
	}
	return urls
}

func scrapeRequest(page string) (*http.Response, error) {
	base := &http.Client{}
	req, _ := http.NewRequest("GET", page, nil)
	req.Header.Set("User-Agent", agent)

	res, err := base.Do(req)
	if res.StatusCode != 200 {
		err := fmt.Errorf("the status code is not 200")
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return res, nil
}

func parseResult(res *http.Response, counter int) ([]SearchResult, error) {
	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		return nil, err
	}

	results := []SearchResult{}
	sel := doc.Find("div.g")
	counter++
	for i := range sel.Nodes {
		item := sel.Eq(i)
		link, _ := item.Find("a").Attr("href")
		desc := item.Find("span").Text()
		title := item.Find("h3").Text()
		link = strings.Trim(link, " ")

		if link != "" && link != "#" && !strings.HasPrefix(link, "/") {
			result := SearchResult{
				counter,
				link,
				title,
				desc,
			}
			results = append(results, result)
			counter++
		}
	}
	return results, err
}

func GoogleSearch(term string, pages int, resultCount int) ([]SearchResult, error) {
	s := spinner.New(spinner.CharSets[43], 150*time.Millisecond)
	s.Suffix = " Searching..."
	s.Color("green", "bold")
	s.Start()

	results := []SearchResult{}
	resCount := 0
	urls := buildURL(term, pages, resultCount)

	for _, page := range urls {
		res, err := scrapeRequest(page)
		if err != nil {
			s.Stop()
			return nil, err
		}
		data, err := parseResult(res, resCount)
		if err != nil {
			s.Stop()
			return nil, err
		}
		resCount += len(data)
		results = append(results, data...)
		time.Sleep(time.Duration(10) * time.Second)
	}
	s.Stop()

	return results, nil
}
