package notify

import (
	"encoding/json"
	"net/http"
	"net/smtp"
	"notifications/config"
	"strconv"
)

type Mail struct {
	Mail    string `json:"mail"`
	Message string `json:"message"`
	Subject string `json:"subject"`
}

var (
	conf    = config.GetConfig()
	mStruct Mail
)

func connectSmtp() smtp.Auth {
	return smtp.PlainAuth("", conf.Mail.Mail, conf.Mail.Password, conf.Mail.Smtp)
}

func SendNotifyMail(request *http.Request) error {
	err := json.NewDecoder(request.Body).Decode(&mStruct)
	if err != nil {
		return err
	}
	auth := connectSmtp()
	connSmtp := conf.Mail.Smtp + ":" + strconv.Itoa(conf.Mail.Port)
	to := []string{mStruct.Mail}
	msg := []byte("To: " + mStruct.Mail + "\r\n" +
		"Subject: " + mStruct.Subject + "\r\n" +
		"\r\n" + mStruct.Message + "\r\n")

	return smtp.SendMail(connSmtp, auth, conf.Mail.Mail, to, msg)
}
