package main

import (
	"fmt"

	"github.com/iceship/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	dictionary["hello"] = "hello"
	fmt.Println(dictionary)
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	errAdd := dictionary.Add("second", "Second word")
	if errAdd != nil {
		fmt.Println(errAdd)
	}

	errUpdate := dictionary.Update("111", "First word update")
	if errUpdate != nil {
		fmt.Println(errUpdate)
	}

	word, _ := dictionary.Search("first")
	fmt.Println(word)

	dictionary.Delete("first")
	word, errSearch := dictionary.Search("first")
	if errSearch != nil {
		fmt.Println(errSearch)
	} else {
		fmt.Println(word)
	}

}
