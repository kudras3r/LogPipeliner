package collect

// Every delay secs run pipeline that export data from tenzir storage
// and save it in out local storage.

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

const delayInSecs = 3

func RunCollector() {
	for {
		time.Sleep(delayInSecs * time.Second)

		fmt.Println("< INFO > Update storage")

		cmd := exec.Command("tenzir", "-f", "scripts/getLogs.tql")
		if err := cmd.Run(); err != nil {
			fmt.Printf("< ERROR > Cant start pipeline: %s\n", err)
			os.Exit(0)
		}
	}
}
