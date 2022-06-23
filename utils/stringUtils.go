package utils

func ToSqlQuery(m map[string]string) string {
	var result string
	for i := 0; i < len(m); i++ {
		if i > 0 {
			result += ","
		}
	}
	return ""
}
