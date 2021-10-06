package main

import (
	"encoding/json"
	"fmt"
)

const json_example = `{"title":"the hitchhiker's guide to the galaxy","author":"douglas adams","published":"1980","isbn":"1400052920"}`

// book struct represents book data
// Fields must start with capitol letter to unmarshal json.
// Note: use json tag if stuct field is different from json field name
type book struct {
	Title  string
	Author string
	Date   int `json:"published,string,omitempty"`
	Isbn10 int `json:"isbn,string,omitempty"`
}

func main() {
	fmt.Println("Json example string: ", json_example)

	// init book struct
	var bk book

	// unmarshal onto a pointer bk
	err := json.Unmarshal([]byte(json_example), &bk)
	if err != nil {
		fmt.Println(err)
		return
	}

	// print struct key and values
	fmt.Println("Book struct post unmarhsal: ", bk)
	fmt.Println("Book title: ", bk.Title)
	fmt.Println("Book author: ", bk.Author)
	fmt.Println("Book publish date: ", bk.Date)
	fmt.Println("Book ISBN Num: ", bk.Isbn10)
}
