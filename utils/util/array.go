package util

func ContainsForArrayString(target string, list []string) bool {
	for _, s := range list {
		if s == target {
			return true
		}
	}
	return false
}
