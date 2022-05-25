package main

import "fmt"

// Abstracciones de un Método de notificación

// INotificationFactory : Interfaz de Notificación que puede ser PUSH, EMAIL o SMS
type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

// Método SMS

type SMSNotification struct {
}

func (SMSNotification) SendNotification() {
	fmt.Println("Enviando notificación por SMS")
}

func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

// Email

type EmailNotification struct {
}

func (EmailNotification) SendNotification() {
	fmt.Println("Enviando notificación por Email")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "EMAIL"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "SendGrid"
}

// Fábricas concretas

func getNotificationFactory(notificationType string) (INotificationFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}
	if notificationType == "EMAIL" {
		return &EmailNotification{}, nil
	}
	return nil, fmt.Errorf("No Notificacion type")
}

func SendNotification(f INotificationFactory) {
	f.SendNotification()
}

func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("EMAIL")

	SendNotification(smsFactory)
	SendNotification(emailFactory)

	getMethod(smsFactory)
	getMethod(emailFactory)

}
