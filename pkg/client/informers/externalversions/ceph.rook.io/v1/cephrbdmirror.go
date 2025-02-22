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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	cephrookiov1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
	versioned "github.com/rook/rook/pkg/client/clientset/versioned"
	internalinterfaces "github.com/rook/rook/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/rook/rook/pkg/client/listers/ceph.rook.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CephRBDMirrorInformer provides access to a shared informer and lister for
// CephRBDMirrors.
type CephRBDMirrorInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.CephRBDMirrorLister
}

type cephRBDMirrorInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCephRBDMirrorInformer constructs a new informer for CephRBDMirror type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCephRBDMirrorInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCephRBDMirrorInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCephRBDMirrorInformer constructs a new informer for CephRBDMirror type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCephRBDMirrorInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CephV1().CephRBDMirrors(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CephV1().CephRBDMirrors(namespace).Watch(context.TODO(), options)
			},
		},
		&cephrookiov1.CephRBDMirror{},
		resyncPeriod,
		indexers,
	)
}

func (f *cephRBDMirrorInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCephRBDMirrorInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *cephRBDMirrorInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cephrookiov1.CephRBDMirror{}, f.defaultInformer)
}

func (f *cephRBDMirrorInformer) Lister() v1.CephRBDMirrorLister {
	return v1.NewCephRBDMirrorLister(f.Informer().GetIndexer())
}
