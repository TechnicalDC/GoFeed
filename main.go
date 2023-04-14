package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

func main() {
	url := "https://lukesmith.xyz/rss"
	fmt.Println(url)
	var body xml.Decoder
	name := xml.Name { }

	item := xml.StartElement{
		Name `xml:"item`,
	}

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	} else {
		body.DecodeElement(resp.Body, &item)
		fmt.Println(body)
	}
	defer resp.Body.Close()

}
