package datavalidation

import (
	"fmt"
	"strings"
)

func ReplaceTextMessage(messageBase string, s *TypeString, value *string) string {
	var message = messageBase
	message = strings.ReplaceAll(message, "{{label}}", *s.label)
	if s.min != nil {
		message = strings.ReplaceAll(message, "{{min}}", fmt.Sprint(*s.min))
	}
	if s.max != nil {
		message = strings.ReplaceAll(message, "{{max}}", fmt.Sprint(*s.max))
	}
	message = strings.ReplaceAll(message, "{{value}}", fmt.Sprint(*value))
	return message
}
