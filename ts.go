package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

var (
	layout        = "2006/01/02-15:04:05.000000"
	flagFromStart = flag.Bool("s", false, "calc time from start")
	flagFromLast  = flag.Bool("i", false, "calc time from last line")
	isShowDoc     = flag.Bool("h", false, "show help")
	lastTime      = time.Now()
	T             = map[rune]int64{
		's': 1e9,
		'm': 1e9 * 60,
		'h': 1e9 * 60 * 60,
		'd': 1e9 * 60 * 60 * 24,
	}
)

const doc = `
NAME
       ts - timestamp input

SYNOPSIS
       ts [-i | -s] [format]

DESCRIPTION
       ts adds a timestamp to the beginning of each line of input.

       The optional format parameter controls how the timestamp is formatted, as used by go time format. The default format is "2006/01/02-15:04:05.000000". 

       If the -i or -s switch is passed, ts timestamps incrementally instead. In case of -i, every timestamp will be the time elapsed since the last timestamp.
       In case of -s, the time elapsed since start of the program is used.  The default format changes to "{ts}", it will show time used like "1d20h30m40.555555s".
`

func calcTime(cost int64) (ret string) {
	for _, r := range "dhm" {
		if t := cost / T[rune(r)]; t > 0 {
			ret += fmt.Sprintf("%d%c", t, r)
			cost -= t * T[rune(r)]
		}
	}
	ret += fmt.Sprintf("%.6fs", float64(cost)/float64(T['s']))
	ret = strings.Replace(layout, "{ts}", ret, 1)
	return
}

func doFormat(line string) {
	var fmtTime string
	if *flagFromStart {
		fmtTime = calcTime(time.Now().Sub(lastTime).Nanoseconds())
	} else if *flagFromLast {
		now := time.Now()
		fmtTime = calcTime(now.Sub(lastTime).Nanoseconds())
		lastTime = now
	} else {
		fmtTime = time.Now().Format(layout)
	}
	fmt.Printf("%s %s", fmtTime, line)
}

func main() {
	flag.Parse()
	if *isShowDoc {
		fmt.Println(doc)
		return
	}
	if *flagFromStart || *flagFromLast {
		layout = "{ts}"
	}
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
