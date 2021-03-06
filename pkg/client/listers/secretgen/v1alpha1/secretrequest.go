// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/vmware-tanzu/carvel-secretgen-controller/pkg/apis/secretgen/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SecretRequestLister helps list SecretRequests.
type SecretRequestLister interface {
	// List lists all SecretRequests in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SecretRequest, err error)
	// SecretRequests returns an object that can list and get SecretRequests.
	SecretRequests(namespace string) SecretRequestNamespaceLister
	SecretRequestListerExpansion
}

// secretRequestLister implements the SecretRequestLister interface.
type secretRequestLister struct {
	indexer cache.Indexer
}

// NewSecretRequestLister returns a new SecretRequestLister.
func NewSecretRequestLister(indexer cache.Indexer) SecretRequestLister {
	return &secretRequestLister{indexer: indexer}
}

// List lists all SecretRequests in the indexer.
func (s *secretRequestLister) List(selector labels.Selector) (ret []*v1alpha1.SecretRequest, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SecretRequest))
	})
	return ret, err
}

// SecretRequests returns an object that can list and get SecretRequests.
func (s *secretRequestLister) SecretRequests(namespace string) SecretRequestNamespaceLister {
	return secretRequestNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SecretRequestNamespaceLister helps list and get SecretRequests.
type SecretRequestNamespaceLister interface {
	// List lists all SecretRequests in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SecretRequest, err error)
	// Get retrieves the SecretRequest from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SecretRequest, error)
	SecretRequestNamespaceListerExpansion
}

// secretRequestNamespaceLister implements the SecretRequestNamespaceLister
// interface.
type secretRequestNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SecretRequests in the indexer for a given namespace.
func (s secretRequestNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SecretRequest, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SecretRequest))
	})
	return ret, err
}

// Get retrieves the SecretRequest from the indexer for a given namespace and name.
func (s secretRequestNamespaceLister) Get(name string) (*v1alpha1.SecretRequest, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("secretrequest"), name)
	}
	return obj.(*v1alpha1.SecretRequest), nil
}
