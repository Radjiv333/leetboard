package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

var PORT int = 8080
var SERVER string = "127.0.0.1"
var USAGEMSG string = "$ ./1337b04rd --help\nhacker board\n\nUsage:\n\t1337b04rd [--port <N>]\n\t1337b04rd --help\n\nOptions:\n\t--help       Show this screen.\n\t--port N     Port number."

func Base(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

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

	// Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Base)

	addr := fmt.Sprintf("%s:%d", SERVER, *port)
	fmt.Printf("Starting the server on port: %d\n", *port)
	http.ListenAndServe(addr, mux)
}
