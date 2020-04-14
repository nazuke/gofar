package gofar

import "fmt"

// FindAspectRatio returns an aspect ratio string
func FindAspectRatio(width int, height int) string {

	var lcd int = findLowestCommonDenominator(width, height)
	var x int = width / lcd
	var y int = height / lcd

	var ar string = fmt.Sprintf("%d:%d", x, y)

	return ar
}

func findLowestCommonDenominator(a int, b int) int {

	if b == 0 {
		return a
	}

	var divisor int = a % b

	return (findLowestCommonDenominator(b, divisor))

}
