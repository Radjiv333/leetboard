package api

import (
	"encoding/json"
	"fmt"
	"io"
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

func FetchRickAndMortyCharacters() ([]RickAndMortyCharacters, RickAndMortyInfo) {
	url := "https://rickandmortyapi.com/api/character"
	client := &http.Client{Timeout: 10 * time.Second} // Use http.Client instead of http.Get to configure the connection settings (like timeout)
	var allCharacters []RickAndMortyCharacters
	var info RickAndMortyInfo
	for url != "" {
		resp, err := client.Get(url)
		if err != nil {
			fmt.Println("could not get rick_and_morty api:", err)
			os.Exit(2)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("could not read response body:", err)
			os.Exit(2)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Println("unexpected status:", resp.Status)
			fmt.Println("body:", string(body))
			os.Exit(2)
		}

		response := RickAndMortyResponse{}
		err = json.Unmarshal(body, &response)
		if err != nil {
			fmt.Println("could not unmarshal body:", err)
			os.Exit(2)
		}
		fmt.Println(response.Results)
		allCharacters = append(allCharacters, response.Results...)
		if response.Info.NextURL != nil {
			url = *response.Info.NextURL
		} else {
			info = response.Info
			url = ""
			break
		}
		time.Sleep(500 * time.Millisecond) // Introduced delay so that we dont overwhelm the 'ram' api server
	}
	return allCharacters, info
}
