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
		events := fromFirehose()
		metrics := toOpenMetrics(events)
		store(os.Stdout, metrics)
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

// toOpenMetrics takes a JSON formatted kubectl result of a list of events
// and turns it into a sequence of OpenMetrics lines in the format:
//
// # HELP pod_count_all Number of pods in any state (running, terminating, etc.)
// # TYPE pod_count_all gauge
// pod_count_all{namespace="krs"} 4 1538675211
func toOpenMetrics(events string) string {
	return events
}

// store takes OpenMetrics lines as input and stores it in the target file
// which could be, for example, stdout
func store(target, metrics string) {
	fmt.Printf("%v", metrics)
}

func log(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "\x1b[91m%v\x1b[0m\n", err)
}
