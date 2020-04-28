package misc

import (
	"fmt"
	"strings"
)

func ProgressBar(ch <-chan int, rows int) {
	fmt.Println()
	for i := 1; i <= rows; i++ {
		<-ch
		pct := 100 * float64(i) / float64(rows)
		filled := (80 * i) / rows
		bar := strings.Repeat("=", filled) + strings.Repeat("-", 80-filled)
		fmt.Printf("\r[%s] %.2f%%", bar, pct)
	}
	fmt.Println()
}
