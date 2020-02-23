package commons

import (
	"net/http"
	"regexp"
	"usermanagement/models"

	"golang.org/x/crypto/bcrypt"
)

func IsEmailValid(email string) bool {
	regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return regex.MatchString(email)
}

func ComparePassword(hashedPassword string, password []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), password)
	if err != nil {
		return false
	}

	return true
}

func ValidateNameEmailPass(writer http.ResponseWriter, user models.User) bool {
	if len(user.Name) < 3 {
		Response(writer, http.StatusBadRequest, Message(false, "Name cannot be empty and al least 3 character."))
		return false
	}

	valid := ValidateEmailPass(writer, user)
	if !(valid) {
		return false
	}

	return true
}

func ValidateEmailPass(writer http.ResponseWriter, user models.User) bool {
	if len(user.Email) < 1 {
		Response(writer, http.StatusBadRequest, Message(false, "Email cannot be empty."))
		return false
	}

	if !(IsEmailValid(user.Email)) {
		Response(writer, http.StatusBadRequest, Message(false, "Please insert a valid email address."))
		return false
	}

	if len(user.Password) < 8 {
		Response(writer, http.StatusBadRequest, Message(false, "Password cannot be empty and at least 8 character."))
		return false
	}

	return true
}
