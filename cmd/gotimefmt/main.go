package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/yudppp/gotimefmt/timefmt"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Required arguments")
		fmt.Println("ex) gotimefmt YYYYMMDD HH:mm")
		os.Exit(1)
	}
	input := strings.Join(args, " ")
	timefmt := timefmt.Format(input)
	fmt.Println("------------------------------------")
	fmt.Println("format:", timefmt)
	fmt.Println("now:   ", time.Now().Format(timefmt))
	fmt.Println("------------------------------------")
	os.Exit(0)
}
