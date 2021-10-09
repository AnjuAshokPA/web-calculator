package main

import (
	"log"

	"github.com/AnjuAshokPA/web-calculator/routers"
)

func main() {
	var err error
	err = routers.Routers()
	log.Fatal(err)
}
