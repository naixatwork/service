package handlers

import (
	"encoding/json"
	"github.com/dimfeld/httptreemux/v5"
	"go.uber.org/zap"
	"net/http"
	"os"
)

type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *zap.SugaredLogger
}

func APIMux(cfg APIMuxConfig) http.Handler {
	mux := httptreemux.NewContextMux()

	h := func(w http.ResponseWriter, r *http.Request) {
		status := struct {
			Status string
		}{
			Status: "OK",
		}

		if err := json.NewEncoder(w).Encode(status); err != nil {
			cfg.Log.Errorw("error in encoding", "status", status)
		}
	}
	mux.GET("/test", h)

	return mux
}
