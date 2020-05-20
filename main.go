package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	res, err := http.Get("https://rt.pornhub.com/video/webmasterss")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatal("Status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(doc.Find("item").Text())
	doc.Find("item").Each(func(i int, s *goquery.Selection) {
		title := s.Find("title").Text()
		img := s.Find("thumb").Text()
		imgLrg := s.Find("thumd_large").Text()
		video := s.Find("embed").Text()

		fmt.Println("title ", title)
		fmt.Println("Image: ", img)
		fmt.Println("Large image: ", imgLrg)
		fmt.Println("Video: ", video)
	})

}

func scraper() {

}
