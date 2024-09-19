/*
The PaymentProcessor interface defines a contract for payment methods, allowing different payment methods (e.g., CreditCard, PayPal) to implement their own Process method.
The PaymentService class is now closed for modification because you don’t need to change it to support new payment methods. It's open for extension since you can add new payment types (e.g., Bitcoin, Google Pay) by simply creating new structs that implement the PaymentProcessor interface.
*/

package main

import "fmt"

// PaymentProcessor is an interface that all payment methods must implement.
type PaymentProcessor interface {
	Process(amount float64)
}

// CreditCard is one type of payment method.
type CreditCard struct{}

func (c CreditCard) Process(amount float64) {
	fmt.Printf("Processing credit card payment of $%.2f\n", amount)
}

// PayPal is another type of payment method.
type PayPal struct{}

func (p PayPal) Process(amount float64) {
	fmt.Printf("Processing PayPal payment of $%.2f\n", amount)
}

// PaymentService is open for extension but closed for modification.
type PaymentService struct{}

// ProcessPayment accepts any payment processor that implements the PaymentProcessor interface.
func (p PaymentService) ProcessPayment(processor PaymentProcessor, amount float64) {
	processor.Process(amount) // Delegate processing to the specific payment processor
}

func main() {
	service := PaymentService{}
	creditCard := CreditCard{}
	payPal := PayPal{}

	service.ProcessPayment(creditCard, 100.00)  // No need to modify service for new types
	service.ProcessPayment(payPal, 200.00)
}
