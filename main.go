package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	batchSize   int
	outDir      string
	kakaoApiKey string
)

func main() {
	flag.IntVar(&batchSize, "b", 1, "batch size")
	flag.StringVar(&outDir, "o", ".", "output directory")
	flag.Parse()

	kakaoApiKey = os.Getenv("KAKAO_REST_API_KEY")
	if kakaoApiKey == "" {
		log.Fatal("env 'KAKAO_REST_API_KEY' is not set")
	}

	prompt := strings.Join(flag.Args(), " ")
	if prompt == "" {
		log.Fatal("prompt is empty")
	}

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
