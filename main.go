package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/pengsrc/go-shared/check"
	"github.com/spf13/cobra"

	"github.com/Xuanwo/xuanzang/config"
	"github.com/Xuanwo/xuanzang/constants"
	"github.com/Xuanwo/xuanzang/contexts"
	"github.com/Xuanwo/xuanzang/index"
	"github.com/Xuanwo/xuanzang/search"
)

var application = &cobra.Command{
	Use:   constants.Name,
	Short: constants.Name,
	Long:  constants.Name,
	Run: func(cmd *cobra.Command, args []string) {
		if flagVersion {
			fmt.Printf("%s version %s\n", constants.Name, constants.Version)
			return
		}

		// Initialize config.
		c := config.New()
		check.ErrorForExit(constants.Name, c.LoadFromFilePath(flagConfig))
		check.ErrorForExit(constants.Name, c.Check())

		// Setup contexts.
		check.ErrorForExit(constants.Name, contexts.SetupContexts(c))
		defer contexts.Searcher.Close()
		defer contexts.DB.Close()

		wg := sync.WaitGroup{}

		wg.Add(1)
		go func() {
			defer wg.Done()

			// Do index firstly.
			go index.Index()

			// Execute index for every Source.Duration seconds.
			ticker := time.NewTicker(time.Duration(c.Source.Duration) * time.Second)

			for _ = range ticker.C {
				go index.Index()
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()

			r := chi.NewRouter()
			r.Use(middleware.Logger)
			r.Use(middleware.Recoverer)

			r.Get("/", search.Search)

			addr := fmt.Sprintf("%s:%d", *c.Host, *c.Port)
			http.ListenAndServe(addr, r)
		}()

		wg.Wait()
	},
}

var (
	flagVersion bool
	flagConfig  string
)

func init() {
	application.Flags().BoolVarP(
		&flagVersion, "version", "v", false, "Show version",
	)
	application.Flags().StringVarP(
		&flagConfig, "config", "c", "xuanzang.yaml", "Specify config file path",
	)
}

func main() {
	application.Execute()
}
