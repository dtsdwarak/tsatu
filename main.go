package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"

	"tsatu/constants"
	"tsatu/controller"
	"tsatu/model"
	"tsatu/util"

	"mvdan.cc/xurls/v2"
)

func main() {

	var result model.URLSet

	if xmlBytes, err := util.GetHTTPContent(constants.SeenUnseenURL); err != nil {
		fmt.Printf("Failed to get XML: %v", err)
	} else {
		xml.Unmarshal(xmlBytes, &result)
	}

	var urlsToParse []string

	for _, url := range result.URLS {
		if strings.Contains(url.Location, "episodes") {
			urlsToParse = append(urlsToParse, url.Location)
		}
	}

	rxRelaxed := xurls.Relaxed()
	var identifiedURLS []string

	for i := 0; i < len(urlsToParse); i++ {
		fmt.Printf("\n Currently parsing URL - %s", urlsToParse[i])
		if urlContent, err := util.GetHTTPContent(urlsToParse[i]); err != nil {
			fmt.Printf("Failed to get XML: %v", err)
		} else {
			identifiedURLS = append(identifiedURLS, rxRelaxed.FindAllString(string(urlContent), -1)...)
		}
	}

	urlMap := make(map[string]int)

	fmt.Printf("\n Currently parsing URL contents")

	for _, url := range identifiedURLS {

		if strings.Contains(url, "https://www.amazon.in") ||
			strings.Contains(url, "https://amazon.in") ||
			strings.Contains(url, "https://amazon.com") ||
			strings.Contains(url, "https://www.amazon.com") ||
			strings.Contains(url, "https://amzn.eu") ||
			strings.Contains(url, "https://amzn.to") {

			if strings.Contains(url, "https://amzn.eu") || strings.Contains(url, "https://amzn.to") {
				unshortenedURL, err := util.UnshortenURL(url)
				if err != nil {
					fmt.Printf("Unable to shorten url - %s", url)
				}
				url = unshortenedURL
			}

			sanitizedURL, _ := util.SanitizeURL(url)

			if count, ok := urlMap[sanitizedURL]; !ok {
				urlMap[sanitizedURL] = 1
			} else {
				urlMap[sanitizedURL] = count + 1
			}
		}
	}

	var books []model.BookCount
	var authors []model.AuthorCount
	for url, count := range urlMap {

		// This should ideally not be required. But finding this on a huge scale. Hence, cancelling.
		if strings.Compare(url, "https://www.amazon.in/s") == 0 {
			continue
		}

		if strings.Contains(url, "/e/") || strings.Contains(url, "/author/") {
			authors = append(authors, model.AuthorCount{AuthorURL: url, Count: count})
		} else {
			books = append(books, model.BookCount{BookURL: url, Count: count})
		}
	}

	// bookData, err := json.Marshal(books)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Println(string(bookData))

	// authorData, err := json.Marshal(authors)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Println(string(authorData))

	var products []model.Product
	for _, book := range books {
		product, err := controller.FetchAmazonProductDetails(book.BookURL)
		if err != nil {
			fmt.Errorf("Error fetching amazon product details. Error - %w", err)
		} else {
			products = append(products, *product)
		}
	}

	amazonBooks := model.AmazonProducts{Books: products}
	amznbookData, err := json.Marshal(amazonBooks)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(amznbookData))

}
