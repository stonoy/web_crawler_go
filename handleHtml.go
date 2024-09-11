package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse HTML: %v", err)
	}

	// parse html
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, err
	}

	return traverseDoc(doc, baseURL), nil
}

func traverseDoc(n *html.Node, baseUrl *url.URL) []string {

	innerFinal := []string{}
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, tag := range n.Attr {
			if tag.Key == "href" {
				parsedUrl, err := url.Parse(tag.Val)
				if err != nil {
					fmt.Printf("couldn't parse href '%v': %v\n", tag.Val, err)
					continue
				}

				resolvedUrl := baseUrl.ResolveReference(parsedUrl)
				innerFinal = append(innerFinal, resolvedUrl.String())

			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		innerFinal = append(innerFinal, traverseDoc(c, baseUrl)...)
	}

	return innerFinal
}

func getHTML(rawUrl string) (string, error) {
	// make a http get request to any url given
	res, err := http.Get(rawUrl)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	// check response status code and content type
	if res.StatusCode > 399 || !strings.Contains(res.Header["Content-Type"][0], "text/html") {
		return "", fmt.Errorf("response status code : %v and header : %v", res.StatusCode, res.Header["Content-Type"][0])
	}

	// read response body which is a io.ReadCloser
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// return html string
	return string(data), nil
}
