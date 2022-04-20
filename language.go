package pirsch

import (
	iso6391 "github.com/emvi/iso-639-1"
	"net/http"
	"strings"
)

func getLanguage(r *http.Request) string {
	lang := r.Header.Get("Accept-Language")

	if lang != "" {
		left, _, _ := strings.Cut(lang, ";")
		left, _, _ = strings.Cut(left, ",")
		left, _, _ = strings.Cut(left, "-")
		code := strings.ToLower(strings.TrimSpace(left))

		if iso6391.ValidCode(code) {
			return code
		}
	}

	return ""
}
