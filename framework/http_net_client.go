package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	httpC := &http.Client{}
	res, err := httpC.Get("https://www.yousuu.com/bookstore/?channel")
	if err != nil {
		fmt.Errorf("QueryAlarmConfig http request error-(%s)", err.Error())
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Errorf("QueryAlarmConfig read error-(%s)", err.Error())
	}
	defer res.Body.Close()

	fmt.Println(string(body))
}
