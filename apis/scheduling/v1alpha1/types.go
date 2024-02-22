/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MetricProviderType is a "string" type.
type MetricProviderType string

const (
	KubernetesMetricsServer MetricProviderType = "KubernetesMetricsServer"
	Prometheus              MetricProviderType = "Prometheus"
	SignalFx                MetricProviderType = "SignalFx"
)

// Denote the spec of the metric provider
type MetricProviderSpec struct {
	// Types of the metric provider
	Type MetricProviderType
	// The address of the metric provider
	Address string
	// The authentication token of the metric provider
	Token string
	// Whether to enable the InsureSkipVerify options for https requests on Metric Providers.
	InsecureSkipVerify bool
}

// TrimaranSpec holds common parameters for trimaran plugins
type TrimaranSpec struct {
	// Metric Provider to use when using load watcher as a library
	MetricProvider MetricProviderSpec
	// Address of load watcher service
	WatcherAddress string
}

// LowRiskOverCommitmentArgs holds arguments used to configure LowRiskOverCommitment plugin.
type LowRiskOverCommitmentArgs struct {
	metav1.TypeMeta

	// Common parameters for trimaran plugins
	TrimaranSpec
	// The number of windows over which usage data metrics are smoothed
	SmoothingWindowSize int64
	// Resources fractional weight of risk due to limits specification [0,1]
	RiskLimitWeights map[v1.ResourceName]float64
}