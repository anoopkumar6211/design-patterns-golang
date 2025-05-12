package main

import "fmt"

// Concrete class. 
type FileLogger struct{}

func (fl FileLogger) Log(message string) {
	fmt.Println("Logging to file:", message)
}

// High-level module
type Service struct {
  //Directly concrete class dependency injected. Now, it's tightly coupled with fileLogger only. Not possible to change the logger here. 
	logger FileLogger
}

func (s Service) PerformAction() {
	s.logger.Log("Action performed!")
}

func main() {
	// High-level module directly depends on concrete class(Not interface)
	logger := FileLogger{}
	service := Service{logger: logger}
	service.PerformAction()
}
