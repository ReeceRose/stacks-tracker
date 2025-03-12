package main

import "fmt"

var (
	Version        = "dev"
	CommitHash     = "n/a"
	BuildTimestamp = "n/a"
)

func main() {
	fmt.Println("Version:", Version)
	fmt.Println("Commit Hash:", CommitHash)
	fmt.Println("Build Timestamp:", BuildTimestamp)
}
