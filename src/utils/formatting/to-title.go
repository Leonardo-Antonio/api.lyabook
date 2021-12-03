package formatting

import "strings"

func ToTitle(value string) string {
	return strings.Title(strings.ToLower(value))
}
