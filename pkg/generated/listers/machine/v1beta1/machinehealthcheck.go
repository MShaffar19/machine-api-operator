/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2019 Red Hat, Inc.
 *
 */
// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/openshift/machine-api-operator/pkg/apis/machine/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MachineHealthCheckLister helps list MachineHealthChecks.
type MachineHealthCheckLister interface {
	// List lists all MachineHealthChecks in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.MachineHealthCheck, err error)
	// MachineHealthChecks returns an object that can list and get MachineHealthChecks.
	MachineHealthChecks(namespace string) MachineHealthCheckNamespaceLister
	MachineHealthCheckListerExpansion
}

// machineHealthCheckLister implements the MachineHealthCheckLister interface.
type machineHealthCheckLister struct {
	indexer cache.Indexer
}

// NewMachineHealthCheckLister returns a new MachineHealthCheckLister.
func NewMachineHealthCheckLister(indexer cache.Indexer) MachineHealthCheckLister {
	return &machineHealthCheckLister{indexer: indexer}
}

// List lists all MachineHealthChecks in the indexer.
func (s *machineHealthCheckLister) List(selector labels.Selector) (ret []*v1beta1.MachineHealthCheck, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.MachineHealthCheck))
	})
	return ret, err
}

// MachineHealthChecks returns an object that can list and get MachineHealthChecks.
func (s *machineHealthCheckLister) MachineHealthChecks(namespace string) MachineHealthCheckNamespaceLister {
	return machineHealthCheckNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MachineHealthCheckNamespaceLister helps list and get MachineHealthChecks.
type MachineHealthCheckNamespaceLister interface {
	// List lists all MachineHealthChecks in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.MachineHealthCheck, err error)
	// Get retrieves the MachineHealthCheck from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.MachineHealthCheck, error)
	MachineHealthCheckNamespaceListerExpansion
}

// machineHealthCheckNamespaceLister implements the MachineHealthCheckNamespaceLister
// interface.
type machineHealthCheckNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MachineHealthChecks in the indexer for a given namespace.
func (s machineHealthCheckNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.MachineHealthCheck, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.MachineHealthCheck))
	})
	return ret, err
}

// Get retrieves the MachineHealthCheck from the indexer for a given namespace and name.
func (s machineHealthCheckNamespaceLister) Get(name string) (*v1beta1.MachineHealthCheck, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("machinehealthcheck"), name)
	}
	return obj.(*v1beta1.MachineHealthCheck), nil
}