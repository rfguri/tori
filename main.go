package main

import (
	"fmt"
	"github.com/assistai/tori/crawler"
	"os"
	"strconv"
)

type articles struct {
	begin int
	end   int
}

func usage() {
	fmt.Println("[Tori] Usage: tori waitTime range")
	fmt.Println("[Tori] Example: tori 0 1 10")
	fmt.Println("[Tori] Example: tori 1 20 35")
	fmt.Println("[Tori] Example: tori 5 1 1000")
}

func main() {
	wait := 0
	argv := os.Args[1:]
	argc := len(os.Args[1:])
	if argc != 3 {
		usage()
		return
	}
	wait, _ = strconv.Atoi(argv[0])
	begin, _ := strconv.Atoi(argv[1])
	end, _ := strconv.Atoi(argv[2])
	a := articles{begin, end}
	urls := []string{}
	for i := a.begin; i <= a.end; i++ {
		url := fmt.Sprintf("https://www.airbnb.com/help/article/%d", i)
		urls = append(urls, url)
	}
	qa := crawler.Get(urls, wait)
	fmt.Println("\n[Tori] Found", len(qa), "unique questions:")
	for q, a := range qa {
		fmt.Println("Q: " + q)
		fmt.Println("A: " + a)
		fmt.Println()
	}
}
