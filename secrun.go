// SecRun
// by Sculptex
// Run bash script every (n) second(s)

package main

import (
	"fmt"
	"time"
	"flag"
	"os/exec"
	"os"
)

const VERSION = "0.0.1"
const DEFAULTFREQ = 1

func main() {

    var cmd string
    var version bool
    var freq int64
    var max int64

    flag.BoolVar(&version, "version", false, "Show Version info")
    flag.StringVar(&cmd, "cmd", cmd, "Bash Command To Run")
    flag.Int64Var(&freq, "freq", DEFAULTFREQ, "Frequency in Seconds")
    flag.Int64Var(&max, "max", 0, "Max Duration in Seconds, (0 infinite)")


    flag.Parse()

	if(version) {
		fmt.Printf("VERSION %s\n", VERSION)
		os.Exit(0)
	}

	go crun(cmd, freq)

	maxsecs , _:= time.ParseDuration(fmt.Sprintf("%ds", max))

	if(maxsecs > 0) {
		time.Sleep( maxsecs )
	} else {
		for {
			// NOWT
		}
	}
}

func crun(cmd string, freq int64) {
	var cmdarray []string
	var o []byte

	cmdarray = []string{
			"bash",
			cmd}
	head := cmdarray[0]
	parts := cmdarray[1:len(cmdarray)]

	freqsecs , _ := time.ParseDuration(fmt.Sprintf("%ds", freq))
	for range time.Tick( freqsecs ) {
		o , _ = exec.Command(head,parts...).Output()
		fmt.Println( string(o) )
	}
}
