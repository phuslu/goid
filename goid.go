// Package iploc provides fast way to get goroutine id.
//
//      package main
//
//      import (
//      	"fmt"
//      	"github.com/phuslu/goid"
//      )
//
//      func main() {
//      	fmt.Println(goid.Goid())
//      }
//
//      // Output: 1

//go:build amd64 || arm64 || arm || 386 || mipsle || riscv64
// +build amd64 arm64 arm 386 mipsle riscv64

package goid

func goid() int

// Goid returns the current goroutine id.
// It exactly matches goroutine id of the stack trace.
func Goid() int64 {
	return int64(goid())
}
