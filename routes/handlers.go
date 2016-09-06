package routes

import (
	"net/http"
	"strconv"
)

func Route404(w http.ResponseWriter, r *http.Request) {
	sendReponse(w, http.StatusNotFound, map[string]interface{}{
		"error": "Given path was not found",
	})
}

func HomeRoute(w http.ResponseWriter, r *http.Request) {
	sendReponse(w, 200, map[string]interface{}{
		"status":  "Ok",
		"version": "1.0",
	})
}

func ListRoute(w http.ResponseWriter, r *http.Request) {
	m := make([]struct {
		Id   int    `json:"id"`
		Name string `json:"string"`
	}, len(allItems))

	for i, item := range allItems {
		m[i].Id = item.Id
		m[i].Name = item.Name
	}

	sendReponse(w, 200, map[string]interface{}{
		"items": m,
	})
}

func GetRoute(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if len(query["id"]) == 0 {
		sendReponse(w, http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid ID",
		})
		return
	}

	id := query["id"][0]

	idx, err := strconv.Atoi(id)
	if err != nil {
		sendReponse(w, http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid ID",
		})
		return
	}

	item, ok := mapItems[idx]
	if !ok {
		sendReponse(w, http.StatusNotFound, map[string]interface{}{
			"error": "Item not found",
		})
		return
	}

	sendReponse(w, 200, map[string]interface{}{
		"item": item,
	})
}

func MixPotionRoute(w http.ResponseWriter, r *http.Request) {
	params, err := parseJSONBody(r.Body)
	if err != nil {
		sendReponse(w, http.StatusBadRequest, map[string]interface{}{
			"error": "could not parse json body",
		})
		return
	}

	if params["ids"] == nil || len(params["ids"].([]interface{})) < 2 {
		sendReponse(w, http.StatusBadRequest, map[string]interface{}{
			"error": "Need at least 2 IDs",
		})
		return
	}

	items := make([]Item, len(params["ids"].([]interface{})))
	for i, id := range params["ids"].([]interface{}) {
		var idStr string
		var idx int

		if t, ok := id.(string); ok {
			idStr = t
			idx, err = strconv.Atoi(idStr)
			if err != nil {
				sendReponse(w, http.StatusBadRequest, map[string]interface{}{
					"error": "Invalid ID " + idStr,
				})
				return
			}
		} else if t, ok := id.(float64); ok {
			idx = int(t)
			idStr = strconv.Itoa(idx)
		} else {
			sendReponse(w, http.StatusBadRequest, map[string]interface{}{
				"error": "Invalid ID",
			})
			return
		}

		item, ok := mapItems[idx]
		if !ok {
			sendReponse(w, http.StatusNotFound, map[string]interface{}{
				"error": "Item not found for ID " + idStr,
			})
			return
		}

		items[i] = item
	}

	sendReponse(w, 200, map[string]interface{}{
		"toxicity": mixPotionItems(items...),
	})
}
