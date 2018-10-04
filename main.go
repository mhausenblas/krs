package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mhausenblas/kubecuddler"
)

const (

	// ScrapeDelayInSec defines how long to wait between sampling events:
	ScrapeDelayInSec = 5
)

func main() {
	for {
		res, err := kubecuddler.Kubectl(false, false, "", "get", "events", "--all-namespaces", "--output=json")
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "\x1b[91m%v\x1b[0m\n", err)
		}
		fmt.Printf("%v", res)
		time.Sleep(ScrapeDelayInSec * time.Second)
	}

}
