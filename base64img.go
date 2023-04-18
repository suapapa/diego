package main

import (
	"encoding/base64"

	"github.com/pkg/errors"
)

func base64ToBytes(base64String string) ([]byte, error) {
	// Decode base64 string to image
	decoded, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return nil, errors.Wrap(err, "fail to decoding base64 string")
	}
	return decoded, err
}

// func imageToBase64(img image.Image) string {
// 	// Encode image to base64 string
// 	var buf bytes.Buffer
// 	err := png.Encode(&buf, img)
// 	if err != nil {
// 		fmt.Println("Error encoding image:", err)
// 	}
// 	return base64.StdEncoding.EncodeToString(buf.Bytes())
// }
