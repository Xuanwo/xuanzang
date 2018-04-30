package constants

import "errors"

var (
	// ErrCrawlerLoadFailed is returned while crawler request failed.
	ErrCrawlerLoadFailed = errors.New("crawler load failed")
	// ErrSitemapLoadFailed is returned while sitemap load failed.
	ErrSitemapLoadFailed = errors.New("sitemap load failed")
)
