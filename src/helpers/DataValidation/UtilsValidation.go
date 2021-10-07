package datavalidation

import (
	"fmt"
	"strings"
)

type Generic interface {
	GetValue() string
}

type TextString struct {
	Value string
}

func (t TextString) GetValue() string {
	return t.Value
}

type TextInteger struct {
	Value int
}

func (t TextInteger) GetValue() string {
	return fmt.Sprint(t.Value)
}

func ReplaceTextMessage(messageBase string, s *TypeString, text Generic) string {
	var message = messageBase
	message = strings.ReplaceAll(message, "{{label}}", *s.label)
	if s.min != nil {
		message = strings.ReplaceAll(message, "{{min}}", fmt.Sprint(*s.min))
	}
	if s.max != nil {
		message = strings.ReplaceAll(message, "{{max}}", fmt.Sprint(*s.max))
	}
	message = strings.ReplaceAll(message, "{{value}}", text.GetValue())
	return message
}
