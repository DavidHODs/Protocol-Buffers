package main

import (
	"fmt"
	"log"

	"github.com/DavidHODs/buffers/model"
	"google.golang.org/protobuf/proto"
)

func main() {
	book := &model.Book{
		Id:       1,
		Title:    "The Road to Wigan Pier",
		Author:   &model.Authors{
			Id:   0,
			Name: "Oscar Wilde",
		},
		Category: model.Category_Novel,
	}
	data, err := proto.Marshal(book)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)

	newData := &model.Book{}

	err = proto.Unmarshal(data, newData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(newData.GetAuthor(), newData.GetCategory(), newData.GetTitle())
}

