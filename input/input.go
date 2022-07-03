package input

import (
	"email-checker-tool/entity"
	"errors"
	"regexp"
	"strings"
)

const REGEX_VALIDATE_EMAIL = "^[-!#$%&'*+/0-9=?A-Z^_a-z`{|}~](\\.?[-!#$%&'*+/0-9=?A-Z^_a-z`{|}~])*@[a-zA-Z0-9](-*\\.?[a-zA-Z0-9])*\\.[a-zA-Z](-?[a-zA-Z0-9])+$"

type EmailInput struct {
	Email string
}

func NewEmailInput(email string) EmailInput {
	return EmailInput{Email: email}
}

func (i EmailInput) NewEmail() (*entity.Email, error) {
	r, err := regexp.Compile(REGEX_VALIDATE_EMAIL)

	if err != nil {
		return nil, err
	}

	if !r.MatchString(i.Email) {
		return nil, errors.New("Invalid email format.")
	}

	localAndDomain := strings.Split(i.Email, "@")
	domain, err := entity.NewDomain(localAndDomain[1])

	if err != nil {
		return nil, err
	}

	return &entity.Email{
		Local:  localAndDomain[0],
		Domain: *domain,
	}, nil
}
