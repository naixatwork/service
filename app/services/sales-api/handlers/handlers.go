package handlers

import (
	"github.com/naixatwork/service/app/services/sales-api/handlers/v1/testgrp"
	"github.com/naixatwork/service/business/web/v1/mid"
	"github.com/naixatwork/service/foundation/web"
	"go.uber.org/zap"
	"net/http"
	"os"
)

type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *zap.SugaredLogger
}

func APIMux(cfg APIMuxConfig) *web.App {
	app := web.NewApp(cfg.Shutdown, mid.Logger(cfg.Log), mid.Errors(cfg.Log), mid.Panics())

	app.Handle(http.MethodGet, "/test", testgrp.Test)

	return app
}
