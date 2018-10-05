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
	for {
		events := fromFirehose("krs")
		metrics := toOpenMetrics(events)
		if metrics != "" {
			store(os.Stdout, metrics)
		}
		time.Sleep(ScrapeDelayInSec * time.Second)
	}

}

// fromFirehose uses kubectl to query for events
// and returns them in JSON format as a list.
// If the namespace param is non-empty then the
// specified namespace will be watched, otherwise
// cluster-wide.
func fromFirehose(namespace string) string {
	events, err := kubecuddler.Kubectl(false, false, "", "get", "--namespace="+namespace, "events", "--output=json")
	if err != nil {
		log(err)
	}
	return events
}

// store takes OpenMetrics lines as input and stores it in the target file
// which could be, for example, stdout
func store(target io.Writer, metrics string) {
	fmt.Fprintf(target, "%v", metrics)
}

func log(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "\x1b[91m%v\x1b[0m\n", err)
}
