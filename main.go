package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var link = "https://rt.pornhub.com/video/webmasterss"

func main() {
	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatal("Status code error: %d %s", res.StatusCode, res.Status)
	}

	parse(res.Body)

}

func parse(b io.Reader) {
	doc, err := goquery.NewDocumentFromReader(b)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(doc.Find("item").Text())
	doc.Find("item").Each(func(i int, s *goquery.Selection) {
		title := s.Find("title").Text()
		img := s.Find("thumb").Text()
		imgLrg := s.Find("thumb_large").Text()
		videoLink := s.Find("iframe").Text()

		title = cropTitle(title)

		// video := Video{
		// 	title:      title,
		// 	image:      img,
		// 	largeImage: imgLrg,
		// 	videoLink:  videoLink,
		// }

		// Videos := append(Videos, video)

		logOutput(title, img, imgLrg, videoLink)
	})
}

func cropTitle(title string) string {
	extraExclamation := strings.Index(title, "!")
	if extraExclamation > -1 {
		title = title[extraExclamation+1:]
	}
	title = strings.ReplaceAll(title, "[CDATA[", "")
	title = strings.ReplaceAll(title, "]]>", "")
	return title
}
