package main

import (
	"log"

	"github.com/lapeko/k12t_lab/apps/subscriptions/src/api"
)

func main() {
	a := api.New()
	a.Setup()
	log.Fatalf("%e\n", a.Listen())
}
