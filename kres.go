package main

import (
	"fmt"
	"strings"
)

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

const (
	Pod        = "Pod"
	Deployment = "Deployment"
	Service    = "Service"
)

var (
	supportedres map[string]string
)

// initres sets the supported resources
func initres() {
	supportedres = map[string]string{
		"pods":        Pod,
		"deployments": Deployment,
		"services":    Service,
	}
}

// isvalidres checks if a given resource is supported
func isvalidres(resource string) bool {
	_, ok := supportedres[resource]
	return ok
}

// listres outputs supported resources
func listres() (res string) {
	for k := range supportedres {
		res += fmt.Sprintf("%v, ", k)
	}
	res = strings.TrimSuffix(res, "'")
	return
}

// parseres checks if we're dealing with a valid resource targets string
// and if so, extracts the potentially comma-separated list of resource(s)
// we're want to track. For example, this is valid: 'pods,services'.
// Note that unsupported ones will be silently dropped.
func parseres(targets string) (tresources []string, err error) {
	if !strings.Contains(targets, ",") {
		if isvalidres(targets) {
			return []string{supportedres[targets]}, nil
		}
		return []string{}, fmt.Errorf("%v is not supported, valid ones are: %v", targets, listres())
	}
	rawtres := strings.Split(targets, ",")
	for _, tres := range rawtres {
		if isvalidres(tres) {
			tresources = append(tresources, supportedres[targets])
		}
	}
	if len(tresources) == 0 {
		return []string{}, fmt.Errorf("No supported resources found, valid ones are: %v", targets, listres())
	}
	return
}
