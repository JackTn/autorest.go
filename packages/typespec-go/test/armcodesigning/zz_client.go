// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armcodesigning

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// Client - Code Signing resource provider api.
// Don't use this type directly, use NewClient() instead.
type Client struct {
	internal       *arm.Client
	subscriptionID string
}

// NewClient creates a new instance of Client with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*Client, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &Client{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewAccountsClient creates a new instance of [AccountsClient].
func (client *Client) NewAccountsClient() *AccountsClient {
	return &AccountsClient{
		internal:       client.internal,
		subscriptionID: client.subscriptionID,
	}
}

// NewCertificateProfilesClient creates a new instance of [CertificateProfilesClient].
func (client *Client) NewCertificateProfilesClient() *CertificateProfilesClient {
	return &CertificateProfilesClient{
		internal:       client.internal,
		subscriptionID: client.subscriptionID,
	}
}

// NewOperationsClient creates a new instance of [OperationsClient].
func (client *Client) NewOperationsClient() *OperationsClient {
	return &OperationsClient{
		internal:       client.internal,
		subscriptionID: client.subscriptionID,
	}
}
