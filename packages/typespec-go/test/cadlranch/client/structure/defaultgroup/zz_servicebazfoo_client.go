// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package defaultgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// ServiceBazFooClient contains the methods for the Client.Structure.Service namespace.
// Don't use this type directly, use [ServiceBazClient.NewServiceBazFooClient] instead.
type ServiceBazFooClient struct {
	internal *azcore.Client
	endpoint string
	client   ClientType
}

// - options - ServiceBazFooClientSevenOptions contains the optional parameters for the ServiceBazFooClient.Seven method.
func (client *ServiceBazFooClient) Seven(ctx context.Context, options *ServiceBazFooClientSevenOptions) (ServiceBazFooClientSevenResponse, error) {
	var err error
	const operationName = "ServiceBazFooClient.Seven"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.sevenCreateRequest(ctx, options)
	if err != nil {
		return ServiceBazFooClientSevenResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceBazFooClientSevenResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ServiceBazFooClientSevenResponse{}, err
	}
	return ServiceBazFooClientSevenResponse{}, nil
}

// sevenCreateRequest creates the Seven request.
func (client *ServiceBazFooClient) sevenCreateRequest(ctx context.Context, _ *ServiceBazFooClientSevenOptions) (*policy.Request, error) {
	host := "{endpoint}/client/structure/{client}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	host = strings.ReplaceAll(host, "{client}", string(client.client))
	urlPath := "/seven"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}
