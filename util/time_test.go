package util

import "testing"

func TestStringToTime(testing *testing.T) {
	var convertedTime = StringToTime("2006-01-02T15:04:05")

	if convertedTime.Year() != 2006 {
		testing.Errorf("Espera que o ano seja 2006")
	}

	if convertedTime.Month() != 1 {
		testing.Errorf("Espera que o mes seja 1")
	}

	if convertedTime.Day() != 2 {
		testing.Errorf("Espera que o dia seja 2")
	}

	if convertedTime.Hour() != 15 {
		testing.Errorf("Espera que a hora seja 15")
	}

	if convertedTime.Minute() != 4 {
		testing.Errorf("Espera que o minuto seja 4")
	}

	if convertedTime.Second() != 5 {
		testing.Errorf("Espera que o segundo seja 5")
	}
}
