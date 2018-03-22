package primitives

import "fmt"

// Format float
func Ff(num interface{}) string {
	return fmt.Sprintf("%.1f", num)
}
