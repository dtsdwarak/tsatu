package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"tsatu/constants"
	"tsatu/model"
	"tsatu/util"
)

func main() {

	var result model.URLSet

	if xmlBytes, err := util.GetXML(constants.SeenUnseenURL); err != nil {
		log.Printf("Failed to get XML: %v", err)
	} else {
		xml.Unmarshal(xmlBytes, &result)
	}

	fmt.Println(result.String())

}
