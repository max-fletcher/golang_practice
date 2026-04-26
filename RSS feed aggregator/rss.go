package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

// types for the keys in the RSS feeds that we are going to scrape
type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"` // defines a slice of items. So this is an array of a slice within a slice
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func urlToFeed(url string) (RSSFeed, error) {
	httpClient := http.Client{ // creates an http client for making API requests
		Timeout: 10 * time.Second,
	}

	res, err := httpClient.Get(url)
	if err != nil {
		return RSSFeed{}, err
	}

	defer res.Body.Close() // closes the http client on getting a response

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RSSFeed{}, err
	}

	rssFeed := RSSFeed{} // An empty rss feed struct that we will use to store the response xml data after unmarshalling

	xml.Unmarshal(data, &rssFeed) // store unmarshalled data in rssFeed

	return rssFeed, nil
}
