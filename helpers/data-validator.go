package helpers

import (
	"errors"
	"regexp"
)

func ValidateEmail(email string) error {
	emailRgx := regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]`)

	if !emailRgx.MatchString(email) {
		// return fmt.Errorf("email %s no es válido", email)
		return errors.New("email no es válido")
	}

	return nil
}

func ValidatePassword(password string) error {
	passwordLen := 6

	if len(password) < passwordLen {
		return errors.New("la contraseña es demasiado corta")
	}

	return nil
}
