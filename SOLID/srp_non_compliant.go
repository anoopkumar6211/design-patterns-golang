package main

import "fmt"

// UserService manages users and also sends emails (violating SRP).
type UserService struct {
	users []string
}

func (s *UserService) AddUser(name string) {
	s.users = append(s.users, name)
	fmt.Printf("User %s added\n", name)
	s.sendWelcomeEmail(name) // User service is handling email, violating SRP
}

func (s *UserService) sendWelcomeEmail(name string) {
	fmt.Printf("Sending welcome email to %s\n", name)
}

func main() {
	service := &UserService{}
	service.AddUser("John")  // User management and email are mixed
}
