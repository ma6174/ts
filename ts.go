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
	layoutCurrent = "2006/01/02-15:04:05.000000"
	flagFromStart = flag.Bool("s", false, "calc time from start")
	flagFromLast  = flag.Bool("i", false, "calc time from last line")
	lastTime      = time.Now()
)

func doFormat(line string) {
	var fmtTime string
	if *flagFromStart {
		cost := time.Now().Sub(lastTime).Seconds()
		fmtTime = fmt.Sprintf("%.6fs", cost)
	} else if *flagFromLast {
		now := time.Now()
		cost := now.Sub(lastTime).Seconds()
		fmtTime = fmt.Sprintf("%.6fs", cost)
		lastTime = now
	} else {
		fmtTime = time.Now().Format(layoutCurrent)
	}
	fmt.Printf("%s %s", fmtTime, line)
}

func main() {
	flag.Parse()
	if flag.NArg() > 0 {
		layoutCurrent = flag.Arg(0)
	}
	reader := bufio.NewReaderSize(os.Stdin, 0)
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
