package main

import (
	"encoding/json"
	"fmt"
	"io"
	"leetboard/internal/adapters/api"
	"net/http"
	"os"
	"time"
)

var PORT int = 8080
var SERVER string = "127.0.0.1"
var USAGEMSG string = "$ ./1337b04rd --help\nhacker board\n\nUsage:\n\t1337b04rd [--port <N>]\n\t1337b04rd --help\n\nOptions:\n\t--help       Show this screen.\n\t--port N     Port number."

func main() {
	// Flags
	// flagSet := flag.NewFlagSet("flagSet", flag.ContinueOnError) // Using FlagSet to write custom error messages
	// flagSet.SetOutput(io.Discard)                               // This code makes it so that flagSet.Parse is not allowed to write its own error messages (at all)
	// port := flagSet.Int("port", PORT, "selects port on which the server will listen")
	// help := flagSet.Bool("help", false, "shows usage of the program")
	// err := flagSet.Parse(os.Args[1:])
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(2)
	// }
	// if *help == true {
	// 	fmt.Println(USAGEMSG)
	// 	return
	// }

	url := "https://rickandmortyapi.com/api/character"
	client := &http.Client{Timeout: 10 * time.Second}
	var allCharacters []api.RickAndMortyCharacters
	var info api.RickAndMortyInfo
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
			fmt.Printf("unexpected status: %s\n", resp.Status)
			fmt.Println("body:", string(body))
			os.Exit(2)
		}

		response := api.RickAndMortyResponse{}
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
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println(info)
	fmt.Println(allCharacters)

	// var characters []api.RickAndMortyCharacters = make([]api.RickAndMortyCharacters, 0)
	// err = json.Unmarshal(body, &characters)
	// if err != nil {
	// 	fmt.Println("could not unmarshal body:", err)
	// 	os.Exit(2)
	// }
	// fmt.Println(characters[0])

	// Mux
	// mux := http.NewServeMux()
	// mux.HandleFunc("/catalog", handler.Catalog)
	// mux.HandleFunc("/create-post", handler.CreatePost)
	// mux.HandleFunc("/archive", handler.Archive)

	// // Server
	// addr := fmt.Sprintf("%s:%d", SERVER, *port)
	// fmt.Printf("Starting the server on port: %d\n", *port)
	// http.ListenAndServe(addr, mux)
}
