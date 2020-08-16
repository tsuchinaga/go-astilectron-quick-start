package main

import (
	"log"
	"os"

	"github.com/asticode/go-astilectron"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	// Initialize astilectron
	a, err := astilectron.New(log.New(os.Stderr, "", 0), astilectron.Options{
		AppName:           "go-astilectron-quick-start",
		BaseDirectoryPath: "example",
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer a.Close()

	// // Handle signals
	// a.HandleSignals()

	// Start astilectron
	if err = a.Start(); err != nil {
		log.Fatalln(err)
	}

	center := true
	height := 700
	width := 700
	// Create a new window
	var w *astilectron.Window
	if w, err = a.NewWindow("index.html", &astilectron.WindowOptions{
		Center: &center,
		Height: &height,
		Width:  &width,
	}); err != nil {
		log.Fatalln(err)
	}

	// Create windows
	if err = w.Create(); err != nil {
		log.Fatalln(err)
	}

	// Blocking pattern
	a.Wait()
}
