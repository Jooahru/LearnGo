package main

import (
	"fmt"
	"time"
)

func main() {
	go sexyCount("nico") //go를 붙이면 메인 함수가 실행 되는 동안 goroutines가 활성화된다. main함수는 다른 goroutines을 기다려 주지 않는다
	sexyCount("flynn")
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
