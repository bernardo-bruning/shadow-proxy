package filters

import "net/http"

func NopFilter(r *http.Request) bool {
	return false
}
