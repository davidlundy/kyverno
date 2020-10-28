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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/kyverno/kyverno/pkg/client/clientset/versioned/typed/policyreport/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakePolicyV1alpha1 struct {
	*testing.Fake
}

func (c *FakePolicyV1alpha1) ClusterPolicyReports() v1alpha1.ClusterPolicyReportInterface {
	return &FakeClusterPolicyReports{c}
}

func (c *FakePolicyV1alpha1) ClusterReportChangeRequests() v1alpha1.ClusterReportChangeRequestInterface {
	return &FakeClusterReportChangeRequests{c}
}

func (c *FakePolicyV1alpha1) PolicyReports(namespace string) v1alpha1.PolicyReportInterface {
	return &FakePolicyReports{c, namespace}
}

func (c *FakePolicyV1alpha1) ReportChangeRequests(namespace string) v1alpha1.ReportChangeRequestInterface {
	return &FakeReportChangeRequests{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakePolicyV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
