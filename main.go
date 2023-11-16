package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"strings"

	"tsatu/constants"
	"tsatu/model"
	"tsatu/util"

	"mvdan.cc/xurls/v2"
)

func main() {

	var result model.URLSet

	fmt.Println("Parse Started")

	if xmlBytes, err := util.GetHTTPContent(constants.SeenUnseenURL); err != nil {
		log.Printf("Failed to get XML: %v", err)
	} else {
		xml.Unmarshal(xmlBytes, &result)
	}

	var urlsToParse []string

	for _, url := range result.URLS {
		if strings.Contains(url.Location, "episodes") {
			urlsToParse = append(urlsToParse, url.Location)
			// fmt.Println(url.Location)
		}
	}

	rxRelaxed := xurls.Relaxed()
	var identifiedURLS []string

	for i := 0; i < len(urlsToParse); i++ {
		if urlContent, err := util.GetHTTPContent(urlsToParse[i]); err != nil {
			log.Printf("Failed to get XML: %v", err)
		} else {
			identifiedURLS = append(identifiedURLS, rxRelaxed.FindAllString(string(urlContent), -1)...)
		}
	}

	urlMap := make(map[string]int)

	for _, url := range identifiedURLS {
		if strings.Contains(url, "https://www.amazon.in") || strings.Contains(url, "https://amazon.in") || strings.Contains(url, "https://amazon.com") || strings.Contains(url, "https://www.amazon.com") {

			sanitizedURL, _ := util.SanitizeURL(url)

			if count, ok := urlMap[sanitizedURL]; !ok {
				urlMap[sanitizedURL] = 1
			} else {
				urlMap[sanitizedURL] = count + 1
			}

		}

	}

	for url, count := range urlMap {
		fmt.Println(url, count)
	}
}
