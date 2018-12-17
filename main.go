package main

import (
	"flag"
	"fmt"
)

func main() {
	versionFlag := flag.Bool("v", false, "display version info")
	flag.Parse()
	if *versionFlag {
		fmt.Println(green(LOGO))
		fmt.Printf(SOURCE, VERSION)
		return
	}

	fmt.Println(green(LOGO))
}

func green(str string) string {
	return fmt.Sprintf("\x1b[0;32m%s\x1b[0m", str)
}
