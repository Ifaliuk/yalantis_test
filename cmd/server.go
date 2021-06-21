package main

import (
	"log"
	"os"
	"yalantis/internal"
)

func main() {

	f, err := os.OpenFile("/logs/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	//wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(f)

	srv := internal.NewHttpServer(80)
	srv.Start()


}

