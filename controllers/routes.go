package controllers

import (
	"time"

	"github.com/go-chi/chi"
	ChiMiddlewares "github.com/go-chi/chi/middleware"
	"github.com/khengsaurus/mahjong-cms/middlewares"
)

var MahjongContent = func(mj_router chi.Router) {
	mj_router.Use(middlewares.VerifyHeader("source", "mj-sg-"))
	mj_router.Use(middlewares.SetHeader("Cache-Control", "public, max-age=21600"))
	mj_router.Use(ChiMiddlewares.Timeout(30 * time.Second))
	mj_router.Get("/{key}", GetContent)
}
