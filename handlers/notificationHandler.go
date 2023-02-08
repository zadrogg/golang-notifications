package handlers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"notifications/notify"
)

const Mail string = "mail"
const Telegram string = "telegram"

func SendNotification(write http.ResponseWriter, request *http.Request) {
	typeNotify := chi.URLParam(request, "type")
	switch typeNotify {
	case Mail:
		notify.SendNotifyMail()
	case Telegram:
		notify.SendNotifyTelegram()
	}
}
