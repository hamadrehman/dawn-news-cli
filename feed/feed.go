package feed

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
)

type Feed struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Items []Item `xml:"item"`
}

type Item struct {
	Title string `xml:"title"`
}

func GetNews() Feed {
	var rss Feed
	resp, err := http.Get("https://www.dawn.com/feeds/home/")
	if err != nil {
		log.Fatal(err)
	}
	xmlBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	feedErr := xml.Unmarshal(xmlBytes, &rss)
	if feedErr != nil {
		log.Fatal(feedErr)
	}
	return rss
}
