package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// namespaceStats holds the stats for
// a specific namespace across the tracked
// resources, each key represents a resource
// kind, such as Pod, Deployment, etc.
type namespaceStats struct {
	Resources map[string]*resourceMetric
	Namespace string
}

// resourceMetric holds the number of
// resources in a namespaces of a kind
// over the observation period.
type resourceMetric struct {
	Number int
}

func initStats(namespace string) (ns namespaceStats) {
	ns = namespaceStats{
		Namespace: namespace,
	}
	ns.Resources = map[string]*resourceMetric{}
	for _, r := range supportedres {
		ns.Resources[r] = &resourceMetric{Number: 0}
	}
	return
}

// toOpenMetrics takes the result of a `kubectl get events` as a
// JSON formatted string as an input and turns it into a
// sequence of OpenMetrics lines.
func toOpenMetrics(namespace, rawkres string) string {
	kres := K8sList{}
	err := json.Unmarshal([]byte(rawkres), &kres)
	if err != nil {
		log(err)
	}
	if len(kres.Items) == 0 {
		return ""
	}
	// set up list of supported resources:
	nsstats := initStats(namespace)
	// gather stats:
	for _, kr := range kres.Items {
		if isvalidkind(kr.Kind) {
			nsstats.Resources[kr.Kind].Number++
		}
	}
	// serialize in OpenMetrics format
	var oml string
	for reskind, val := range nsstats.Resources {
		labels := map[string]string{"namespace": nsstats.Namespace}
		switch reskind {
		case Pod:
			oml += ometricsline("pods",
				"gauge",
				"Number of pods in any state, for example running",
				fmt.Sprintf("%v", val.Number),
				labels)
		case Deployment:
			oml += ometricsline("deployments",
				"gauge",
				"Number of deployments",
				fmt.Sprintf("%v", val.Number),
				labels)
		case Service:
			oml += ometricsline("services",
				"gauge",
				"Number of services",
				fmt.Sprintf("%v", val.Number),
				labels)
		}
	}
	return oml
}

// ometricsline creates an OpenMetrics compliant line, for example:
// # HELP pod_count_all Number of pods in any state (running, terminating, etc.)
// # TYPE pod_count_all gauge
// pod_count_all{namespace="krs"} 4 1538675211
func ometricsline(metric, mtype, mdesc, value string, labels map[string]string) (line string) {
	line = fmt.Sprintf("# HELP %v %v\n", metric, mdesc)
	line += fmt.Sprintf("# TYPE %v %v\n", metric, mtype)
	// add labels:
	line += fmt.Sprintf("%v{", metric)
	for k, v := range labels {
		line += fmt.Sprintf("%v=\"%v\"", k, v)
		line += ","
	}
	// make sure that we get rid of trialing comma:
	line = strings.TrimSuffix(line, ",")
	// now add value and we're done:
	line += fmt.Sprintf("} %v\n", value)
	return
}
