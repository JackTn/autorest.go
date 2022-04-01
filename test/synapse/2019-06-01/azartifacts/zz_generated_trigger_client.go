//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

type triggerClient struct {
	endpoint string
	pl       runtime.Pipeline
}

// newTriggerClient creates a new instance of triggerClient with the specified values.
// endpoint - The workspace development endpoint, for example https://myworkspace.dev.azuresynapse.net.
// pl - the pipeline used for sending requests and handling responses.
func newTriggerClient(endpoint string, pl runtime.Pipeline) *triggerClient {
	client := &triggerClient{
		endpoint: endpoint,
		pl:       pl,
	}
	return client
}

// BeginCreateOrUpdateTrigger - Creates or updates a trigger.
// If the operation fails it returns an *azcore.ResponseError type.
// triggerName - The trigger name.
// trigger - Trigger resource definition.
// options - triggerClientBeginCreateOrUpdateTriggerOptions contains the optional parameters for the triggerClient.BeginCreateOrUpdateTrigger
// method.
func (client *triggerClient) BeginCreateOrUpdateTrigger(ctx context.Context, triggerName string, trigger TriggerResource, options *triggerClientBeginCreateOrUpdateTriggerOptions) (*runtime.Poller[triggerClientCreateOrUpdateTriggerResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdateTrigger(ctx, triggerName, trigger, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[triggerClientCreateOrUpdateTriggerResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[triggerClientCreateOrUpdateTriggerResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdateTrigger - Creates or updates a trigger.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *triggerClient) createOrUpdateTrigger(ctx context.Context, triggerName string, trigger TriggerResource, options *triggerClientBeginCreateOrUpdateTriggerOptions) (*http.Response, error) {
	req, err := client.createOrUpdateTriggerCreateRequest(ctx, triggerName, trigger, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateTriggerCreateRequest creates the CreateOrUpdateTrigger request.
func (client *triggerClient) createOrUpdateTriggerCreateRequest(ctx context.Context, triggerName string, trigger TriggerResource, options *triggerClientBeginCreateOrUpdateTriggerOptions) (*policy.Request, error) {
	urlPath := "/triggers/{triggerName}"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, trigger)
}

// BeginDeleteTrigger - Deletes a trigger.
// If the operation fails it returns an *azcore.ResponseError type.
// triggerName - The trigger name.
// options - triggerClientBeginDeleteTriggerOptions contains the optional parameters for the triggerClient.BeginDeleteTrigger
// method.
func (client *triggerClient) BeginDeleteTrigger(ctx context.Context, triggerName string, options *triggerClientBeginDeleteTriggerOptions) (*runtime.Poller[triggerClientDeleteTriggerResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteTrigger(ctx, triggerName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[triggerClientDeleteTriggerResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[triggerClientDeleteTriggerResponse](options.ResumeToken, client.pl, nil)
	}
}

// DeleteTrigger - Deletes a trigger.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *triggerClient) deleteTrigger(ctx context.Context, triggerName string, options *triggerClientBeginDeleteTriggerOptions) (*http.Response, error) {
	req, err := client.deleteTriggerCreateRequest(ctx, triggerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteTriggerCreateRequest creates the DeleteTrigger request.
func (client *triggerClient) deleteTriggerCreateRequest(ctx context.Context, triggerName string, options *triggerClientBeginDeleteTriggerOptions) (*policy.Request, error) {
	urlPath := "/triggers/{triggerName}"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// GetEventSubscriptionStatus - Get a trigger's event subscription status.
// If the operation fails it returns an *azcore.ResponseError type.
// triggerName - The trigger name.
// options - triggerClientGetEventSubscriptionStatusOptions contains the optional parameters for the triggerClient.GetEventSubscriptionStatus
// method.
func (client *triggerClient) GetEventSubscriptionStatus(ctx context.Context, triggerName string, options *triggerClientGetEventSubscriptionStatusOptions) (triggerClientGetEventSubscriptionStatusResponse, error) {
	req, err := client.getEventSubscriptionStatusCreateRequest(ctx, triggerName, options)
	if err != nil {
		return triggerClientGetEventSubscriptionStatusResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return triggerClientGetEventSubscriptionStatusResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return triggerClientGetEventSubscriptionStatusResponse{}, runtime.NewResponseError(resp)
	}
	return client.getEventSubscriptionStatusHandleResponse(resp)
}

// getEventSubscriptionStatusCreateRequest creates the GetEventSubscriptionStatus request.
func (client *triggerClient) getEventSubscriptionStatusCreateRequest(ctx context.Context, triggerName string, options *triggerClientGetEventSubscriptionStatusOptions) (*policy.Request, error) {
	urlPath := "/triggers/{triggerName}/getEventSubscriptionStatus"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getEventSubscriptionStatusHandleResponse handles the GetEventSubscriptionStatus response.
func (client *triggerClient) getEventSubscriptionStatusHandleResponse(resp *http.Response) (triggerClientGetEventSubscriptionStatusResponse, error) {
	result := triggerClientGetEventSubscriptionStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TriggerSubscriptionOperationStatus); err != nil {
		return triggerClientGetEventSubscriptionStatusResponse{}, err
	}
	return result, nil
}

// GetTrigger - Gets a trigger.
// If the operation fails it returns an *azcore.ResponseError type.
// triggerName - The trigger name.
// options - triggerClientGetTriggerOptions contains the optional parameters for the triggerClient.GetTrigger method.
func (client *triggerClient) GetTrigger(ctx context.Context, triggerName string, options *triggerClientGetTriggerOptions) (triggerClientGetTriggerResponse, error) {
	req, err := client.getTriggerCreateRequest(ctx, triggerName, options)
	if err != nil {
		return triggerClientGetTriggerResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return triggerClientGetTriggerResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNotModified) {
		return triggerClientGetTriggerResponse{}, runtime.NewResponseError(resp)
	}
	return client.getTriggerHandleResponse(resp)
}

// getTriggerCreateRequest creates the GetTrigger request.
func (client *triggerClient) getTriggerCreateRequest(ctx context.Context, triggerName string, options *triggerClientGetTriggerOptions) (*policy.Request, error) {
	urlPath := "/triggers/{triggerName}"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getTriggerHandleResponse handles the GetTrigger response.
func (client *triggerClient) getTriggerHandleResponse(resp *http.Response) (triggerClientGetTriggerResponse, error) {
	result := triggerClientGetTriggerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TriggerResource); err != nil {
		return triggerClientGetTriggerResponse{}, err
	}
	return result, nil
}

// GetTriggersByWorkspace - Lists triggers.
// If the operation fails it returns an *azcore.ResponseError type.
// options - triggerClientGetTriggersByWorkspaceOptions contains the optional parameters for the triggerClient.GetTriggersByWorkspace
// method.
func (client *triggerClient) GetTriggersByWorkspace(options *triggerClientGetTriggersByWorkspaceOptions) *runtime.Pager[triggerClientGetTriggersByWorkspaceResponse] {
	return runtime.NewPager(runtime.PageProcessor[triggerClientGetTriggersByWorkspaceResponse]{
		More: func(page triggerClientGetTriggersByWorkspaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *triggerClientGetTriggersByWorkspaceResponse) (triggerClientGetTriggersByWorkspaceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.getTriggersByWorkspaceCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return triggerClientGetTriggersByWorkspaceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return triggerClientGetTriggersByWorkspaceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return triggerClientGetTriggersByWorkspaceResponse{}, runtime.NewResponseError(resp)
			}
			return client.getTriggersByWorkspaceHandleResponse(resp)
		},
	})
}

// getTriggersByWorkspaceCreateRequest creates the GetTriggersByWorkspace request.
func (client *triggerClient) getTriggersByWorkspaceCreateRequest(ctx context.Context, options *triggerClientGetTriggersByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/triggers"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getTriggersByWorkspaceHandleResponse handles the GetTriggersByWorkspace response.
func (client *triggerClient) getTriggersByWorkspaceHandleResponse(resp *http.Response) (triggerClientGetTriggersByWorkspaceResponse, error) {
	result := triggerClientGetTriggersByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TriggerListResponse); err != nil {
		return triggerClientGetTriggersByWorkspaceResponse{}, err
	}
	return result, nil
}

// BeginStartTrigger - Starts a trigger.
// If the operation fails it returns an *azcore.ResponseError type.
// triggerName - The trigger name.
// options - triggerClientBeginStartTriggerOptions contains the optional parameters for the triggerClient.BeginStartTrigger
// method.
func (client *triggerClient) BeginStartTrigger(ctx context.Context, triggerName string, options *triggerClientBeginStartTriggerOptions) (*runtime.Poller[triggerClientStartTriggerResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.startTrigger(ctx, triggerName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[triggerClientStartTriggerResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[triggerClientStartTriggerResponse](options.ResumeToken, client.pl, nil)
	}
}

// StartTrigger - Starts a trigger.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *triggerClient) startTrigger(ctx context.Context, triggerName string, options *triggerClientBeginStartTriggerOptions) (*http.Response, error) {
	req, err := client.startTriggerCreateRequest(ctx, triggerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// startTriggerCreateRequest creates the StartTrigger request.
func (client *triggerClient) startTriggerCreateRequest(ctx context.Context, triggerName string, options *triggerClientBeginStartTriggerOptions) (*policy.Request, error) {
	urlPath := "/triggers/{triggerName}/start"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginStopTrigger - Stops a trigger.
// If the operation fails it returns an *azcore.ResponseError type.
// triggerName - The trigger name.
// options - triggerClientBeginStopTriggerOptions contains the optional parameters for the triggerClient.BeginStopTrigger
// method.
func (client *triggerClient) BeginStopTrigger(ctx context.Context, triggerName string, options *triggerClientBeginStopTriggerOptions) (*runtime.Poller[triggerClientStopTriggerResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.stopTrigger(ctx, triggerName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[triggerClientStopTriggerResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[triggerClientStopTriggerResponse](options.ResumeToken, client.pl, nil)
	}
}

// StopTrigger - Stops a trigger.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *triggerClient) stopTrigger(ctx context.Context, triggerName string, options *triggerClientBeginStopTriggerOptions) (*http.Response, error) {
	req, err := client.stopTriggerCreateRequest(ctx, triggerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// stopTriggerCreateRequest creates the StopTrigger request.
func (client *triggerClient) stopTriggerCreateRequest(ctx context.Context, triggerName string, options *triggerClientBeginStopTriggerOptions) (*policy.Request, error) {
	urlPath := "/triggers/{triggerName}/stop"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginSubscribeTriggerToEvents - Subscribe event trigger to events.
// If the operation fails it returns an *azcore.ResponseError type.
// triggerName - The trigger name.
// options - triggerClientBeginSubscribeTriggerToEventsOptions contains the optional parameters for the triggerClient.BeginSubscribeTriggerToEvents
// method.
func (client *triggerClient) BeginSubscribeTriggerToEvents(ctx context.Context, triggerName string, options *triggerClientBeginSubscribeTriggerToEventsOptions) (*runtime.Poller[triggerClientSubscribeTriggerToEventsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.subscribeTriggerToEvents(ctx, triggerName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[triggerClientSubscribeTriggerToEventsResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[triggerClientSubscribeTriggerToEventsResponse](options.ResumeToken, client.pl, nil)
	}
}

// SubscribeTriggerToEvents - Subscribe event trigger to events.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *triggerClient) subscribeTriggerToEvents(ctx context.Context, triggerName string, options *triggerClientBeginSubscribeTriggerToEventsOptions) (*http.Response, error) {
	req, err := client.subscribeTriggerToEventsCreateRequest(ctx, triggerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// subscribeTriggerToEventsCreateRequest creates the SubscribeTriggerToEvents request.
func (client *triggerClient) subscribeTriggerToEventsCreateRequest(ctx context.Context, triggerName string, options *triggerClientBeginSubscribeTriggerToEventsOptions) (*policy.Request, error) {
	urlPath := "/triggers/{triggerName}/subscribeToEvents"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginUnsubscribeTriggerFromEvents - Unsubscribe event trigger from events.
// If the operation fails it returns an *azcore.ResponseError type.
// triggerName - The trigger name.
// options - triggerClientBeginUnsubscribeTriggerFromEventsOptions contains the optional parameters for the triggerClient.BeginUnsubscribeTriggerFromEvents
// method.
func (client *triggerClient) BeginUnsubscribeTriggerFromEvents(ctx context.Context, triggerName string, options *triggerClientBeginUnsubscribeTriggerFromEventsOptions) (*runtime.Poller[triggerClientUnsubscribeTriggerFromEventsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.unsubscribeTriggerFromEvents(ctx, triggerName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[triggerClientUnsubscribeTriggerFromEventsResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[triggerClientUnsubscribeTriggerFromEventsResponse](options.ResumeToken, client.pl, nil)
	}
}

// UnsubscribeTriggerFromEvents - Unsubscribe event trigger from events.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *triggerClient) unsubscribeTriggerFromEvents(ctx context.Context, triggerName string, options *triggerClientBeginUnsubscribeTriggerFromEventsOptions) (*http.Response, error) {
	req, err := client.unsubscribeTriggerFromEventsCreateRequest(ctx, triggerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// unsubscribeTriggerFromEventsCreateRequest creates the UnsubscribeTriggerFromEvents request.
func (client *triggerClient) unsubscribeTriggerFromEventsCreateRequest(ctx context.Context, triggerName string, options *triggerClientBeginUnsubscribeTriggerFromEventsOptions) (*policy.Request, error) {
	urlPath := "/triggers/{triggerName}/unsubscribeFromEvents"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}
