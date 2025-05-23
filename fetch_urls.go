package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchUrls(urls[] string) map[string]string {
	response := make(chan struct {
		url string
		body string
	}, 3)

	for _,url := range urls {
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				response <- struct { url string
				body string } {url, fmt.Sprint("error: %v /n ", err)}
				return
			}else{
				defer resp.Body.Close()
				bodyBytes, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					response <- struct {
						url  string
						body string
					}{url, fmt.Sprintf("read error: %v /n", err)}
					return
				}else{
					response <- struct { url string
					body string} {url, string(bodyBytes[:100])}
				}
			}
		}(url)
	}


	results := make(map[string]string)

	for i := 0; i < len(urls) ; i++ {
		res := <-response
		results[res.url] = res.body
	}

	return results

}


func main() {
	urlsToFetch := []string{
	    "https://www.google.com",
	    "https://www.example.com",
	    "https://nonexistent.domain.xyz", // This one should result in an error
	}
	results := FetchUrls(urlsToFetch)
	for url, response := range results {
		fmt.Printf("url : %s \n and the response was : %s \n",url,response)
	}
}