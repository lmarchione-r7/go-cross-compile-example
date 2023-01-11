// https://dev.to/tidalcloud/how-to-cross-compile-go-app-for-apple-silicon-m1-27l6
// hello.go

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf(
		"Hello world from %s/%s\n",
		runtime.GOOS,
		runtime.GOARCH,
	)
}
