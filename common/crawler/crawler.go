package crawler

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/rs/zerolog/log"

	"github.com/Xuanwo/xuanzang/constants"
)

// LoadContent will load content form current page.
func LoadContent(url string) (title, body string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Error().Msgf("HTTP Request failed for %v.", err)
		return "", "", constants.ErrCrawlerLoadFailed
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Error().Msgf("HTTP Response is %d, %s.", resp.StatusCode, resp.Status)
		return "", "", constants.ErrCrawlerLoadFailed
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Error().Msgf("HTTP Document load failed for %v.", err)
		return "", "", constants.ErrCrawlerLoadFailed
	}

	return doc.Find("title").Text(), doc.Find("body").Text(), nil
}
