package formatters

import (
	"dummy/internal/database"
	"time"

	"github.com/google/uuid"
)

// NOTE: Exporting type fields
// Turns out any type that has fields that are not exported as pascal-case(or at least starts with a capital case) is
// not exported. So if you used id instead of ID, the returned struct(and consequently the JSON) will be missing that field.
type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	APIKey    string    `json:"api_key"`
}

// NOTE: Exporting functions from a folder(a.k.a a package as different folders are each considered a package in go)
// Turns out any exported function from a folder/package needs to be pascal-case(or at least starts with a capital case)
// else, the function is not made available when imported and used. This is why the convention for writing golang often
// is to use pascal-case.

// This is a function that converts struct keys from pascal-case to use camel-case. This is so we get camel-cased JSON.
// Remember that database.go has the type for the user being fetched from database and that is being passed here.
func DatabaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		APIKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// This is a function that converts struct keys from pascal-case to use camel-case. This is so we get camel-cased JSON.
// Remember that database.go has the type for the user being fetched from database and that is being passed here.
func DatabaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
	}
}

func DatabaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}

	// remember that when looping through an array of slices using range, the 1st val is index and the 2nd value are the individual slices
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, Feed{
			ID:        dbFeed.ID,
			Name:      dbFeed.Name,
			Url:       dbFeed.Url,
			UserID:    dbFeed.UserID,
			CreatedAt: dbFeed.CreatedAt,
			UpdatedAt: dbFeed.UpdatedAt,
		})

		// This also works if you want to reuse functionality i.e replace above block with the line below
		// feeds = append(feeds, DatabaseFeedToFeed(dbFeed))
	}

	return feeds
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	FeedID    uuid.UUID `json:"feed_id"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func DatabaseFeedFollowToFeedFollow(dbFeed database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeed.ID,
		FeedID:    dbFeed.FeedID,
		UserID:    dbFeed.UserID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
	}
}

func DatabaseFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {
	feedFollows := []FeedFollow{}

	for _, dbFeedFollow := range dbFeedFollows {
		feedFollows = append(feedFollows, FeedFollow{
			ID:        dbFeedFollow.ID,
			FeedID:    dbFeedFollow.FeedID,
			UserID:    dbFeedFollow.UserID,
			CreatedAt: dbFeedFollow.CreatedAt,
			UpdatedAt: dbFeedFollow.UpdatedAt,
		})

		// This also works if you want to reuse functionality i.e replace above block with the line below
		// feeds = append(feeds, DatabaseFeedFollowToFeedFollow(dbFeed))
	}

	return feedFollows
}
