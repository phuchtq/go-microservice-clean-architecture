package helper

import (
	"architecture_template/common_dtos/request"
	mailconst "architecture_template/constants/mailConst"
	"architecture_template/constants/notis"
	"bytes"
	"errors"
	"html/template"
	"log"
	"os"
	"strconv"

	"gopkg.in/mail.v2"
)

func SendMail(req request.SendMailReqDto) error {
	var body bytes.Buffer

	template, err := template.ParseFiles(req.TemplatePath)
	if err != nil {
		log.Print(notis.MailHelperMsg + "SendMail - " + err.Error())
		return errors.New(notis.InternalErr)
	}
	template.Execute(&body, req.Body)

	var serviceEmail string = os.Getenv(mailconst.ServiceEmail)
	var securityPass string = os.Getenv(mailconst.SecurityPassword)
	var host string = os.Getenv(mailconst.Host)
	port, err := strconv.Atoi(os.Getenv(mailconst.MailPort))
	if err != nil {
		port = 587
	}

	var m = mail.NewMessage()
	m.SetHeader("From", serviceEmail)
	m.SetHeader("To", req.Body.Email)
	m.SetHeader("Subject", req.Subject)
	m.SetBody("text/html", body.String())

	diabler := mail.NewDialer(host, port, serviceEmail, securityPass)

	if err := diabler.DialAndSend(m); err != nil {
		log.Print(notis.MailHelperMsg + "SendMail - " + err.Error())
		return errors.New(notis.GenerateMailWarnMsg)
	}

	return nil
}
