/*
Copyright 2023.

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

package v1

import (
	v1 "github.com/apache/doris-operator/api/doris/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DorisClusterLister helps list DorisClusters.
// All objects returned here must be treated as read-only.
type DorisClusterLister interface {
	// List lists all DorisClusters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.DorisCluster, err error)
	// DorisClusters returns an object that can list and get DorisClusters.
	DorisClusters(namespace string) DorisClusterNamespaceLister
	DorisClusterListerExpansion
}

// dorisClusterLister implements the DorisClusterLister interface.
type dorisClusterLister struct {
	indexer cache.Indexer
}

// NewDorisClusterLister returns a new DorisClusterLister.
func NewDorisClusterLister(indexer cache.Indexer) DorisClusterLister {
	return &dorisClusterLister{indexer: indexer}
}

// List lists all DorisClusters in the indexer.
func (s *dorisClusterLister) List(selector labels.Selector) (ret []*v1.DorisCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.DorisCluster))
	})
	return ret, err
}

// DorisClusters returns an object that can list and get DorisClusters.
func (s *dorisClusterLister) DorisClusters(namespace string) DorisClusterNamespaceLister {
	return dorisClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DorisClusterNamespaceLister helps list and get DorisClusters.
// All objects returned here must be treated as read-only.
type DorisClusterNamespaceLister interface {
	// List lists all DorisClusters in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.DorisCluster, err error)
	// Get retrieves the DorisCluster from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.DorisCluster, error)
	DorisClusterNamespaceListerExpansion
}

// dorisClusterNamespaceLister implements the DorisClusterNamespaceLister
// interface.
type dorisClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DorisClusters in the indexer for a given namespace.
func (s dorisClusterNamespaceLister) List(selector labels.Selector) (ret []*v1.DorisCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.DorisCluster))
	})
	return ret, err
}

// Get retrieves the DorisCluster from the indexer for a given namespace and name.
func (s dorisClusterNamespaceLister) Get(name string) (*v1.DorisCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("doriscluster"), name)
	}
	return obj.(*v1.DorisCluster), nil
}
