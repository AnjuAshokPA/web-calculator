package main

import (
	"log"
	"webcalculator/routers"
)

func main() {
	var err error
	err = routers.Routers()
	log.Fatal(err)
}
