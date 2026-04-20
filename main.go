package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func read(v *http.Response) []byte {
	defer v.Body.Close()

	b, err := io.ReadAll(v.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	return b
}

func worker(wg *sync.WaitGroup, ch <-chan string) {

	defer wg.Done()

	for url := range ch {

		v, err := http.Get(url)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		fmt.Printf("url: %+v | byte: %+v \n", url, len(read(v)))

	}

}

func main() {

	sUrl := []string{
		"https://google.com",
		"https://youtube.com",
		"https://github.com",
		"https://Facebook.com",
		"https://Instagram.com",
		"https://chatgpt.com",
	}

	wg := sync.WaitGroup{}

	chUrl := make(chan string)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, v := range sUrl {
			chUrl <- v
		}

		close(chUrl)

	}()

	wg.Add(1)
	go worker(&wg, chUrl)
	wg.Add(1)
	go worker(&wg, chUrl)
	wg.Add(1)
	go worker(&wg, chUrl)

	wg.Wait()

}
