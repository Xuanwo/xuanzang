package index

import (
	"time"

	"github.com/huichen/wukong/types"
	"github.com/rs/zerolog/log"

	"github.com/Xuanwo/xuanzang/common/crawler"
	"github.com/Xuanwo/xuanzang/common/sitemap"
	"github.com/Xuanwo/xuanzang/contexts"
	"github.com/Xuanwo/xuanzang/model"
)

// Index will do the index job.
func Index() {
	sm, err := sitemap.LoadSitemap(contexts.Source.URL)
	if err != nil {
		log.Error().Msgf("Load sitemap failed.")
		return
	}

	for _, u := range sm.URL {
		doc, err := model.GetDoc(u.Loc)
		if err != nil {
			log.Error().Msgf("Get model url failed for %v.", err)
			return
		}
		if doc != nil && doc.UpdatedAt >= u.LastMod.Unix() {
			continue
		}

		if doc == nil {
			doc, err = model.CreateDoc(u.Loc, 0)
			if err != nil {
				log.Error().Msgf("URL create failed for %v.", err)
				return
			}
		}

		go indexDocument(doc)
	}

	// TODO: we need to remove not exist documents here.

	contexts.Searcher.FlushIndex()
}

// indexDocument will index the every document.
func indexDocument(doc *model.Doc) {
	title, body, err := crawler.LoadContent(doc.URL, contexts.Source.TitleTag)
	if err != nil {
		log.Error().Msgf("Index document failed for %v.", err)
		return
	}

	contexts.Searcher.IndexDocument(doc.ID, types.DocumentIndexData{
		Content: body,
	}, false)

	doc.UpdatedAt = time.Now().Unix()
	doc.Title = title
	err = doc.Save()
	if err != nil {
		log.Error().Msgf("URL save failed for %v.", err)
		return
	}

	log.Info().Msgf("URL %s indexed.", doc.URL)
}
