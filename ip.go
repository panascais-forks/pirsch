package pirsch

import (
	"net"
	"net/http"
	"strings"
)

// Headers and corresponding parser to look up the real client IP.
// They will be check in order, the first non-empty one will be picked,
// or else the remote address is selected.
// CF-Connecting-IP is a header added by Cloudflare: https://support.cloudflare.com/hc/en-us/articles/206776727-What-is-True-Client-IP-
var ipHeaders = []ipHeader{
	{"CF-Connecting-IP", parseXForwardedForHeader},
	{"True-Client-IP", parseXForwardedForHeader},
	{"X-Forwarded-For", parseXForwardedForHeader},
	{"Forwarded", parseForwardedHeader},
	{"X-Real-IP", parseXRealIPHeader},
}

type ipHeader struct {
	header string
	parser func(string) string
}

// getIP returns the IP from given request.
// It will try to extract the real client IP from headers if possible.
func getIP(r *http.Request) string {
	ip := r.RemoteAddr

	for _, header := range ipHeaders {
		value := r.Header.Get(header.header)

		if value != "" {
			parsedIP := header.parser(value)

			if parsedIP != "" {
				ip = parsedIP
				break
			}
		}
	}

	if strings.Contains(ip, ":") {
		host, _, err := net.SplitHostPort(ip)

		if err != nil {
			return ip
		}

		return host
	}

	return ip
}

func parseForwardedHeader(value string) string {
	left, _, _ := strings.Cut(value, ",")
	parts := strings.Split(left, ";")

	for _, part := range parts {
		k, v, found := strings.Cut(part, "=")

		if found && k == "for" {
			return v
		}
	}

	return ""
}

func parseXForwardedForHeader(value string) string {
	left, _, _ := strings.Cut(value, ",")
	return strings.TrimSpace(left)
}

func parseXRealIPHeader(value string) string {
	return value
}
