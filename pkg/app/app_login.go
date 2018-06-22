package app

import (
	"mps-ui/pkg/view"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"api_url": apiURL,
	}
	view.Login(w, data)
}
