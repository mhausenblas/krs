package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// toOpenMetrics takes the result of a `kubectl get events` as a
// JSON formatted string as an input and turns it into a
// sequence of OpenMetrics lines.
func toOpenMetrics(rawevents string) string {
	events := K8sEvents{}
	err := json.Unmarshal([]byte(rawevents), &events)
	if err != nil {
		log(err)
	}
	var oml string
	for _, event := range events.Items {
		if event.InvolvedObjectRef.Kind == "Pod" {
			labels := map[string]string{"name": event.InvolvedObjectRef.Name, "namespace": event.InvolvedObjectRef.Namespace}
			oml += ometricsline("pod_count_all", "gauge", "Number of pods in any state (running, terminating, etc.)", "1", labels)
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
	line += fmt.Sprintf("%v{", metric)
	for k, v := range labels {
		line += fmt.Sprintf("%v=\"%v\"", k, v)
		line += ","
	}
	// make sure that we get rid of trialing comma:
	line = strings.TrimSuffix(line, ",")

	line += fmt.Sprintf("} %v %v\n", value, time.Now().UnixNano())
	return
}
