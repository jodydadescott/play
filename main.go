package main

import (
	"fmt"
	"runtime/debug"
)

func main() {

	commit := func() string {
		if info, ok := debug.ReadBuildInfo(); ok {
			for _, setting := range info.Settings {
				if setting.Key == "vcs.revision" {
					return setting.Value
				}
			}
		}

		return ""
	}()

	fmt.Println(commit)

}
