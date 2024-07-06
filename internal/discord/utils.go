package discord

import (
	"strings"
)

func isBadRequest(err error) bool {
	return strings.Contains(err.Error(), "HTTP 400 Bad Request")
}
