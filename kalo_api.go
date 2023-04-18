package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"path"

	"github.com/pkg/errors"
)

type KaloRequest struct {
	Prompt struct {
		Text      string `json:"text"`
		BatchSize int    `json:"batch_size"`
	} `json:"prompt"`
}

type KaloResponse struct {
	ID           string `json:"id"`
	ModelVersion string `json:"model_version"`
	Images       []struct {
		ID    string `json:"id"`
		Image string `json:"image"`
		Nsfw  bool   `json:"nsfw"`
	} `json:"images"`
}

/*
POST /v1/inference/karlo/t2i
Host: api.kakaobrain.com
Authorization: KakaoAK ${REST_API_KEY}
Content-Type: application/json
*/

func makeRequest(prompt string, batchSize int) (*http.Request, error) {
	var kr KaloRequest
	kr.Prompt.Text = prompt
	kr.Prompt.BatchSize = batchSize

	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(&kr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to encode request body")
	}

	req, err := http.NewRequest("POST", "https://api.kakaobrain.com/v1/inference/karlo/t2i", buf)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request")
	}

	apiKey := os.Getenv("KAKAO_REST_API_KEY")
	if apiKey == "" {
		return nil, errors.New("KAKAO_REST_API_KEY is not set")
	}
	req.Header.Add("Authorization", "KakaoAK "+apiKey)
	req.Header.Add("Content-Type", "application/json")

	return req, nil
}

func saveResp(resp *http.Response, outDir string) error {
	var kr KaloResponse
	err := json.NewDecoder(resp.Body).Decode(&kr)
	if err != nil {
		return errors.Wrap(err, "failed to decode response body")
	}

	for _, b64Img := range kr.Images {
		imgBytes, err := base64ToBytes(b64Img.Image)
		if err != nil {
			return errors.Wrap(err, "failed to decode base64 image")
		}
		w, err := os.Create(path.Join(outDir, b64Img.ID+".webp"))
		if err != nil {
			return errors.Wrap(err, "failed to create file")
		}
		defer w.Close()
		_, err = w.Write(imgBytes)
		if err != nil {
			return errors.Wrap(err, "failed to write file")
		}
	}

	return nil
}
