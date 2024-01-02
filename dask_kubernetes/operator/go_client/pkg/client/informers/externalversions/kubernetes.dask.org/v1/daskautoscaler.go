// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	kubernetesdaskorgv1 "github.com/dask/dask-kubernetes/v2024/dask_kubernetes/operator/go_client/pkg/apis/kubernetes.dask.org/v1"
	versioned "github.com/dask/dask-kubernetes/v2024/dask_kubernetes/operator/go_client/pkg/client/clientset/versioned"
	internalinterfaces "github.com/dask/dask-kubernetes/v2024/dask_kubernetes/operator/go_client/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/dask/dask-kubernetes/v2024/dask_kubernetes/operator/go_client/pkg/client/listers/kubernetes.dask.org/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DaskAutoscalerInformer provides access to a shared informer and lister for
// DaskAutoscalers.
type DaskAutoscalerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.DaskAutoscalerLister
}

type daskAutoscalerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDaskAutoscalerInformer constructs a new informer for DaskAutoscaler type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDaskAutoscalerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDaskAutoscalerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDaskAutoscalerInformer constructs a new informer for DaskAutoscaler type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDaskAutoscalerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubernetesV1().DaskAutoscalers(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubernetesV1().DaskAutoscalers(namespace).Watch(context.TODO(), options)
			},
		},
		&kubernetesdaskorgv1.DaskAutoscaler{},
		resyncPeriod,
		indexers,
	)
}

func (f *daskAutoscalerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDaskAutoscalerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *daskAutoscalerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kubernetesdaskorgv1.DaskAutoscaler{}, f.defaultInformer)
}

func (f *daskAutoscalerInformer) Lister() v1.DaskAutoscalerLister {
	return v1.NewDaskAutoscalerLister(f.Informer().GetIndexer())
}
