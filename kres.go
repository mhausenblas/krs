package main

// K8sList represents a list of Kubernetes
// resources, for example, as a result of executing:
// `kubectl get all --output=json`
type K8sList struct {
	APIVersion string        `json:"apiVersion"`
	Kind       string        `json:"kind"`
	Items      []K8sResource `json:"items"`
}

// K8sResource represents a single Kubernetes resource
type K8sResource struct {
	APIVersion string  `json:"apiVersion"`
	Kind       string  `json:"kind"`
	Meta       K8SMeta `json:"metadata"`
}

// K8SMeta represents the metadata, common to all Kubernetes resources
type K8SMeta struct {
	Created   string `json:"creationTimestamp"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	UID       string `json:"uid"`
}
