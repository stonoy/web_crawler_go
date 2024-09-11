package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentUrl string) {
	// pass empty value in channel to occupy its space
	cfg.concurrencyControl <- struct{}{}

	defer func() {
		cfg.wg.Done()
		// receive those values from channel to allow next go routines to proceed
		<-cfg.concurrencyControl
	}()

	// stop when max number of page reach
	if cfg.getLengthPagesMap() >= cfg.maxPage {
		return
	}

	// parse raw urls to url struct
	baseUrl, err := url.Parse(cfg.rawBaseUrl)
	if err != nil {
		return
	}

	currentUrl, err := url.Parse(rawCurrentUrl)
	if err != nil {
		return
	}

	// check domain is same , if not -> return pages
	if baseUrl.Host != currentUrl.Host {
		return
	}

	// get the normalised version of current url
	normalisedUrl, err := normalizeUrl(rawCurrentUrl)
	if err != nil {
		return
	}

	// add or increment the normalised current url count in pages map
	isUrlAlreadyPresent := cfg.addPagesMap(normalisedUrl)
	if isUrlAlreadyPresent {
		return
	}

	// show which page we are going to call now...
	fmt.Printf("Calling %v ...\n", rawCurrentUrl)

	// get html from current url
	htmlBody, err := getHTML(rawCurrentUrl)
	if err != nil {
		return
	}

	// get urls from the html
	urls, err := getURLsFromHTML(htmlBody, rawCurrentUrl)
	if err != nil {
		return
	}

	// recursively call each url
	for _, url := range urls {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)

	}

}
