package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/khengsaurus/mahjong-cms/utils"
)

const last_checked_count = 21_000
const counter_url = "https://firestore.googleapis.com/v1/projects/mahjong-sg/databases/(default)/documents/metrics/local-counter"

type BadgeIOPayload struct {
	SchemaVersion int    `json:"schemaVersion"`
	Label         string `json:"label"`
	Message       string `json:"message"`
	Color         string `json:"color"`
}

type FirebaseRes struct {
	Name       string `json:"name"`
	Fields     Fields `json:"fields"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

type Fields struct {
	Count Count `json:"count"`
}

type Count struct {
	IntegerValue string `json:"integerValue"`
}

func GetCount(w http.ResponseWriter, r *http.Request) {
	badge := r.URL.Query().Get("badge") == "true"
	games := last_checked_count

	resp, err := http.Get(counter_url)
	if err == nil {

		if body, err := io.ReadAll(resp.Body); err == nil {
			var res FirebaseRes
			if err := json.Unmarshal([]byte(body), &res); err == nil {
				updatedGames, convErr := strconv.Atoi(res.Fields.Count.IntegerValue)
				if convErr == nil {
					games = updatedGames
				}
			} else {
				fmt.Printf("GetCount controller: unexpected response from firebase\n%v\n", err)
			}
		}
	}

	if badge {
		utils.Json200(&BadgeIOPayload{
			SchemaVersion: 1,
			Label:         "Games played",
			Message:       utils.GetStr(games),
			Color:         "orange",
		}, w)
	} else {
		utils.Json200(games, w)
	}
}
