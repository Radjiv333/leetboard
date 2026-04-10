package main

import (
	"flag"
	"fmt"
	"io"
	"leetboard/internal/adapters/api"
	"leetboard/internal/adapters/handler"
	"leetboard/internal/adapters/s3"
	"log/slog"
	"net/http"
	"os"
)

var PORT int = 8080
var SERVER string = "0.0.0.0" // Docker container now uses all available interfaces
var USAGEMSG string = "$ ./1337b04rd --help\nhacker board\n\nUsage:\n\t1337b04rd [--port <N>]\n\t1337b04rd --help\n\nOptions:\n\t--help       Show this screen.\n\t--port N     Port number."

func main() {
	// Logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	// Flags
	flagSet := flag.NewFlagSet("flagSet", flag.ContinueOnError) // Using FlagSet to write custom error messages
	flagSet.SetOutput(io.Discard)                               // This code makes it so that flagSet.Parse is not allowed to write its own error messages (at all)
	port := flagSet.Int("port", PORT, "selects port on which the server will listen")
	help := flagSet.Bool("help", false, "shows usage of the program")
	err := flagSet.Parse(os.Args[1:])
	if err != nil {
		slog.Error("Could not parse arguments", "error", err.Error())
		os.Exit(2)
	}
	if *help == true {
		fmt.Println(USAGEMSG)
		return
	}

	// Rick And Morty API
	slog.Info("Loading characters from API...")
	apiLogger := logger.With(slog.String("service", "api"))
	rickAndMortyCharacters, rickAndMortyInfo := api.FetchRickAndMortyCharacters(apiLogger)
	// api.FetchRickAndMortyCharacters()
	fmt.Println(rickAndMortyCharacters)
	fmt.Println(rickAndMortyInfo)

	// S3
	s3Logger := logger.With(slog.String("service", "s3")) // Created s3 logger that will have "service":"s3" even if i dont write it explicitly
	err = s3.UploadImages(rickAndMortyCharacters, s3Logger)
	if err != nil {
		slog.Error(err.Error())
	}

	// Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/catalog", handler.Catalog)
	mux.HandleFunc("/create-post", handler.CreatePost)
	mux.HandleFunc("/archive", handler.Archive)

	// Server
	addr := fmt.Sprintf("%s:%d", SERVER, *port)
	slog.Info("Starting the server...", "port", *port)
	http.ListenAndServe(addr, mux)
}
