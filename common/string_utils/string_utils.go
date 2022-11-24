package string_utils

func IsEmpty(s *string) bool {
	if s == nil {
		return true
	}
	if len(*s) == 0 {
		return true
	}
	return false
}
