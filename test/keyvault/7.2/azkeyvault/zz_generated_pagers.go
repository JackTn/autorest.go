//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azkeyvault

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// ClientGetCertificateIssuersPager provides operations for iterating over paged responses.
type ClientGetCertificateIssuersPager struct {
	client    *Client
	current   ClientGetCertificateIssuersResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClientGetCertificateIssuersResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ClientGetCertificateIssuersPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ClientGetCertificateIssuersPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.CertificateIssuerListResult.NextLink == nil || len(*p.current.CertificateIssuerListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.getCertificateIssuersHandleError(resp)
		return false
	}
	result, err := p.client.getCertificateIssuersHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ClientGetCertificateIssuersResponse page.
func (p *ClientGetCertificateIssuersPager) PageResponse() ClientGetCertificateIssuersResponse {
	return p.current
}

// ClientGetCertificateVersionsPager provides operations for iterating over paged responses.
type ClientGetCertificateVersionsPager struct {
	client    *Client
	current   ClientGetCertificateVersionsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClientGetCertificateVersionsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ClientGetCertificateVersionsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ClientGetCertificateVersionsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.CertificateListResult.NextLink == nil || len(*p.current.CertificateListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.getCertificateVersionsHandleError(resp)
		return false
	}
	result, err := p.client.getCertificateVersionsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ClientGetCertificateVersionsResponse page.
func (p *ClientGetCertificateVersionsPager) PageResponse() ClientGetCertificateVersionsResponse {
	return p.current
}

// ClientGetCertificatesPager provides operations for iterating over paged responses.
type ClientGetCertificatesPager struct {
	client    *Client
	current   ClientGetCertificatesResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClientGetCertificatesResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ClientGetCertificatesPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ClientGetCertificatesPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.CertificateListResult.NextLink == nil || len(*p.current.CertificateListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.getCertificatesHandleError(resp)
		return false
	}
	result, err := p.client.getCertificatesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ClientGetCertificatesResponse page.
func (p *ClientGetCertificatesPager) PageResponse() ClientGetCertificatesResponse {
	return p.current
}

// ClientGetDeletedCertificatesPager provides operations for iterating over paged responses.
type ClientGetDeletedCertificatesPager struct {
	client    *Client
	current   ClientGetDeletedCertificatesResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClientGetDeletedCertificatesResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ClientGetDeletedCertificatesPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ClientGetDeletedCertificatesPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DeletedCertificateListResult.NextLink == nil || len(*p.current.DeletedCertificateListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.getDeletedCertificatesHandleError(resp)
		return false
	}
	result, err := p.client.getDeletedCertificatesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ClientGetDeletedCertificatesResponse page.
func (p *ClientGetDeletedCertificatesPager) PageResponse() ClientGetDeletedCertificatesResponse {
	return p.current
}

// ClientGetDeletedKeysPager provides operations for iterating over paged responses.
type ClientGetDeletedKeysPager struct {
	client    *Client
	current   ClientGetDeletedKeysResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClientGetDeletedKeysResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ClientGetDeletedKeysPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ClientGetDeletedKeysPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DeletedKeyListResult.NextLink == nil || len(*p.current.DeletedKeyListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.getDeletedKeysHandleError(resp)
		return false
	}
	result, err := p.client.getDeletedKeysHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ClientGetDeletedKeysResponse page.
func (p *ClientGetDeletedKeysPager) PageResponse() ClientGetDeletedKeysResponse {
	return p.current
}

// ClientGetDeletedSasDefinitionsPager provides operations for iterating over paged responses.
type ClientGetDeletedSasDefinitionsPager struct {
	client    *Client
	current   ClientGetDeletedSasDefinitionsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClientGetDeletedSasDefinitionsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ClientGetDeletedSasDefinitionsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ClientGetDeletedSasDefinitionsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DeletedSasDefinitionListResult.NextLink == nil || len(*p.current.DeletedSasDefinitionListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.getDeletedSasDefinitionsHandleError(resp)
		return false
	}
	result, err := p.client.getDeletedSasDefinitionsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ClientGetDeletedSasDefinitionsResponse page.
func (p *ClientGetDeletedSasDefinitionsPager) PageResponse() ClientGetDeletedSasDefinitionsResponse {
	return p.current
}

// ClientGetDeletedSecretsPager provides operations for iterating over paged responses.
type ClientGetDeletedSecretsPager struct {
	client    *Client
	current   ClientGetDeletedSecretsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClientGetDeletedSecretsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ClientGetDeletedSecretsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ClientGetDeletedSecretsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DeletedSecretListResult.NextLink == nil || len(*p.current.DeletedSecretListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.getDeletedSecretsHandleError(resp)
		return false
	}
	result, err := p.client.getDeletedSecretsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ClientGetDeletedSecretsResponse page.
func (p *ClientGetDeletedSecretsPager) PageResponse() ClientGetDeletedSecretsResponse {
	return p.current
}

// ClientGetDeletedStorageAccountsPager provides operations for iterating over paged responses.
type ClientGetDeletedStorageAccountsPager struct {
	client    *Client
	current   ClientGetDeletedStorageAccountsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClientGetDeletedStorageAccountsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ClientGetDeletedStorageAccountsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ClientGetDeletedStorageAccountsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DeletedStorageListResult.NextLink == nil || len(*p.current.DeletedStorageListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.getDeletedStorageAccountsHandleError(resp)
		return false
	}
	result, err := p.client.getDeletedStorageAccountsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ClientGetDeletedStorageAccountsResponse page.
func (p *ClientGetDeletedStorageAccountsPager) PageResponse() ClientGetDeletedStorageAccountsResponse {
	return p.current
}

// ClientGetKeyVersionsPager provides operations for iterating over paged responses.
type ClientGetKeyVersionsPager struct {
	client    *Client
	current   ClientGetKeyVersionsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClientGetKeyVersionsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ClientGetKeyVersionsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ClientGetKeyVersionsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.KeyListResult.NextLink == nil || len(*p.current.KeyListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.getKeyVersionsHandleError(resp)
		return false
	}
	result, err := p.client.getKeyVersionsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ClientGetKeyVersionsResponse page.
func (p *ClientGetKeyVersionsPager) PageResponse() ClientGetKeyVersionsResponse {
	return p.current
}

// ClientGetKeysPager provides operations for iterating over paged responses.
type ClientGetKeysPager struct {
	client    *Client
	current   ClientGetKeysResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClientGetKeysResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ClientGetKeysPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ClientGetKeysPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.KeyListResult.NextLink == nil || len(*p.current.KeyListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.getKeysHandleError(resp)
		return false
	}
	result, err := p.client.getKeysHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ClientGetKeysResponse page.
func (p *ClientGetKeysPager) PageResponse() ClientGetKeysResponse {
	return p.current
}

// ClientGetSasDefinitionsPager provides operations for iterating over paged responses.
type ClientGetSasDefinitionsPager struct {
	client    *Client
	current   ClientGetSasDefinitionsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClientGetSasDefinitionsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ClientGetSasDefinitionsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ClientGetSasDefinitionsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SasDefinitionListResult.NextLink == nil || len(*p.current.SasDefinitionListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.getSasDefinitionsHandleError(resp)
		return false
	}
	result, err := p.client.getSasDefinitionsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ClientGetSasDefinitionsResponse page.
func (p *ClientGetSasDefinitionsPager) PageResponse() ClientGetSasDefinitionsResponse {
	return p.current
}

// ClientGetSecretVersionsPager provides operations for iterating over paged responses.
type ClientGetSecretVersionsPager struct {
	client    *Client
	current   ClientGetSecretVersionsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClientGetSecretVersionsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ClientGetSecretVersionsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ClientGetSecretVersionsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SecretListResult.NextLink == nil || len(*p.current.SecretListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.getSecretVersionsHandleError(resp)
		return false
	}
	result, err := p.client.getSecretVersionsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ClientGetSecretVersionsResponse page.
func (p *ClientGetSecretVersionsPager) PageResponse() ClientGetSecretVersionsResponse {
	return p.current
}

// ClientGetSecretsPager provides operations for iterating over paged responses.
type ClientGetSecretsPager struct {
	client    *Client
	current   ClientGetSecretsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClientGetSecretsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ClientGetSecretsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ClientGetSecretsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SecretListResult.NextLink == nil || len(*p.current.SecretListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.getSecretsHandleError(resp)
		return false
	}
	result, err := p.client.getSecretsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ClientGetSecretsResponse page.
func (p *ClientGetSecretsPager) PageResponse() ClientGetSecretsResponse {
	return p.current
}

// ClientGetStorageAccountsPager provides operations for iterating over paged responses.
type ClientGetStorageAccountsPager struct {
	client    *Client
	current   ClientGetStorageAccountsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClientGetStorageAccountsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ClientGetStorageAccountsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ClientGetStorageAccountsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.StorageListResult.NextLink == nil || len(*p.current.StorageListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.getStorageAccountsHandleError(resp)
		return false
	}
	result, err := p.client.getStorageAccountsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ClientGetStorageAccountsResponse page.
func (p *ClientGetStorageAccountsPager) PageResponse() ClientGetStorageAccountsResponse {
	return p.current
}

// RoleAssignmentsClientListForScopePager provides operations for iterating over paged responses.
type RoleAssignmentsClientListForScopePager struct {
	client    *RoleAssignmentsClient
	current   RoleAssignmentsClientListForScopeResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, RoleAssignmentsClientListForScopeResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *RoleAssignmentsClientListForScopePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *RoleAssignmentsClientListForScopePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RoleAssignmentListResult.NextLink == nil || len(*p.current.RoleAssignmentListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listForScopeHandleError(resp)
		return false
	}
	result, err := p.client.listForScopeHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current RoleAssignmentsClientListForScopeResponse page.
func (p *RoleAssignmentsClientListForScopePager) PageResponse() RoleAssignmentsClientListForScopeResponse {
	return p.current
}

// RoleDefinitionsClientListPager provides operations for iterating over paged responses.
type RoleDefinitionsClientListPager struct {
	client    *RoleDefinitionsClient
	current   RoleDefinitionsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, RoleDefinitionsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *RoleDefinitionsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *RoleDefinitionsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RoleDefinitionListResult.NextLink == nil || len(*p.current.RoleDefinitionListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current RoleDefinitionsClientListResponse page.
func (p *RoleDefinitionsClientListPager) PageResponse() RoleDefinitionsClientListResponse {
	return p.current
}
