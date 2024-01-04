package model

import (
	"encoding/xml"
	"fmt"
)

type URLSet struct {
	XMLName        xml.Name `xml:"urlset"`
	XHTML          string   `xml:"xmlns:xhtml,attr"`
	XSI            string   `xml:"xmlns:xsi,attr"`
	SchemaLocation string   `xsi:"xsi:schemaLocation,attr"`
	URLS           []URL    `xml:"url"`
}

func (v *URLSet) String() string {
	if v == nil {
		return "<nil>"
	}

	var urlSetString string

	for _, url := range v.URLS {
		urlSetString = urlSetString + "," + url.String()
	}

	return fmt.Sprintf("URLSet := = %s", urlSetString)
}

type URL struct {
	XMLName  xml.Name `xml:"url"`
	Location string   `xml:"loc"`
	LastMod  string   `xml:"lastmod"`
}

func (v *URL) String() string {
	if v == nil {
		return "<nil>"
	}
	return fmt.Sprintf("URL := Location = %s, LastMod = %s", v.Location, v.LastMod)
}

type Product struct {
	Title       string
	Description string
	Author      string
	Rating      string
	Images      []string
	URL         string
}

type BookCount struct {
	BookURL string `json:"bookURL"`
	Count   int    `json:"count"`
}

type AuthorCount struct {
	AuthorURL string `json:"authorURL"`
	Count     int    `json:"count"`
}
