package main

// K8sEvents represents a list of events,
// result of `kubectl events--output=json`
type K8sEvents struct {
	APIVersion string     `json:"apiVersion"`
	Kind       string     `json:"kind"`
	Items      []K8sEvent `json:"items"`
}

// K8sEvent is a single event entry
type K8sEvent struct {
	APIVersion        string         `json:"apiVersion"`
	Kind              string         `json:"kind"`
	FirstTS           string         `json:"firstTimestamp"`
	LastTS            string         `json:"lastTimestamp"`
	Message           string         `json:"message"`
	InvolvedObjectRef InvolvedObject `json:"involvedObject"`
}

// InvolvedObject is the target of the event,
// that is, what the event is about
type InvolvedObject struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Namespace  string `json:"namespace"`
	UID        string `json:"uid"`
}
