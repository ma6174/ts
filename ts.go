package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var (
	layout   = "2006/01/02-15:04:05.000000"
	lastTime = time.Now()
)

func doFormat(line string) {
	fmtTime := time.Now().Format(layout)
	fmt.Printf("%s %s", fmtTime, line)
}

func main() {
	flag.Parse()
	if flag.NArg() > 0 {
		layout = flag.Arg(0)
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				if line != "" {
					doFormat(line)
				}
				break
			}
			log.Fatal(err)
		}
		doFormat(line)
	}
}
