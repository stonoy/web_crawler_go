package main

import (
	"fmt"
	"sort"
)

func (cfg *config) addPagesMap(currentNormalizedUrl string) bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, ok := cfg.pages[currentNormalizedUrl]; !ok {
		cfg.pages[currentNormalizedUrl] = 1
		return false

	} else {
		cfg.pages[currentNormalizedUrl]++
		// cfg.concurrencyControl <- struct{}{}
		return true
	}
}

func (cfg *config) getLengthPagesMap() int {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	return len(cfg.pages)
}

func printReport(pages map[string]int) {
	// create a required struct type
	type report struct {
		url   string
		count int
	}

	// initiate a empty slice of struct
	reports := []report{}

	// loop through the pages map and append the slice
	for k, v := range pages {
		theReport := report{
			url:   k,
			count: v,
		}

		reports = append(reports, theReport)
	}

	// sort the slice what required
	sort.Slice(reports, func(i, j int) bool { return reports[i].url < reports[j].url })
	sort.Slice(reports, func(i, j int) bool { return reports[i].count < reports[j].count })

	// print report
	for _, report := range reports {
		fmt.Printf("Found %v internal links to %v\n", report.count, report.url)
	}
}
