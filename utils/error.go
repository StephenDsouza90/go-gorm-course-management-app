package utils

import "fmt"

// CheckErr : ...
func CheckErr(err error) {
	if err != nil {
		fmt.Printf(err.Error())
	}
}
