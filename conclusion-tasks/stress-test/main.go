package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Report struct {
	TotalRequests      int
	SuccessfulRequests int
	StatusCodeCount    map[int]int
	TotalTime          time.Duration
}

func main() {
	url := flag.String("url", "", "Service URL")
	requests := flag.Int("requests", 0, "Total request quantity")
	concurrency := flag.Int("concurrency", 0, "Total of simultanious request quantity")
	flag.Parse()

	if *url == "" {
		fmt.Println("URL is required")
		return
	}

	reports := executeLoadTest(*url, *requests, *concurrency)
	printReport(reports)
}

func executeLoadTest(url string, totalRequests, concurrency int) Report {
	var wg sync.WaitGroup
	report := Report{
		TotalRequests:   totalRequests,
		StatusCodeCount: make(map[int]int),
	}

	start := time.Now()
	requestsChan := make(chan int, totalRequests)
	results := make(chan map[int]int, concurrency)

	makeConcurrency(&wg, concurrency, requestsChan, url, results)
	getTotalRequests(totalRequests, requestsChan)

	wg.Wait()
	close(results)

	report.TotalTime = time.Since(start)

	statusCodeCount, successfullRequests := getStatusResults(results)
	report.StatusCodeCount = statusCodeCount
	report.SuccessfulRequests = successfullRequests

	return report
}

func makeConcurrency(wg *sync.WaitGroup, concurrency int, requestsChan chan int, url string, results chan map[int]int) {
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			statusCount := make(map[int]int)

			for range requestsChan {
				statusCode := makeRequest(url)
				statusCount[statusCode]++
			}

			results <- statusCount
		}()
	}
}

func getTotalRequests(totalRequests int, requestsChan chan int) {
	for i := 0; i < totalRequests; i++ {
		requestsChan <- i
	}
	close(requestsChan)
}

func getStatusResults(results chan map[int]int) (map[int]int, int) {
	statusCodeCount := make(map[int]int)
	var successfullRequests int

	for result := range results {
		for status, count := range result {
			statusCodeCount[status] += count

			if status == 200 {
				successfullRequests += count
			}

		}
	}

	return statusCodeCount, successfullRequests
}

func makeRequest(url string) int {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("err", err)
		return 0
	}

	defer res.Body.Close()
	return res.StatusCode
}

func printReport(report Report) {
	fmt.Printf("\nRelatório de Teste de Carga:\n")
	fmt.Printf("Tempo total gasto: %v\n", report.TotalTime)
	fmt.Printf("Total de requests: %d\n", report.TotalRequests)
	fmt.Printf("Requests com status HTTP 200: %d\n", report.SuccessfulRequests)
	fmt.Println("Distribuição dos códigos de status:")

	for statusCode, count := range report.StatusCodeCount {
		fmt.Printf("Status %d: %d\n", statusCode, count)
	}
}
