// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	secretgenv1alpha1 "github.com/k14s/secretgen-controller/pkg/apis/secretgen/v1alpha1"
	versioned "github.com/k14s/secretgen-controller/pkg/client/clientset/versioned"
	internalinterfaces "github.com/k14s/secretgen-controller/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/k14s/secretgen-controller/pkg/client/listers/secretgen/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SecretExportInformer provides access to a shared informer and lister for
// SecretExports.
type SecretExportInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.SecretExportLister
}

type secretExportInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSecretExportInformer constructs a new informer for SecretExport type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSecretExportInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSecretExportInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSecretExportInformer constructs a new informer for SecretExport type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSecretExportInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SecretgenV1alpha1().SecretExports(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SecretgenV1alpha1().SecretExports(namespace).Watch(options)
			},
		},
		&secretgenv1alpha1.SecretExport{},
		resyncPeriod,
		indexers,
	)
}

func (f *secretExportInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSecretExportInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *secretExportInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&secretgenv1alpha1.SecretExport{}, f.defaultInformer)
}

func (f *secretExportInformer) Lister() v1alpha1.SecretExportLister {
	return v1alpha1.NewSecretExportLister(f.Informer().GetIndexer())
}