package main

import (
	"fmt"
	"runtime/debug"
)

var Ref string
var RefName string
var RefType string

func main() {
	fmt.Printf("Ref: %q\n", Ref)
	fmt.Printf("RefName: %q\n", RefName)
	fmt.Printf("RefType: %q\n", RefType)
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
