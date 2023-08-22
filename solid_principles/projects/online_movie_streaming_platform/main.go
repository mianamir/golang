package online_movie_streaming_platform

import "fmt"

func main() {
	movie := Movie{
		Title:       "Inception",
		Description: "Mind-bending thriller",
		Category:    "Science Fiction",
		Price:       9.99,
	}

	basicUser := User{
		ID:      1,
		Name:    "Alice",
		Email:   "alice@example.com",
		Subtype: "basic",
	}

	premiumUser := User{
		ID:      2,
		Name:    "Bob",
		Email:   "bob@example.com",
		Subtype: "premium",
	}

	// ... Initialize user subscriptions

	// ... Process billing

	fmt.Println("Movie:", movie)
	fmt.Println("Basic User:", basicUser)
	fmt.Println("Premium User:", premiumUser)
}
