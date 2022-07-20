package main

import "fmt"

// Pass if test is a successful test and the name of the test
func getTestConclusion(isSuccess bool, testName string) string {
	// if TRUE
	if isSuccess {
		return fmt.Sprintf("✅\t%v\n", testName)
	}
	// if FALSE
	return fmt.Sprintf("❌\t%v\n", testName)
}

func getReason(err error) string {
	return fmt.Sprintf("❓REASON:\n%v", err)
}
