package main

import "fmt"

// Interfaces are now split. Whoever need Email sending functionality, can just extend/implement this interface.
type EmailSender interface {
	SendEmail(to string, msg string)
}

// Interfaces are now split. Whoever need SMS sending functionality, can just extend/implement this interface.
type SMSSender interface {
	SendSMS(to string, msg string)
}

// Interfaces are now split. Whoever need Push(PN) sending functionality, can just extend/implement this interface.
type PushSender interface {
	SendPush(to string, msg string)
}

// Concrete implementations. Only implementing EmailSender interface. 
type EmailNotifier struct{}

func (e EmailNotifier) SendEmail(to string, msg string) {
	fmt.Printf("Email to %s: %s\n", to, msg)
}

// Concrete implementations. Only implementing SMSSender interface. 
type SMSNotifier struct{}

func (s SMSNotifier) SendSMS(to string, msg string) {
	fmt.Printf("SMS to %s: %s\n", to, msg)
}

// Concrete implementations. Only implementing PushSender interface. 
type PushNotifier struct{}

func (p PushNotifier) SendPush(to string, msg string) {
	fmt.Printf("Push to %s: %s\n", to, msg)
}

// UserService only depends on what it needs
type UserService struct {
	emailNotifier EmailSender
}

func NewUserService(emailNotifier EmailSender) *UserService {
	return &UserService{emailNotifier: emailNotifier}
}

func (us *UserService) WelcomeUser(email string) {
	us.emailNotifier.SendEmail(email, "Welcome to our service!")
}

func main() {
	emailService := EmailNotifier{}
	userService := NewUserService(emailService)

	userService.WelcomeUser("user@example.com")
}

/*
Further improvements in this:
Create Notification interface
expose one method: sendNotification(Map<Data>);

Have respective implementation of Notification interface.

Have a factory class which asks for NotificationType and gives the object of that.

call notification.sendNotification(Map<Data>);

Above one is more generic in nature. 
*/

