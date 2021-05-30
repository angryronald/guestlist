package main

import (
	"fmt"
	netHttp "net/http"

	"github.com/angryronald/guestlist/cmd/container"
	"github.com/angryronald/guestlist/cmd/http"
	"github.com/angryronald/guestlist/config"
	"github.com/angryronald/guestlist/database"
	"github.com/go-chi/chi"
	"github.com/oklog/oklog/pkg/group"
	"github.com/sirupsen/logrus"
)

func main() {
	var g = group.Group{}
	var logger = config.GetLogger()

	container.NewIOC()

	runMigration()
	injectSeed()
	runHTTP(&g, *logger)

	logger.Fatal("exit", g.Run())
}

func runMigration() {
	db := config.DB()

	err := database.Migrate(db)
	if nil != err {
		panic(fmt.Sprintf("Error on migrating database: %v", err))
	}
}

func injectSeed() {
	db := config.DB()

	_ = database.Seed(db)
}

func runHTTP(
	g *group.Group,
	logger logrus.Logger,
) {
	port := config.GetEnv(config.HTTP_ADDRESS)

	if len(port) < 1 {
		panic(fmt.Sprintf("Environment Missing!\n*%s* is required", port))
	}

	var router *chi.Mux
	router = chi.NewRouter()

	router.Mount("/v1", http.MakeHandlerV1(router))

	server := &netHttp.Server{
		Addr:    port,
		Handler: router,
	}

	fmtLog := logger.WithFields(logrus.Fields{
		"transport": "debug/HTTP",
		"addr":      port,
	})

	g.Add(
		func() error {
			fmtLog.Info("HTTP transport run at ", port)
			return server.ListenAndServe()
		},
		func(err error) {
			if nil != err {
				fmtLog.Warn("Error Occurred ", err.Error())
				panic(err)
			}
		},
	)
}
