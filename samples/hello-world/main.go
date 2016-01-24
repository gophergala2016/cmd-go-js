package main

import (
	"fmt"
	"go/build"
	"runtime"
)

func main() {
	fmt.Printf("Hello brave new world! It is working on %v %v/%v!", runtime.Version(), build.Default.GOOS, build.Default.GOARCH)
	if build.Default.GOARCH == "js" {
		fmt.Print(" That means you can execute it in browsers.")
	}
	fmt.Println()
}
