package main

import "fmt"

func JoinFloat(values []float64, sep string) string {
	output := ""
	for i, value := range values {
		if i == 0 {
			output = fmt.Sprintf("%.4f", value)
		} else {
			output = fmt.Sprintf("%s%s%.4f", output, sep, value)
		}
	}
	return output
}
