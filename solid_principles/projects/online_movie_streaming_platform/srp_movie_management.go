package online_movie_streaming_platform

type Movie struct {
	Title       string
	Description string
	Category    string
	Price       float64
}

type Category struct {
	Name        string
	Description string
}
