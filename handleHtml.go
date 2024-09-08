package main

import (
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	final := []string{}
	// parse html
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, err
	}

	final = traverseDoc(doc, rawBaseURL)

	return final, nil
}

func traverseDoc(n *html.Node, baseUrl string) []string {
	innerFinal := []string{}
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, tag := range n.Attr {
			if tag.Key == "href" {
				if strings.HasPrefix(tag.Val, "https://") || strings.HasPrefix(tag.Val, "http://") {
					innerFinal = append(innerFinal, tag.Val)
				} else {
					fullUrl := baseUrl + tag.Val
					innerFinal = append(innerFinal, fullUrl)
				}

			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		innerFinal = append(innerFinal, traverseDoc(c, baseUrl)...)
	}

	return innerFinal
}
