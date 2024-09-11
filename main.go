package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

type config struct {
	pages              map[string]int
	rawBaseUrl         string
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPage            int
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} else {
		fmt.Println("=====================")
		fmt.Printf("REPORT for %v\n", os.Args[1])
		fmt.Println("=====================\n")
	}

	theUrl := os.Args[1]

	maxConcurrency, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Error parsing string from command line argument")
	}

	maxPageInt, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Error parsing string from command line argument")
	}

	configObj := config{
		pages:              map[string]int{},
		rawBaseUrl:         os.Args[1],
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPage:            maxPageInt,
	}

	configObj.wg.Add(1)

	go configObj.crawlPage(theUrl)

	configObj.wg.Wait()

	fmt.Println("\n\n")

	printReport(configObj.pages)

	fmt.Println("\n\n")
	fmt.Println("Thank You!")

}
