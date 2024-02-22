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

// +k8s:defaulter-gen=true

package v1beta1

import (
	//"strconv"

	v1 "k8s.io/api/core/v1"
	// "k8s.io/apimachinery/pkg/api/resource"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schedulerconfig "k8s.io/kube-scheduler/config/v1"

	// pluginConfig "sigs.k8s.io/scheduler-plugins/apis/config"
)

var (

	// Defaults for LowRiskOverCommitment plugin

	// The default number of windows over which usage data metrics are smoothed.
	// DefaultSmoothingWindowSize is 5 (used by Prometheus)
	DefaultSmoothingWindowSize int64 = 5
	// The default weight of risk due to limit for a resource
	DefaultRiskLimitWeight float64 = 0.5
	// Resources fractional weight of risk due to limits specification [0,1]
	DefaultRiskLimitWeights = map[v1.ResourceName]float64{
		v1.ResourceCPU:    DefaultRiskLimitWeight,
		v1.ResourceMemory: DefaultRiskLimitWeight,
	}

	// DefaultMetricProviderType is the Kubernetes metrics server
	DefaultMetricProviderType = KubernetesMetricsServer

	defaultResourceSpec = []schedulerconfig.ResourceSpec{
		{Name: string(v1.ResourceCPU), Weight: 1},
		{Name: string(v1.ResourceMemory), Weight: 1},
	}
)

// SetDefaultTrimaranSpec sets the default parameters for common Trimaran plugins
func SetDefaultTrimaranSpec(args *TrimaranSpec) {
	if args.WatcherAddress == nil && args.MetricProvider.Type == "" {
		args.MetricProvider.Type = DefaultMetricProviderType
	}
}

// SetDefaults_LowRiskOverCommitmentArgs sets the default parameters for LowRiskOverCommitment plugin
func SetDefaults_LowRiskOverCommitmentArgs(args *LowRiskOverCommitmentArgs) {
	SetDefaultTrimaranSpec(&args.TrimaranSpec)
	if args.SmoothingWindowSize == nil || *args.SmoothingWindowSize <= 0 {
		args.SmoothingWindowSize = &DefaultSmoothingWindowSize
	}
	if args.RiskLimitWeights == nil || len(args.RiskLimitWeights) == 0 {
		args.RiskLimitWeights = DefaultRiskLimitWeights
	} else {
		for r, w := range args.RiskLimitWeights {
			if w < 0 || w > 1 {
				(args.RiskLimitWeights)[r] = DefaultRiskLimitWeight
			}
		}
	}
}
