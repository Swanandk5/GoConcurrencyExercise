package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"os"
)

func main() {
	data, err := os.ReadFile("urls.json")
	if err != nil {
		panic(err)
	}
	var urls []string
	unmarshalErr := jsoniter.Unmarshal(data, &urls)
	if unmarshalErr != nil {
		fmt.Println(unmarshalErr.Error())
		return
	}
	//FetchImagesUsingChannel(urls)
	FetchImagesUsingWaitGroup(urls)
}
