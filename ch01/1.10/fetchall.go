package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	start := time.Now()
	ch := make(chan string)
	for i, url := range os.Args[1:] {
		// go fetch(url, ch)
		go fetch(url, ch, "out."+strconv.Itoa(i))
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	wg.Wait()
}

func fetch(url string, ch chan<- string, ofile string) {
	start := time.Now()
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	// nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	fmt.Println("StatusCode: ", resp.StatusCode)
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		// ch <- fmt.Sprintf("while reading %s: %v", url, err)
		ch <- fmt.Sprintf("fetch: reading %s: %v\n", url, err)
		return
	}

	wg.Add(1)
	go func() {
		file, err := os.Create(ofile)
		if err != nil {
			ch <- fmt.Sprint(err)
			return
		}
		file.Write(b)
		file.Close()
		wg.Done()
	}()

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, len(b), url)
}
