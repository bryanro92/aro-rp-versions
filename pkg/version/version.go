package version

import (
	"fmt"
	"runtime"
)

var (
	Version = "v0.0.1"
)

func Print() {
	fmt.Printf("aro-rp-versions\nversion %v\ngo version: %v\n",
		Version, runtime.Version())
}
