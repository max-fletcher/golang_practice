package main

import (
	"context"
	"database/sql"
	"dummy/internal/database"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

// a scraper that starts when our go code is executed. This runs every "timeBetweenRequest" so if "timeBetweenRequest = 1 min"
// it will run every 1 minute
func startScraping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Scraping on %v goroutines every %s duration", concurrency, timeBetweenRequest)

	// creating a ticker object. A ticker is performs an operation at fixed intervals
	ticker := time.NewTicker(timeBetweenRequest)
	// ticker.C is a channel. If the ticker is set for 1 min, a value will be sent across the channel to here.
	// One of the benefits of doing "for ; ; <-ticker.C {" is that it will fire immdiately once and then.
	// If we used "for range <-ticker.C {" we would have to wait 1 min for the first execution and would continue per interval(1 minute in this case).
	for ; ; <-ticker.C {
		// this is using a background context as opposed to a scoped context. Background context is globally available to use while scoped context
		// is for one-off operations like http requests.
		// casting concurrency to int32 since sqlc uses it by default for limit params
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Println("Error fetching feeds", feeds)
			continue // using continue here since this function should always be running. It uses an infinite loop(see above).
		}

		// *IMPORTANT: We are storing a pointer/reference to a waitgroup inside wg. The reason is because we will be passing this to some goroutines.
		// Passing a value instead of a pointer/reference to a waitgroup will cause a copy of the waitgroup to be passed, which in turn will
		// cause deadlock/hang because using wg.Done from inside the function/goroutine to which it was passed will not decrement the waitgroup, rather
		// it will decrement the count from the copy of the waitgroup, leading to the deadlock/hang.
		wg := &sync.WaitGroup{}
		// For all the feeds, we are spawning a goroutine and each will handle processing the feeds independently.
		for _, feed := range feeds {
			wg.Add(1) // we add a 1 to the waitgroup everytime we spawn a new goroutine

			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()

	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error marking feed as fetched", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("Error fetching feed", err)
		return
	}

	// looping through the rssFeeds and inserting then into the database.
	// #TODO: Try using batch inserts later
	for _, item := range rssFeed.Channel.Item {
		// log.Println("Found post", item.Title, " on feed ", feed.Name)

		// Creating a sql null string object(from sql package). Otherwise, go 's type system doesn't comply with nullable values.
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}
		pubAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("Error parsing published at date %v : %v", item.PubDate, err)
		}

		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			FeedID:      feed.ID,
			Title:       item.Title,
			Description: description,
			PublishedAt: pubAt,
			Url:         item.Link,
		})
		if err != nil {
			// "strings.Contains" - if the string contains a substring. Returns a boolean.
			//  If error contains the substring "duplicate key", then continue the loop instead of logging it.
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Printf("Error inserting posts into database: %v", err)
		}
	}
	log.Printf("Feed %v collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
