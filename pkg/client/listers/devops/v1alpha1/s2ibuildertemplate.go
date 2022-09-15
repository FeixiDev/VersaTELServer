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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "kubesphere.io/api/devops/v1alpha1"
)

// S2iBuilderTemplateLister helps list S2iBuilderTemplates.
// All objects returned here must be treated as read-only.
type S2iBuilderTemplateLister interface {
	// List lists all S2iBuilderTemplates in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.S2iBuilderTemplate, err error)
	// Get retrieves the S2iBuilderTemplate from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.S2iBuilderTemplate, error)
	S2iBuilderTemplateListerExpansion
}

// s2iBuilderTemplateLister implements the S2iBuilderTemplateLister interface.
type s2iBuilderTemplateLister struct {
	indexer cache.Indexer
}

// NewS2iBuilderTemplateLister returns a new S2iBuilderTemplateLister.
func NewS2iBuilderTemplateLister(indexer cache.Indexer) S2iBuilderTemplateLister {
	return &s2iBuilderTemplateLister{indexer: indexer}
}

// List lists all S2iBuilderTemplates in the indexer.
func (s *s2iBuilderTemplateLister) List(selector labels.Selector) (ret []*v1alpha1.S2iBuilderTemplate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.S2iBuilderTemplate))
	})
	return ret, err
}

// Get retrieves the S2iBuilderTemplate from the index for a given name.
func (s *s2iBuilderTemplateLister) Get(name string) (*v1alpha1.S2iBuilderTemplate, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("s2ibuildertemplate"), name)
	}
	return obj.(*v1alpha1.S2iBuilderTemplate), nil
}
