package handlers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"notifications/notify"
)

const Mail string = "mail"
const Telegram string = "telegram"

var err error

func SendNotification(write http.ResponseWriter, request *http.Request) {
	typeNotify := chi.URLParam(request, "type")

	switch typeNotify {
	case Mail:
		err = notify.SendNotifyMail(request)
	case Telegram:
		notify.SendNotifyTelegram()
	}

	if err != nil {
		Throw400(write, err)
		return
	}
}
