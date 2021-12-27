package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Failed to request From server")

type resultRequest struct {
	url    string
	status string
}

func main() {
	results := map[string]string{} // 초기화 해줘야 값을 넣을 수 있음
	c := make(chan resultRequest)

	// var resultRequest = make(map[string]string)표현도 가능

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++ {
		// fmt.Print("Checking:", i)
		// fmt.Println(<-c)
		result := <-c
		results[result.url] = result.status
	}
	fmt.Println(results)
	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan<- resultRequest) { // chan <-(direction) 표현하면 Send Only

	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- resultRequest{url: url, status: status}
}
