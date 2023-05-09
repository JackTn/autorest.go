// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package paramgroupinggroup

import (
	"context"
	"generatortests"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/stretchr/testify/require"
)

func newParameterGroupingClient(t *testing.T) *ParameterGroupingClient {
	client, err := NewParameterGroupingClient(nil)
	require.NoError(t, err)
	return client
}

func NewParameterGroupingClient(options *azcore.ClientOptions) (*ParameterGroupingClient, error) {
	client, err := azcore.NewClient("paramgroupinggroup.ParameterGroupingClient", generatortests.ModuleVersion, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	return &ParameterGroupingClient{internal: client}, nil
}

// PostMultiParamGroups - Post parameters from multiple different parameter groups
func TestPostMultiParamGroups(t *testing.T) {
	client := newParameterGroupingClient(t)
	result, err := client.PostMultiParamGroups(context.Background(), nil, nil, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// PostOptional - Post a bunch of optional parameters grouped
func TestPostOptional(t *testing.T) {
	client := newParameterGroupingClient(t)
	result, err := client.PostOptional(context.Background(), nil, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// PostRequired - Post a bunch of required parameters grouped
func TestPostRequired(t *testing.T) {
	client := newParameterGroupingClient(t)
	result, err := client.PostRequired(context.Background(), ParameterGroupingClientPostRequiredParameters{
		Body: 1234,
		Path: "path",
	}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// PostSharedParameterGroupObject - Post parameters with a shared parameter group object
func TestPostSharedParameterGroupObject(t *testing.T) {
	client := newParameterGroupingClient(t)
	result, err := client.PostSharedParameterGroupObject(context.Background(), nil, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}