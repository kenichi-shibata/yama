package main

import "github.com/kenichi-shibata/yama/cmd"

var (
	VERSION = "0.0.2"
)
func main() {
	cmd.Execute(VERSION)
}
