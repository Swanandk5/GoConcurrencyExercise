package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func FetchImagesUsingWaitGroup(urls []string) {
	for i := 0; i < len(urls); {
		var slice = make([]string, 4)
		for j := 0; j < 4; j++ {
			slice[j] = urls[i]
			i++
		}
		//this will be started 5 times since the slice size is 4
		startSubtaskForImages(slice)
	}
}

// starts a new goroutine for each slice passed
func startSubtaskForImages(urls []string) {
	var wg sync.WaitGroup
	for i := 0; i < len(urls); i++ {
		wg.Add(1)
		go fetchImageForWaitGroup(urls[i], &wg)
		wg.Wait()
	}
}

func fetchImageForWaitGroup(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	imageData, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(imageData.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(body))

}
