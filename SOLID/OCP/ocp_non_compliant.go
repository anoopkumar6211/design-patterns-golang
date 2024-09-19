/*
In the following example, every time a new payment method is added, the ProcessPayment function needs to be modified. This violates the OCP because we have to keep modifying the existing code to support new functionality.
Adding a new payment method, e.g., Bitcoin or Google Pay, would require modifying the ProcessPayment method to add more else if statements. This breaks the "closed for modification" rule of OCP.
*/

package main

import "fmt"

// PaymentService handles multiple payment methods.
type PaymentService struct{}

func (p PaymentService) ProcessPayment(paymentType string, amount float64) {
	// Modification required every time a new payment type is introduced
	if paymentType == "credit" {
		fmt.Printf("Processing credit card payment of $%.2f\n", amount)
	} else if paymentType == "paypal" {
		fmt.Printf("Processing PayPal payment of $%.2f\n", amount)
	} else {
		fmt.Println("Unknown payment method")
	}
}

func main() {
	service := PaymentService{}
	service.ProcessPayment("credit", 100.00)  // Needs modification for new types
	service.ProcessPayment("paypal", 200.00)  // Same here
}
