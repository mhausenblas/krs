package main

import (
	"flag"
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

var (
	releaseVersion string
	kubectlbin     string
	verbose        bool
)

func main() {
	// ns defines where resources should be captured:
	ns := flag.String("namespace", "default", "defines the namespace to capture")
	// targetresources defines what resources to capture:
	targetresources := flag.String("resources", "pods,svc,deploy", "defines the kind of resources to capture")
	// if we have an argument, we interpret it as the namespace:
	if len(os.Args) > 1 {
		if os.Args[1] == "version" {
			fmt.Printf("This is the Kubernetes Resource Stats (krs) tool in version %v\n", releaseVersion)
			fmt.Println("Usage: [KRS_KUBECTL_BIN=...] krs [namespace] [--resources=...]")
			os.Exit(0)
		}
	}
	// get params and env variables:
	flag.Parse()
	if kb := os.Getenv("KRS_KUBECTL_BIN"); kb != "" {
		kubectlbin = kb
	}
	if v := os.Getenv("KRS_VERBOSE"); v != "" {
		verbose = true
	}
	// populate the lookup table for supported resources
	initres()
	tres, err := parseres(*targetresources)
	if err != nil {
		log(err)
		os.Exit(1)
	}
	if verbose {
		info(fmt.Sprintf("Parsed targets: %v", tres))
	}
	// start main processing loop:
	for {
		// use kubectl to capture resources:
		allres := captures(*ns)
		// convert the string representation
		// of the JSON result from kubectl
		// into OpenMetrics lines:
		metrics := toOpenMetrics(*ns, allres, tres)
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
	res, err := kubecuddler.Kubectl(verbose, verbose, kubectlbin, "get", "--namespace="+namespace, "all", "--output=json")
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

func info(msg string) {
	_, _ = fmt.Fprintf(os.Stderr, "\x1b[92m%v\x1b[0m\n", msg)
}
