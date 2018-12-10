package main

import (
	"bytes"
	"fmt"
	"github.com/k0kubun/pp"
	"html/template"
	"log"
	"net/smtp"
)

const (
	MIME     = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	SERVER   = "in-v3.mailjet.com"
	PORT     = 587
	EMAIL    = "e90ba227bb65caf680fa65204f129c93"
	PASSWORD = "211cf1ee79d96d220d88db3efb026a7f"
)

type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

func NewRequest(to []string, subject string) *Request {
	return &Request{
		to:      to,
		subject: subject,
	}
}

func (r *Request) parseTemplate(fileName string, data map[string]interface{}) error {
	t, err := template.ParseFiles(fileName)
	if err != nil {
		return err
	}
	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, data); err != nil {
		return err
	}
	r.body = buffer.String()
	return nil
}

func (r *Request) SendMail() error {
	body := "From: CITA Utec <cita@utec.edu.pe>\r\nTo: " + r.to[0] + "\r\nSubject: " + r.subject + "\r\n" + MIME + "\r\n" + r.body
	// body := "To: " + r.to[0] + "\r\nSubject: " + r.subject + "\r\n" + MIME + "\r\n" + r.body
	SMTP := fmt.Sprintf("%s:%d", SERVER, PORT)
	auther := smtp.PlainAuth("", EMAIL, PASSWORD, SERVER)

	if err := smtp.SendMail(SMTP, auther, EMAIL, r.to, []byte(body)); err != nil {
		pp.Println("ERROR", err.Error())
		return err
	}
	return nil
}

func (r *Request) Send(templateName string, items map[string]interface{}) error {
	err := r.parseTemplate(templateName, items)
	if err != nil {
		log.Fatal(err)
	}
	if err := r.SendMail(); err != nil {
		log.Printf("Failed to send the email to %s\n", r.to)
		return err
	}
	log.Printf("Email has been sent to %s\n", r.to)
	return nil
}
