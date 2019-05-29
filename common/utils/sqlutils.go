package utils

import (
	"strings"
)

func SqlBuild(data map[string]interface{}, sep string) (string, string, []interface{}) {
	ks := []string(nil)
	vs := []interface{}(nil)
	placeholder := "?,"
	for k, v := range data {
		ks = append(ks, k+sep)
		vs = append(vs, v)
	}
	return strings.Join(ks, ","),
		strings.Trim(strings.Repeat(placeholder, len(vs)), ","),
		vs
}
