package main

import (
	"container/list"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func isurlok(uri *url.URL) bool {
	if strings.HasPrefix(uri.Hostname(), "www.zhihu.com"){
		return true
	}
	return false
}

func fetch(target string) ([]string, string, error) {
	uri, err := url.Parse(target)
	if err != nil {
		return nil, "", err
	}
	if !isurlok(uri) {
		return nil, "", errors.New("skip " + target)
	}
	resp, err := http.Get(target)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, "", err
	}
	var urls []string
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, ok := s.Attr("href")
		if !ok {
			return
		}
		if len(link) > 0 && link[0] == '/' {
			u := uri
			u.Path = link
			u.RawQuery = ""
			u.Fragment = ""
			urls = append(urls, u.String())
		} else {
			urls = append(urls, link)
		}
	})
	body, err := doc.Html()
	if err != nil {
		return urls, "", nil
	}
	return urls, body, nil
}

func main() {
	root := os.Args[1]
	visited := make(map[string]bool)
	l := list.New()
	l.PushBack(root)
	for l.Len() != 0 {
		front := l.Front()
		l.Remove(front)
		url_ok := front.Value.(string)
		if visited[url_ok] {
			continue
		}
		visited[url_ok] = true
		urls, body, err := fetch(url_ok)
		if err != nil{
			// log.Print(err)
			continue
		}
		fmt.Printf("%s %0.2fk\n", url_ok, float32(len(body))/1024.0)
		for _, url_ok := range urls{
			l.PushBack(url_ok)
		}
	}

}
