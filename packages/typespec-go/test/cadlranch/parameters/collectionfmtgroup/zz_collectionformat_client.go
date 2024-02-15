// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package collectionfmtgroup

import "github.com/Azure/azure-sdk-for-go/sdk/azcore"

// CollectionFormatClient - Test for collectionFormat.
// Don't use this type directly, use a constructor function instead.
type CollectionFormatClient struct {
	internal *azcore.Client
}

// NewHeaderClient creates a new instance of [HeaderClient].
func (client *CollectionFormatClient) NewHeaderClient() *HeaderClient {
	return &HeaderClient{
		internal: client.internal,
	}
}

// NewQueryClient creates a new instance of [QueryClient].
func (client *CollectionFormatClient) NewQueryClient() *QueryClient {
	return &QueryClient{
		internal: client.internal,
	}
}
