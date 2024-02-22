// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1beta1

import (
	unsafe "unsafe"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	// configv1 "k8s.io/kube-scheduler/config/v1"
	// config "sigs.k8s.io/scheduler-plugins/pkg/apis/config"
	config "sigs.k8s.io/scheduler-plugins/apis/config"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*LowRiskOverCommitmentArgs)(nil), (*config.LowRiskOverCommitmentArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_LowRiskOverCommitmentArgs_To_config_LowRiskOverCommitmentArgs(a.(*LowRiskOverCommitmentArgs), b.(*config.LowRiskOverCommitmentArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.LowRiskOverCommitmentArgs)(nil), (*LowRiskOverCommitmentArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_LowRiskOverCommitmentArgs_To_v1_LowRiskOverCommitmentArgs(a.(*config.LowRiskOverCommitmentArgs), b.(*LowRiskOverCommitmentArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MetricProviderSpec)(nil), (*config.MetricProviderSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_MetricProviderSpec_To_config_MetricProviderSpec(a.(*MetricProviderSpec), b.(*config.MetricProviderSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.MetricProviderSpec)(nil), (*MetricProviderSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_MetricProviderSpec_To_v1_MetricProviderSpec(a.(*config.MetricProviderSpec), b.(*MetricProviderSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*TrimaranSpec)(nil), (*config.TrimaranSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_TrimaranSpec_To_config_TrimaranSpec(a.(*TrimaranSpec), b.(*config.TrimaranSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.TrimaranSpec)(nil), (*TrimaranSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_TrimaranSpec_To_v1_TrimaranSpec(a.(*config.TrimaranSpec), b.(*TrimaranSpec), scope)
	}); err != nil {
		return err
	}
	return nil
}



func autoConvert_v1_LowRiskOverCommitmentArgs_To_config_LowRiskOverCommitmentArgs(in *LowRiskOverCommitmentArgs, out *config.LowRiskOverCommitmentArgs, s conversion.Scope) error {
	if err := Convert_v1_TrimaranSpec_To_config_TrimaranSpec(&in.TrimaranSpec, &out.TrimaranSpec, s); err != nil {
		return err
	}
	if err := metav1.Convert_Pointer_int64_To_int64(&in.SmoothingWindowSize, &out.SmoothingWindowSize, s); err != nil {
		return err
	}
	out.RiskLimitWeights = *(*map[corev1.ResourceName]float64)(unsafe.Pointer(&in.RiskLimitWeights))
	return nil
}

// Convert_v1_LowRiskOverCommitmentArgs_To_config_LowRiskOverCommitmentArgs is an autogenerated conversion function.
func Convert_v1_LowRiskOverCommitmentArgs_To_config_LowRiskOverCommitmentArgs(in *LowRiskOverCommitmentArgs, out *config.LowRiskOverCommitmentArgs, s conversion.Scope) error {
	return autoConvert_v1_LowRiskOverCommitmentArgs_To_config_LowRiskOverCommitmentArgs(in, out, s)
}

func autoConvert_config_LowRiskOverCommitmentArgs_To_v1_LowRiskOverCommitmentArgs(in *config.LowRiskOverCommitmentArgs, out *LowRiskOverCommitmentArgs, s conversion.Scope) error {
	if err := Convert_config_TrimaranSpec_To_v1_TrimaranSpec(&in.TrimaranSpec, &out.TrimaranSpec, s); err != nil {
		return err
	}
	if err := metav1.Convert_int64_To_Pointer_int64(&in.SmoothingWindowSize, &out.SmoothingWindowSize, s); err != nil {
		return err
	}
	out.RiskLimitWeights = *(*map[corev1.ResourceName]float64)(unsafe.Pointer(&in.RiskLimitWeights))
	return nil
}

// Convert_config_LowRiskOverCommitmentArgs_To_v1_LowRiskOverCommitmentArgs is an autogenerated conversion function.
func Convert_config_LowRiskOverCommitmentArgs_To_v1_LowRiskOverCommitmentArgs(in *config.LowRiskOverCommitmentArgs, out *LowRiskOverCommitmentArgs, s conversion.Scope) error {
	return autoConvert_config_LowRiskOverCommitmentArgs_To_v1_LowRiskOverCommitmentArgs(in, out, s)
}

func autoConvert_v1_MetricProviderSpec_To_config_MetricProviderSpec(in *MetricProviderSpec, out *config.MetricProviderSpec, s conversion.Scope) error {
	out.Type = config.MetricProviderType(in.Type)
	if err := metav1.Convert_Pointer_string_To_string(&in.Address, &out.Address, s); err != nil {
		return err
	}
	if err := metav1.Convert_Pointer_string_To_string(&in.Token, &out.Token, s); err != nil {
		return err
	}
	if err := metav1.Convert_Pointer_bool_To_bool(&in.InsecureSkipVerify, &out.InsecureSkipVerify, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_MetricProviderSpec_To_config_MetricProviderSpec is an autogenerated conversion function.
func Convert_v1_MetricProviderSpec_To_config_MetricProviderSpec(in *MetricProviderSpec, out *config.MetricProviderSpec, s conversion.Scope) error {
	return autoConvert_v1_MetricProviderSpec_To_config_MetricProviderSpec(in, out, s)
}

func autoConvert_config_MetricProviderSpec_To_v1_MetricProviderSpec(in *config.MetricProviderSpec, out *MetricProviderSpec, s conversion.Scope) error {
	out.Type = MetricProviderType(in.Type)
	if err := metav1.Convert_string_To_Pointer_string(&in.Address, &out.Address, s); err != nil {
		return err
	}
	if err := metav1.Convert_string_To_Pointer_string(&in.Token, &out.Token, s); err != nil {
		return err
	}
	if err := metav1.Convert_bool_To_Pointer_bool(&in.InsecureSkipVerify, &out.InsecureSkipVerify, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_MetricProviderSpec_To_v1_MetricProviderSpec is an autogenerated conversion function.
func Convert_config_MetricProviderSpec_To_v1_MetricProviderSpec(in *config.MetricProviderSpec, out *MetricProviderSpec, s conversion.Scope) error {
	return autoConvert_config_MetricProviderSpec_To_v1_MetricProviderSpec(in, out, s)
}

func autoConvert_v1_TrimaranSpec_To_config_TrimaranSpec(in *TrimaranSpec, out *config.TrimaranSpec, s conversion.Scope) error {
	if err := Convert_v1_MetricProviderSpec_To_config_MetricProviderSpec(&in.MetricProvider, &out.MetricProvider, s); err != nil {
		return err
	}
	if err := metav1.Convert_Pointer_string_To_string(&in.WatcherAddress, &out.WatcherAddress, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_TrimaranSpec_To_config_TrimaranSpec is an autogenerated conversion function.
func Convert_v1_TrimaranSpec_To_config_TrimaranSpec(in *TrimaranSpec, out *config.TrimaranSpec, s conversion.Scope) error {
	return autoConvert_v1_TrimaranSpec_To_config_TrimaranSpec(in, out, s)
}

func autoConvert_config_TrimaranSpec_To_v1_TrimaranSpec(in *config.TrimaranSpec, out *TrimaranSpec, s conversion.Scope) error {
	if err := Convert_config_MetricProviderSpec_To_v1_MetricProviderSpec(&in.MetricProvider, &out.MetricProvider, s); err != nil {
		return err
	}
	if err := metav1.Convert_string_To_Pointer_string(&in.WatcherAddress, &out.WatcherAddress, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_TrimaranSpec_To_v1_TrimaranSpec is an autogenerated conversion function.
func Convert_config_TrimaranSpec_To_v1_TrimaranSpec(in *config.TrimaranSpec, out *TrimaranSpec, s conversion.Scope) error {
	return autoConvert_config_TrimaranSpec_To_v1_TrimaranSpec(in, out, s)
}
