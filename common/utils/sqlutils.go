package utils

import (
	"strings"
)

func SqlBuild(data map[string]interface{}, sep string) (string, string, []interface{}) {
	ks := []string(nil)
	vs := []interface{}(nil)
	placeholder := "?,"
	for k, v := range data {
		ks = append(ks, k)
		vs = append(vs, v)
	}
	return strings.Join(ks, sep),
		strings.Trim(strings.Repeat(placeholder, len(vs)), ","),
		vs
}
