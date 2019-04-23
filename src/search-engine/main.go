package main

import (
	"log"
	"os"

	_ "search-engine/matchers"
	"search-engine/search"
)

// init is called prior to main.
func init() {
	// Change the device for loggin to stdout
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program
func main() {
	// Perform the search for the specified term
	search.Run("president")
}
