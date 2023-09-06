//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearning

// BatchDeploymentsClientBeginCreateOrUpdateOptions contains the optional parameters for the BatchDeploymentsClient.BeginCreateOrUpdate
// method.
type BatchDeploymentsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BatchDeploymentsClientBeginDeleteOptions contains the optional parameters for the BatchDeploymentsClient.BeginDelete method.
type BatchDeploymentsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BatchDeploymentsClientBeginUpdateOptions contains the optional parameters for the BatchDeploymentsClient.BeginUpdate method.
type BatchDeploymentsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BatchDeploymentsClientGetOptions contains the optional parameters for the BatchDeploymentsClient.Get method.
type BatchDeploymentsClientGetOptions struct {
	// placeholder for future optional parameters
}

// BatchDeploymentsClientListOptions contains the optional parameters for the BatchDeploymentsClient.NewListPager method.
type BatchDeploymentsClientListOptions struct {
	// Ordering of list.
	OrderBy *string

	// Continuation token for pagination.
	Skip *string

	// Top of list.
	Top *int32
}

// BatchEndpointsClientBeginCreateOrUpdateOptions contains the optional parameters for the BatchEndpointsClient.BeginCreateOrUpdate
// method.
type BatchEndpointsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BatchEndpointsClientBeginDeleteOptions contains the optional parameters for the BatchEndpointsClient.BeginDelete method.
type BatchEndpointsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BatchEndpointsClientBeginUpdateOptions contains the optional parameters for the BatchEndpointsClient.BeginUpdate method.
type BatchEndpointsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BatchEndpointsClientGetOptions contains the optional parameters for the BatchEndpointsClient.Get method.
type BatchEndpointsClientGetOptions struct {
	// placeholder for future optional parameters
}

// BatchEndpointsClientListKeysOptions contains the optional parameters for the BatchEndpointsClient.ListKeys method.
type BatchEndpointsClientListKeysOptions struct {
	// placeholder for future optional parameters
}

// BatchEndpointsClientListOptions contains the optional parameters for the BatchEndpointsClient.NewListPager method.
type BatchEndpointsClientListOptions struct {
	// Number of endpoints to be retrieved in a page of results.
	Count *int32

	// Continuation token for pagination.
	Skip *string
}

// CodeContainersClientCreateOrUpdateOptions contains the optional parameters for the CodeContainersClient.CreateOrUpdate
// method.
type CodeContainersClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// CodeContainersClientDeleteOptions contains the optional parameters for the CodeContainersClient.Delete method.
type CodeContainersClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// CodeContainersClientGetOptions contains the optional parameters for the CodeContainersClient.Get method.
type CodeContainersClientGetOptions struct {
	// placeholder for future optional parameters
}

// CodeContainersClientListOptions contains the optional parameters for the CodeContainersClient.NewListPager method.
type CodeContainersClientListOptions struct {
	// Continuation token for pagination.
	Skip *string
}

// CodeVersionsClientCreateOrUpdateOptions contains the optional parameters for the CodeVersionsClient.CreateOrUpdate method.
type CodeVersionsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// CodeVersionsClientDeleteOptions contains the optional parameters for the CodeVersionsClient.Delete method.
type CodeVersionsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// CodeVersionsClientGetOptions contains the optional parameters for the CodeVersionsClient.Get method.
type CodeVersionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// CodeVersionsClientListOptions contains the optional parameters for the CodeVersionsClient.NewListPager method.
type CodeVersionsClientListOptions struct {
	// Ordering of list.
	OrderBy *string

	// Continuation token for pagination.
	Skip *string

	// Maximum number of records to return.
	Top *int32
}

// ComponentContainersClientCreateOrUpdateOptions contains the optional parameters for the ComponentContainersClient.CreateOrUpdate
// method.
type ComponentContainersClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ComponentContainersClientDeleteOptions contains the optional parameters for the ComponentContainersClient.Delete method.
type ComponentContainersClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ComponentContainersClientGetOptions contains the optional parameters for the ComponentContainersClient.Get method.
type ComponentContainersClientGetOptions struct {
	// placeholder for future optional parameters
}

// ComponentContainersClientListOptions contains the optional parameters for the ComponentContainersClient.NewListPager method.
type ComponentContainersClientListOptions struct {
	// View type for including/excluding (for example) archived entities.
	ListViewType *ListViewType

	// Continuation token for pagination.
	Skip *string
}

// ComponentVersionsClientCreateOrUpdateOptions contains the optional parameters for the ComponentVersionsClient.CreateOrUpdate
// method.
type ComponentVersionsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ComponentVersionsClientDeleteOptions contains the optional parameters for the ComponentVersionsClient.Delete method.
type ComponentVersionsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ComponentVersionsClientGetOptions contains the optional parameters for the ComponentVersionsClient.Get method.
type ComponentVersionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ComponentVersionsClientListOptions contains the optional parameters for the ComponentVersionsClient.NewListPager method.
type ComponentVersionsClientListOptions struct {
	// View type for including/excluding (for example) archived entities.
	ListViewType *ListViewType

	// Ordering of list.
	OrderBy *string

	// Continuation token for pagination.
	Skip *string

	// Maximum number of records to return.
	Top *int32
}

// ComputeClientBeginCreateOrUpdateOptions contains the optional parameters for the ComputeClient.BeginCreateOrUpdate method.
type ComputeClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ComputeClientBeginDeleteOptions contains the optional parameters for the ComputeClient.BeginDelete method.
type ComputeClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ComputeClientBeginRestartOptions contains the optional parameters for the ComputeClient.BeginRestart method.
type ComputeClientBeginRestartOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ComputeClientBeginStartOptions contains the optional parameters for the ComputeClient.BeginStart method.
type ComputeClientBeginStartOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ComputeClientBeginStopOptions contains the optional parameters for the ComputeClient.BeginStop method.
type ComputeClientBeginStopOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ComputeClientBeginUpdateOptions contains the optional parameters for the ComputeClient.BeginUpdate method.
type ComputeClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ComputeClientGetOptions contains the optional parameters for the ComputeClient.Get method.
type ComputeClientGetOptions struct {
	// placeholder for future optional parameters
}

// ComputeClientListKeysOptions contains the optional parameters for the ComputeClient.ListKeys method.
type ComputeClientListKeysOptions struct {
	// placeholder for future optional parameters
}

// ComputeClientListNodesOptions contains the optional parameters for the ComputeClient.NewListNodesPager method.
type ComputeClientListNodesOptions struct {
	// placeholder for future optional parameters
}

// ComputeClientListOptions contains the optional parameters for the ComputeClient.NewListPager method.
type ComputeClientListOptions struct {
	// Continuation token for pagination.
	Skip *string
}

// DataContainersClientCreateOrUpdateOptions contains the optional parameters for the DataContainersClient.CreateOrUpdate
// method.
type DataContainersClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// DataContainersClientDeleteOptions contains the optional parameters for the DataContainersClient.Delete method.
type DataContainersClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// DataContainersClientGetOptions contains the optional parameters for the DataContainersClient.Get method.
type DataContainersClientGetOptions struct {
	// placeholder for future optional parameters
}

// DataContainersClientListOptions contains the optional parameters for the DataContainersClient.NewListPager method.
type DataContainersClientListOptions struct {
	// View type for including/excluding (for example) archived entities.
	ListViewType *ListViewType

	// Continuation token for pagination.
	Skip *string
}

// DataVersionsClientCreateOrUpdateOptions contains the optional parameters for the DataVersionsClient.CreateOrUpdate method.
type DataVersionsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// DataVersionsClientDeleteOptions contains the optional parameters for the DataVersionsClient.Delete method.
type DataVersionsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// DataVersionsClientGetOptions contains the optional parameters for the DataVersionsClient.Get method.
type DataVersionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// DataVersionsClientListOptions contains the optional parameters for the DataVersionsClient.NewListPager method.
type DataVersionsClientListOptions struct {
	// [ListViewType.ActiveOnly, ListViewType.ArchivedOnly, ListViewType.All]View type for including/excluding (for example) archived
	// entities.
	ListViewType *ListViewType

	// Please choose OrderBy value from ['createdtime', 'modifiedtime']
	OrderBy *string

	// Continuation token for pagination.
	Skip *string

	// Comma-separated list of tag names (and optionally values). Example: tag1,tag2=value2
	Tags *string

	// Top count of results, top count cannot be greater than the page size. If topCount > page size, results with be default
	// page size count will be returned
	Top *int32
}

// DatastoresClientCreateOrUpdateOptions contains the optional parameters for the DatastoresClient.CreateOrUpdate method.
type DatastoresClientCreateOrUpdateOptions struct {
	// Flag to skip validation.
	SkipValidation *bool
}

// DatastoresClientDeleteOptions contains the optional parameters for the DatastoresClient.Delete method.
type DatastoresClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// DatastoresClientGetOptions contains the optional parameters for the DatastoresClient.Get method.
type DatastoresClientGetOptions struct {
	// placeholder for future optional parameters
}

// DatastoresClientListOptions contains the optional parameters for the DatastoresClient.NewListPager method.
type DatastoresClientListOptions struct {
	// Maximum number of results to return.
	Count *int32

	// Filter down to the workspace default datastore.
	IsDefault *bool

	// Names of datastores to return.
	Names []string

	// Order by property (createdtime | modifiedtime | name).
	OrderBy *string

	// Order by property in ascending order.
	OrderByAsc *bool

	// Text to search for in the datastore names.
	SearchText *string

	// Continuation token for pagination.
	Skip *string
}

// DatastoresClientListSecretsOptions contains the optional parameters for the DatastoresClient.ListSecrets method.
type DatastoresClientListSecretsOptions struct {
	// placeholder for future optional parameters
}

// EnvironmentContainersClientCreateOrUpdateOptions contains the optional parameters for the EnvironmentContainersClient.CreateOrUpdate
// method.
type EnvironmentContainersClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// EnvironmentContainersClientDeleteOptions contains the optional parameters for the EnvironmentContainersClient.Delete method.
type EnvironmentContainersClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// EnvironmentContainersClientGetOptions contains the optional parameters for the EnvironmentContainersClient.Get method.
type EnvironmentContainersClientGetOptions struct {
	// placeholder for future optional parameters
}

// EnvironmentContainersClientListOptions contains the optional parameters for the EnvironmentContainersClient.NewListPager
// method.
type EnvironmentContainersClientListOptions struct {
	// View type for including/excluding (for example) archived entities.
	ListViewType *ListViewType

	// Continuation token for pagination.
	Skip *string
}

// EnvironmentVersionsClientCreateOrUpdateOptions contains the optional parameters for the EnvironmentVersionsClient.CreateOrUpdate
// method.
type EnvironmentVersionsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// EnvironmentVersionsClientDeleteOptions contains the optional parameters for the EnvironmentVersionsClient.Delete method.
type EnvironmentVersionsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// EnvironmentVersionsClientGetOptions contains the optional parameters for the EnvironmentVersionsClient.Get method.
type EnvironmentVersionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// EnvironmentVersionsClientListOptions contains the optional parameters for the EnvironmentVersionsClient.NewListPager method.
type EnvironmentVersionsClientListOptions struct {
	// View type for including/excluding (for example) archived entities.
	ListViewType *ListViewType

	// Ordering of list.
	OrderBy *string

	// Continuation token for pagination.
	Skip *string

	// Maximum number of records to return.
	Top *int32
}

// JobsClientBeginDeleteOptions contains the optional parameters for the JobsClient.BeginDelete method.
type JobsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// JobsClientCancelOptions contains the optional parameters for the JobsClient.Cancel method.
type JobsClientCancelOptions struct {
	// placeholder for future optional parameters
}

// JobsClientCreateOrUpdateOptions contains the optional parameters for the JobsClient.CreateOrUpdate method.
type JobsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// JobsClientGetOptions contains the optional parameters for the JobsClient.Get method.
type JobsClientGetOptions struct {
	// placeholder for future optional parameters
}

// JobsClientListOptions contains the optional parameters for the JobsClient.NewListPager method.
type JobsClientListOptions struct {
	// Type of job to be returned.
	JobType *string

	// View type for including/excluding (for example) archived entities.
	ListViewType *ListViewType

	// The scheduled id for listing the job triggered from
	ScheduleID *string

	// Indicator whether the job is scheduled job.
	Scheduled *bool

	// Continuation token for pagination.
	Skip *string

	// Jobs returned will have this tag key.
	Tag *string
}

// ModelContainersClientCreateOrUpdateOptions contains the optional parameters for the ModelContainersClient.CreateOrUpdate
// method.
type ModelContainersClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ModelContainersClientDeleteOptions contains the optional parameters for the ModelContainersClient.Delete method.
type ModelContainersClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ModelContainersClientGetOptions contains the optional parameters for the ModelContainersClient.Get method.
type ModelContainersClientGetOptions struct {
	// placeholder for future optional parameters
}

// ModelContainersClientListOptions contains the optional parameters for the ModelContainersClient.NewListPager method.
type ModelContainersClientListOptions struct {
	// Maximum number of results to return.
	Count *int32

	// View type for including/excluding (for example) archived entities.
	ListViewType *ListViewType

	// Continuation token for pagination.
	Skip *string
}

// ModelVersionsClientCreateOrUpdateOptions contains the optional parameters for the ModelVersionsClient.CreateOrUpdate method.
type ModelVersionsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ModelVersionsClientDeleteOptions contains the optional parameters for the ModelVersionsClient.Delete method.
type ModelVersionsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ModelVersionsClientGetOptions contains the optional parameters for the ModelVersionsClient.Get method.
type ModelVersionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ModelVersionsClientListOptions contains the optional parameters for the ModelVersionsClient.NewListPager method.
type ModelVersionsClientListOptions struct {
	// Model description.
	Description *string

	// Name of the feed.
	Feed *string

	// View type for including/excluding (for example) archived entities.
	ListViewType *ListViewType

	// Number of initial results to skip.
	Offset *int32

	// Ordering of list.
	OrderBy *string

	// Comma-separated list of property names (and optionally values). Example: prop1,prop2=value2
	Properties *string

	// Continuation token for pagination.
	Skip *string

	// Comma-separated list of tag names (and optionally values). Example: tag1,tag2=value2
	Tags *string

	// Maximum number of records to return.
	Top *int32

	// Model version.
	Version *string
}

// OnlineDeploymentsClientBeginCreateOrUpdateOptions contains the optional parameters for the OnlineDeploymentsClient.BeginCreateOrUpdate
// method.
type OnlineDeploymentsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// OnlineDeploymentsClientBeginDeleteOptions contains the optional parameters for the OnlineDeploymentsClient.BeginDelete
// method.
type OnlineDeploymentsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// OnlineDeploymentsClientBeginUpdateOptions contains the optional parameters for the OnlineDeploymentsClient.BeginUpdate
// method.
type OnlineDeploymentsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// OnlineDeploymentsClientGetLogsOptions contains the optional parameters for the OnlineDeploymentsClient.GetLogs method.
type OnlineDeploymentsClientGetLogsOptions struct {
	// placeholder for future optional parameters
}

// OnlineDeploymentsClientGetOptions contains the optional parameters for the OnlineDeploymentsClient.Get method.
type OnlineDeploymentsClientGetOptions struct {
	// placeholder for future optional parameters
}

// OnlineDeploymentsClientListOptions contains the optional parameters for the OnlineDeploymentsClient.NewListPager method.
type OnlineDeploymentsClientListOptions struct {
	// Ordering of list.
	OrderBy *string

	// Continuation token for pagination.
	Skip *string

	// Top of list.
	Top *int32
}

// OnlineDeploymentsClientListSKUsOptions contains the optional parameters for the OnlineDeploymentsClient.NewListSKUsPager
// method.
type OnlineDeploymentsClientListSKUsOptions struct {
	// Number of Skus to be retrieved in a page of results.
	Count *int32

	// Continuation token for pagination.
	Skip *string
}

// OnlineEndpointsClientBeginCreateOrUpdateOptions contains the optional parameters for the OnlineEndpointsClient.BeginCreateOrUpdate
// method.
type OnlineEndpointsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// OnlineEndpointsClientBeginDeleteOptions contains the optional parameters for the OnlineEndpointsClient.BeginDelete method.
type OnlineEndpointsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// OnlineEndpointsClientBeginRegenerateKeysOptions contains the optional parameters for the OnlineEndpointsClient.BeginRegenerateKeys
// method.
type OnlineEndpointsClientBeginRegenerateKeysOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// OnlineEndpointsClientBeginUpdateOptions contains the optional parameters for the OnlineEndpointsClient.BeginUpdate method.
type OnlineEndpointsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// OnlineEndpointsClientGetOptions contains the optional parameters for the OnlineEndpointsClient.Get method.
type OnlineEndpointsClientGetOptions struct {
	// placeholder for future optional parameters
}

// OnlineEndpointsClientGetTokenOptions contains the optional parameters for the OnlineEndpointsClient.GetToken method.
type OnlineEndpointsClientGetTokenOptions struct {
	// placeholder for future optional parameters
}

// OnlineEndpointsClientListKeysOptions contains the optional parameters for the OnlineEndpointsClient.ListKeys method.
type OnlineEndpointsClientListKeysOptions struct {
	// placeholder for future optional parameters
}

// OnlineEndpointsClientListOptions contains the optional parameters for the OnlineEndpointsClient.NewListPager method.
type OnlineEndpointsClientListOptions struct {
	// EndpointComputeType to be filtered by.
	ComputeType *EndpointComputeType

	// Number of endpoints to be retrieved in a page of results.
	Count *int32

	// Name of the endpoint.
	Name *string

	// The option to order the response.
	OrderBy *OrderString

	// A set of properties with which to filter the returned models. It is a comma separated string of properties key and/or properties
	// key=value Example: propKey1,propKey2,propKey3=value3 .
	Properties *string

	// Continuation token for pagination.
	Skip *string

	// A set of tags with which to filter the returned models. It is a comma separated string of tags key or tags key=value. Example:
	// tagKey1,tagKey2,tagKey3=value3 .
	Tags *string
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientCreateOrUpdateOptions contains the optional parameters for the PrivateEndpointConnectionsClient.CreateOrUpdate
// method.
type PrivateEndpointConnectionsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientDeleteOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Delete
// method.
type PrivateEndpointConnectionsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientGetOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Get
// method.
type PrivateEndpointConnectionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientListOptions contains the optional parameters for the PrivateEndpointConnectionsClient.NewListPager
// method.
type PrivateEndpointConnectionsClientListOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkResourcesClientListOptions contains the optional parameters for the PrivateLinkResourcesClient.List method.
type PrivateLinkResourcesClientListOptions struct {
	// placeholder for future optional parameters
}

// QuotasClientListOptions contains the optional parameters for the QuotasClient.NewListPager method.
type QuotasClientListOptions struct {
	// placeholder for future optional parameters
}

// QuotasClientUpdateOptions contains the optional parameters for the QuotasClient.Update method.
type QuotasClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// UsagesClientListOptions contains the optional parameters for the UsagesClient.NewListPager method.
type UsagesClientListOptions struct {
	// placeholder for future optional parameters
}

// VirtualMachineSizesClientListOptions contains the optional parameters for the VirtualMachineSizesClient.List method.
type VirtualMachineSizesClientListOptions struct {
	// placeholder for future optional parameters
}

// WorkspaceConnectionsClientCreateOptions contains the optional parameters for the WorkspaceConnectionsClient.Create method.
type WorkspaceConnectionsClientCreateOptions struct {
	// placeholder for future optional parameters
}

// WorkspaceConnectionsClientDeleteOptions contains the optional parameters for the WorkspaceConnectionsClient.Delete method.
type WorkspaceConnectionsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// WorkspaceConnectionsClientGetOptions contains the optional parameters for the WorkspaceConnectionsClient.Get method.
type WorkspaceConnectionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// WorkspaceConnectionsClientListOptions contains the optional parameters for the WorkspaceConnectionsClient.NewListPager
// method.
type WorkspaceConnectionsClientListOptions struct {
	// Category of the workspace connection.
	Category *string

	// Target of the workspace connection.
	Target *string
}

// WorkspaceFeaturesClientListOptions contains the optional parameters for the WorkspaceFeaturesClient.NewListPager method.
type WorkspaceFeaturesClientListOptions struct {
	// placeholder for future optional parameters
}

// WorkspacesClientBeginCreateOrUpdateOptions contains the optional parameters for the WorkspacesClient.BeginCreateOrUpdate
// method.
type WorkspacesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// WorkspacesClientBeginDeleteOptions contains the optional parameters for the WorkspacesClient.BeginDelete method.
type WorkspacesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// WorkspacesClientBeginDiagnoseOptions contains the optional parameters for the WorkspacesClient.BeginDiagnose method.
type WorkspacesClientBeginDiagnoseOptions struct {
	// The parameter of diagnosing workspace health
	Parameters *DiagnoseWorkspaceParameters

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// WorkspacesClientBeginPrepareNotebookOptions contains the optional parameters for the WorkspacesClient.BeginPrepareNotebook
// method.
type WorkspacesClientBeginPrepareNotebookOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// WorkspacesClientBeginResyncKeysOptions contains the optional parameters for the WorkspacesClient.BeginResyncKeys method.
type WorkspacesClientBeginResyncKeysOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// WorkspacesClientBeginUpdateOptions contains the optional parameters for the WorkspacesClient.BeginUpdate method.
type WorkspacesClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// WorkspacesClientGetOptions contains the optional parameters for the WorkspacesClient.Get method.
type WorkspacesClientGetOptions struct {
	// placeholder for future optional parameters
}

// WorkspacesClientListByResourceGroupOptions contains the optional parameters for the WorkspacesClient.NewListByResourceGroupPager
// method.
type WorkspacesClientListByResourceGroupOptions struct {
	// Continuation token for pagination.
	Skip *string
}

// WorkspacesClientListBySubscriptionOptions contains the optional parameters for the WorkspacesClient.NewListBySubscriptionPager
// method.
type WorkspacesClientListBySubscriptionOptions struct {
	// Continuation token for pagination.
	Skip *string
}

// WorkspacesClientListKeysOptions contains the optional parameters for the WorkspacesClient.ListKeys method.
type WorkspacesClientListKeysOptions struct {
	// placeholder for future optional parameters
}

// WorkspacesClientListNotebookAccessTokenOptions contains the optional parameters for the WorkspacesClient.ListNotebookAccessToken
// method.
type WorkspacesClientListNotebookAccessTokenOptions struct {
	// placeholder for future optional parameters
}

// WorkspacesClientListNotebookKeysOptions contains the optional parameters for the WorkspacesClient.ListNotebookKeys method.
type WorkspacesClientListNotebookKeysOptions struct {
	// placeholder for future optional parameters
}

// WorkspacesClientListOutboundNetworkDependenciesEndpointsOptions contains the optional parameters for the WorkspacesClient.ListOutboundNetworkDependenciesEndpoints
// method.
type WorkspacesClientListOutboundNetworkDependenciesEndpointsOptions struct {
	// placeholder for future optional parameters
}

// WorkspacesClientListStorageAccountKeysOptions contains the optional parameters for the WorkspacesClient.ListStorageAccountKeys
// method.
type WorkspacesClientListStorageAccountKeysOptions struct {
	// placeholder for future optional parameters
}
