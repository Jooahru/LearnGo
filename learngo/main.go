package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Failed to request From server")

func main() {
	results := map[string]string{} // 초기화 해줘야 값을 넣을 수 있음
	// var result = make(map[string]string)표현도 가능

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
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	fmt.Println("checking:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}

	return nil
}
