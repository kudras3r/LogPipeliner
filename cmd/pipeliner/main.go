package main

import (
	"fmt"

	"github.com/kudras3r/LogPipeliner/internal/logs/parse"
)

const (
	logsDir  = "/home/kud/Code/go/src/LogPipeliner/test_logs"
	logsType = "dpkg"
)

func main() {
	// TODO ->

	// init parser
	logs, err := parse.Dpkg(logsDir)
	if err != nil {
		fmt.Printf("< ERROR > error when reading the dir : %s\n", err)
	}
	for _, l := range logs {
		fmt.Println(l.Timestamp, l.Content)
	}

	// init sender

	// init collector
	// collect.RunCollector()

}
