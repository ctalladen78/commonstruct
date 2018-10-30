package stdreq

import (
	"strings"
)

//NormalizeEmail normalizes email address
func NormalizeEmail(s string) string {
	s = strings.ToLower(NormalizeString(s)) // lower case and trim space
	return strings.TrimRight(s, ".")        //trim . at the end
}
