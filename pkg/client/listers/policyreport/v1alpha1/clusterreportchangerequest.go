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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kyverno/kyverno/pkg/api/policyreport/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterReportChangeRequestLister helps list ClusterReportChangeRequests.
type ClusterReportChangeRequestLister interface {
	// List lists all ClusterReportChangeRequests in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterReportChangeRequest, err error)
	// Get retrieves the ClusterReportChangeRequest from the index for a given name.
	Get(name string) (*v1alpha1.ClusterReportChangeRequest, error)
	ClusterReportChangeRequestListerExpansion
}

// clusterReportChangeRequestLister implements the ClusterReportChangeRequestLister interface.
type clusterReportChangeRequestLister struct {
	indexer cache.Indexer
}

// NewClusterReportChangeRequestLister returns a new ClusterReportChangeRequestLister.
func NewClusterReportChangeRequestLister(indexer cache.Indexer) ClusterReportChangeRequestLister {
	return &clusterReportChangeRequestLister{indexer: indexer}
}

// List lists all ClusterReportChangeRequests in the indexer.
func (s *clusterReportChangeRequestLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterReportChangeRequest, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterReportChangeRequest))
	})
	return ret, err
}

// Get retrieves the ClusterReportChangeRequest from the index for a given name.
func (s *clusterReportChangeRequestLister) Get(name string) (*v1alpha1.ClusterReportChangeRequest, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clusterreportchangerequest"), name)
	}
	return obj.(*v1alpha1.ClusterReportChangeRequest), nil
}
