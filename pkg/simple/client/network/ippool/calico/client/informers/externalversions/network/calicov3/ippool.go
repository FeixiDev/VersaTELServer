/*
Copyright 2020 The KubeSphere Authors.

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

package calicov3

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"

	networkcalicov3 "kubesphere.io/api/network/calicov3"

	versioned "kubesphere.io/kubesphere/pkg/simple/client/network/ippool/calico/client/clientset/versioned"
	internalinterfaces "kubesphere.io/kubesphere/pkg/simple/client/network/ippool/calico/client/informers/externalversions/internalinterfaces"
	calicov3 "kubesphere.io/kubesphere/pkg/simple/client/network/ippool/calico/client/listers/network/calicov3"
)

// IPPoolInformer provides access to a shared informer and lister for
// IPPools.
type IPPoolInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() calicov3.IPPoolLister
}

type iPPoolInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewIPPoolInformer constructs a new informer for IPPool type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewIPPoolInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredIPPoolInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredIPPoolInformer constructs a new informer for IPPool type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredIPPoolInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CrdCalicov3().IPPools().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CrdCalicov3().IPPools().Watch(context.TODO(), options)
			},
		},
		&networkcalicov3.IPPool{},
		resyncPeriod,
		indexers,
	)
}

func (f *iPPoolInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredIPPoolInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *iPPoolInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&networkcalicov3.IPPool{}, f.defaultInformer)
}

func (f *iPPoolInformer) Lister() calicov3.IPPoolLister {
	return calicov3.NewIPPoolLister(f.Informer().GetIndexer())
}
