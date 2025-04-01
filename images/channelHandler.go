package images

import (
	"fmt"
	"io"
	"net/http"
)

func FetchImagesUsingChannel(urls []string) {
	imageChannel := make(chan int, 5)
	for i := 0; i < len(urls); {
		select {
		case imageChannel <- i:
			go fetchImageForChannel(urls[i], imageChannel)
			i++
		}
	}
}

func fetchImageForChannel(url string, imageChannel chan int) {
	//fetch image
	fmt.Println("starting goroutine")
	imageData, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(imageData.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(body))
	fmt.Println("ending goroutine")
	<-imageChannel
}
