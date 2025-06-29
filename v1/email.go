package mytools

import "net/mail"

// IsEmailValid chacks if the email in parameter is a valid email address
func IsEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
