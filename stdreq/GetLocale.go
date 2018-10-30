package stdreq

import (
	"github.com/goodsign/monday"
)

//GetLocale get locale for a language from monday package
func GetLocale(language string) monday.Locale {
	if language == "en" {
		return monday.LocaleEnUS
	}
	return monday.LocaleIdID
}
