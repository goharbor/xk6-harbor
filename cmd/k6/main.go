package main

import (
	k6cmd "github.com/loadimpact/k6/cmd"

	_ "github.com/heww/xk6-harbor"
	_ "github.com/mstoykov/xk6-counter"
)

func main() {
	k6cmd.Execute()
}
