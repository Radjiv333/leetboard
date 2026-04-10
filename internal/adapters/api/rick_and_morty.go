package api

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type RickAndMortyResponse struct {
	Info    RickAndMortyInfo         `json:"info"`
	Results []RickAndMortyCharacters `json:"results"`
}

type RickAndMortyCharacters struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image"`
}

type RickAndMortyInfo struct {
	Count   int     `json:"count"`
	Pages   int     `json:"pages"`
	NextURL *string `json:"next"`
}

func FetchRickAndMortyCharacters(logger *slog.Logger) ([]RickAndMortyCharacters, RickAndMortyInfo) {
	url := "https://rickandmortyapi.com/api/character"
	client := &http.Client{Timeout: 10 * time.Second} // Use http.Client instead of http.Get to configure the connection settings (like timeout)
	var allCharacters []RickAndMortyCharacters
	response := RickAndMortyResponse{}
	for url != "" {
		resp, err := client.Get(url)
		logger.Debug("response status code check", "status code", resp.Status)
		if err != nil {
			logger.Error("Could not get rick_and_morty api", "error", err.Error())
			os.Exit(2)
		}
		if resp.StatusCode != http.StatusOK {
			logger.Error("Could not get rick_and_morty api. Invalid status code", "status code", resp.Status)
			os.Exit(2)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.Error("Could not read response body", "error", err.Error())
			os.Exit(2)
		}
		defer resp.Body.Close()

		err = json.Unmarshal(body, &response)
		if err != nil {
			logger.Error("Could not unmarshal body", "error", err.Error())
			os.Exit(2)
		}

		allCharacters = append(allCharacters, response.Results...)
		if *response.Info.NextURL != "https://rickandmortyapi.com/api/character?page=2" { // Restraining to loading only 1 page. Otherwise should be != nil
			url = *response.Info.NextURL
		} else {
			url = ""
			break
		}
		// time.Sleep(500 * time.Millisecond) // Introduced delay so that we dont overwhelm the 'ram' api server
	}
	return allCharacters, response.Info
}
