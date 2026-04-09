package main

import (
	"flag"
	"fmt"
	"io"
	"leetboard/internal/adapters/api"
	"leetboard/internal/adapters/handler"
	"net/http"
	"os"
)

var PORT int = 8080
var SERVER string = "0.0.0.0"
var USAGEMSG string = "$ ./1337b04rd --help\nhacker board\n\nUsage:\n\t1337b04rd [--port <N>]\n\t1337b04rd --help\n\nOptions:\n\t--help       Show this screen.\n\t--port N     Port number."

func main() {
	// Flags
	flagSet := flag.NewFlagSet("flagSet", flag.ContinueOnError) // Using FlagSet to write custom error messages
	flagSet.SetOutput(io.Discard)                               // This code makes it so that flagSet.Parse is not allowed to write its own error messages (at all)
	port := flagSet.Int("port", PORT, "selects port on which the server will listen")
	help := flagSet.Bool("help", false, "shows usage of the program")
	err := flagSet.Parse(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	if *help == true {
		fmt.Println(USAGEMSG)
		return
	}

	// Rick And Morty API
	fmt.Println("Loading characters from API...")
	rickAndMortyCharacters, rickAndMortyInfo := api.FetchRickAndMortyCharacters()
	fmt.Println(rickAndMortyCharacters)
	fmt.Println(rickAndMortyInfo)

	
	// Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/catalog", handler.Catalog)
	mux.HandleFunc("/create-post", handler.CreatePost)
	mux.HandleFunc("/archive", handler.Archive)

	// Server
	addr := fmt.Sprintf("%s:%d", SERVER, *port)
	fmt.Printf("Starting the server on port: %d\n", *port)
	http.ListenAndServe(addr, mux)
}
