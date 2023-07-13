package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
)

const baseWeatherAPI = "http://api.openweathermap.org/data/2.5/weather?"
const openWeatherAPIKey = "{API-KEY}"

func main() {
	r := chi.NewRouter()
	r.Get("/current-weather", CurrentWeather)

	portToListenOn := "3000"
	// Start the server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", portToListenOn),
		Handler: r,
	}

	fmt.Printf("Server listening on port: %s", portToListenOn)
	err := srv.ListenAndServe()
	if err != nil {
		panic("Error while starting server: " + err.Error())
	}
}

// CurrentWeather function to get the current weather of a specific city
func CurrentWeather(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		http.Error(w, "City is required", http.StatusBadRequest)
		return
	}

	weatherAPI := fmt.Sprintf("%sq=%s&appid=%s", baseWeatherAPI, city, openWeatherAPIKey)
	resp, err := http.Get(weatherAPI)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		http.Error(w, "Error fetching weather", http.StatusBadGateway)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
