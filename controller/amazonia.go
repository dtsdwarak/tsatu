package controller

import (
	"encoding/json"
	"tsatu/model"
	"tsatu/util"

	"github.com/pkg/errors"
)

func FetchAmazonProductDetails(url string) (*model.Product, error) {

	amazoniaURL := "http://localhost:8080/product?requestURL=" + url

	bytes, err := util.GetHTTPContent(amazoniaURL)
	if err != nil {
		return nil, errors.Wrapf(err, "Error fetching Amazon product details.")
	}

	var product model.Product
	err = json.Unmarshal(bytes, &product)
	if err != nil {
		return nil, errors.Wrapf(err, "Error unmarshalling amazon product details.")
	}

	return &product, nil
}
