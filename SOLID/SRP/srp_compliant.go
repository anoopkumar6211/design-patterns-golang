package main

import "fmt"

// UserService is responsible for managing users.
type UserService struct {
	users []string
	emailService EmailService // Dependency on the EmailService
}

// AddUser adds a user and delegates email notification to EmailService.
func (s *UserService) AddUser(name string) {
	s.users = append(s.users, name)
	fmt.Printf("User %s added\n", name)
	s.emailService.SendWelcomeEmail(name) // Delegate email sending
}

// EmailService is responsible for sending emails.
type EmailService struct{}

func (e EmailService) SendWelcomeEmail(name string) {
	fmt.Printf("Sending welcome email to %s\n", name)
}

func main() {
	emailService := EmailService{}
	userService := UserService{emailService: emailService}
	userService.AddUser("John")  // User management and email services are separated
}
