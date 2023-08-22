package online_movie_streaming_platform

type Playable interface {
	Play()
}

type Billable interface {
	Pay(amount float64)
}
