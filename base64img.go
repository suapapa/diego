package main

import (
	"bytes"
	"encoding/base64"
	"image"

	"github.com/pkg/errors"
)

func base64ToImage(base64String string) (image.Image, error) {
	// Decode base64 string to image
	decoded, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return nil, errors.Wrap(err, "fail to decoding base64 string")
	}
	// Decode image from bytes
	img, _, err := image.Decode(bytes.NewReader(decoded))
	if err != nil {
		return nil, errors.Wrap(err, "fail to decoding image")
	}
	return img, nil
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
