package sitemap

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/Xuanwo/xuanzang/constants"
)

// Sitemap is a structure of sitemap
type Sitemap struct {
	XMLName xml.Name `xml:"urlset"`
	URL     []URL    `xml:"url"`
}

// URL is the base element for sitemap.
type URL struct {
	Loc     string    `xml:"loc"`
	LastMod time.Time `xml:"lastmod"`
}

// LoadSitemap will parse a sitemap.
func LoadSitemap(url string) (*Sitemap, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Error().Msgf("HTTP Request failed for %v.", err)
		return nil, constants.ErrSitemapLoadFailed
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Error().Msgf("HTTP Response is %d, %s.", resp.StatusCode, resp.Status)
		return nil, constants.ErrSitemapLoadFailed
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error().Msgf("HTTP Body read failed for %v.", err)
		return nil, constants.ErrSitemapLoadFailed
	}

	s := &Sitemap{}
	err = xml.Unmarshal(data, s)
	return s, err
}
