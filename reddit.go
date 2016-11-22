package main

import (
    "fmt"

    "github.com/jzelinskie/geddit"
)

// Please don't handle errors this way.
func main() {
    // Login to reddit
    session, _ := geddit.NewLoginSession(
        "novelty_account",
        "password",
        "gedditAgent v1",
    )

    // Set listing options
    subOpts := geddit.ListingOptions{
        Limit: 10,
    }

    // Get reddit's default frontpage
    submissions, _ := session.DefaultFrontpage(geddit.DefaultPopularity, subOpts)

    // Get our own personal frontpage
    submissions, _ = session.Frontpage(geddit.DefaultPopularity, subOpts)

    // Get specific subreddit submissions, sorted by new
    submissions, _ = session.SubredditSubmissions("hockey", geddit.NewSubmissions, subOpts)

    // Print title and author of each submission
    for _, s := range submissions {
        fmt.Printf("Title: %s\nAuthor: %s\n\n", s.Title, s.Author)
    }

    // Upvote the first post
    session.Vote(submissions[0], geddit.UpVote)
}
