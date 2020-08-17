package to_lower_case

func toLowerCase(str string) string {
	offset := 'a' - 'A'
	result := ""
	for _, r := range str {
		if r >= 'A' && r <= 'Z' {
			result += string(r + offset)
		} else {
			result += string(r)
		}
	}
	return result
}
