package util

import (
	"encoding/base64"
)

func DeCodeBase64ToS(in string) string {
	out := DeCodeBase64(in)
	return string(out)
}

func DeCodeBase64(in string) []byte {
	out := make([]byte, base64.StdEncoding.DecodedLen(len(in)))
	n, err := base64.StdEncoding.Decode(out, []byte(in))
	if err != nil {
		return nil
	}
	return out[0:n]
}
