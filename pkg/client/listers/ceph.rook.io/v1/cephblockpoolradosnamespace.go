/*
Copyright 2018 The Rook Authors. All rights reserved.

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
	v1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// CephBlockPoolRadosNamespaceLister helps list CephBlockPoolRadosNamespaces.
// All objects returned here must be treated as read-only.
type CephBlockPoolRadosNamespaceLister interface {
	// List lists all CephBlockPoolRadosNamespaces in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.CephBlockPoolRadosNamespace, err error)
	// CephBlockPoolRadosNamespaces returns an object that can list and get CephBlockPoolRadosNamespaces.
	CephBlockPoolRadosNamespaces(namespace string) CephBlockPoolRadosNamespaceNamespaceLister
	CephBlockPoolRadosNamespaceListerExpansion
}

// cephBlockPoolRadosNamespaceLister implements the CephBlockPoolRadosNamespaceLister interface.
type cephBlockPoolRadosNamespaceLister struct {
	listers.ResourceIndexer[*v1.CephBlockPoolRadosNamespace]
}

// NewCephBlockPoolRadosNamespaceLister returns a new CephBlockPoolRadosNamespaceLister.
func NewCephBlockPoolRadosNamespaceLister(indexer cache.Indexer) CephBlockPoolRadosNamespaceLister {
	return &cephBlockPoolRadosNamespaceLister{listers.New[*v1.CephBlockPoolRadosNamespace](indexer, v1.Resource("cephblockpoolradosnamespace"))}
}

// CephBlockPoolRadosNamespaces returns an object that can list and get CephBlockPoolRadosNamespaces.
func (s *cephBlockPoolRadosNamespaceLister) CephBlockPoolRadosNamespaces(namespace string) CephBlockPoolRadosNamespaceNamespaceLister {
	return cephBlockPoolRadosNamespaceNamespaceLister{listers.NewNamespaced[*v1.CephBlockPoolRadosNamespace](s.ResourceIndexer, namespace)}
}

// CephBlockPoolRadosNamespaceNamespaceLister helps list and get CephBlockPoolRadosNamespaces.
// All objects returned here must be treated as read-only.
type CephBlockPoolRadosNamespaceNamespaceLister interface {
	// List lists all CephBlockPoolRadosNamespaces in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.CephBlockPoolRadosNamespace, err error)
	// Get retrieves the CephBlockPoolRadosNamespace from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.CephBlockPoolRadosNamespace, error)
	CephBlockPoolRadosNamespaceNamespaceListerExpansion
}

// cephBlockPoolRadosNamespaceNamespaceLister implements the CephBlockPoolRadosNamespaceNamespaceLister
// interface.
type cephBlockPoolRadosNamespaceNamespaceLister struct {
	listers.ResourceIndexer[*v1.CephBlockPoolRadosNamespace]
}
