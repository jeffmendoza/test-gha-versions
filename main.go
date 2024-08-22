package main

import (
	"fmt"
	"runtime/debug"
)

var Branch string
var Tag string
var Commit string
var Date string
var TS string
var State string

func main() {
	fmt.Printf("Branch: %q\n", Branch)
	fmt.Printf("Tag: %q\n", Tag)
	fmt.Printf("Commit: %q\n", Commit)
	fmt.Printf("Date: %q\n", Date)
	fmt.Printf("TS: %q\n", TS)
	fmt.Printf("State: %q\n", State)
	fmt.Println("---------------")
	if b, ok := debug.ReadBuildInfo(); ok {
		for _, s := range b.Settings {
			if s.Key == "vcs.revision" {
				fmt.Printf("Revision: %q\n", s.Value[:7])
			}
			if s.Key == "vcs.time" {
				fmt.Printf("Time: %q\n", s.Value)
			}
			if s.Key == "vcs.modified" {
				fmt.Printf("Modified: %q\n", s.Value)
			}
		}
	}
}
