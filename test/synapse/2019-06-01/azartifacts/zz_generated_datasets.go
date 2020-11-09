// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

type datasetsClient struct {
	con *connection
}

// Pipeline returns the pipeline associated with this client.
func (client *datasetsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// Rename - Renames a dataset.
func (client *datasetsClient) Rename(ctx context.Context, datasetName string, request ArtifactRenameRequest, options *DatasetsRenameOptions) (*azcore.Response, error) {
	req, err := client.RenameCreateRequest(ctx, datasetName, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.RenameHandleError(resp)
	}
	return resp, nil
}

// RenameCreateRequest creates the Rename request.
func (client *datasetsClient) RenameCreateRequest(ctx context.Context, datasetName string, request ArtifactRenameRequest, options *DatasetsRenameOptions) (*azcore.Request, error) {
	urlPath := "/datasets/{datasetName}/rename"
	urlPath = strings.ReplaceAll(urlPath, "{datasetName}", url.PathEscape(datasetName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(request)
}

// RenameHandleError handles the Rename error response.
func (client *datasetsClient) RenameHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
