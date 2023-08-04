package main

import (
	_ "github.com/goharbor/xk6-harbor"
	_ "github.com/mstoykov/xk6-counter"
	k6cmd "go.k6.io/k6/cmd"
)

func main() {
	k6cmd.Execute()
}
