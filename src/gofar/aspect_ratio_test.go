package gofar

import (
	"testing"
)

func TestFindAspectRatio(t *testing.T) {

	var width int = 1920
	var height int = 1080
	var ar string = FindAspectRatio(width, height)

	if ar != "16:9" {
		t.Error(`FindAspectRatio did not produce the expected 16:9 aspect ration with 1920x1080`)
	}

}

func TestFindLowestCommonDenominator(t *testing.T) {

	var width int = 1920
	var height int = 1080
	var lcd int = findLowestCommonDenominator(width, height)

	if lcd != 120 {
		t.Error(`findLowestCommonDenominator did not produce the expected value of 120 for 1920x1080`)
	}

}
