package main

import (
	"alexdenkk/words/app"
	"log"
	"flag"
)

var addr string

func main() {
	parseFlags()

	wordsApp := app.New(addr)
	log.Println("app initialized")

	log.Fatal(wordsApp.Run())
}

func parseFlags() {
	flag.StringVar(&addr, "address", ":8080", "address and(or) port for app")
	flag.Parse()
}
