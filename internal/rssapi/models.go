package rssapi

import "fmt"

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func (r *RSSFeed) PrintDetails() {
	fmt.Printf("- Title: %s\n", r.Channel.Title)
	fmt.Printf("- Description:  %s\n", r.Channel.Description)
	fmt.Printf("- URL: %s\n", r.Channel.Link)
	fmt.Println("- Posts: ")
	for i := range r.Channel.Item {
		fmt.Printf("Post %v\n", i+1)
		r.Channel.Item[i].PrintDetails()
	}
}

func (item *RSSItem) PrintDetails() {
	fmt.Printf("- Title: %s\n", item.Title)
	fmt.Printf("- Pub Date: %s\n", item.PubDate)
	fmt.Printf("- Description: %s\n", item.Description)
}
