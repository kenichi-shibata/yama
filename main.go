package main

import "./cmd"

var (
	VERSION = "0.0.2"
)
func main() {
	cmd.Execute(VERSION)
}
