package handlers

import (
	"github.com/dimfeld/httptreemux/v5"
	"github.com/naixatwork/service/app/services/sales-api/handlers/v1/testgrp"
	"go.uber.org/zap"
	"os"
)

type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *zap.SugaredLogger
}

func APIMux(cfg APIMuxConfig) *httptreemux.ContextMux {
	mux := httptreemux.NewContextMux()

	mux.GET("/test", testgrp.Test)

	return mux
}
