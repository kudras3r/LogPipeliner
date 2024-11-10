package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
)

func main() {
	// TODO -> ? logger ?

	// start node + pipeline
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		cmd := exec.Command("tenzir-node")
		if err := cmd.Run(); err != nil {
			fmt.Printf("< ERROR > cant start node: %s\n", err)
			os.Exit(0)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		cmd := exec.Command("tenzir", "-f", "pipeline.tql")
		if err := cmd.Run(); err != nil {
			fmt.Printf("< ERROR > cant start pipeline: %s\n", err)
			os.Exit(0)
		}
	}()

	fmt.Print("< INFO > node; pipeline are up\n")

	// init CLI (only requests the path to the dir with logs)
	// init validator (check .log extension)
	// init sender (send logs to api)

	wg.Wait()
}
