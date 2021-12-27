package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string) // 채널만드는법 c는 변수명
	people := [5]string{"nico", "flynn", "will", "doge", "rsr"}
	for _, person := range people {
		go isSexy(person, c)
	}

	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}

	// // result := <-c
	// // fmt.Println(result)
	// resultOne := <-c //<-c 받아오는 갯수보다 많으면 무한대기
	// fmt.Println("Waiting for messages")
	// fmt.Println("Received this message:", resultOne) //<-c 채널에서 메세지를 받아오는 표현
	// fmt.Println("Received this message:", <-c)
}

func isSexy(person string, c chan string) { // c chan[채널임을알려주고] string[타입을 지정해줘야한다]
	time.Sleep(time.Second * 5)
	c <- person + "is Sexy"
}
