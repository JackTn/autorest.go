//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azspark

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// SessionClient contains the methods for the SparkSession group.
// Don't use this type directly, use a constructor function instead.
type SessionClient struct {
	internal       *azcore.Client
	endpoint       string
	livyAPIVersion string
	sparkPoolName  string
}

// CancelSparkSession - Cancels a running spark session.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - sessionID - Identifier for the session.
//   - options - SessionClientCancelSparkSessionOptions contains the optional parameters for the SessionClient.CancelSparkSession
//     method.
func (client *SessionClient) CancelSparkSession(ctx context.Context, sessionID int32, options *SessionClientCancelSparkSessionOptions) (SessionClientCancelSparkSessionResponse, error) {
	req, err := client.cancelSparkSessionCreateRequest(ctx, sessionID, options)
	if err != nil {
		return SessionClientCancelSparkSessionResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SessionClientCancelSparkSessionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SessionClientCancelSparkSessionResponse{}, runtime.NewResponseError(resp)
	}
	return SessionClientCancelSparkSessionResponse{}, nil
}

// cancelSparkSessionCreateRequest creates the CancelSparkSession request.
func (client *SessionClient) cancelSparkSessionCreateRequest(ctx context.Context, sessionID int32, options *SessionClientCancelSparkSessionOptions) (*policy.Request, error) {
	urlPath := "/livyApi/versions/{livyApiVersion}/sparkPools/{sparkPoolName}/sessions/{sessionId}"
	urlPath = strings.ReplaceAll(urlPath, "{livyApiVersion}", client.livyAPIVersion)
	urlPath = strings.ReplaceAll(urlPath, "{sparkPoolName}", client.sparkPoolName)
	urlPath = strings.ReplaceAll(urlPath, "{sessionId}", url.PathEscape(strconv.FormatInt(int64(sessionID), 10)))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// CancelSparkStatement - Kill a statement within a session.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - sessionID - Identifier for the session.
//   - statementID - Identifier for the statement.
//   - options - SessionClientCancelSparkStatementOptions contains the optional parameters for the SessionClient.CancelSparkStatement
//     method.
func (client *SessionClient) CancelSparkStatement(ctx context.Context, sessionID int32, statementID int32, options *SessionClientCancelSparkStatementOptions) (SessionClientCancelSparkStatementResponse, error) {
	req, err := client.cancelSparkStatementCreateRequest(ctx, sessionID, statementID, options)
	if err != nil {
		return SessionClientCancelSparkStatementResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SessionClientCancelSparkStatementResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SessionClientCancelSparkStatementResponse{}, runtime.NewResponseError(resp)
	}
	return client.cancelSparkStatementHandleResponse(resp)
}

// cancelSparkStatementCreateRequest creates the CancelSparkStatement request.
func (client *SessionClient) cancelSparkStatementCreateRequest(ctx context.Context, sessionID int32, statementID int32, options *SessionClientCancelSparkStatementOptions) (*policy.Request, error) {
	urlPath := "/livyApi/versions/{livyApiVersion}/sparkPools/{sparkPoolName}/sessions/{sessionId}/statements/{statementId}/cancel"
	urlPath = strings.ReplaceAll(urlPath, "{livyApiVersion}", client.livyAPIVersion)
	urlPath = strings.ReplaceAll(urlPath, "{sparkPoolName}", client.sparkPoolName)
	urlPath = strings.ReplaceAll(urlPath, "{sessionId}", url.PathEscape(strconv.FormatInt(int64(sessionID), 10)))
	urlPath = strings.ReplaceAll(urlPath, "{statementId}", url.PathEscape(strconv.FormatInt(int64(statementID), 10)))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// cancelSparkStatementHandleResponse handles the CancelSparkStatement response.
func (client *SessionClient) cancelSparkStatementHandleResponse(resp *http.Response) (SessionClientCancelSparkStatementResponse, error) {
	result := SessionClientCancelSparkStatementResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StatementCancellationResult); err != nil {
		return SessionClientCancelSparkStatementResponse{}, err
	}
	return result, nil
}

// CreateSparkSession - Create new spark session.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - sparkSessionOptions - Livy compatible batch job request payload.
//   - options - SessionClientCreateSparkSessionOptions contains the optional parameters for the SessionClient.CreateSparkSession
//     method.
func (client *SessionClient) CreateSparkSession(ctx context.Context, sparkSessionOptions SessionOptions, options *SessionClientCreateSparkSessionOptions) (SessionClientCreateSparkSessionResponse, error) {
	req, err := client.createSparkSessionCreateRequest(ctx, sparkSessionOptions, options)
	if err != nil {
		return SessionClientCreateSparkSessionResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SessionClientCreateSparkSessionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SessionClientCreateSparkSessionResponse{}, runtime.NewResponseError(resp)
	}
	return client.createSparkSessionHandleResponse(resp)
}

// createSparkSessionCreateRequest creates the CreateSparkSession request.
func (client *SessionClient) createSparkSessionCreateRequest(ctx context.Context, sparkSessionOptions SessionOptions, options *SessionClientCreateSparkSessionOptions) (*policy.Request, error) {
	urlPath := "/livyApi/versions/{livyApiVersion}/sparkPools/{sparkPoolName}/sessions"
	urlPath = strings.ReplaceAll(urlPath, "{livyApiVersion}", client.livyAPIVersion)
	urlPath = strings.ReplaceAll(urlPath, "{sparkPoolName}", client.sparkPoolName)
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Detailed != nil {
		reqQP.Set("detailed", strconv.FormatBool(*options.Detailed))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, sparkSessionOptions)
}

// createSparkSessionHandleResponse handles the CreateSparkSession response.
func (client *SessionClient) createSparkSessionHandleResponse(resp *http.Response) (SessionClientCreateSparkSessionResponse, error) {
	result := SessionClientCreateSparkSessionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Session); err != nil {
		return SessionClientCreateSparkSessionResponse{}, err
	}
	return result, nil
}

// CreateSparkStatement - Create statement within a spark session.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - sessionID - Identifier for the session.
//   - sparkStatementOptions - Livy compatible batch job request payload.
//   - options - SessionClientCreateSparkStatementOptions contains the optional parameters for the SessionClient.CreateSparkStatement
//     method.
func (client *SessionClient) CreateSparkStatement(ctx context.Context, sessionID int32, sparkStatementOptions StatementOptions, options *SessionClientCreateSparkStatementOptions) (SessionClientCreateSparkStatementResponse, error) {
	req, err := client.createSparkStatementCreateRequest(ctx, sessionID, sparkStatementOptions, options)
	if err != nil {
		return SessionClientCreateSparkStatementResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SessionClientCreateSparkStatementResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SessionClientCreateSparkStatementResponse{}, runtime.NewResponseError(resp)
	}
	return client.createSparkStatementHandleResponse(resp)
}

// createSparkStatementCreateRequest creates the CreateSparkStatement request.
func (client *SessionClient) createSparkStatementCreateRequest(ctx context.Context, sessionID int32, sparkStatementOptions StatementOptions, options *SessionClientCreateSparkStatementOptions) (*policy.Request, error) {
	urlPath := "/livyApi/versions/{livyApiVersion}/sparkPools/{sparkPoolName}/sessions/{sessionId}/statements"
	urlPath = strings.ReplaceAll(urlPath, "{livyApiVersion}", client.livyAPIVersion)
	urlPath = strings.ReplaceAll(urlPath, "{sparkPoolName}", client.sparkPoolName)
	urlPath = strings.ReplaceAll(urlPath, "{sessionId}", url.PathEscape(strconv.FormatInt(int64(sessionID), 10)))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, sparkStatementOptions)
}

// createSparkStatementHandleResponse handles the CreateSparkStatement response.
func (client *SessionClient) createSparkStatementHandleResponse(resp *http.Response) (SessionClientCreateSparkStatementResponse, error) {
	result := SessionClientCreateSparkStatementResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Statement); err != nil {
		return SessionClientCreateSparkStatementResponse{}, err
	}
	return result, nil
}

// GetSparkSession - Gets a single spark session.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - sessionID - Identifier for the session.
//   - options - SessionClientGetSparkSessionOptions contains the optional parameters for the SessionClient.GetSparkSession method.
func (client *SessionClient) GetSparkSession(ctx context.Context, sessionID int32, options *SessionClientGetSparkSessionOptions) (SessionClientGetSparkSessionResponse, error) {
	req, err := client.getSparkSessionCreateRequest(ctx, sessionID, options)
	if err != nil {
		return SessionClientGetSparkSessionResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SessionClientGetSparkSessionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SessionClientGetSparkSessionResponse{}, runtime.NewResponseError(resp)
	}
	return client.getSparkSessionHandleResponse(resp)
}

// getSparkSessionCreateRequest creates the GetSparkSession request.
func (client *SessionClient) getSparkSessionCreateRequest(ctx context.Context, sessionID int32, options *SessionClientGetSparkSessionOptions) (*policy.Request, error) {
	urlPath := "/livyApi/versions/{livyApiVersion}/sparkPools/{sparkPoolName}/sessions/{sessionId}"
	urlPath = strings.ReplaceAll(urlPath, "{livyApiVersion}", client.livyAPIVersion)
	urlPath = strings.ReplaceAll(urlPath, "{sparkPoolName}", client.sparkPoolName)
	urlPath = strings.ReplaceAll(urlPath, "{sessionId}", url.PathEscape(strconv.FormatInt(int64(sessionID), 10)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Detailed != nil {
		reqQP.Set("detailed", strconv.FormatBool(*options.Detailed))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getSparkSessionHandleResponse handles the GetSparkSession response.
func (client *SessionClient) getSparkSessionHandleResponse(resp *http.Response) (SessionClientGetSparkSessionResponse, error) {
	result := SessionClientGetSparkSessionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Session); err != nil {
		return SessionClientGetSparkSessionResponse{}, err
	}
	return result, nil
}

// GetSparkSessions - List all spark sessions which are running under a particular spark pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - options - SessionClientGetSparkSessionsOptions contains the optional parameters for the SessionClient.GetSparkSessions
//     method.
func (client *SessionClient) GetSparkSessions(ctx context.Context, options *SessionClientGetSparkSessionsOptions) (SessionClientGetSparkSessionsResponse, error) {
	req, err := client.getSparkSessionsCreateRequest(ctx, options)
	if err != nil {
		return SessionClientGetSparkSessionsResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SessionClientGetSparkSessionsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SessionClientGetSparkSessionsResponse{}, runtime.NewResponseError(resp)
	}
	return client.getSparkSessionsHandleResponse(resp)
}

// getSparkSessionsCreateRequest creates the GetSparkSessions request.
func (client *SessionClient) getSparkSessionsCreateRequest(ctx context.Context, options *SessionClientGetSparkSessionsOptions) (*policy.Request, error) {
	urlPath := "/livyApi/versions/{livyApiVersion}/sparkPools/{sparkPoolName}/sessions"
	urlPath = strings.ReplaceAll(urlPath, "{livyApiVersion}", client.livyAPIVersion)
	urlPath = strings.ReplaceAll(urlPath, "{sparkPoolName}", client.sparkPoolName)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.From != nil {
		reqQP.Set("from", strconv.FormatInt(int64(*options.From), 10))
	}
	if options != nil && options.Size != nil {
		reqQP.Set("size", strconv.FormatInt(int64(*options.Size), 10))
	}
	if options != nil && options.Detailed != nil {
		reqQP.Set("detailed", strconv.FormatBool(*options.Detailed))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getSparkSessionsHandleResponse handles the GetSparkSessions response.
func (client *SessionClient) getSparkSessionsHandleResponse(resp *http.Response) (SessionClientGetSparkSessionsResponse, error) {
	result := SessionClientGetSparkSessionsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SessionCollection); err != nil {
		return SessionClientGetSparkSessionsResponse{}, err
	}
	return result, nil
}

// GetSparkStatement - Gets a single statement within a spark session.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - sessionID - Identifier for the session.
//   - statementID - Identifier for the statement.
//   - options - SessionClientGetSparkStatementOptions contains the optional parameters for the SessionClient.GetSparkStatement
//     method.
func (client *SessionClient) GetSparkStatement(ctx context.Context, sessionID int32, statementID int32, options *SessionClientGetSparkStatementOptions) (SessionClientGetSparkStatementResponse, error) {
	req, err := client.getSparkStatementCreateRequest(ctx, sessionID, statementID, options)
	if err != nil {
		return SessionClientGetSparkStatementResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SessionClientGetSparkStatementResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SessionClientGetSparkStatementResponse{}, runtime.NewResponseError(resp)
	}
	return client.getSparkStatementHandleResponse(resp)
}

// getSparkStatementCreateRequest creates the GetSparkStatement request.
func (client *SessionClient) getSparkStatementCreateRequest(ctx context.Context, sessionID int32, statementID int32, options *SessionClientGetSparkStatementOptions) (*policy.Request, error) {
	urlPath := "/livyApi/versions/{livyApiVersion}/sparkPools/{sparkPoolName}/sessions/{sessionId}/statements/{statementId}"
	urlPath = strings.ReplaceAll(urlPath, "{livyApiVersion}", client.livyAPIVersion)
	urlPath = strings.ReplaceAll(urlPath, "{sparkPoolName}", client.sparkPoolName)
	urlPath = strings.ReplaceAll(urlPath, "{sessionId}", url.PathEscape(strconv.FormatInt(int64(sessionID), 10)))
	urlPath = strings.ReplaceAll(urlPath, "{statementId}", url.PathEscape(strconv.FormatInt(int64(statementID), 10)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getSparkStatementHandleResponse handles the GetSparkStatement response.
func (client *SessionClient) getSparkStatementHandleResponse(resp *http.Response) (SessionClientGetSparkStatementResponse, error) {
	result := SessionClientGetSparkStatementResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Statement); err != nil {
		return SessionClientGetSparkStatementResponse{}, err
	}
	return result, nil
}

// GetSparkStatements - Gets a list of statements within a spark session.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - sessionID - Identifier for the session.
//   - options - SessionClientGetSparkStatementsOptions contains the optional parameters for the SessionClient.GetSparkStatements
//     method.
func (client *SessionClient) GetSparkStatements(ctx context.Context, sessionID int32, options *SessionClientGetSparkStatementsOptions) (SessionClientGetSparkStatementsResponse, error) {
	req, err := client.getSparkStatementsCreateRequest(ctx, sessionID, options)
	if err != nil {
		return SessionClientGetSparkStatementsResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SessionClientGetSparkStatementsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SessionClientGetSparkStatementsResponse{}, runtime.NewResponseError(resp)
	}
	return client.getSparkStatementsHandleResponse(resp)
}

// getSparkStatementsCreateRequest creates the GetSparkStatements request.
func (client *SessionClient) getSparkStatementsCreateRequest(ctx context.Context, sessionID int32, options *SessionClientGetSparkStatementsOptions) (*policy.Request, error) {
	urlPath := "/livyApi/versions/{livyApiVersion}/sparkPools/{sparkPoolName}/sessions/{sessionId}/statements"
	urlPath = strings.ReplaceAll(urlPath, "{livyApiVersion}", client.livyAPIVersion)
	urlPath = strings.ReplaceAll(urlPath, "{sparkPoolName}", client.sparkPoolName)
	urlPath = strings.ReplaceAll(urlPath, "{sessionId}", url.PathEscape(strconv.FormatInt(int64(sessionID), 10)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getSparkStatementsHandleResponse handles the GetSparkStatements response.
func (client *SessionClient) getSparkStatementsHandleResponse(resp *http.Response) (SessionClientGetSparkStatementsResponse, error) {
	result := SessionClientGetSparkStatementsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StatementCollection); err != nil {
		return SessionClientGetSparkStatementsResponse{}, err
	}
	return result, nil
}

// ResetSparkSessionTimeout - Sends a keep alive call to the current session to reset the session timeout.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-12-01
//   - sessionID - Identifier for the session.
//   - options - SessionClientResetSparkSessionTimeoutOptions contains the optional parameters for the SessionClient.ResetSparkSessionTimeout
//     method.
func (client *SessionClient) ResetSparkSessionTimeout(ctx context.Context, sessionID int32, options *SessionClientResetSparkSessionTimeoutOptions) (SessionClientResetSparkSessionTimeoutResponse, error) {
	req, err := client.resetSparkSessionTimeoutCreateRequest(ctx, sessionID, options)
	if err != nil {
		return SessionClientResetSparkSessionTimeoutResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SessionClientResetSparkSessionTimeoutResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SessionClientResetSparkSessionTimeoutResponse{}, runtime.NewResponseError(resp)
	}
	return SessionClientResetSparkSessionTimeoutResponse{}, nil
}

// resetSparkSessionTimeoutCreateRequest creates the ResetSparkSessionTimeout request.
func (client *SessionClient) resetSparkSessionTimeoutCreateRequest(ctx context.Context, sessionID int32, options *SessionClientResetSparkSessionTimeoutOptions) (*policy.Request, error) {
	urlPath := "/livyApi/versions/{livyApiVersion}/sparkPools/{sparkPoolName}/sessions/{sessionId}/reset-timeout"
	urlPath = strings.ReplaceAll(urlPath, "{livyApiVersion}", client.livyAPIVersion)
	urlPath = strings.ReplaceAll(urlPath, "{sparkPoolName}", client.sparkPoolName)
	urlPath = strings.ReplaceAll(urlPath, "{sessionId}", url.PathEscape(strconv.FormatInt(int64(sessionID), 10)))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}