package stdreq

// CheckLanguage check supported language
func CheckLanguage(language string) string {
	supportedLanguage := []string{"id", "en"}
	found := false
	for _, lang := range supportedLanguage {
		if lang == language {
			found = true
			break
		}
	}
	if !found {
		return "id"
	}
	return language
}
