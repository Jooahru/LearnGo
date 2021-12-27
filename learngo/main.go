package main

import (
	"fmt"

	"github.com/will/learngo/mydict"
)

func main() {
	// Delete Method 확인
	dictionary := mydict.Dictionary{}
	word := "hello"
	dictionary.Add(word, "First")
	fmt.Println(dictionary)
	err := dictionary.Delete("word")
	fmt.Println(err)
	fmt.Println(dictionary)
	dictionary.Delete("word")
	fmt.Println(err)
	fmt.Println(dictionary)

	// Update Method 확인
	// dictionary := mydict.Dictionary{}
	// word := "hello"
	// dictionary.Add(word, "First")
	// err := dictionary.Update("dsgf", "Second")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// word2, _ := dictionary.Search(word)
	// fmt.Println(word2)

	// Add Method 확인
	// dictionary := mydict.Dictionary{}
	// word := "hello"
	// definition := "Greeting"
	// err := dictionary.Add(word, definition)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// hello, _ := dictionary.Search(word)
	// fmt.Println(hello)
	// err2 := dictionary.Add(word, definition)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }

	// Search Method 확인
	// dictionary := mydict.Dictionary{"first": "First word"}
	// definition, err := dictionary.Search("first")
	// //definition, err := dictionary.Search("second")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }

	// dictionary :=mydict.Dictionary{"first":"First word"}
	// fmt.Println(dictionary["first"])

	// dictionary := mydict.Dictionary{}
	// dictionary["hello"] ="hello"
	// fmt.Println(dictionary)
}
