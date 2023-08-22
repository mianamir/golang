package online_movie_streaming_platform

type User struct {
	ID      int
	Name    string
	Email   string
	Subtype string // Can be "basic" or "premium"
	// ...
}

type Subscription interface {
	GetSubscriptionType() string
	CalculateMonthlyFee() float64
}
