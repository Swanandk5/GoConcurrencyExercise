package api

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io"
	"net/http"
	"sync"
)

var userAgentURL = "https://httpbin.org/user-agent"
var uuidURL = "https://httpbin.org/uuid"
var IPURL = "https://httpbin.org/ip"

func FetchData() {
	wg := sync.WaitGroup{}
	result := Result{}
	wg.Add(3)
	go getData(&wg, userAgentURL, func(userAgent *UserAgent) {
		result.UserAgent = userAgent.UserAgent
	})
	go getData(&wg, IPURL, func(ip *IP) {
		result.Ip = ip.IP
	})
	go getData(&wg, uuidURL, func(uuid *UUID) {
		result.UUID = uuid.UUID
	})
	wg.Wait()
	jsonString, err := jsoniter.MarshalIndent(result, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonString))
}

func getData[T any](wg *sync.WaitGroup, url string, assign func(response *T)) {
	defer wg.Done()
	ipResponse, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	jsonData, err := io.ReadAll(ipResponse.Body)
	if err != nil {
		panic(err)
	}
	var data T
	err = jsoniter.Unmarshal(jsonData, &data)
	if err != nil {
		panic(err)
	}
	assign(&data)
}
