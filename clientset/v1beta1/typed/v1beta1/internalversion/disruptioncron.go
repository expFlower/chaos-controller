// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2025 Datadog, Inc.
// Code generated by client-gen. DO NOT EDIT.

package internalversion

import (
	"context"
	"time"

	v1beta1 "github.com/DataDog/chaos-controller/api/v1beta1"
	scheme "github.com/DataDog/chaos-controller/clientset/v1beta1/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DisruptionCronsGetter has a method to return a DisruptionCronInterface.
// A group's client should implement this interface.
type DisruptionCronsGetter interface {
	DisruptionCrons(namespace string) DisruptionCronInterface
}

// DisruptionCronInterface has methods to work with DisruptionCron resources.
type DisruptionCronInterface interface {
	Create(ctx context.Context, disruptionCron *v1beta1.DisruptionCron, opts v1.CreateOptions) (*v1beta1.DisruptionCron, error)
	Update(ctx context.Context, disruptionCron *v1beta1.DisruptionCron, opts v1.UpdateOptions) (*v1beta1.DisruptionCron, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.DisruptionCron, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.DisruptionCronList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	DisruptionCronExpansion
}

// disruptionCrons implements DisruptionCronInterface
type disruptionCrons struct {
	client rest.Interface
	ns     string
}

// newDisruptionCrons returns a DisruptionCrons
func newDisruptionCrons(c *ChaosClient, namespace string) *disruptionCrons {
	return &disruptionCrons{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the disruptionCron, and returns the corresponding disruptionCron object, and an error if there is any.
func (c *disruptionCrons) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.DisruptionCron, err error) {
	result = &v1beta1.DisruptionCron{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("disruptioncrons").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DisruptionCrons that match those selectors.
func (c *disruptionCrons) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.DisruptionCronList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.DisruptionCronList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("disruptioncrons").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested disruptionCrons.
func (c *disruptionCrons) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("disruptioncrons").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a disruptionCron and creates it.  Returns the server's representation of the disruptionCron, and an error, if there is any.
func (c *disruptionCrons) Create(ctx context.Context, disruptionCron *v1beta1.DisruptionCron, opts v1.CreateOptions) (result *v1beta1.DisruptionCron, err error) {
	result = &v1beta1.DisruptionCron{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("disruptioncrons").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(disruptionCron).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a disruptionCron and updates it. Returns the server's representation of the disruptionCron, and an error, if there is any.
func (c *disruptionCrons) Update(ctx context.Context, disruptionCron *v1beta1.DisruptionCron, opts v1.UpdateOptions) (result *v1beta1.DisruptionCron, err error) {
	result = &v1beta1.DisruptionCron{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("disruptioncrons").
		Name(disruptionCron.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(disruptionCron).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the disruptionCron and deletes it. Returns an error if one occurs.
func (c *disruptionCrons) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("disruptioncrons").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}
