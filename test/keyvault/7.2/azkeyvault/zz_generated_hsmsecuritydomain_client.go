//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azkeyvault

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// HSMSecurityDomainClient contains the methods for the HSMSecurityDomain group.
// Don't use this type directly, use NewHSMSecurityDomainClient() instead.
type HSMSecurityDomainClient struct {
	pl runtime.Pipeline
}

// NewHSMSecurityDomainClient creates a new instance of HSMSecurityDomainClient with the specified values.
// pl - the pipeline used for sending requests and handling responses.
func NewHSMSecurityDomainClient(pl runtime.Pipeline) *HSMSecurityDomainClient {
	client := &HSMSecurityDomainClient{
		pl: pl,
	}
	return client
}

// BeginDownload - Retrieves the Security Domain from the managed HSM. Calling this endpoint can be used to activate a provisioned
// managed HSM resource.
// If the operation fails it returns the *Error error type.
// vaultBaseURL - The vault name, for example https://myvault.vault.azure.net.
// certificateInfoObject - The Security Domain download operation requires customer to provide N certificates (minimum 3 and
// maximum 10) containing a public key in JWK format.
// options - HSMSecurityDomainClientBeginDownloadOptions contains the optional parameters for the HSMSecurityDomainClient.BeginDownload
// method.
func (client *HSMSecurityDomainClient) BeginDownload(ctx context.Context, vaultBaseURL string, certificateInfoObject CertificateInfoObject, options *HSMSecurityDomainClientBeginDownloadOptions) (HSMSecurityDomainClientDownloadPollerResponse, error) {
	resp, err := client.download(ctx, vaultBaseURL, certificateInfoObject, options)
	if err != nil {
		return HSMSecurityDomainClientDownloadPollerResponse{}, err
	}
	result := HSMSecurityDomainClientDownloadPollerResponse{
		RawResponse: resp,
	}
	pt, err := runtime.NewPoller("HSMSecurityDomainClient.Download", resp, client.pl, client.downloadHandleError)
	if err != nil {
		return HSMSecurityDomainClientDownloadPollerResponse{}, err
	}
	result.Poller = &HSMSecurityDomainClientDownloadPoller{
		pt: pt,
	}
	return result, nil
}

// Download - Retrieves the Security Domain from the managed HSM. Calling this endpoint can be used to activate a provisioned
// managed HSM resource.
// If the operation fails it returns the *Error error type.
func (client *HSMSecurityDomainClient) download(ctx context.Context, vaultBaseURL string, certificateInfoObject CertificateInfoObject, options *HSMSecurityDomainClientBeginDownloadOptions) (*http.Response, error) {
	req, err := client.downloadCreateRequest(ctx, vaultBaseURL, certificateInfoObject, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, client.downloadHandleError(resp)
	}
	return resp, nil
}

// downloadCreateRequest creates the Download request.
func (client *HSMSecurityDomainClient) downloadCreateRequest(ctx context.Context, vaultBaseURL string, certificateInfoObject CertificateInfoObject, options *HSMSecurityDomainClientBeginDownloadOptions) (*policy.Request, error) {
	host := "{vaultBaseUrl}"
	host = strings.ReplaceAll(host, "{vaultBaseUrl}", vaultBaseURL)
	urlPath := "/securitydomain/download"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "7.2")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, certificateInfoObject)
}

// downloadHandleError handles the Download error response.
func (client *HSMSecurityDomainClient) downloadHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// DownloadPending - Retrieves the Security Domain download operation status
// If the operation fails it returns the *Error error type.
// vaultBaseURL - The vault name, for example https://myvault.vault.azure.net.
// options - HSMSecurityDomainClientDownloadPendingOptions contains the optional parameters for the HSMSecurityDomainClient.DownloadPending
// method.
func (client *HSMSecurityDomainClient) DownloadPending(ctx context.Context, vaultBaseURL string, options *HSMSecurityDomainClientDownloadPendingOptions) (HSMSecurityDomainClientDownloadPendingResponse, error) {
	req, err := client.downloadPendingCreateRequest(ctx, vaultBaseURL, options)
	if err != nil {
		return HSMSecurityDomainClientDownloadPendingResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HSMSecurityDomainClientDownloadPendingResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HSMSecurityDomainClientDownloadPendingResponse{}, client.downloadPendingHandleError(resp)
	}
	return client.downloadPendingHandleResponse(resp)
}

// downloadPendingCreateRequest creates the DownloadPending request.
func (client *HSMSecurityDomainClient) downloadPendingCreateRequest(ctx context.Context, vaultBaseURL string, options *HSMSecurityDomainClientDownloadPendingOptions) (*policy.Request, error) {
	host := "{vaultBaseUrl}"
	host = strings.ReplaceAll(host, "{vaultBaseUrl}", vaultBaseURL)
	urlPath := "/securitydomain/download/pending"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// downloadPendingHandleResponse handles the DownloadPending response.
func (client *HSMSecurityDomainClient) downloadPendingHandleResponse(resp *http.Response) (HSMSecurityDomainClientDownloadPendingResponse, error) {
	result := HSMSecurityDomainClientDownloadPendingResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityDomainOperationStatus); err != nil {
		return HSMSecurityDomainClientDownloadPendingResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// downloadPendingHandleError handles the DownloadPending error response.
func (client *HSMSecurityDomainClient) downloadPendingHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// TransferKey - Retrieve Security Domain transfer key
// If the operation fails it returns the *Error error type.
// vaultBaseURL - The vault name, for example https://myvault.vault.azure.net.
// options - HSMSecurityDomainClientTransferKeyOptions contains the optional parameters for the HSMSecurityDomainClient.TransferKey
// method.
func (client *HSMSecurityDomainClient) TransferKey(ctx context.Context, vaultBaseURL string, options *HSMSecurityDomainClientTransferKeyOptions) (HSMSecurityDomainClientTransferKeyResponse, error) {
	req, err := client.transferKeyCreateRequest(ctx, vaultBaseURL, options)
	if err != nil {
		return HSMSecurityDomainClientTransferKeyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HSMSecurityDomainClientTransferKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HSMSecurityDomainClientTransferKeyResponse{}, client.transferKeyHandleError(resp)
	}
	return client.transferKeyHandleResponse(resp)
}

// transferKeyCreateRequest creates the TransferKey request.
func (client *HSMSecurityDomainClient) transferKeyCreateRequest(ctx context.Context, vaultBaseURL string, options *HSMSecurityDomainClientTransferKeyOptions) (*policy.Request, error) {
	host := "{vaultBaseUrl}"
	host = strings.ReplaceAll(host, "{vaultBaseUrl}", vaultBaseURL)
	urlPath := "/securitydomain/upload"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "7.2")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// transferKeyHandleResponse handles the TransferKey response.
func (client *HSMSecurityDomainClient) transferKeyHandleResponse(resp *http.Response) (HSMSecurityDomainClientTransferKeyResponse, error) {
	result := HSMSecurityDomainClientTransferKeyResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TransferKey); err != nil {
		return HSMSecurityDomainClientTransferKeyResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// transferKeyHandleError handles the TransferKey error response.
func (client *HSMSecurityDomainClient) transferKeyHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginUpload - Restore the provided Security Domain.
// If the operation fails it returns the *Error error type.
// vaultBaseURL - The vault name, for example https://myvault.vault.azure.net.
// securityDomain - The Security Domain to be restored.
// options - HSMSecurityDomainClientBeginUploadOptions contains the optional parameters for the HSMSecurityDomainClient.BeginUpload
// method.
func (client *HSMSecurityDomainClient) BeginUpload(ctx context.Context, vaultBaseURL string, securityDomain SecurityDomainObject, options *HSMSecurityDomainClientBeginUploadOptions) (HSMSecurityDomainClientUploadPollerResponse, error) {
	resp, err := client.upload(ctx, vaultBaseURL, securityDomain, options)
	if err != nil {
		return HSMSecurityDomainClientUploadPollerResponse{}, err
	}
	result := HSMSecurityDomainClientUploadPollerResponse{
		RawResponse: resp,
	}
	pt, err := runtime.NewPoller("HSMSecurityDomainClient.Upload", resp, client.pl, client.uploadHandleError)
	if err != nil {
		return HSMSecurityDomainClientUploadPollerResponse{}, err
	}
	result.Poller = &HSMSecurityDomainClientUploadPoller{
		pt: pt,
	}
	return result, nil
}

// Upload - Restore the provided Security Domain.
// If the operation fails it returns the *Error error type.
func (client *HSMSecurityDomainClient) upload(ctx context.Context, vaultBaseURL string, securityDomain SecurityDomainObject, options *HSMSecurityDomainClientBeginUploadOptions) (*http.Response, error) {
	req, err := client.uploadCreateRequest(ctx, vaultBaseURL, securityDomain, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.uploadHandleError(resp)
	}
	return resp, nil
}

// uploadCreateRequest creates the Upload request.
func (client *HSMSecurityDomainClient) uploadCreateRequest(ctx context.Context, vaultBaseURL string, securityDomain SecurityDomainObject, options *HSMSecurityDomainClientBeginUploadOptions) (*policy.Request, error) {
	host := "{vaultBaseUrl}"
	host = strings.ReplaceAll(host, "{vaultBaseUrl}", vaultBaseURL)
	urlPath := "/securitydomain/upload"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, securityDomain)
}

// uploadHandleError handles the Upload error response.
func (client *HSMSecurityDomainClient) uploadHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// UploadPending - Get Security Domain upload operation status
// If the operation fails it returns the *Error error type.
// vaultBaseURL - The vault name, for example https://myvault.vault.azure.net.
// options - HSMSecurityDomainClientUploadPendingOptions contains the optional parameters for the HSMSecurityDomainClient.UploadPending
// method.
func (client *HSMSecurityDomainClient) UploadPending(ctx context.Context, vaultBaseURL string, options *HSMSecurityDomainClientUploadPendingOptions) (HSMSecurityDomainClientUploadPendingResponse, error) {
	req, err := client.uploadPendingCreateRequest(ctx, vaultBaseURL, options)
	if err != nil {
		return HSMSecurityDomainClientUploadPendingResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HSMSecurityDomainClientUploadPendingResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HSMSecurityDomainClientUploadPendingResponse{}, client.uploadPendingHandleError(resp)
	}
	return client.uploadPendingHandleResponse(resp)
}

// uploadPendingCreateRequest creates the UploadPending request.
func (client *HSMSecurityDomainClient) uploadPendingCreateRequest(ctx context.Context, vaultBaseURL string, options *HSMSecurityDomainClientUploadPendingOptions) (*policy.Request, error) {
	host := "{vaultBaseUrl}"
	host = strings.ReplaceAll(host, "{vaultBaseUrl}", vaultBaseURL)
	urlPath := "/securitydomain/upload/pending"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// uploadPendingHandleResponse handles the UploadPending response.
func (client *HSMSecurityDomainClient) uploadPendingHandleResponse(resp *http.Response) (HSMSecurityDomainClientUploadPendingResponse, error) {
	result := HSMSecurityDomainClientUploadPendingResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityDomainOperationStatus); err != nil {
		return HSMSecurityDomainClientUploadPendingResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// uploadPendingHandleError handles the UploadPending error response.
func (client *HSMSecurityDomainClient) uploadPendingHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
