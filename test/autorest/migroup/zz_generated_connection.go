// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package migroup

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

const telemetryInfo = "azsdk-go-migroup/<version>"

// ConnectionOptions contains configuration settings for the connection's pipeline.
type ConnectionOptions struct {
	// HTTPClient sets the transport for making HTTP requests.
	HTTPClient azcore.Transport
	// Retry configures the built-in retry policy behavior.
	Retry azcore.RetryOptions
	// Telemetry configures the built-in telemetry policy behavior.
	Telemetry azcore.TelemetryOptions
}

// DefaultConnectionOptions creates a ConnectionOptions type initialized with default values.
func DefaultConnectionOptions() ConnectionOptions {
	return ConnectionOptions{
		Retry:     azcore.DefaultRetryOptions(),
		Telemetry: azcore.DefaultTelemetryOptions(),
	}
}

func (c *ConnectionOptions) telemetryOptions() *azcore.TelemetryOptions {
	to := c.Telemetry
	if to.Value == "" {
		to.Value = telemetryInfo
	} else {
		to.Value = fmt.Sprintf("%s %s", telemetryInfo, to.Value)
	}
	return &to
}

// Connection - Service client for multiinheritance client testing
type Connection struct {
	u string
	p azcore.Pipeline
}

// DefaultEndpoint is the default service endpoint.
const DefaultEndpoint = "http://localhost:3000"

// NewDefaultConnection creates an instance of the Connection type using the DefaultEndpoint.
func NewDefaultConnection(options *ConnectionOptions) *Connection {
	return NewConnection(DefaultEndpoint, options)
}

// NewConnection creates an instance of the Connection type with the specified endpoint.
func NewConnection(endpoint string, options *ConnectionOptions) *Connection {
	if options == nil {
		o := DefaultConnectionOptions()
		options = &o
	}
	p := azcore.NewPipeline(options.HTTPClient,
		azcore.NewTelemetryPolicy(options.telemetryOptions()),
		azcore.NewRetryPolicy(&options.Retry),
		azcore.NewLogPolicy(nil))
	return NewConnectionWithPipeline(endpoint, p)
}

// NewConnectionWithPipeline creates an instance of the Connection type with the specified endpoint and pipeline.
func NewConnectionWithPipeline(endpoint string, p azcore.Pipeline) *Connection {
	return &Connection{u: endpoint, p: p}
}

// Endpoint returns the connection's endpoint.
func (c *Connection) Endpoint() string {
	return c.u
}

// Pipeline returns the connection's pipeline.
func (c *Connection) Pipeline() azcore.Pipeline {
	return c.p
}
