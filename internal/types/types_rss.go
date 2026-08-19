package types

import (
	"fmt"
	"html"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Items       []RSSItem `xml:"item"`
	} `xml:"channel"`
}

func (rf *RSSFeed) String() string {
	return fmt.Sprintf("RSS Feed Title: %s\nRSS Feed Link: %s\nRSS Feed Description: %s\nItems:\n%v", html.UnescapeString(rf.Channel.Title), rf.Channel.Link, html.UnescapeString(rf.Channel.Description), rf.Channel.Items)
}

type RSSItem struct {
	Title         string `xml:"title"`
	Link          string `xml:"link"`
	Description   string `xml:"description"`
	PublishedDate string `xml:"pubDate"`
}

func (ri RSSItem) String() string {
	return fmt.Sprintf("=====\nTitle: %s\nLink: %s\n Description: %s\nPublished Date: %s\n", html.UnescapeString(ri.Title), ri.Link, html.UnescapeString(ri.Description), ri.PublishedDate)
}
