package main

import "practice/internal/runner"

func main() {
	runner.Run(map[string]func(){
		"grades":       demoGrades,
		"word-count":   demoWordCount,
		"bank-account": demoBankAccount,
	})
}
