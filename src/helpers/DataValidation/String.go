package datavalidation

import (
	"errors"
	"fmt"
	"regexp"
)

var MessageNotNull = "Um valor deve ser informado para o campo {{label}}"
var MessageNotEmpty = "O campo {{label}} não pode estar vazio"
var MessageMin = "O campo {{label}} deve ter no mínimo {{min}} letras"
var MessageMax = "O campo {{label}} deve ter no maximo {{max}} letras"
var MessageEmail = "O e-mail \"{{value}}\" informado não é válido"

type TypeString struct {
	label    *string
	notEmpty *bool
	notNull  *bool
	min      *int
	max      *int
	isEmail  *bool
}

func String(label string) *TypeString {
	var obj TypeString
	neg := false
	obj.label = &label
	obj.isEmail = &neg
	obj.notEmpty = &neg
	obj.notNull = &neg
	return &obj
}

func (s *TypeString) Min(value int) *TypeString {
	s.min = &value
	return s
}

func (s *TypeString) Max(value int) *TypeString {
	s.max = &value
	return s
}

func (s *TypeString) NotEmpty(value bool) *TypeString {
	s.notEmpty = &value
	return s
}

func (s *TypeString) NotNull(value bool) *TypeString {
	s.notNull = &value
	return s
}

func (s *TypeString) IsEmail(value bool) *TypeString {
	s.isEmail = &value
	return s
}

func (s *TypeString) Validate(value *string) error {

	varString := value

	fmt.Println(fmt.Sprint(varString))

	if value == nil {
		if *s.notNull == true {
			return errors.New(ReplaceTextMessage(MessageNotNull, s, varString))
		}
	} else {
		if s.min != nil {
			if len(*value) < *s.min {
				return errors.New(ReplaceTextMessage(MessageMin, s, value))
			}
		}
		if s.max != nil {
			if len(*value) > *s.max {
				return errors.New(ReplaceTextMessage(MessageMax, s, value))
			}
		}
		if *s.notEmpty == true {
			if *value == "" {
				return errors.New(ReplaceTextMessage(MessageNotEmpty, s, value))
			}
		}
		if *s.isEmail == true {
			if !isEmailValid(*value) {
				return errors.New(ReplaceTextMessage(MessageEmail, s, value))
			}
		}
	}

	return nil
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}
