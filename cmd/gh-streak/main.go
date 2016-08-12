package main

import (
	"fmt"
	"net/http"

	"os"
	"strconv"
	"time"

	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get("https://github.com/users/gomachan46/contributions")
	if err != nil {
		fmt.Println("http get error", err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("body read error", err)
	}

	var count, streak int
	var start string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "rect" {
			for _, attr := range n.Attr {
				if start == "" {
					if attr.Key == "data-date" {
						start = attr.Val
					}
				}
				if attr.Key == "data-count" {
					c, err := strconv.Atoi(attr.Val)
					if err != nil {
						fmt.Println("Atoi error", err)
					}
					count += c
					if c == 0 {
						streak = 0
					} else {
						streak++
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	fmt.Fprint(os.Stdout, "start,end,contributions,current-streak\n")
	fmt.Fprintf(os.Stdout, "%s,%s,%d,%d\n", start, time.Now().Format("2006-01-02"), count, streak)
}
