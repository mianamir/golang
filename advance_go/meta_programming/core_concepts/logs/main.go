package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/oschwald/geoip2-golang"
)

var db *geoip2.Reader

func init() {
	var err error
	db, err = geoip2.Open("GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
}

func getCurrentUser(r *http.Request) string {
	// Replace this function with your actual implementation to get the current user.
	return "JohnDoe"
}

func getClientLocation(ipAddress string) string {
	ip := net.ParseIP(ipAddress)
	record, err := db.City(ip)
	if err != nil {
		log.Printf("Error fetching location for IP %s: %v", ipAddress, err)
		return "Unknown"
	}
	location := fmt.Sprintf("%s, %s, %s", record.City.Names["en"], record.Subdivisions[0].Names["en"], record.Country.Names["en"])
	return location
}

func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		user := getCurrentUser(r)
		apiURL := r.URL.Path
		clientIP := r.RemoteAddr
		userAgent := r.UserAgent()

		// Parse the user-agent string to extract browser and OS information (use a library for more accurate results)
		browser := "Unknown"
		os := "Unknown"

		// Get client location based on IP address
		location := getClientLocation(clientIP)

		log.Printf("Logged: %s accessed by %s at %s", apiURL, user, startTime.Format("2006-01-02 15:04:05"))
		log.Printf("API URL: %s", apiURL)
		log.Printf("Time: %s", startTime.Format("2006-01-02 15:04:05"))
		log.Printf("IP: %s", clientIP)
		log.Printf("Browser: %s", browser)
		log.Printf("OS: %s", os)
		log.Printf("Location: %s", location)

		next.ServeHTTP(w, r)
	})
}

func homepageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to our website!")
}

func aboutPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Learn more about us here.")
}

func main() {
	r := mux.NewRouter()

	// Applying the logRequest middleware to all routes
	r.Use(logRequest)

	// Define your FastAPI-style routes
	r.HandleFunc("/", homepageHandler)
	r.HandleFunc("/about", aboutPageHandler)

	http.Handle("/", r)

	fmt.Println("Listening on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
