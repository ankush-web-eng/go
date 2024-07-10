package main

import (
	"fmt"
	"io"
	"net/http"
)

func throwErr(err error) {
	if err != nil {
		panic(err)
	}
}

var url = "https://www.wikipedia.org/"

func main() {
	response, err := http.Get(url)
	throwErr(err)

	defer response.Body.Close()

	result, _ := io.ReadAll(response.Body)
	fmt.Printf("%s", string(result))
}
