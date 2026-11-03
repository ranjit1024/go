package main

import "concepts/internal/runner"

func main() {
	runner.Run(map[string]func(){
		"json":   demo_json,
		"xml":    demo_xml,
		"test":   epoch_demo,
		"time":   time_formatiing,
		"number": r_number,
		"parse":  parse_number,
		"url":    url_parsing,
	})
}
