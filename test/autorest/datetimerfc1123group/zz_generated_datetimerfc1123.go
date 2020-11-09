// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package datetimerfc1123group

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"time"
)

// Datetimerfc1123Operations contains the methods for the Datetimerfc1123 group.
type Datetimerfc1123Operations interface {
	// GetInvalid - Get invalid datetime value
	GetInvalid(ctx context.Context, options *Datetimerfc1123GetInvalidOptions) (*TimeResponse, error)
	// GetNull - Get null datetime value
	GetNull(ctx context.Context, options *Datetimerfc1123GetNullOptions) (*TimeResponse, error)
	// GetOverflow - Get overflow datetime value
	GetOverflow(ctx context.Context, options *Datetimerfc1123GetOverflowOptions) (*TimeResponse, error)
	// GetUTCLowercaseMaxDateTime - Get max datetime value fri, 31 dec 9999 23:59:59 gmt
	GetUTCLowercaseMaxDateTime(ctx context.Context, options *Datetimerfc1123GetUTCLowercaseMaxDateTimeOptions) (*TimeResponse, error)
	// GetUTCMinDateTime - Get min datetime value Mon, 1 Jan 0001 00:00:00 GMT
	GetUTCMinDateTime(ctx context.Context, options *Datetimerfc1123GetUTCMinDateTimeOptions) (*TimeResponse, error)
	// GetUTCUppercaseMaxDateTime - Get max datetime value FRI, 31 DEC 9999 23:59:59 GMT
	GetUTCUppercaseMaxDateTime(ctx context.Context, options *Datetimerfc1123GetUTCUppercaseMaxDateTimeOptions) (*TimeResponse, error)
	// GetUnderflow - Get underflow datetime value
	GetUnderflow(ctx context.Context, options *Datetimerfc1123GetUnderflowOptions) (*TimeResponse, error)
	// PutUTCMaxDateTime - Put max datetime value Fri, 31 Dec 9999 23:59:59 GMT
	PutUTCMaxDateTime(ctx context.Context, datetimeBody time.Time, options *Datetimerfc1123PutUTCMaxDateTimeOptions) (*http.Response, error)
	// PutUTCMinDateTime - Put min datetime value Mon, 1 Jan 0001 00:00:00 GMT
	PutUTCMinDateTime(ctx context.Context, datetimeBody time.Time, options *Datetimerfc1123PutUTCMinDateTimeOptions) (*http.Response, error)
}

// Datetimerfc1123Client implements the Datetimerfc1123Operations interface.
// Don't use this type directly, use NewDatetimerfc1123Client() instead.
type Datetimerfc1123Client struct {
	con *Connection
}

// NewDatetimerfc1123Client creates a new instance of Datetimerfc1123Client with the specified values.
func NewDatetimerfc1123Client(con *Connection) Datetimerfc1123Operations {
	return &Datetimerfc1123Client{con: con}
}

// Pipeline returns the pipeline associated with this client.
func (client *Datetimerfc1123Client) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// GetInvalid - Get invalid datetime value
func (client *Datetimerfc1123Client) GetInvalid(ctx context.Context, options *Datetimerfc1123GetInvalidOptions) (*TimeResponse, error) {
	req, err := client.GetInvalidCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetInvalidHandleError(resp)
	}
	result, err := client.GetInvalidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetInvalidCreateRequest creates the GetInvalid request.
func (client *Datetimerfc1123Client) GetInvalidCreateRequest(ctx context.Context, options *Datetimerfc1123GetInvalidOptions) (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/invalid"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetInvalidHandleResponse handles the GetInvalid response.
func (client *Datetimerfc1123Client) GetInvalidHandleResponse(resp *azcore.Response) (*TimeResponse, error) {
	var aux *timeRFC1123
	err := resp.UnmarshalAsJSON(&aux)
	return &TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, err
}

// GetInvalidHandleError handles the GetInvalid error response.
func (client *Datetimerfc1123Client) GetInvalidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetNull - Get null datetime value
func (client *Datetimerfc1123Client) GetNull(ctx context.Context, options *Datetimerfc1123GetNullOptions) (*TimeResponse, error) {
	req, err := client.GetNullCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetNullHandleError(resp)
	}
	result, err := client.GetNullHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetNullCreateRequest creates the GetNull request.
func (client *Datetimerfc1123Client) GetNullCreateRequest(ctx context.Context, options *Datetimerfc1123GetNullOptions) (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/null"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetNullHandleResponse handles the GetNull response.
func (client *Datetimerfc1123Client) GetNullHandleResponse(resp *azcore.Response) (*TimeResponse, error) {
	var aux *timeRFC1123
	err := resp.UnmarshalAsJSON(&aux)
	return &TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, err
}

// GetNullHandleError handles the GetNull error response.
func (client *Datetimerfc1123Client) GetNullHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetOverflow - Get overflow datetime value
func (client *Datetimerfc1123Client) GetOverflow(ctx context.Context, options *Datetimerfc1123GetOverflowOptions) (*TimeResponse, error) {
	req, err := client.GetOverflowCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetOverflowHandleError(resp)
	}
	result, err := client.GetOverflowHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetOverflowCreateRequest creates the GetOverflow request.
func (client *Datetimerfc1123Client) GetOverflowCreateRequest(ctx context.Context, options *Datetimerfc1123GetOverflowOptions) (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/overflow"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetOverflowHandleResponse handles the GetOverflow response.
func (client *Datetimerfc1123Client) GetOverflowHandleResponse(resp *azcore.Response) (*TimeResponse, error) {
	var aux *timeRFC1123
	err := resp.UnmarshalAsJSON(&aux)
	return &TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, err
}

// GetOverflowHandleError handles the GetOverflow error response.
func (client *Datetimerfc1123Client) GetOverflowHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetUTCLowercaseMaxDateTime - Get max datetime value fri, 31 dec 9999 23:59:59 gmt
func (client *Datetimerfc1123Client) GetUTCLowercaseMaxDateTime(ctx context.Context, options *Datetimerfc1123GetUTCLowercaseMaxDateTimeOptions) (*TimeResponse, error) {
	req, err := client.GetUTCLowercaseMaxDateTimeCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetUTCLowercaseMaxDateTimeHandleError(resp)
	}
	result, err := client.GetUTCLowercaseMaxDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUTCLowercaseMaxDateTimeCreateRequest creates the GetUTCLowercaseMaxDateTime request.
func (client *Datetimerfc1123Client) GetUTCLowercaseMaxDateTimeCreateRequest(ctx context.Context, options *Datetimerfc1123GetUTCLowercaseMaxDateTimeOptions) (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/max/lowercase"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetUTCLowercaseMaxDateTimeHandleResponse handles the GetUTCLowercaseMaxDateTime response.
func (client *Datetimerfc1123Client) GetUTCLowercaseMaxDateTimeHandleResponse(resp *azcore.Response) (*TimeResponse, error) {
	var aux *timeRFC1123
	err := resp.UnmarshalAsJSON(&aux)
	return &TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, err
}

// GetUTCLowercaseMaxDateTimeHandleError handles the GetUTCLowercaseMaxDateTime error response.
func (client *Datetimerfc1123Client) GetUTCLowercaseMaxDateTimeHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetUTCMinDateTime - Get min datetime value Mon, 1 Jan 0001 00:00:00 GMT
func (client *Datetimerfc1123Client) GetUTCMinDateTime(ctx context.Context, options *Datetimerfc1123GetUTCMinDateTimeOptions) (*TimeResponse, error) {
	req, err := client.GetUTCMinDateTimeCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetUTCMinDateTimeHandleError(resp)
	}
	result, err := client.GetUTCMinDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUTCMinDateTimeCreateRequest creates the GetUTCMinDateTime request.
func (client *Datetimerfc1123Client) GetUTCMinDateTimeCreateRequest(ctx context.Context, options *Datetimerfc1123GetUTCMinDateTimeOptions) (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/min"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetUTCMinDateTimeHandleResponse handles the GetUTCMinDateTime response.
func (client *Datetimerfc1123Client) GetUTCMinDateTimeHandleResponse(resp *azcore.Response) (*TimeResponse, error) {
	var aux *timeRFC1123
	err := resp.UnmarshalAsJSON(&aux)
	return &TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, err
}

// GetUTCMinDateTimeHandleError handles the GetUTCMinDateTime error response.
func (client *Datetimerfc1123Client) GetUTCMinDateTimeHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetUTCUppercaseMaxDateTime - Get max datetime value FRI, 31 DEC 9999 23:59:59 GMT
func (client *Datetimerfc1123Client) GetUTCUppercaseMaxDateTime(ctx context.Context, options *Datetimerfc1123GetUTCUppercaseMaxDateTimeOptions) (*TimeResponse, error) {
	req, err := client.GetUTCUppercaseMaxDateTimeCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetUTCUppercaseMaxDateTimeHandleError(resp)
	}
	result, err := client.GetUTCUppercaseMaxDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUTCUppercaseMaxDateTimeCreateRequest creates the GetUTCUppercaseMaxDateTime request.
func (client *Datetimerfc1123Client) GetUTCUppercaseMaxDateTimeCreateRequest(ctx context.Context, options *Datetimerfc1123GetUTCUppercaseMaxDateTimeOptions) (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/max/uppercase"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetUTCUppercaseMaxDateTimeHandleResponse handles the GetUTCUppercaseMaxDateTime response.
func (client *Datetimerfc1123Client) GetUTCUppercaseMaxDateTimeHandleResponse(resp *azcore.Response) (*TimeResponse, error) {
	var aux *timeRFC1123
	err := resp.UnmarshalAsJSON(&aux)
	return &TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, err
}

// GetUTCUppercaseMaxDateTimeHandleError handles the GetUTCUppercaseMaxDateTime error response.
func (client *Datetimerfc1123Client) GetUTCUppercaseMaxDateTimeHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetUnderflow - Get underflow datetime value
func (client *Datetimerfc1123Client) GetUnderflow(ctx context.Context, options *Datetimerfc1123GetUnderflowOptions) (*TimeResponse, error) {
	req, err := client.GetUnderflowCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetUnderflowHandleError(resp)
	}
	result, err := client.GetUnderflowHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUnderflowCreateRequest creates the GetUnderflow request.
func (client *Datetimerfc1123Client) GetUnderflowCreateRequest(ctx context.Context, options *Datetimerfc1123GetUnderflowOptions) (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/underflow"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetUnderflowHandleResponse handles the GetUnderflow response.
func (client *Datetimerfc1123Client) GetUnderflowHandleResponse(resp *azcore.Response) (*TimeResponse, error) {
	var aux *timeRFC1123
	err := resp.UnmarshalAsJSON(&aux)
	return &TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, err
}

// GetUnderflowHandleError handles the GetUnderflow error response.
func (client *Datetimerfc1123Client) GetUnderflowHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutUTCMaxDateTime - Put max datetime value Fri, 31 Dec 9999 23:59:59 GMT
func (client *Datetimerfc1123Client) PutUTCMaxDateTime(ctx context.Context, datetimeBody time.Time, options *Datetimerfc1123PutUTCMaxDateTimeOptions) (*http.Response, error) {
	req, err := client.PutUTCMaxDateTimeCreateRequest(ctx, datetimeBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutUTCMaxDateTimeHandleError(resp)
	}
	return resp.Response, nil
}

// PutUTCMaxDateTimeCreateRequest creates the PutUTCMaxDateTime request.
func (client *Datetimerfc1123Client) PutUTCMaxDateTimeCreateRequest(ctx context.Context, datetimeBody time.Time, options *Datetimerfc1123PutUTCMaxDateTimeOptions) (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/max"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	aux := timeRFC1123(datetimeBody)
	return req, req.MarshalAsJSON(aux)
}

// PutUTCMaxDateTimeHandleError handles the PutUTCMaxDateTime error response.
func (client *Datetimerfc1123Client) PutUTCMaxDateTimeHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutUTCMinDateTime - Put min datetime value Mon, 1 Jan 0001 00:00:00 GMT
func (client *Datetimerfc1123Client) PutUTCMinDateTime(ctx context.Context, datetimeBody time.Time, options *Datetimerfc1123PutUTCMinDateTimeOptions) (*http.Response, error) {
	req, err := client.PutUTCMinDateTimeCreateRequest(ctx, datetimeBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutUTCMinDateTimeHandleError(resp)
	}
	return resp.Response, nil
}

// PutUTCMinDateTimeCreateRequest creates the PutUTCMinDateTime request.
func (client *Datetimerfc1123Client) PutUTCMinDateTimeCreateRequest(ctx context.Context, datetimeBody time.Time, options *Datetimerfc1123PutUTCMinDateTimeOptions) (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/min"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	aux := timeRFC1123(datetimeBody)
	return req, req.MarshalAsJSON(aux)
}

// PutUTCMinDateTimeHandleError handles the PutUTCMinDateTime error response.
func (client *Datetimerfc1123Client) PutUTCMinDateTimeHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
