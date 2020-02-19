// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package headergroup

import (
	"context"
	azinternal "generatortests/autorest/generated/headergroup/internal/headergroup"
	"time"
)

// HeaderOperations contains the methods for the Header group.
type HeaderOperations interface {
	// CustomRequestID - Send x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 in the header of the request 
	CustomRequestID(ctx context.Context) (*HeaderCustomRequestIDResponse, error)
	// ParamBool - Send a post request with header values "scenario": "true", "value": true or "scenario": "false", "value": false 
	ParamBool(ctx context.Context, scenario string, value bool) (*HeaderParamBoolResponse, error)
	// ParamByte - Send a post request with header values "scenario": "valid", "value": "啊齄丂狛狜隣郎隣兀﨩" 
	ParamByte(ctx context.Context, scenario string, value []byte) (*HeaderParamByteResponse, error)
	// ParamDate - Send a post request with header values "scenario": "valid", "value": "2010-01-01" or "scenario": "min", "value": "0001-01-01" 
	ParamDate(ctx context.Context, scenario string, value time.Time) (*HeaderParamDateResponse, error)
	// ParamDatetime - Send a post request with header values "scenario": "valid", "value": "2010-01-01T12:34:56Z" or "scenario": "min", "value": "0001-01-01T00:00:00Z" 
	ParamDatetime(ctx context.Context, scenario string, value time.Time) (*HeaderParamDatetimeResponse, error)
	// ParamDatetimeRFC1123 - Send a post request with header values "scenario": "valid", "value": "Wed, 01 Jan 2010 12:34:56 GMT" or "scenario": "min", "value": "Mon, 01 Jan 0001 00:00:00 GMT" 
	ParamDatetimeRFC1123(ctx context.Context, scenario string, value time.Time) (*HeaderParamDatetimeRFC1123Response, error)
	// ParamDouble - Send a post request with header values "scenario": "positive", "value": 7e120 or "scenario": "negative", "value": -3.0 
	ParamDouble(ctx context.Context, scenario string, value float64) (*HeaderParamDoubleResponse, error)
	// ParamDuration - Send a post request with header values "scenario": "valid", "value": "P123DT22H14M12.011S" 
	ParamDuration(ctx context.Context, scenario string, value time.Duration) (*HeaderParamDurationResponse, error)
	// ParamEnum - Send a post request with header values "scenario": "valid", "value": "GREY" or "scenario": "null", "value": null 
	ParamEnum(ctx context.Context, scenario string, value GreyscaleColors) (*HeaderParamEnumResponse, error)
	// ParamExistingKey - Send a post request with header value "User-Agent": "overwrite" 
	ParamExistingKey(ctx context.Context, userAgent string) (*HeaderParamExistingKeyResponse, error)
	// ParamFloat - Send a post request with header values "scenario": "positive", "value": 0.07 or "scenario": "negative", "value": -3.0 
	ParamFloat(ctx context.Context, scenario string, value float32) (*HeaderParamFloatResponse, error)
	// ParamInteger - Send a post request with header values "scenario": "positive", "value": 1 or "scenario": "negative", "value": -2  
	ParamInteger(ctx context.Context, scenario string, value int32) (*HeaderParamIntegerResponse, error)
	// ParamLong - Send a post request with header values "scenario": "positive", "value": 105 or "scenario": "negative", "value": -2  
	ParamLong(ctx context.Context, scenario string, value int64) (*HeaderParamLongResponse, error)
	// ParamProtectedKey - Send a post request with header value "Content-Type": "text/html" 
	ParamProtectedKey(ctx context.Context, contentType string) (*HeaderParamProtectedKeyResponse, error)
	// ParamString - Send a post request with header values "scenario": "valid", "value": "The quick brown fox jumps over the lazy dog" or "scenario": "null", "value": null or "scenario": "empty", "value": "" 
	ParamString(ctx context.Context, scenario string, value string) (*HeaderParamStringResponse, error)
	// ResponseBool - Get a response with header value "value": true or false 
	ResponseBool(ctx context.Context, scenario string) (*HeaderResponseBoolResponse, error)
	// ResponseByte - Get a response with header values "啊齄丂狛狜隣郎隣兀﨩" 
	ResponseByte(ctx context.Context, scenario string) (*HeaderResponseByteResponse, error)
	// ResponseDate - Get a response with header values "2010-01-01" or "0001-01-01" 
	ResponseDate(ctx context.Context, scenario string) (*HeaderResponseDateResponse, error)
	// ResponseDatetime - Get a response with header values "2010-01-01T12:34:56Z" or "0001-01-01T00:00:00Z" 
	ResponseDatetime(ctx context.Context, scenario string) (*HeaderResponseDatetimeResponse, error)
	// ResponseDatetimeRFC1123 - Get a response with header values "Wed, 01 Jan 2010 12:34:56 GMT" or "Mon, 01 Jan 0001 00:00:00 GMT" 
	ResponseDatetimeRFC1123(ctx context.Context, scenario string) (*HeaderResponseDatetimeRFC1123Response, error)
	// ResponseDouble - Get a response with header value "value": 7e120 or -3.0 
	ResponseDouble(ctx context.Context, scenario string) (*HeaderResponseDoubleResponse, error)
	// ResponseDuration - Get a response with header values "P123DT22H14M12.011S" 
	ResponseDuration(ctx context.Context, scenario string) (*HeaderResponseDurationResponse, error)
	// ResponseEnum - Get a response with header values "GREY" or null 
	ResponseEnum(ctx context.Context, scenario string) (*HeaderResponseEnumResponse, error)
	// ResponseExistingKey - Get a response with header value "User-Agent": "overwrite" 
	ResponseExistingKey(ctx context.Context) (*HeaderResponseExistingKeyResponse, error)
	// ResponseFloat - Get a response with header value "value": 0.07 or -3.0 
	ResponseFloat(ctx context.Context, scenario string) (*HeaderResponseFloatResponse, error)
	// ResponseInteger - Get a response with header value "value": 1 or -2 
	ResponseInteger(ctx context.Context, scenario string) (*HeaderResponseIntegerResponse, error)
	// ResponseLong - Get a response with header value "value": 105 or -2 
	ResponseLong(ctx context.Context, scenario string) (*HeaderResponseLongResponse, error)
	// ResponseProtectedKey - Get a response with header value "Content-Type": "text/html" 
	ResponseProtectedKey(ctx context.Context) (*HeaderResponseProtectedKeyResponse, error)
	// ResponseString - Get a response with header values "The quick brown fox jumps over the lazy dog" or null or "" 
	ResponseString(ctx context.Context, scenario string) (*HeaderResponseStringResponse, error)
}

type headerOperations struct {
	*Client
	azinternal.HeaderOperations
}

// CustomRequestID - Send x-ms-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 in the header of the request 
func (client *headerOperations) CustomRequestID(ctx context.Context) (*HeaderCustomRequestIDResponse, error) {
	req, err := client.CustomRequestIDCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.CustomRequestIDHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ParamBool - Send a post request with header values "scenario": "true", "value": true or "scenario": "false", "value": false 
func (client *headerOperations) ParamBool(ctx context.Context, scenario string, value bool) (*HeaderParamBoolResponse, error) {
	req, err := client.ParamBoolCreateRequest(*client.u, scenario, value)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ParamBoolHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ParamByte - Send a post request with header values "scenario": "valid", "value": "啊齄丂狛狜隣郎隣兀﨩" 
func (client *headerOperations) ParamByte(ctx context.Context, scenario string, value []byte) (*HeaderParamByteResponse, error) {
	req, err := client.ParamByteCreateRequest(*client.u, scenario, value)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ParamByteHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ParamDate - Send a post request with header values "scenario": "valid", "value": "2010-01-01" or "scenario": "min", "value": "0001-01-01" 
func (client *headerOperations) ParamDate(ctx context.Context, scenario string, value time.Time) (*HeaderParamDateResponse, error) {
	req, err := client.ParamDateCreateRequest(*client.u, scenario, value)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ParamDateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ParamDatetime - Send a post request with header values "scenario": "valid", "value": "2010-01-01T12:34:56Z" or "scenario": "min", "value": "0001-01-01T00:00:00Z" 
func (client *headerOperations) ParamDatetime(ctx context.Context, scenario string, value time.Time) (*HeaderParamDatetimeResponse, error) {
	req, err := client.ParamDatetimeCreateRequest(*client.u, scenario, value)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ParamDatetimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ParamDatetimeRFC1123 - Send a post request with header values "scenario": "valid", "value": "Wed, 01 Jan 2010 12:34:56 GMT" or "scenario": "min", "value": "Mon, 01 Jan 0001 00:00:00 GMT" 
func (client *headerOperations) ParamDatetimeRFC1123(ctx context.Context, scenario string, value time.Time) (*HeaderParamDatetimeRFC1123Response, error) {
	req, err := client.ParamDatetimeRFC1123CreateRequest(*client.u, scenario, value)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ParamDatetimeRFC1123HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ParamDouble - Send a post request with header values "scenario": "positive", "value": 7e120 or "scenario": "negative", "value": -3.0 
func (client *headerOperations) ParamDouble(ctx context.Context, scenario string, value float64) (*HeaderParamDoubleResponse, error) {
	req, err := client.ParamDoubleCreateRequest(*client.u, scenario, value)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ParamDoubleHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ParamDuration - Send a post request with header values "scenario": "valid", "value": "P123DT22H14M12.011S" 
func (client *headerOperations) ParamDuration(ctx context.Context, scenario string, value time.Duration) (*HeaderParamDurationResponse, error) {
	req, err := client.ParamDurationCreateRequest(*client.u, scenario, value)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ParamDurationHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ParamEnum - Send a post request with header values "scenario": "valid", "value": "GREY" or "scenario": "null", "value": null 
func (client *headerOperations) ParamEnum(ctx context.Context, scenario string, value GreyscaleColors) (*HeaderParamEnumResponse, error) {
	req, err := client.ParamEnumCreateRequest(*client.u, scenario, value)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ParamEnumHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ParamExistingKey - Send a post request with header value "User-Agent": "overwrite" 
func (client *headerOperations) ParamExistingKey(ctx context.Context, userAgent string) (*HeaderParamExistingKeyResponse, error) {
	req, err := client.ParamExistingKeyCreateRequest(*client.u, userAgent)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ParamExistingKeyHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ParamFloat - Send a post request with header values "scenario": "positive", "value": 0.07 or "scenario": "negative", "value": -3.0 
func (client *headerOperations) ParamFloat(ctx context.Context, scenario string, value float32) (*HeaderParamFloatResponse, error) {
	req, err := client.ParamFloatCreateRequest(*client.u, scenario, value)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ParamFloatHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ParamInteger - Send a post request with header values "scenario": "positive", "value": 1 or "scenario": "negative", "value": -2  
func (client *headerOperations) ParamInteger(ctx context.Context, scenario string, value int32) (*HeaderParamIntegerResponse, error) {
	req, err := client.ParamIntegerCreateRequest(*client.u, scenario, value)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ParamIntegerHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ParamLong - Send a post request with header values "scenario": "positive", "value": 105 or "scenario": "negative", "value": -2  
func (client *headerOperations) ParamLong(ctx context.Context, scenario string, value int64) (*HeaderParamLongResponse, error) {
	req, err := client.ParamLongCreateRequest(*client.u, scenario, value)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ParamLongHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ParamProtectedKey - Send a post request with header value "Content-Type": "text/html" 
func (client *headerOperations) ParamProtectedKey(ctx context.Context, contentType string) (*HeaderParamProtectedKeyResponse, error) {
	req, err := client.ParamProtectedKeyCreateRequest(*client.u, contentType)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ParamProtectedKeyHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ParamString - Send a post request with header values "scenario": "valid", "value": "The quick brown fox jumps over the lazy dog" or "scenario": "null", "value": null or "scenario": "empty", "value": "" 
func (client *headerOperations) ParamString(ctx context.Context, scenario string, value string) (*HeaderParamStringResponse, error) {
	req, err := client.ParamStringCreateRequest(*client.u, scenario, value)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ParamStringHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ResponseBool - Get a response with header value "value": true or false 
func (client *headerOperations) ResponseBool(ctx context.Context, scenario string) (*HeaderResponseBoolResponse, error) {
	req, err := client.ResponseBoolCreateRequest(*client.u, scenario)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ResponseBoolHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ResponseByte - Get a response with header values "啊齄丂狛狜隣郎隣兀﨩" 
func (client *headerOperations) ResponseByte(ctx context.Context, scenario string) (*HeaderResponseByteResponse, error) {
	req, err := client.ResponseByteCreateRequest(*client.u, scenario)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ResponseByteHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ResponseDate - Get a response with header values "2010-01-01" or "0001-01-01" 
func (client *headerOperations) ResponseDate(ctx context.Context, scenario string) (*HeaderResponseDateResponse, error) {
	req, err := client.ResponseDateCreateRequest(*client.u, scenario)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ResponseDateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ResponseDatetime - Get a response with header values "2010-01-01T12:34:56Z" or "0001-01-01T00:00:00Z" 
func (client *headerOperations) ResponseDatetime(ctx context.Context, scenario string) (*HeaderResponseDatetimeResponse, error) {
	req, err := client.ResponseDatetimeCreateRequest(*client.u, scenario)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ResponseDatetimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ResponseDatetimeRFC1123 - Get a response with header values "Wed, 01 Jan 2010 12:34:56 GMT" or "Mon, 01 Jan 0001 00:00:00 GMT" 
func (client *headerOperations) ResponseDatetimeRFC1123(ctx context.Context, scenario string) (*HeaderResponseDatetimeRFC1123Response, error) {
	req, err := client.ResponseDatetimeRFC1123CreateRequest(*client.u, scenario)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ResponseDatetimeRFC1123HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ResponseDouble - Get a response with header value "value": 7e120 or -3.0 
func (client *headerOperations) ResponseDouble(ctx context.Context, scenario string) (*HeaderResponseDoubleResponse, error) {
	req, err := client.ResponseDoubleCreateRequest(*client.u, scenario)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ResponseDoubleHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ResponseDuration - Get a response with header values "P123DT22H14M12.011S" 
func (client *headerOperations) ResponseDuration(ctx context.Context, scenario string) (*HeaderResponseDurationResponse, error) {
	req, err := client.ResponseDurationCreateRequest(*client.u, scenario)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ResponseDurationHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ResponseEnum - Get a response with header values "GREY" or null 
func (client *headerOperations) ResponseEnum(ctx context.Context, scenario string) (*HeaderResponseEnumResponse, error) {
	req, err := client.ResponseEnumCreateRequest(*client.u, scenario)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ResponseEnumHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ResponseExistingKey - Get a response with header value "User-Agent": "overwrite" 
func (client *headerOperations) ResponseExistingKey(ctx context.Context) (*HeaderResponseExistingKeyResponse, error) {
	req, err := client.ResponseExistingKeyCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ResponseExistingKeyHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ResponseFloat - Get a response with header value "value": 0.07 or -3.0 
func (client *headerOperations) ResponseFloat(ctx context.Context, scenario string) (*HeaderResponseFloatResponse, error) {
	req, err := client.ResponseFloatCreateRequest(*client.u, scenario)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ResponseFloatHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ResponseInteger - Get a response with header value "value": 1 or -2 
func (client *headerOperations) ResponseInteger(ctx context.Context, scenario string) (*HeaderResponseIntegerResponse, error) {
	req, err := client.ResponseIntegerCreateRequest(*client.u, scenario)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ResponseIntegerHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ResponseLong - Get a response with header value "value": 105 or -2 
func (client *headerOperations) ResponseLong(ctx context.Context, scenario string) (*HeaderResponseLongResponse, error) {
	req, err := client.ResponseLongCreateRequest(*client.u, scenario)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ResponseLongHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ResponseProtectedKey - Get a response with header value "Content-Type": "text/html" 
func (client *headerOperations) ResponseProtectedKey(ctx context.Context) (*HeaderResponseProtectedKeyResponse, error) {
	req, err := client.ResponseProtectedKeyCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ResponseProtectedKeyHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ResponseString - Get a response with header values "The quick brown fox jumps over the lazy dog" or null or "" 
func (client *headerOperations) ResponseString(ctx context.Context, scenario string) (*HeaderResponseStringResponse, error) {
	req, err := client.ResponseStringCreateRequest(*client.u, scenario)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.ResponseStringHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

var _ HeaderOperations = (*headerOperations)(nil)
