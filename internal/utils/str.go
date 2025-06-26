package utils

func StrOrEmpty(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
