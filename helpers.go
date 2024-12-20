package utils

// Contains проверяет наличие элемента в срезе строк
func Contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
