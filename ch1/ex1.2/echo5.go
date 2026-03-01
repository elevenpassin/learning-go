package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	start := time.Now()
	s := ""

	for i, arg := range os.Args {
		s += fmt.Sprintf("%d %s\n", i, arg)
	}

	fmt.Print(s)

	fmt.Printf("%.4ds elapsed\n", time.Since(start).Microseconds())
}
