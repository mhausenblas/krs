package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/mhausenblas/kubecuddler"
)

const (
	// ScrapeDelayInSec defines how long to wait between
	// sampling events, called the observation period:
	ScrapeDelayInSec = 5
)

func main() {
	var ns string
	// if we have an argument, we interpret it as the namespace:
	if len(os.Args) > 0 {
		ns = os.Args[1]
	}
	for {
		// use kubectl to capture resources:
		res := captures(ns)
		// convert the string representation
		// of the JSON result from kubectl
		// into OpenMetrics lines:
		metrics := toOpenMetrics(ns, res)
		// if we got something to report,
		// write it to stdout:
		if metrics != "" {
			store(os.Stdout, metrics)
		}
		time.Sleep(ScrapeDelayInSec * time.Second)
	}
}

// captures uses kubectl to query for resources
// and returns them as a JSON format list string.
func captures(namespace string) string {
	if namespace == "" {
		namespace = "default"
	}
	res, err := kubecuddler.Kubectl(false, false, "", "get", "--namespace="+namespace, "all", "--output=json")
	if err != nil {
		log(err)
	}
	return res
}

// store takes OpenMetrics lines as input and stores it in the target file
// which could be, for example, stdout
func store(target io.Writer, metrics string) {
	fmt.Fprintf(target, "%v", metrics)
}

func log(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "\x1b[91m%v\x1b[0m\n", err)
}
