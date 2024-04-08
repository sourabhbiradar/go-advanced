package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func first(url string) {

	defer wg.Done()

	fmt.Println("url :", url)

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Step2: ", url)
	defer response.Body.Close()

	fmt.Println("Step3: ", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Step4: ", len(body))
}

func main() {
	wg.Add(2)
	go first("https://www.google.com")
	go first("https://www.facebook.com")
	//go first("https://www.go.dev")
	wg.Wait()
}
