package main

import (
	"errors"
	"fmt"
)

type User struct {
	Phone   string
	Name    string
	Balance float64
}

var users = make(map[string]*User)

func RegisterUser(phone, name string) error {
	if _, exists := users[phone]; exists {
		return errors.New("user already exists")
	}
	users[phone] = &User{
		Phone:   phone,
		Name:    name,
		Balance: 0.0,
	}
	return nil
}

func AddMoney(phone string, amount float64) error {
	user, exists := users[phone]
	if !exists {
		return errors.New("user not found")
	}
	user.Balance += amount
	return nil
}

func TransferMoney(from, to string, amount float64) error {
	sender, exists := users[from]
	if !exists {
		return errors.New("sender not found")
	}
	receiver, exists := users[to]
	if !exists {
		return errors.New("receiver not found")
	}
	if sender.Balance < amount {
		return errors.New("insufficient balance")
	}
	sender.Balance -= amount
	receiver.Balance += amount
	fmt.Printf("Transferred ₹%.2f from %s to %s\n", amount, sender.Name, receiver.Name)
	return nil
}

func ShowBalance(phone string) {
	if user, exists := users[phone]; exists {
		fmt.Printf("%s (%s) - Balance: ₹%.2f\n", user.Name, user.Phone, user.Balance)
	} else {
		fmt.Println("User not found")
	}
}

func main() {
	_ = RegisterUser("9876543210", "Alice")
	_ = RegisterUser("9123456780", "Bob")

	_ = AddMoney("9876543210", 1000)
	_ = AddMoney("9123456780", 500)

	ShowBalance("9876543210")
	ShowBalance("9123456780")

	_ = TransferMoney("9876543210", "9123456780", 300)

	ShowBalance("9876543210")
	ShowBalance("9123456780")
}
