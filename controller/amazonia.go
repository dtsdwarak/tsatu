package controller

import (
	"fmt"
	"tsatu/util"

	"github.com/pkg/errors"
)

func FetchAmazonProductDetails(url string) error {

	amazoniaURL := "http://localhost:8080/product?requestURL=" + url

	bytes, err := util.GetHTTPContent(amazoniaURL)
	if err != nil {
		return errors.Wrapf(err, "Error fetching Amazon product details.")
	}

	fmt.Println("String =", string(bytes))

	return nil
}
