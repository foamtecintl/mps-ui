package app

import (
	"mps-ui/pkg/view"
	"net/http"
)

func mpsHome(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"api_url": apiURL,
	}
	view.MpsHome(w, data)
}

func mpsSelectForecast(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"api_url": apiURL,
	}
	search := r.URL.Query().Get("search")
	data["search"] = search
	view.MpsSelectForecast(w, data)
}

func mpsCreateForecast(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"api_url": apiURL,
	}
	id := r.URL.Query().Get("id")
	dateT := r.URL.Query().Get("dateT")
	data["id"] = id
	data["dateT"] = dateT
	view.MpsCreateForecast(w, data)
}

func mpsCreateDate(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"api_url": apiURL,
	}
	id := r.URL.Query().Get("id")
	data["id"] = id
	view.MpsCreateDate(w, data)
}

func mpsCreateGroup(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"api_url": apiURL,
	}
	search := r.URL.Query().Get("search")
	data["search"] = search
	view.MpsCreateGroup(w, data)
}

func mpsCreatePart(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"api_url": apiURL,
	}
	search := r.URL.Query().Get("search")
	data["search"] = search
	groupID := r.URL.Query().Get("id")
	if len(groupID) != 0 {
		data["id"] = groupID
		view.MpsCreatePart(w, data)
	}
}
