package decoder

//go:generate bunsan-mockgen -gofile=$GOFILE

import "encoding/base64"

type HumanDecoder interface {
	DecodeToText(data []byte) (string, error)
	DecodeBase64ToText(data string) (string, error)
}

func decodeBase64ToText(decoder HumanDecoder, data string) (string, error) {
	binData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	return decoder.DecodeToText(binData)
}
