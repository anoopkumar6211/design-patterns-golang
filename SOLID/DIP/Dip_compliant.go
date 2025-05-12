/* Service depends on interface rather than class.
   Very easy to switch the logger implementation by passing the respective implementation during class initialisation. 
   This offers decoupling of dependency injections. Depend on abstractions (interfaces) rather than concrete implementations.
*/

package main

import "fmt"

// Abstraction (Interface). This logger interface can be extended by any new logger implementation. 
type Logger interface {
	Log(message string)
}

// Low-level module: FileLogger implements Logger interface
type FileLogger struct{}

func (fl FileLogger) Log(message string) {
	fmt.Println("Logging to file:", message)
}

// Low-level module: ConsoleLogger implements Logger interface
type ConsoleLogger struct{}

func (cl ConsoleLogger) Log(message string) {
	fmt.Println("Logging to console:", message)
}

// High-level module: Service depends on Logger interface, not FileLogger directly so easy to switch whenever needed and that will just be dependency change rather than thinking through this again.
type Service struct {
	logger Logger
}

func (s Service) PerformAction() {
	s.logger.Log("Action performed!")
}

func main() {
	// You can now inject any Logger implementation and system will just start using that respective logger implementation. 
	var logger Logger

	// Using FileLogger
	logger = FileLogger{}
	service := Service{logger: logger}
	service.PerformAction()

	// Switching to ConsoleLogger without modifying Service
	logger = ConsoleLogger{}
	service = Service{logger: logger}
	service.PerformAction()
}

