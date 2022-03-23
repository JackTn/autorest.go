//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azblob

import (
	"context"
	"encoding/base64"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"io"
	"net/http"
	"strconv"
	"time"
)

type appendBlobClient struct {
	endpoint string
	version  Enum2
	pl       runtime.Pipeline
}

// newAppendBlobClient creates a new instance of appendBlobClient with the specified values.
// endpoint - The URL of the service account, container, or blob that is the targe of the desired operation.
// version - Specifies the version of the operation to use for this request.
// pl - the pipeline used for sending requests and handling responses.
func newAppendBlobClient(endpoint string, version Enum2, pl runtime.Pipeline) *appendBlobClient {
	client := &appendBlobClient{
		endpoint: endpoint,
		version:  version,
		pl:       pl,
	}
	return client
}

// AppendBlock - The Append Block operation commits a new block of data to the end of an existing append blob. The Append
// Block operation is permitted only if the blob was created with x-ms-blob-type set to
// AppendBlob. Append Block is supported only on version 2015-02-21 version or later.
// If the operation fails it returns an *azcore.ResponseError type.
// contentLength - The length of the request.
// body - Initial data
// options - appendBlobClientAppendBlockOptions contains the optional parameters for the appendBlobClient.AppendBlock method.
// LeaseAccessConditions - LeaseAccessConditions contains a group of parameters for the containerClient.GetProperties method.
// AppendPositionAccessConditions - AppendPositionAccessConditions contains a group of parameters for the appendBlobClient.AppendBlock
// method.
// CpkInfo - CpkInfo contains a group of parameters for the client.Download method.
// CpkScopeInfo - CpkScopeInfo contains a group of parameters for the client.SetMetadata method.
// ModifiedAccessConditions - ModifiedAccessConditions contains a group of parameters for the containerClient.Delete method.
func (client *appendBlobClient) AppendBlock(ctx context.Context, comp Enum38, contentLength int64, body io.ReadSeekCloser, options *appendBlobClientAppendBlockOptions, leaseAccessConditions *LeaseAccessConditions, appendPositionAccessConditions *AppendPositionAccessConditions, cpkInfo *CpkInfo, cpkScopeInfo *CpkScopeInfo, modifiedAccessConditions *ModifiedAccessConditions) (appendBlobClientAppendBlockResponse, error) {
	req, err := client.appendBlockCreateRequest(ctx, comp, contentLength, body, options, leaseAccessConditions, appendPositionAccessConditions, cpkInfo, cpkScopeInfo, modifiedAccessConditions)
	if err != nil {
		return appendBlobClientAppendBlockResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return appendBlobClientAppendBlockResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return appendBlobClientAppendBlockResponse{}, runtime.NewResponseError(resp)
	}
	return client.appendBlockHandleResponse(resp)
}

// appendBlockCreateRequest creates the AppendBlock request.
func (client *appendBlobClient) appendBlockCreateRequest(ctx context.Context, comp Enum38, contentLength int64, body io.ReadSeekCloser, options *appendBlobClientAppendBlockOptions, leaseAccessConditions *LeaseAccessConditions, appendPositionAccessConditions *AppendPositionAccessConditions, cpkInfo *CpkInfo, cpkScopeInfo *CpkScopeInfo, modifiedAccessConditions *ModifiedAccessConditions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodPut, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("comp", string(comp))
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Content-Length", strconv.FormatInt(contentLength, 10))
	if options != nil && options.TransactionalContentMD5 != nil {
		req.Raw().Header.Set("Content-MD5", base64.StdEncoding.EncodeToString(options.TransactionalContentMD5))
	}
	if options != nil && options.TransactionalContentCRC64 != nil {
		req.Raw().Header.Set("x-ms-content-crc64", base64.StdEncoding.EncodeToString(options.TransactionalContentCRC64))
	}
	if leaseAccessConditions != nil && leaseAccessConditions.LeaseID != nil {
		req.Raw().Header.Set("x-ms-lease-id", *leaseAccessConditions.LeaseID)
	}
	if appendPositionAccessConditions != nil && appendPositionAccessConditions.MaxSize != nil {
		req.Raw().Header.Set("x-ms-blob-condition-maxsize", strconv.FormatInt(*appendPositionAccessConditions.MaxSize, 10))
	}
	if appendPositionAccessConditions != nil && appendPositionAccessConditions.AppendPosition != nil {
		req.Raw().Header.Set("x-ms-blob-condition-appendpos", strconv.FormatInt(*appendPositionAccessConditions.AppendPosition, 10))
	}
	if cpkInfo != nil && cpkInfo.EncryptionKey != nil {
		req.Raw().Header.Set("x-ms-encryption-key", *cpkInfo.EncryptionKey)
	}
	if cpkInfo != nil && cpkInfo.EncryptionKeySHA256 != nil {
		req.Raw().Header.Set("x-ms-encryption-key-sha256", *cpkInfo.EncryptionKeySHA256)
	}
	if cpkInfo != nil && cpkInfo.EncryptionAlgorithm != nil {
		req.Raw().Header.Set("x-ms-encryption-algorithm", "AES256")
	}
	if cpkScopeInfo != nil && cpkScopeInfo.EncryptionScope != nil {
		req.Raw().Header.Set("x-ms-encryption-scope", *cpkScopeInfo.EncryptionScope)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfModifiedSince != nil {
		req.Raw().Header.Set("If-Modified-Since", modifiedAccessConditions.IfModifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfUnmodifiedSince != nil {
		req.Raw().Header.Set("If-Unmodified-Since", modifiedAccessConditions.IfUnmodifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *modifiedAccessConditions.IfMatch)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfNoneMatch != nil {
		req.Raw().Header.Set("If-None-Match", *modifiedAccessConditions.IfNoneMatch)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfTags != nil {
		req.Raw().Header.Set("x-ms-if-tags", *modifiedAccessConditions.IfTags)
	}
	req.Raw().Header.Set("x-ms-version", string(client.version))
	if options != nil && options.RequestID != nil {
		req.Raw().Header.Set("x-ms-client-request-id", *options.RequestID)
	}
	req.Raw().Header.Set("Accept", "application/xml")
	return req, req.SetBody(body, "application/octet-stream")
}

// appendBlockHandleResponse handles the AppendBlock response.
func (client *appendBlobClient) appendBlockHandleResponse(resp *http.Response) (appendBlobClientAppendBlockResponse, error) {
	result := appendBlobClientAppendBlockResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if val := resp.Header.Get("Last-Modified"); val != "" {
		lastModified, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return appendBlobClientAppendBlockResponse{}, err
		}
		result.LastModified = &lastModified
	}
	if val := resp.Header.Get("Content-MD5"); val != "" {
		contentMD5, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return appendBlobClientAppendBlockResponse{}, err
		}
		result.ContentMD5 = contentMD5
	}
	if val := resp.Header.Get("x-ms-content-crc64"); val != "" {
		xMSContentCRC64, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return appendBlobClientAppendBlockResponse{}, err
		}
		result.XMSContentCRC64 = xMSContentCRC64
	}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return appendBlobClientAppendBlockResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-blob-append-offset"); val != "" {
		result.BlobAppendOffset = &val
	}
	if val := resp.Header.Get("x-ms-blob-committed-block-count"); val != "" {
		blobCommittedBlockCount32, err := strconv.ParseInt(val, 10, 32)
		blobCommittedBlockCount := int32(blobCommittedBlockCount32)
		if err != nil {
			return appendBlobClientAppendBlockResponse{}, err
		}
		result.BlobCommittedBlockCount = &blobCommittedBlockCount
	}
	if val := resp.Header.Get("x-ms-request-server-encrypted"); val != "" {
		isServerEncrypted, err := strconv.ParseBool(val)
		if err != nil {
			return appendBlobClientAppendBlockResponse{}, err
		}
		result.IsServerEncrypted = &isServerEncrypted
	}
	if val := resp.Header.Get("x-ms-encryption-key-sha256"); val != "" {
		result.EncryptionKeySHA256 = &val
	}
	if val := resp.Header.Get("x-ms-encryption-scope"); val != "" {
		result.EncryptionScope = &val
	}
	return result, nil
}

// AppendBlockFromURL - The Append Block operation commits a new block of data to the end of an existing append blob where
// the contents are read from a source url. The Append Block operation is permitted only if the blob was
// created with x-ms-blob-type set to AppendBlob. Append Block is supported only on version 2015-02-21 version or later.
// If the operation fails it returns an *azcore.ResponseError type.
// sourceURL - Specify a URL to the copy source.
// contentLength - The length of the request.
// options - appendBlobClientAppendBlockFromURLOptions contains the optional parameters for the appendBlobClient.AppendBlockFromURL
// method.
// CpkInfo - CpkInfo contains a group of parameters for the client.Download method.
// CpkScopeInfo - CpkScopeInfo contains a group of parameters for the client.SetMetadata method.
// LeaseAccessConditions - LeaseAccessConditions contains a group of parameters for the containerClient.GetProperties method.
// AppendPositionAccessConditions - AppendPositionAccessConditions contains a group of parameters for the appendBlobClient.AppendBlock
// method.
// ModifiedAccessConditions - ModifiedAccessConditions contains a group of parameters for the containerClient.Delete method.
// SourceModifiedAccessConditions - SourceModifiedAccessConditions contains a group of parameters for the directoryClient.Rename
// method.
func (client *appendBlobClient) AppendBlockFromURL(ctx context.Context, comp Enum38, sourceURL string, contentLength int64, options *appendBlobClientAppendBlockFromURLOptions, cpkInfo *CpkInfo, cpkScopeInfo *CpkScopeInfo, leaseAccessConditions *LeaseAccessConditions, appendPositionAccessConditions *AppendPositionAccessConditions, modifiedAccessConditions *ModifiedAccessConditions, sourceModifiedAccessConditions *SourceModifiedAccessConditions) (appendBlobClientAppendBlockFromURLResponse, error) {
	req, err := client.appendBlockFromURLCreateRequest(ctx, comp, sourceURL, contentLength, options, cpkInfo, cpkScopeInfo, leaseAccessConditions, appendPositionAccessConditions, modifiedAccessConditions, sourceModifiedAccessConditions)
	if err != nil {
		return appendBlobClientAppendBlockFromURLResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return appendBlobClientAppendBlockFromURLResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return appendBlobClientAppendBlockFromURLResponse{}, runtime.NewResponseError(resp)
	}
	return client.appendBlockFromURLHandleResponse(resp)
}

// appendBlockFromURLCreateRequest creates the AppendBlockFromURL request.
func (client *appendBlobClient) appendBlockFromURLCreateRequest(ctx context.Context, comp Enum38, sourceURL string, contentLength int64, options *appendBlobClientAppendBlockFromURLOptions, cpkInfo *CpkInfo, cpkScopeInfo *CpkScopeInfo, leaseAccessConditions *LeaseAccessConditions, appendPositionAccessConditions *AppendPositionAccessConditions, modifiedAccessConditions *ModifiedAccessConditions, sourceModifiedAccessConditions *SourceModifiedAccessConditions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodPut, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("comp", string(comp))
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("x-ms-copy-source", sourceURL)
	if options != nil && options.SourceRange != nil {
		req.Raw().Header.Set("x-ms-source-range", *options.SourceRange)
	}
	if options != nil && options.SourceContentMD5 != nil {
		req.Raw().Header.Set("x-ms-source-content-md5", base64.StdEncoding.EncodeToString(options.SourceContentMD5))
	}
	if options != nil && options.SourceContentcrc64 != nil {
		req.Raw().Header.Set("x-ms-source-content-crc64", base64.StdEncoding.EncodeToString(options.SourceContentcrc64))
	}
	req.Raw().Header.Set("Content-Length", strconv.FormatInt(contentLength, 10))
	if options != nil && options.TransactionalContentMD5 != nil {
		req.Raw().Header.Set("Content-MD5", base64.StdEncoding.EncodeToString(options.TransactionalContentMD5))
	}
	if cpkInfo != nil && cpkInfo.EncryptionKey != nil {
		req.Raw().Header.Set("x-ms-encryption-key", *cpkInfo.EncryptionKey)
	}
	if cpkInfo != nil && cpkInfo.EncryptionKeySHA256 != nil {
		req.Raw().Header.Set("x-ms-encryption-key-sha256", *cpkInfo.EncryptionKeySHA256)
	}
	if cpkInfo != nil && cpkInfo.EncryptionAlgorithm != nil {
		req.Raw().Header.Set("x-ms-encryption-algorithm", "AES256")
	}
	if cpkScopeInfo != nil && cpkScopeInfo.EncryptionScope != nil {
		req.Raw().Header.Set("x-ms-encryption-scope", *cpkScopeInfo.EncryptionScope)
	}
	if leaseAccessConditions != nil && leaseAccessConditions.LeaseID != nil {
		req.Raw().Header.Set("x-ms-lease-id", *leaseAccessConditions.LeaseID)
	}
	if appendPositionAccessConditions != nil && appendPositionAccessConditions.MaxSize != nil {
		req.Raw().Header.Set("x-ms-blob-condition-maxsize", strconv.FormatInt(*appendPositionAccessConditions.MaxSize, 10))
	}
	if appendPositionAccessConditions != nil && appendPositionAccessConditions.AppendPosition != nil {
		req.Raw().Header.Set("x-ms-blob-condition-appendpos", strconv.FormatInt(*appendPositionAccessConditions.AppendPosition, 10))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfModifiedSince != nil {
		req.Raw().Header.Set("If-Modified-Since", modifiedAccessConditions.IfModifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfUnmodifiedSince != nil {
		req.Raw().Header.Set("If-Unmodified-Since", modifiedAccessConditions.IfUnmodifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *modifiedAccessConditions.IfMatch)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfNoneMatch != nil {
		req.Raw().Header.Set("If-None-Match", *modifiedAccessConditions.IfNoneMatch)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfTags != nil {
		req.Raw().Header.Set("x-ms-if-tags", *modifiedAccessConditions.IfTags)
	}
	if sourceModifiedAccessConditions != nil && sourceModifiedAccessConditions.SourceIfModifiedSince != nil {
		req.Raw().Header.Set("x-ms-source-if-modified-since", sourceModifiedAccessConditions.SourceIfModifiedSince.Format(time.RFC1123))
	}
	if sourceModifiedAccessConditions != nil && sourceModifiedAccessConditions.SourceIfUnmodifiedSince != nil {
		req.Raw().Header.Set("x-ms-source-if-unmodified-since", sourceModifiedAccessConditions.SourceIfUnmodifiedSince.Format(time.RFC1123))
	}
	if sourceModifiedAccessConditions != nil && sourceModifiedAccessConditions.SourceIfMatch != nil {
		req.Raw().Header.Set("x-ms-source-if-match", *sourceModifiedAccessConditions.SourceIfMatch)
	}
	if sourceModifiedAccessConditions != nil && sourceModifiedAccessConditions.SourceIfNoneMatch != nil {
		req.Raw().Header.Set("x-ms-source-if-none-match", *sourceModifiedAccessConditions.SourceIfNoneMatch)
	}
	req.Raw().Header.Set("x-ms-version", string(client.version))
	if options != nil && options.RequestID != nil {
		req.Raw().Header.Set("x-ms-client-request-id", *options.RequestID)
	}
	req.Raw().Header.Set("Accept", "application/xml")
	return req, nil
}

// appendBlockFromURLHandleResponse handles the AppendBlockFromURL response.
func (client *appendBlobClient) appendBlockFromURLHandleResponse(resp *http.Response) (appendBlobClientAppendBlockFromURLResponse, error) {
	result := appendBlobClientAppendBlockFromURLResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if val := resp.Header.Get("Last-Modified"); val != "" {
		lastModified, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return appendBlobClientAppendBlockFromURLResponse{}, err
		}
		result.LastModified = &lastModified
	}
	if val := resp.Header.Get("Content-MD5"); val != "" {
		contentMD5, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return appendBlobClientAppendBlockFromURLResponse{}, err
		}
		result.ContentMD5 = contentMD5
	}
	if val := resp.Header.Get("x-ms-content-crc64"); val != "" {
		xMSContentCRC64, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return appendBlobClientAppendBlockFromURLResponse{}, err
		}
		result.XMSContentCRC64 = xMSContentCRC64
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return appendBlobClientAppendBlockFromURLResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-blob-append-offset"); val != "" {
		result.BlobAppendOffset = &val
	}
	if val := resp.Header.Get("x-ms-blob-committed-block-count"); val != "" {
		blobCommittedBlockCount32, err := strconv.ParseInt(val, 10, 32)
		blobCommittedBlockCount := int32(blobCommittedBlockCount32)
		if err != nil {
			return appendBlobClientAppendBlockFromURLResponse{}, err
		}
		result.BlobCommittedBlockCount = &blobCommittedBlockCount
	}
	if val := resp.Header.Get("x-ms-encryption-key-sha256"); val != "" {
		result.EncryptionKeySHA256 = &val
	}
	if val := resp.Header.Get("x-ms-encryption-scope"); val != "" {
		result.EncryptionScope = &val
	}
	if val := resp.Header.Get("x-ms-request-server-encrypted"); val != "" {
		isServerEncrypted, err := strconv.ParseBool(val)
		if err != nil {
			return appendBlobClientAppendBlockFromURLResponse{}, err
		}
		result.IsServerEncrypted = &isServerEncrypted
	}
	return result, nil
}

// Create - The Create Append Blob operation creates a new append blob.
// If the operation fails it returns an *azcore.ResponseError type.
// contentLength - The length of the request.
// options - appendBlobClientCreateOptions contains the optional parameters for the appendBlobClient.Create method.
// BlobHTTPHeaders - BlobHTTPHeaders contains a group of parameters for the client.SetHTTPHeaders method.
// LeaseAccessConditions - LeaseAccessConditions contains a group of parameters for the containerClient.GetProperties method.
// CpkInfo - CpkInfo contains a group of parameters for the client.Download method.
// CpkScopeInfo - CpkScopeInfo contains a group of parameters for the client.SetMetadata method.
// ModifiedAccessConditions - ModifiedAccessConditions contains a group of parameters for the containerClient.Delete method.
func (client *appendBlobClient) Create(ctx context.Context, contentLength int64, options *appendBlobClientCreateOptions, blobHTTPHeaders *BlobHTTPHeaders, leaseAccessConditions *LeaseAccessConditions, cpkInfo *CpkInfo, cpkScopeInfo *CpkScopeInfo, modifiedAccessConditions *ModifiedAccessConditions) (appendBlobClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, contentLength, options, blobHTTPHeaders, leaseAccessConditions, cpkInfo, cpkScopeInfo, modifiedAccessConditions)
	if err != nil {
		return appendBlobClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return appendBlobClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return appendBlobClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *appendBlobClient) createCreateRequest(ctx context.Context, contentLength int64, options *appendBlobClientCreateOptions, blobHTTPHeaders *BlobHTTPHeaders, leaseAccessConditions *LeaseAccessConditions, cpkInfo *CpkInfo, cpkScopeInfo *CpkScopeInfo, modifiedAccessConditions *ModifiedAccessConditions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodPut, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("x-ms-blob-type", "AppendBlob")
	req.Raw().Header.Set("Content-Length", strconv.FormatInt(contentLength, 10))
	if blobHTTPHeaders != nil && blobHTTPHeaders.BlobContentType != nil {
		req.Raw().Header.Set("x-ms-blob-content-type", *blobHTTPHeaders.BlobContentType)
	}
	if blobHTTPHeaders != nil && blobHTTPHeaders.BlobContentEncoding != nil {
		req.Raw().Header.Set("x-ms-blob-content-encoding", *blobHTTPHeaders.BlobContentEncoding)
	}
	if blobHTTPHeaders != nil && blobHTTPHeaders.BlobContentLanguage != nil {
		req.Raw().Header.Set("x-ms-blob-content-language", *blobHTTPHeaders.BlobContentLanguage)
	}
	if blobHTTPHeaders != nil && blobHTTPHeaders.BlobContentMD5 != nil {
		req.Raw().Header.Set("x-ms-blob-content-md5", base64.StdEncoding.EncodeToString(blobHTTPHeaders.BlobContentMD5))
	}
	if blobHTTPHeaders != nil && blobHTTPHeaders.BlobCacheControl != nil {
		req.Raw().Header.Set("x-ms-blob-cache-control", *blobHTTPHeaders.BlobCacheControl)
	}
	if options != nil && options.Metadata != nil {
		for k, v := range options.Metadata {
			req.Raw().Header.Set("x-ms-meta-"+k, v)
		}
	}
	if leaseAccessConditions != nil && leaseAccessConditions.LeaseID != nil {
		req.Raw().Header.Set("x-ms-lease-id", *leaseAccessConditions.LeaseID)
	}
	if blobHTTPHeaders != nil && blobHTTPHeaders.BlobContentDisposition != nil {
		req.Raw().Header.Set("x-ms-blob-content-disposition", *blobHTTPHeaders.BlobContentDisposition)
	}
	if cpkInfo != nil && cpkInfo.EncryptionKey != nil {
		req.Raw().Header.Set("x-ms-encryption-key", *cpkInfo.EncryptionKey)
	}
	if cpkInfo != nil && cpkInfo.EncryptionKeySHA256 != nil {
		req.Raw().Header.Set("x-ms-encryption-key-sha256", *cpkInfo.EncryptionKeySHA256)
	}
	if cpkInfo != nil && cpkInfo.EncryptionAlgorithm != nil {
		req.Raw().Header.Set("x-ms-encryption-algorithm", "AES256")
	}
	if cpkScopeInfo != nil && cpkScopeInfo.EncryptionScope != nil {
		req.Raw().Header.Set("x-ms-encryption-scope", *cpkScopeInfo.EncryptionScope)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfModifiedSince != nil {
		req.Raw().Header.Set("If-Modified-Since", modifiedAccessConditions.IfModifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfUnmodifiedSince != nil {
		req.Raw().Header.Set("If-Unmodified-Since", modifiedAccessConditions.IfUnmodifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *modifiedAccessConditions.IfMatch)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfNoneMatch != nil {
		req.Raw().Header.Set("If-None-Match", *modifiedAccessConditions.IfNoneMatch)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfTags != nil {
		req.Raw().Header.Set("x-ms-if-tags", *modifiedAccessConditions.IfTags)
	}
	req.Raw().Header.Set("x-ms-version", string(client.version))
	if options != nil && options.RequestID != nil {
		req.Raw().Header.Set("x-ms-client-request-id", *options.RequestID)
	}
	if options != nil && options.BlobTagsString != nil {
		req.Raw().Header.Set("x-ms-tags", *options.BlobTagsString)
	}
	if options != nil && options.ImmutabilityPolicyExpiry != nil {
		req.Raw().Header.Set("x-ms-immutability-policy-until-date", options.ImmutabilityPolicyExpiry.Format(time.RFC1123))
	}
	if options != nil && options.ImmutabilityPolicyMode != nil {
		req.Raw().Header.Set("x-ms-immutability-policy-mode", string(*options.ImmutabilityPolicyMode))
	}
	if options != nil && options.LegalHold != nil {
		req.Raw().Header.Set("x-ms-legal-hold", strconv.FormatBool(*options.LegalHold))
	}
	req.Raw().Header.Set("Accept", "application/xml")
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *appendBlobClient) createHandleResponse(resp *http.Response) (appendBlobClientCreateResponse, error) {
	result := appendBlobClientCreateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if val := resp.Header.Get("Last-Modified"); val != "" {
		lastModified, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return appendBlobClientCreateResponse{}, err
		}
		result.LastModified = &lastModified
	}
	if val := resp.Header.Get("Content-MD5"); val != "" {
		contentMD5, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return appendBlobClientCreateResponse{}, err
		}
		result.ContentMD5 = contentMD5
	}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("x-ms-version-id"); val != "" {
		result.VersionID = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return appendBlobClientCreateResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-request-server-encrypted"); val != "" {
		isServerEncrypted, err := strconv.ParseBool(val)
		if err != nil {
			return appendBlobClientCreateResponse{}, err
		}
		result.IsServerEncrypted = &isServerEncrypted
	}
	if val := resp.Header.Get("x-ms-encryption-key-sha256"); val != "" {
		result.EncryptionKeySHA256 = &val
	}
	if val := resp.Header.Get("x-ms-encryption-scope"); val != "" {
		result.EncryptionScope = &val
	}
	return result, nil
}

// Seal - The Seal operation seals the Append Blob to make it read-only. Seal is supported only on version 2019-12-12 version
// or later.
// If the operation fails it returns an *azcore.ResponseError type.
// options - appendBlobClientSealOptions contains the optional parameters for the appendBlobClient.Seal method.
// LeaseAccessConditions - LeaseAccessConditions contains a group of parameters for the containerClient.GetProperties method.
// ModifiedAccessConditions - ModifiedAccessConditions contains a group of parameters for the containerClient.Delete method.
// AppendPositionAccessConditions - AppendPositionAccessConditions contains a group of parameters for the appendBlobClient.AppendBlock
// method.
func (client *appendBlobClient) Seal(ctx context.Context, comp Enum39, options *appendBlobClientSealOptions, leaseAccessConditions *LeaseAccessConditions, modifiedAccessConditions *ModifiedAccessConditions, appendPositionAccessConditions *AppendPositionAccessConditions) (appendBlobClientSealResponse, error) {
	req, err := client.sealCreateRequest(ctx, comp, options, leaseAccessConditions, modifiedAccessConditions, appendPositionAccessConditions)
	if err != nil {
		return appendBlobClientSealResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return appendBlobClientSealResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return appendBlobClientSealResponse{}, runtime.NewResponseError(resp)
	}
	return client.sealHandleResponse(resp)
}

// sealCreateRequest creates the Seal request.
func (client *appendBlobClient) sealCreateRequest(ctx context.Context, comp Enum39, options *appendBlobClientSealOptions, leaseAccessConditions *LeaseAccessConditions, modifiedAccessConditions *ModifiedAccessConditions, appendPositionAccessConditions *AppendPositionAccessConditions) (*policy.Request, error) {
	req, err := runtime.NewRequest(ctx, http.MethodPut, client.endpoint)
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("comp", string(comp))
	if options != nil && options.Timeout != nil {
		reqQP.Set("timeout", strconv.FormatInt(int64(*options.Timeout), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("x-ms-version", string(client.version))
	if options != nil && options.RequestID != nil {
		req.Raw().Header.Set("x-ms-client-request-id", *options.RequestID)
	}
	if leaseAccessConditions != nil && leaseAccessConditions.LeaseID != nil {
		req.Raw().Header.Set("x-ms-lease-id", *leaseAccessConditions.LeaseID)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfModifiedSince != nil {
		req.Raw().Header.Set("If-Modified-Since", modifiedAccessConditions.IfModifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfUnmodifiedSince != nil {
		req.Raw().Header.Set("If-Unmodified-Since", modifiedAccessConditions.IfUnmodifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *modifiedAccessConditions.IfMatch)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfNoneMatch != nil {
		req.Raw().Header.Set("If-None-Match", *modifiedAccessConditions.IfNoneMatch)
	}
	if appendPositionAccessConditions != nil && appendPositionAccessConditions.AppendPosition != nil {
		req.Raw().Header.Set("x-ms-blob-condition-appendpos", strconv.FormatInt(*appendPositionAccessConditions.AppendPosition, 10))
	}
	req.Raw().Header.Set("Accept", "application/xml")
	return req, nil
}

// sealHandleResponse handles the Seal response.
func (client *appendBlobClient) sealHandleResponse(resp *http.Response) (appendBlobClientSealResponse, error) {
	result := appendBlobClientSealResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if val := resp.Header.Get("Last-Modified"); val != "" {
		lastModified, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return appendBlobClientSealResponse{}, err
		}
		result.LastModified = &lastModified
	}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return appendBlobClientSealResponse{}, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-blob-sealed"); val != "" {
		isSealed, err := strconv.ParseBool(val)
		if err != nil {
			return appendBlobClientSealResponse{}, err
		}
		result.IsSealed = &isSealed
	}
	return result, nil
}
