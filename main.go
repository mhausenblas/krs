package main

import (
	"fmt"
	"io"
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
		events := fromFirehose("krs")
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
// and turns it into a sequence of OpenMetrics lines.
func toOpenMetrics(events string) string {
	labels := map[string]string{"namespace": "krs"}
	oml := omline("pod_count_all", "gauge", "Number of pods in any state (running, terminating, etc.)", "4", labels)
	return oml
}

// omline creates an OpenMetrics compliant line, for example:
// # HELP pod_count_all Number of pods in any state (running, terminating, etc.)
// # TYPE pod_count_all gauge
// pod_count_all{namespace="krs"} 4 1538675211
func omline(metric, mtype, mdesc, value string, labels map[string]string) (line string) {
	line = fmt.Sprintf("# HELP %v %v\n", metric, mdesc)
	line += fmt.Sprintf("# TYPE %v %v\n", metric, mtype)
	line += fmt.Sprintf("%v{", metric)
	for k, v := range labels {
		line += fmt.Sprintf("%v=\"%v\",", k, v)
	}
	line += fmt.Sprintf("} %v %v\n", value, time.Now().UnixNano())
	return
}

// store takes OpenMetrics lines as input and stores it in the target file
// which could be, for example, stdout
func store(target io.Writer, metrics string) {
	fmt.Printf("%v", metrics)
}

func log(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "\x1b[91m%v\x1b[0m\n", err)
}
