/*
Copyright 2015 The Kubernetes Authors.

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

package util

const (
	// LabelHostname is the label for hostname
	LabelHostname = "kubernetes.io/hostname"
	// LabelZoneFailureDomain is the label for zone failure domain
	LabelZoneFailureDomain = "failure-domain.beta.kubernetes.io/zone"
	// LabelMultiZoneDelimiter is the label for multi zone delimiter
	LabelMultiZoneDelimiter = "__"
	// LabelZoneRegion is the label for region failure domain
	LabelZoneRegion = "failure-domain.beta.kubernetes.io/region"

	// LabelInstanceType is the label for instance type
	LabelInstanceType = "beta.kubernetes.io/instance-type"

	// LabelOS is the label for OS
	LabelOS = "beta.kubernetes.io/os"
	// LabelArch is the label for arch
	LabelArch = "beta.kubernetes.io/arch"
)
