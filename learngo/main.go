package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Failed to request From server")

func main() {
	c := make(chan string)

	//results := map[string]string{} // 초기화 해줘야 값을 넣을 수 있음
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
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Print("Checking:",i)
		fmt.Println(<-c)
	}
}

func hitURL(url string, c chan string) {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		c <- "Failed"
	}
	c <- "OK"
}
