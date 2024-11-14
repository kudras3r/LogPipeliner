package main

// Here we parse logs from logsDir, convert it in accepted structs
// and send to apiPoint in json.
// In this case parse - send - collect chain implements in one place.
//
// But this chain implies communication.
// For it 'client' takes over the parse - send subchain.
// And 'server' need to deploy tenzir-node on himself, start it,
// run 'scripts/makeEndPoint.tql' pipeline. This pipeline creates an
// apiPoint that receives logs.
// After 'server' can run 'scripts/getLogs.tql' pipeline that collects
// logs from tenzir local storage and save them in json (in this case).

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/kudras3r/LogPipeliner/internal/logs/collect"
	"github.com/kudras3r/LogPipeliner/internal/logs/convert"
	"github.com/kudras3r/LogPipeliner/internal/logs/parse"
)

// To parse logs you need to take into account their format (dpkg)
const (
	logsDir = "/home/kud/Code/go/src/LogPipeliner/test_logs"
	// logsDir  = "/var/log"
	logsType = "dpkg"
	apiPoint = "http://0.0.0.0:4318/v1/logs"
)

func main() {
	logs, err := parse.Dpkg(logsDir)
	if err != nil {
		fmt.Printf("< ERROR > error when reading the dir : %s\n", err)
	}

	jsonData, err := convert.DpkgsToJson(logs)
	if err != nil {
		fmt.Printf("< ERROR > error when converting logs : %s\n", err)
	}

	resp, err := http.Post(apiPoint, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("< ERROR > error when sending logs : %s\n", err)
	}
	fmt.Printf("< INFO > status code : %d\n", resp.StatusCode)

	collect.RunCollector()
}
