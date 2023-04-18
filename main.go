package main

import (
	"flag"
	"log"
	"net/http"
	"strings"
)

var (
	batchSize int
	outDir    string
)

func main() {
	flag.IntVar(&batchSize, "b", 1, "batch size")
	flag.StringVar(&outDir, "o", ".", "output directory")
	flag.Parse()

	prompt := strings.Join(flag.Args(), " ")
	req, err := makeRequest(prompt, batchSize)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	err = saveResp(resp, outDir)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("done")
}
