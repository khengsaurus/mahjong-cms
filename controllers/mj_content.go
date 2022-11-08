package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/khengsaurus/mahjong-cms/static"
	"github.com/khengsaurus/mahjong-cms/utils"
)

var keys = []string{"about", "help", "notifs", "policy"}

func GetContent(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")

	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !utils.StringInArray(key, keys) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	source := GetSource(key)
	var content map[string]interface{}
	err := json.Unmarshal([]byte(source), &content)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.Json200(content, w)
}

func GetSource(key string) string {
	switch key {
	case "about":
		return static.AboutContent
	case "help":
		return static.HelpContent
	case "notifs":
		return static.NotifsContent
	case "policy":
		return static.PolicyContent
	default:
		return ""
	}
}
