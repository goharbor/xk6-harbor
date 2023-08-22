package main

import (
	_ "github.com/goharbor/xk6-harbor"
	_ "github.com/grafana/xk6-output-prometheus-remote"
	_ "github.com/mstoykov/xk6-counter"
	k6cmd "go.k6.io/k6/cmd"
)

func main() {
	k6cmd.Execute()
}
