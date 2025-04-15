package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/go-ping/ping"
)

func main() {
	requestsCount := flag.Int("requests", 0, "Total number of requests to be made")
	url := flag.String("url", "", "URL to be pinged")
	concurrency := flag.Int("concurrency", 0, "Number of concurrent requests")

	flag.Parse()

	if *requestsCount <= 0 || *url == "" || *concurrency <= 0 {
		fmt.Println("Error: All of the parameters are required.")
		fmt.Println("Example: -requests=10 -url=example.com -concurrency=5")
		os.Exit(1)
	}

	var wg sync.WaitGroup
	sem := make(chan struct{}, *concurrency)

	var successCount, failureCount int
	var mu sync.Mutex
	startTime := time.Now()

	for i := range *requestsCount {
		wg.Add(1)
		sem <- struct{}{}

		go func(requestID int) {
			defer wg.Done()
			defer func() { <-sem }()

			pinger, err := ping.NewPinger(*url)
			if err != nil {
				mu.Lock()
				failureCount++
				mu.Unlock()
				return
			}

			pinger.Count = *requestsCount
			pinger.Interval = time.Second
			pinger.Timeout = time.Second * 2

			err = pinger.Run()
			mu.Lock()
			if err != nil {
				failureCount++
			} else {
				successCount++
			}
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	close(sem)

	totalTime := time.Since(startTime)

	fmt.Println("\nFinal Report:")
	fmt.Printf("URL: %s\n", *url)
	totalTimeMinutes := float64(totalTime) / float64(time.Minute)
	if totalTimeMinutes < 1 {
		totalTimeSeconds := float64(totalTime) / float64(time.Second)
		fmt.Printf("Total spent time on the execution: %.2f seconds\n", totalTimeSeconds)
	} else {
		fmt.Printf("Total spent time on the execution: %.2f minutes\n", totalTimeMinutes)
	}
	fmt.Printf("Total of requests made: %d\n", *requestsCount)
	fmt.Printf("Total of requests with success: %d\n", successCount)
	fmt.Printf("Total of requests with failure: %d\n", failureCount)
}
