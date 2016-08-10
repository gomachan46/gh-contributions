package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	resp, err := http.Get("https://github.com/users/gomachan46/contributions")
	if err != nil {
		fmt.Println("http get error", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("body read error", err)
	}
	fmt.Println(string(body))
}
