package main

import (
	"fmt"
	"net/http"

	"github.com/mrtc0/seccamp-2023/pkg/config"
	"github.com/mrtc0/seccamp-2023/pkg/di"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	cfg, err := config.LoadFromEnv()
	if err != nil {
		logrus.Fatal(err)
	}

	engine := di.Initialize(*cfg)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: engine,
	}

	if err := server.ListenAndServe(); err != nil {
		logrus.Fatal(err)
	}
}
