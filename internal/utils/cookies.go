package utils

import (
	"strings"
)

func ParseCookies(cookieString string) map[string]string {
	cookies := make(map[string]string)
	pairs := strings.Split(cookieString, ";")

	for _, pair := range pairs {
		pair = strings.TrimSpace(pair)
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) == 2 {
			cookies[kv[0]] = kv[1]
		}
	}

	return cookies
}
