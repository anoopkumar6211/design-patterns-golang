package main

import "fmt"

// Fat interface forces all implementers to support all channels
type Notifier interface {
	SendEmail(to string, msg string)
	SendSMS(to string, msg string)
	SendPush(to string, msg string)
}

type EmailOnlyNotifier struct{}

func (e EmailOnlyNotifier) SendEmail(to string, msg string) {
	fmt.Printf("Email to %s: %s\n", to, msg)
}

func (e EmailOnlyNotifier) SendSMS(to string, msg string) {
	// Not supported
	fmt.Println("SMS not supported")
}

func (e EmailOnlyNotifier) SendPush(to string, msg string) {
	// Not supported
	fmt.Println("Push not supported")
}

type SMSOnlyNotifier struct{}

func (e SMSOnlyNotifier) SendEmail(to string, msg string) {
	fmt.Printf("Email not supported")
}

func (e SMSOnlyNotifier) SendSMS(to string, msg string) {
	// Not supported
	fmt.Println("SMS to %s: %s\n", to, msg)
}

func (e SMSOnlyNotifier) SendPush(to string, msg string) {
	// Not supported
	fmt.Println("Push not supported")
}

func main() {
	var notifier Notifier = EmailOnlyNotifier{}
	notifier.SendEmail("user@example.com", "Welcome!")
	notifier.SendSMS("12345", "Hello") // irrelevant call
}

