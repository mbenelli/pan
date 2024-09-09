package main

import (
	"fmt"
	"log"

	"github.com/benellm/pan/jenkins"
)

func main() {

	log.SetPrefix("pan: ")
	log.SetFlags(0)

	jobs := jenkins.JobsList()
	fmt.Println(jobs)
}
