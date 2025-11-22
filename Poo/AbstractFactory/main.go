package main

//tenemos un software que emite notificaciones que no sabemos cuales seran pueden ser push, sms o email

import "fmt"

// interface de  mi fabrica abstracta
type INotificationSenderFactory interface {
	CreateNotificationSender() INotificationSender
}

// interface del producto
type INotificationSender interface {
	SendNotification(message string) string
	GetSenderChannel() string
}

// implementacion concreta del producto
type SmsNotificationSender struct {
}

// metodos del producto concreto
func (s SmsNotificationSender) SendNotification(message string) string {
	return fmt.Sprintf("Sending SMS notification with message: %s", message)
}

func (s SmsNotificationSender) GetSenderChannel() string {
	return "SMS"
}

// implementacion concreta de la fabrica
type SmsNotificationSenderFactory struct {
}

// metodo de la fabrica concreta
func (s SmsNotificationSenderFactory) CreateNotificationSender() INotificationSender {
	return SmsNotificationSender{}
}

// otra implementacion concreta del producto
type EmailNotificationSender struct {
}

func (e EmailNotificationSender) SendNotification(message string) string {
	return fmt.Sprintf("Sending Email notification with message: %s", message)
}

func (e EmailNotificationSender) GetSenderChannel() string {
	return "Email"
}

// otra implementacion concreta de la fabrica
type EmailNotificationSenderFactory struct {
}

func (e EmailNotificationSenderFactory) CreateNotificationSender() INotificationSender {
	return EmailNotificationSender{}
}

// funcion que utiliza la fabrica abstracta
func SendNotification(factory INotificationSenderFactory, message string) {
	sender := factory.CreateNotificationSender()
	result := sender.SendNotification(message)
	fmt.Println(result)
}

// funcion main
func main() {
	smsFactory := SmsNotificationSenderFactory{}
	emailFactory := EmailNotificationSenderFactory{}
	SendNotification(emailFactory, "Hello via Email!")
	SendNotification(smsFactory, "Hello via SMS!")
}
