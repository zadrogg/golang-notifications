package rules

import (
	"net/http"
	"notifications/handlers"
)

func HandBadRequest(w http.ResponseWriter, err error) {
	handlers.Throw400(w, err)
	return
}
