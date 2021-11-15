package main

import (
	"io"
	"log"
	"os"
	"time"

	"nhlApiproject/nhlApi"
)

func main() {
	//For benchmarking the request time
	now := time.Now()
	bolbFile, err := os.OpenFile("bolbs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("error opening the file bolbs.txt: %v", err)
	}

	defer bolbFile.Close()

	wrt := io.MultiWriter(os.Stdout, bolbFile)

	log.SetOutput(wrt)

	teams, err := nhlApi.GetAllTeams()

	if err != nil {
		log.Fatalf("error while getting tall teams: %v", err)
	}

	for _, team := range teams {
		log.Println("-----------------------")
		log.Printf("Name: %s", team.Name)
		log.Println("-----------------------")
	}

	log.Printf("took %v", time.Now().Sub(now).String())
}
