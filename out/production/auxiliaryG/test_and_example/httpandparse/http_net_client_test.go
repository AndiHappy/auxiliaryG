package httpandparse

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestHttpGet(t *testing.T) {
	httpC := &http.Client{}
	res, err := httpC.Get("https://www.yousuu.com/bookstore/?channel")
	if err != nil {
		fmt.Printf("QueryAlarmConfig http request error-(%s)", err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("QueryAlarmConfig read error-(%s)", err)
	}
	defer res.Body.Close()

	fmt.Println(string(body))
}
