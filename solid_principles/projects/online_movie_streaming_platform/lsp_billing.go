package online_movie_streaming_platform

type Biller interface {
	GenerateBill(user User) string
}

type BasicSubscription struct {
	User User
}

type PremiumSubscription struct {
	User User
}
