package appresult

import (
	"net/http"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

func HeaderContentTypeJson() (string, string) {
	return "Content-Type", "application/json"
}
