package crawler

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
	"time"
)

func check(t html.Token) (ok bool, q bool, a bool) {
	for _, div := range t.Attr {
		if div.Key == "class" {
			if div.Val == "space-8" {
				q = true
				ok = true
			} else if div.Val == "text-copy space-8" {
				a = true
				ok = true
			}
		}
	}
	return
}

func crawl(url string, ch chan string, end chan bool) {
	var ok, q, a bool = false, false, false
	var p []string
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	fmt.Println("[HTTP Request] " + url)
	client := http.Client{Transport: transport}
	resp, err := client.Get(url)
	defer func() {
		end <- true
	}()
	if err != nil {
		return
	}
	b := resp.Body
	defer b.Close()
	z := html.NewTokenizer(b)
	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			return
		case tt == html.StartTagToken:
			t := z.Token()
			if !q && !a {
				isDiv := t.Data == "div"
				if !isDiv {
					continue
				}
				ok, q, a = check(t)
				if !ok {
					continue
				}
			}
		case tt == html.EndTagToken:
			t := z.Token()
			if a {
				isDiv := t.Data == "div"
				if !isDiv {
					continue
				}
				fmt.Println("[HTTP Response] " + url)
				ch <- strings.Join(p, " ")
				a = false
				return
			}
		case tt == html.TextToken:
			t := z.Token()
			s := strings.TrimSpace(t.Data)
			if ok && len(s) > 0 {
				if a {
					p = append(p, s)
				} else {
					ch <- s
					q = false
				}
			}
		}
	}
}

// Get gets all the questions and answers of the given urls
func Get(urls []string, wait int) (qa map[string]string) {
	var q string
	qa = make(map[string]string)
	info := make(chan string)
	end := make(chan bool)
	d := time.Duration(wait)
	for _, url := range urls {
		go crawl(url, info, end)
		time.Sleep(d * time.Second)
	}
	for i, u := 0, 0; u < len(urls); {
		select {
		case data := <-info:
			if i == 1 {
				qa[q] = data
				i = 0
			} else {
				q = data
				i++
			}
		case <-end:
			u++
		}
	}
	return
}
