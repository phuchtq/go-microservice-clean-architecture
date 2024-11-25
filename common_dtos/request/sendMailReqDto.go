package request

import "log"

type SendMailReqDto struct {
	Body         MailBody    `json:"mail_body"`
	TemplatePath string      `json:"template_path"`
	Subject      string      `json:"subject"`
	Logger       *log.Logger `json:"logger"`
}
