package benchmark

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// Set it to true to print more verbose output
var VerboseOutput bool = false

var wg sync.WaitGroup
var totalDuration int64

func visitWebsite(url string) {

	defer wg.Done()

	before := time.Now()
	outputMessage := "Visiting " + url

	resp, err := http.Get(url)
	if err != nil {
		return
	}

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	after := time.Now()

	duration := after.Sub(before).Milliseconds()
	totalDuration += duration
	outputMessage = outputMessage + fmt.Sprint(" took ", duration, "ms")
	if VerboseOutput {
		fmt.Println(outputMessage)
	}
}

func Start(url string, repetitions int) (avgDuration int64) {

	totalDuration = 0

	wg.Add(repetitions)

	for i := 0; i < repetitions; i++ {
		go visitWebsite(url)
	}

	wg.Wait()

	avgDuration = totalDuration / int64(repetitions)

	return
}
