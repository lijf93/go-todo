package main

import (
	"flag"
	"fmt"
)

const (
	IconGood = "✔"
	IconBad  = "✗"
)

var version = "v0.1.0"

func main() {
	versionFlag := flag.Bool("v", false, "display version info")
	flag.Parse()
	if *versionFlag {
		fmt.Println(green(logo))
		fmt.Printf(source, version)
		return
	}

	fmt.Println(green(logo))
}

func green(str string) string {
	return fmt.Sprintf("\x1b[0;32m%s\x1b[0m", str)
}
