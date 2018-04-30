package contexts

import (
	"github.com/huichen/wukong/engine"
	"github.com/huichen/wukong/types"
	"github.com/rs/zerolog/log"

	"github.com/Xuanwo/xuanzang/common/db"
	"github.com/Xuanwo/xuanzang/config"
)

// Global variable for xuanzang.
var (
	DB       *db.Database
	Searcher *engine.Engine
	Source   *config.Source
)

// SetupContexts will setup contexts for xuanzang.
func SetupContexts(c *config.Config) (err error) {
	Source = c.Source

	err = c.Logger.Prepare()
	if err != nil {
		log.Fatal().Msgf("Logger init failed for %v.", err)
	}

	DB, err = db.NewDB(&db.DatabaseOptions{Address: *c.DBPath})
	if err != nil {
		log.Fatal().Msgf("Bolt databasae open failed for %v.", err)
	}
	err = DB.Init()
	if err != nil {
		log.Fatal().Msgf("Bolt init failed for %v.", err)
	}

	Searcher = &engine.Engine{}
	Searcher.Init(types.EngineInitOptions{
		SegmenterDictionaries:   *c.Dictionary,
		StopTokenFile:           *c.StopTokens,
		UsePersistentStorage:    true,
		PersistentStorageFolder: *c.IndexPath,
	})

	return
}
