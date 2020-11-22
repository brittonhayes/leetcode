package lower

import (
	"strings"
)

// ToLowerCase converts a string from upper to lowercase
func ToLowerCase(str string) string {
	result := make(chan string)

	go func() {
		lower := strings.ToLower(str)
		result <- lower
	}()

	return <-result
}
