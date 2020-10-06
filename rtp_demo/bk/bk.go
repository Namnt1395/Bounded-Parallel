package bk

import (
	"fmt"
	"net/http"
	"sort"
	"time"
)


type result struct {
	index int
	url   string
	res   *http.Response
	err   error
}


func boundedParallelGet(urls []string) []result {
	resultsChan := make(chan *result)
	defer func() {
		close(resultsChan)
	}()
	for i, url := range urls {
		go func(i int, url string) {
			client := http.Client{
				Timeout: 120 * time.Millisecond,
			}
			res, err := client.Get(url)
			resultsChan <- &result{i, url, res, err}
		}(i, url)
	}
	var resultsSuccess []result
	var result []int
L:
	for {
		select {
		case data := <-resultsChan:
			result = append(result, data.index)
			if data.err == nil {
				resultsSuccess = append(resultsSuccess, *data)
			}
			if len(result) == len(urls) {
				break L
			}
		case <-time.After(time.Millisecond * 120):
			fmt.Println("time.Millisecond * 120", time.Millisecond*120)
			break L
		}
	}
	// let's sort these results real quick
	sort.Slice(resultsSuccess, func(i, j int) bool {
		return resultsSuccess[i].index < resultsSuccess[j].index
	})
	// now we're done we return the results
	return resultsSuccess
}

var urls []string

func init() {
	urls = append(urls, "http://localhost:9083/api-delay-100")
	urls = append(urls, "http://localhost:9083/api-delay-50")
}
func main() {
	results := boundedParallelGet(urls)
	if len(results) > 0 {
		fmt.Println("result...", results)
	} else {
		fmt.Println("pass pack")
	}

}

