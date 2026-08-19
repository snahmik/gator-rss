package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/snahmik/gator-rss/internal/types"
)

func commandAgg(args []string, state *cliState) error {
	//if len(args) != 1 {
	//	//return fmt.Errorf("usage: command_agg <command>")
	//	return fmt.Errorf("missing argument")
	//}
	//
	//feedURL := args[0]

	feedURL := "https://www.wagslane.dev/index.xml"

	feed,err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return err
	}

	fmt.Println(feed)

	return nil
}

func fetchFeed(ctx context.Context, feedURL string) (*types.RSSFeed, error) {
	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Accept", "application/rss+xml")
	req.Header.Set("User-Agent", "gator")

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error fetching feed: %w", err)
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var data types.RSSFeed
	if err := xml.Unmarshal(resBody, &data); err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %w", err)
	}

	return &data, nil
}
