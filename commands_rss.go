package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/snahmik/gator-rss/internal/database"
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

	feed, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("db error: %w", err)
	}

	fmt.Println(feed)

	return nil
}

func commandAddFeed(args []string, state *cliState) error {
	feedName := args[0]
	feedURL := args[1]

	params := database.AddFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedURL,
		Name_2:    state.config.CurrentUserName,
	}
	feed, err := state.db.AddFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("db error: %w", err)
	}

	err = commandFollow([]string{feedURL}, state)
	if err != nil {
		return err
	}

	fmt.Println(feed)
	return nil
}

func commandFeeds(args []string, state *cliState) error {
	feeds, err := state.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("db error: %w", err)
	}

	fmt.Println(feeds)
	return nil
}

func commandFollow(args []string, state *cliState) error {
	url := args[0]

	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      state.config.CurrentUserName,
		Url:       url,
	}

	followedFeed, err := state.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("db error: %w", err)
	}

	fmt.Println(followedFeed)
	return nil
}

func commandFollowing(args []string, state *cliState) error {
	followedFeeds, err := state.db.GetFeedFollowsForUser(context.Background(), state.config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("db error: %w", err)
	}

	fmt.Println(followedFeeds)
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
