package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func lenAndUpper2(name string) (lenght int, uppercase string) {
	// return 값을 지정해주지 않아도 변수 생성을 앞에서 해주면 알아서 리턴된다 (naked return)
	defer fmt.Println("I'm done") // defer 함수가 실행된 다음에 실행된다
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) { // 변수 여러개 받는 방법
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers { // 첫번째는 인덱스이다
		total += number
	}
	// for i :=0;i<len(numbers);i++{
	// 	fmt.Println(numbers[i])
	// }
	return total
}

func canIDrink(age int) bool {
	// if koreanAge := age + 2; koreanAge <= 18 { // if문에서 ;앞에 변수 생성가능
	// 	return false
	// }
	// if문안에 return 있으면 else 없어도 된다.

	switch { // if문처럼 바로 variable생성가능
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false

	}
	return true
}

//struct 만들기
type person struct {
	name string
	age int
	favFood []string
}

func main() {

	//struct
	favFood := []string{"kimchi","ramen"}
	nico := person{name:"nico",age:18,favFood:favFood} //person{"nico", 18 ,favFood} 둘다가능
	fmt.Println(nico)

	// Map
	// nico := map[string]string{"name": "nico", "age": "12"}
	// for _,value := range nico{
	// 	fmt.Println(value)
	// }
	// fmt.Println(nico)

	//array
	// names := []string{"nico", "lynn", "pp"} // array 만들때 [5] 숫자입력시 5개만 입력가능 숫자입력 안할시 크기유동적
	// // names[3] = "lala" //array 크기 알 때는 이렇게 추가 가능
	// // slice array사용할 때 추가시 append 사용
	// names = append(names,"flynn")
	// fmt.Println(names)

	// a := 2
	// b := &a
	// *b = 20 //b를 이용해서 a의 값을 바꿀 수 있다. 현재 b는 a의 주소값을 가지고 있다.
	// a = 10
	// fmt.Println(a,b)
	// fmt.Println(&a,&b) //& 붙이면 변수 a,b 주소
	// fmt.Println(a, *b) //*붙이면 주소의 값 알 수 있다.

	// fmt.Println(canIDrink(16))

	// result:=superAdd(1,2,3,4,5,6,7)
	// fmt.Println(result)

	// tt,xx := lenAndUpper2("will")
	// fmt.Println(tt,xx)

	// repeatMe("will","how","moterh")

	// totalLenght, upperName := lenAndUpper("nico")
	// //totalLenght, _ := lenAndUpper("nico") _두개 값중 하나만 받고 싶을 때 언더스코어 작성
	// fmt.Println(totalLenght, upperName)

	// fmt.Println(multiply(2,2))

	// var name string = "nico" //name := "nico" 축약코드 사용시 Type표기 안해줘두된다.
	// name = "will"
	// fmt.Println(name)
}
