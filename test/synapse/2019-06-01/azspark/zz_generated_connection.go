// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azspark

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"strings"
)

const telemetryInfo = "azsdk-go-azspark/<version>"

// connectionOptions contains configuration settings for the connection's pipeline.
// All zero-value fields will be initialized with their default values.
type connectionOptions struct {
	// HTTPClient sets the transport for making HTTP requests.
	HTTPClient azcore.Transport
	// Retry configures the built-in retry policy behavior.
	Retry azcore.RetryOptions
	// Telemetry configures the built-in telemetry policy behavior.
	Telemetry azcore.TelemetryOptions
	// Logging configures the built-in logging policy behavior.
	Logging azcore.LogOptions
}

func (c *connectionOptions) telemetryOptions() *azcore.TelemetryOptions {
	to := c.Telemetry
	if to.Value == "" {
		to.Value = telemetryInfo
	} else {
		to.Value = fmt.Sprintf("%s %s", telemetryInfo, to.Value)
	}
	return &to
}

type connection struct {
	u string
	p azcore.Pipeline
}

// newConnection creates an instance of the connection type with the specified endpoint.
// Pass nil to accept the default options; this is the same as passing a zero-value options.
func newConnection(endpoint string, livyAPIVersion *string, sparkPoolName string, options *connectionOptions) *connection {
	if options == nil {
		options = &connectionOptions{}
	}
	p := azcore.NewPipeline(options.HTTPClient,
		azcore.NewTelemetryPolicy(options.telemetryOptions()),
		azcore.NewRetryPolicy(&options.Retry),
		azcore.NewLogPolicy(&options.Logging))
	return newConnectionWithPipeline(endpoint, livyAPIVersion, sparkPoolName, p)
}

// newConnectionWithPipeline creates an instance of the connection type with the specified endpoint and pipeline.
func newConnectionWithPipeline(endpoint string, livyAPIVersion *string, sparkPoolName string, p azcore.Pipeline) *connection {
	hostURL := "{endpoint}/livyApi/versions/{livyApiVersion}/sparkPools/{sparkPoolName}"
	hostURL = strings.ReplaceAll(hostURL, "{endpoint}", endpoint)
	if livyAPIVersion == nil {
		defaultValue := "2019-11-01-preview"
		livyAPIVersion = &defaultValue
	}
	hostURL = strings.ReplaceAll(hostURL, "{livyApiVersion}", *livyAPIVersion)
	hostURL = strings.ReplaceAll(hostURL, "{sparkPoolName}", sparkPoolName)
	return &connection{u: hostURL, p: p}
}

// Endpoint returns the connection's endpoint.
func (c *connection) Endpoint() string {
	return c.u
}

// Pipeline returns the connection's pipeline.
func (c *connection) Pipeline() azcore.Pipeline {
	return c.p
}
