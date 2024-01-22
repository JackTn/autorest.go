//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package condreqgroup_test

import (
	"condreqgroup"
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestConditionalRequestClient_PostIfMatch(t *testing.T) {
	client, err := condreqgroup.NewConditionalRequestClient(nil)
	require.NoError(t, err)
	resp, err := client.PostIfMatch(context.Background(), &condreqgroup.ConditionalRequestClientPostIfMatchOptions{
		IfMatch: to.Ptr(`"valid"`),
	})
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestConditionalRequestClient_PostIfNoneMatch(t *testing.T) {
	client, err := condreqgroup.NewConditionalRequestClient(nil)
	require.NoError(t, err)
	resp, err := client.PostIfNoneMatch(context.Background(), &condreqgroup.ConditionalRequestClientPostIfNoneMatchOptions{
		IfNoneMatch: to.Ptr(`"invalid"`),
	})
	require.NoError(t, err)
	require.Zero(t, resp)
}
