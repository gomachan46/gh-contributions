package main

import (
	"net/http"
	"fmt"
)

func main() {
	resp, err := http.Get("https://github.com/users/gomachan46/contributions")
	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Println(resp)
}
