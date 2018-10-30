package stdreq

import "strings"

//NormalizeString normalize all string
func NormalizeString(s string) string {
	return strings.TrimSpace(s)
}
