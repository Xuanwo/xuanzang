package search

import (
	"encoding/json"
	"net/http"

	"github.com/huichen/wukong/types"
	"github.com/rs/zerolog/log"

	"github.com/Xuanwo/xuanzang/contexts"
	"github.com/Xuanwo/xuanzang/model"
)

// Response is the response for search.
type Response struct {
	Tokens []string   `json:"tokens"`
	Docs   []Document `json:"docs"`

	Total int `json:"total"`
}

// Document is the document that scored.
type Document struct {
	Title       string `json:"title"`
	URL         string `json:"url"`
	ContentText string `json:"content_text"`
}

// Search is the method to do search.
func Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	text := ""
	if v := r.URL.Query().Get("text"); v != "" {
		text = v
	}
	if text == "" {
		log.Info().Msgf("Text is required.")
		w.WriteHeader(400)
		return
	}

	result := contexts.Searcher.Search(types.SearchRequest{
		Text: text,
		RankOptions: &types.RankOptions{
			MaxOutputs: 20,
		},
	})

	resp := &Response{
		Tokens: result.Tokens,
		Total:  result.NumDocs,
	}

	resp.Docs = make([]Document, len(result.Docs))
	for k, v := range result.Docs {
		doc, err := model.GetDocByID(v.DocId)
		if err != nil {
			log.Error().Msgf("Get doc by id failed for %v.", err)
			w.WriteHeader(500)
			return
		}
		if doc == nil {
			continue
		}

		d := Document{
			Title: doc.Title,
			URL:   doc.URL,
		}
		resp.Docs[k] = d
	}

	content, err := json.Marshal(resp)
	if err != nil {
		log.Error().Msgf("JSON marshal failed for %v.", err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	w.Write(content)
	return
}
