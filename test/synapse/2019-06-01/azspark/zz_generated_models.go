// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azspark

import (
	"encoding/json"
	"net/http"
	"time"
)

// SparkBatchCancelSparkBatchJobOptions contains the optional parameters for the SparkBatch.CancelSparkBatchJob method.
type SparkBatchCancelSparkBatchJobOptions struct {
	// placeholder for future optional parameters
}

// SparkBatchCreateSparkBatchJobOptions contains the optional parameters for the SparkBatch.CreateSparkBatchJob method.
type SparkBatchCreateSparkBatchJobOptions struct {
	// Optional query param specifying whether detailed response is returned beyond plain livy.
	Detailed *bool
}

// SparkBatchGetSparkBatchJobOptions contains the optional parameters for the SparkBatch.GetSparkBatchJob method.
type SparkBatchGetSparkBatchJobOptions struct {
	// Optional query param specifying whether detailed response is returned beyond plain livy.
	Detailed *bool
}

// SparkBatchGetSparkBatchJobsOptions contains the optional parameters for the SparkBatch.GetSparkBatchJobs method.
type SparkBatchGetSparkBatchJobsOptions struct {
	// Optional query param specifying whether detailed response is returned beyond plain livy.
	Detailed *bool
	// Optional param specifying which index the list should begin from.
	From *int32
	// Optional param specifying the size of the returned list.
	// By default it is 20 and that is the maximum.
	Size *int32
}

type SparkBatchJob struct {
	// The application id of this session
	AppID *string `json:"appId,omitempty"`

	// The detailed application info.
	AppInfo *map[string]string `json:"appInfo,omitempty"`

	// The artifact identifier.
	ArtifactID *string `json:"artifactId,omitempty"`

	// The error information.
	Errors *[]SparkServiceError `json:"errorInfo,omitempty"`

	// The session Id.
	ID *int32 `json:"id,omitempty"`

	// The job type.
	JobType  *SparkJobType       `json:"jobType,omitempty"`
	LivyInfo *SparkBatchJobState `json:"livyInfo,omitempty"`

	// The log lines.
	LogLines *[]string `json:"log,omitempty"`

	// The batch name.
	Name *string `json:"name,omitempty"`

	// The plugin information.
	Plugin *SparkServicePlugin `json:"pluginInfo,omitempty"`

	// The Spark batch job result.
	Result *SparkBatchJobResultType `json:"result,omitempty"`

	// The scheduler information.
	Scheduler *SparkScheduler `json:"schedulerInfo,omitempty"`

	// The Spark pool name.
	SparkPoolName *string `json:"sparkPoolName,omitempty"`

	// The batch state
	State *string `json:"state,omitempty"`

	// The submitter identifier.
	SubmitterID *string `json:"submitterId,omitempty"`

	// The submitter name.
	SubmitterName *string `json:"submitterName,omitempty"`

	// The tags.
	Tags *map[string]string `json:"tags,omitempty"`

	// The workspace name.
	WorkspaceName *string `json:"workspaceName,omitempty"`
}

// Response for batch list operation.
type SparkBatchJobCollection struct {
	// The start index of fetched sessions.
	From *int32 `json:"from,omitempty"`

	// Batch list
	Sessions *[]SparkBatchJob `json:"sessions,omitempty"`

	// Number of sessions fetched.
	Total *int32 `json:"total,omitempty"`
}

// SparkBatchJobCollectionResponse is the response envelope for operations that return a SparkBatchJobCollection type.
type SparkBatchJobCollectionResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Response for batch list operation.
	SparkBatchJobCollection *SparkBatchJobCollection
}

type SparkBatchJobOptions struct {
	Archives   *[]string `json:"archives,omitempty"`
	Arguments  *[]string `json:"args,omitempty"`
	ArtifactID *string   `json:"artifactId,omitempty"`
	ClassName  *string   `json:"className,omitempty"`

	// Dictionary of
	Configuration  *map[string]string `json:"conf,omitempty"`
	DriverCores    *int32             `json:"driverCores,omitempty"`
	DriverMemory   *string            `json:"driverMemory,omitempty"`
	ExecutorCores  *int32             `json:"executorCores,omitempty"`
	ExecutorCount  *int32             `json:"numExecutors,omitempty"`
	ExecutorMemory *string            `json:"executorMemory,omitempty"`
	File           *string            `json:"file,omitempty"`
	Files          *[]string          `json:"files,omitempty"`
	Jars           *[]string          `json:"jars,omitempty"`
	Name           *string            `json:"name,omitempty"`
	PythonFiles    *[]string          `json:"pyFiles,omitempty"`

	// Dictionary of
	Tags *map[string]string `json:"tags,omitempty"`
}

// SparkBatchJobResponse is the response envelope for operations that return a SparkBatchJob type.
type SparkBatchJobResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse   *http.Response
	SparkBatchJob *SparkBatchJob
}

type SparkBatchJobState struct {
	// the Spark job state.
	CurrentState *string `json:"currentState,omitempty"`

	// time that at which "dead" livy state was first seen.
	DeadAt             *time.Time    `json:"deadAt,omitempty"`
	JobCreationRequest *SparkRequest `json:"jobCreationRequest,omitempty"`

	// the time that at which "not_started" livy state was first seen.
	NotStartedAt *time.Time `json:"notStartedAt,omitempty"`

	// the time that at which "recovering" livy state was first seen.
	RecoveringAt *time.Time `json:"recoveringAt,omitempty"`

	// the time that at which "running" livy state was first seen.
	RunningAt *time.Time `json:"runningAt,omitempty"`

	// the time that at which "starting" livy state was first seen.
	StartingAt *time.Time `json:"startingAt,omitempty"`

	// the time that at which "success" livy state was first seen.
	SuccessAt *time.Time `json:"successAt,omitempty"`

	// the time that at which "killed" livy state was first seen.
	TerminatedAt *time.Time `json:"killedAt,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type SparkBatchJobState.
func (s SparkBatchJobState) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if s.CurrentState != nil {
		objectMap["currentState"] = s.CurrentState
	}
	if s.DeadAt != nil {
		objectMap["deadAt"] = (*timeRFC3339)(s.DeadAt)
	}
	if s.JobCreationRequest != nil {
		objectMap["jobCreationRequest"] = s.JobCreationRequest
	}
	if s.NotStartedAt != nil {
		objectMap["notStartedAt"] = (*timeRFC3339)(s.NotStartedAt)
	}
	if s.RecoveringAt != nil {
		objectMap["recoveringAt"] = (*timeRFC3339)(s.RecoveringAt)
	}
	if s.RunningAt != nil {
		objectMap["runningAt"] = (*timeRFC3339)(s.RunningAt)
	}
	if s.StartingAt != nil {
		objectMap["startingAt"] = (*timeRFC3339)(s.StartingAt)
	}
	if s.SuccessAt != nil {
		objectMap["successAt"] = (*timeRFC3339)(s.SuccessAt)
	}
	if s.TerminatedAt != nil {
		objectMap["killedAt"] = (*timeRFC3339)(s.TerminatedAt)
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SparkBatchJobState.
func (s *SparkBatchJobState) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "currentState":
			if val != nil {
				err = json.Unmarshal(*val, &s.CurrentState)
			}
			delete(rawMsg, key)
		case "deadAt":
			if val != nil {
				var aux timeRFC3339
				err = json.Unmarshal(*val, &aux)
				s.DeadAt = (*time.Time)(&aux)
			}
			delete(rawMsg, key)
		case "jobCreationRequest":
			if val != nil {
				err = json.Unmarshal(*val, &s.JobCreationRequest)
			}
			delete(rawMsg, key)
		case "notStartedAt":
			if val != nil {
				var aux timeRFC3339
				err = json.Unmarshal(*val, &aux)
				s.NotStartedAt = (*time.Time)(&aux)
			}
			delete(rawMsg, key)
		case "recoveringAt":
			if val != nil {
				var aux timeRFC3339
				err = json.Unmarshal(*val, &aux)
				s.RecoveringAt = (*time.Time)(&aux)
			}
			delete(rawMsg, key)
		case "runningAt":
			if val != nil {
				var aux timeRFC3339
				err = json.Unmarshal(*val, &aux)
				s.RunningAt = (*time.Time)(&aux)
			}
			delete(rawMsg, key)
		case "startingAt":
			if val != nil {
				var aux timeRFC3339
				err = json.Unmarshal(*val, &aux)
				s.StartingAt = (*time.Time)(&aux)
			}
			delete(rawMsg, key)
		case "successAt":
			if val != nil {
				var aux timeRFC3339
				err = json.Unmarshal(*val, &aux)
				s.SuccessAt = (*time.Time)(&aux)
			}
			delete(rawMsg, key)
		case "killedAt":
			if val != nil {
				var aux timeRFC3339
				err = json.Unmarshal(*val, &aux)
				s.TerminatedAt = (*time.Time)(&aux)
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

type SparkRequest struct {
	Archives  *[]string `json:"archives,omitempty"`
	Arguments *[]string `json:"args,omitempty"`
	ClassName *string   `json:"className,omitempty"`

	// Dictionary of
	Configuration  *map[string]string `json:"conf,omitempty"`
	DriverCores    *int32             `json:"driverCores,omitempty"`
	DriverMemory   *string            `json:"driverMemory,omitempty"`
	ExecutorCores  *int32             `json:"executorCores,omitempty"`
	ExecutorCount  *int32             `json:"numExecutors,omitempty"`
	ExecutorMemory *string            `json:"executorMemory,omitempty"`
	File           *string            `json:"file,omitempty"`
	Files          *[]string          `json:"files,omitempty"`
	Jars           *[]string          `json:"jars,omitempty"`
	Name           *string            `json:"name,omitempty"`
	PythonFiles    *[]string          `json:"pyFiles,omitempty"`
}

type SparkScheduler struct {
	CancellationRequestedAt *time.Time             `json:"cancellationRequestedAt,omitempty"`
	CurrentState            *SchedulerCurrentState `json:"currentState,omitempty"`
	EndedAt                 *time.Time             `json:"endedAt,omitempty"`
	ScheduledAt             *time.Time             `json:"scheduledAt,omitempty"`
	SubmittedAt             *time.Time             `json:"submittedAt,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type SparkScheduler.
func (s SparkScheduler) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if s.CancellationRequestedAt != nil {
		objectMap["cancellationRequestedAt"] = (*timeRFC3339)(s.CancellationRequestedAt)
	}
	if s.CurrentState != nil {
		objectMap["currentState"] = s.CurrentState
	}
	if s.EndedAt != nil {
		objectMap["endedAt"] = (*timeRFC3339)(s.EndedAt)
	}
	if s.ScheduledAt != nil {
		objectMap["scheduledAt"] = (*timeRFC3339)(s.ScheduledAt)
	}
	if s.SubmittedAt != nil {
		objectMap["submittedAt"] = (*timeRFC3339)(s.SubmittedAt)
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SparkScheduler.
func (s *SparkScheduler) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "cancellationRequestedAt":
			if val != nil {
				var aux timeRFC3339
				err = json.Unmarshal(*val, &aux)
				s.CancellationRequestedAt = (*time.Time)(&aux)
			}
			delete(rawMsg, key)
		case "currentState":
			if val != nil {
				err = json.Unmarshal(*val, &s.CurrentState)
			}
			delete(rawMsg, key)
		case "endedAt":
			if val != nil {
				var aux timeRFC3339
				err = json.Unmarshal(*val, &aux)
				s.EndedAt = (*time.Time)(&aux)
			}
			delete(rawMsg, key)
		case "scheduledAt":
			if val != nil {
				var aux timeRFC3339
				err = json.Unmarshal(*val, &aux)
				s.ScheduledAt = (*time.Time)(&aux)
			}
			delete(rawMsg, key)
		case "submittedAt":
			if val != nil {
				var aux timeRFC3339
				err = json.Unmarshal(*val, &aux)
				s.SubmittedAt = (*time.Time)(&aux)
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

type SparkServiceError struct {
	ErrorCode *string           `json:"errorCode,omitempty"`
	Message   *string           `json:"message,omitempty"`
	Source    *SparkErrorSource `json:"source,omitempty"`
}

type SparkServicePlugin struct {
	CleanupStartedAt             *time.Time          `json:"cleanupStartedAt,omitempty"`
	CurrentState                 *PluginCurrentState `json:"currentState,omitempty"`
	MonitoringStartedAt          *time.Time          `json:"monitoringStartedAt,omitempty"`
	PreparationStartedAt         *time.Time          `json:"preparationStartedAt,omitempty"`
	ResourceAcquisitionStartedAt *time.Time          `json:"resourceAcquisitionStartedAt,omitempty"`
	SubmissionStartedAt          *time.Time          `json:"submissionStartedAt,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type SparkServicePlugin.
func (s SparkServicePlugin) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if s.CleanupStartedAt != nil {
		objectMap["cleanupStartedAt"] = (*timeRFC3339)(s.CleanupStartedAt)
	}
	if s.CurrentState != nil {
		objectMap["currentState"] = s.CurrentState
	}
	if s.MonitoringStartedAt != nil {
		objectMap["monitoringStartedAt"] = (*timeRFC3339)(s.MonitoringStartedAt)
	}
	if s.PreparationStartedAt != nil {
		objectMap["preparationStartedAt"] = (*timeRFC3339)(s.PreparationStartedAt)
	}
	if s.ResourceAcquisitionStartedAt != nil {
		objectMap["resourceAcquisitionStartedAt"] = (*timeRFC3339)(s.ResourceAcquisitionStartedAt)
	}
	if s.SubmissionStartedAt != nil {
		objectMap["submissionStartedAt"] = (*timeRFC3339)(s.SubmissionStartedAt)
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SparkServicePlugin.
func (s *SparkServicePlugin) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "cleanupStartedAt":
			if val != nil {
				var aux timeRFC3339
				err = json.Unmarshal(*val, &aux)
				s.CleanupStartedAt = (*time.Time)(&aux)
			}
			delete(rawMsg, key)
		case "currentState":
			if val != nil {
				err = json.Unmarshal(*val, &s.CurrentState)
			}
			delete(rawMsg, key)
		case "monitoringStartedAt":
			if val != nil {
				var aux timeRFC3339
				err = json.Unmarshal(*val, &aux)
				s.MonitoringStartedAt = (*time.Time)(&aux)
			}
			delete(rawMsg, key)
		case "preparationStartedAt":
			if val != nil {
				var aux timeRFC3339
				err = json.Unmarshal(*val, &aux)
				s.PreparationStartedAt = (*time.Time)(&aux)
			}
			delete(rawMsg, key)
		case "resourceAcquisitionStartedAt":
			if val != nil {
				var aux timeRFC3339
				err = json.Unmarshal(*val, &aux)
				s.ResourceAcquisitionStartedAt = (*time.Time)(&aux)
			}
			delete(rawMsg, key)
		case "submissionStartedAt":
			if val != nil {
				var aux timeRFC3339
				err = json.Unmarshal(*val, &aux)
				s.SubmissionStartedAt = (*time.Time)(&aux)
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

type SparkSession struct {
	AppID *string `json:"appId,omitempty"`

	// Dictionary of
	AppInfo    *map[string]string   `json:"appInfo,omitempty"`
	ArtifactID *string              `json:"artifactId,omitempty"`
	Errors     *[]SparkServiceError `json:"errorInfo,omitempty"`
	ID         *int32               `json:"id,omitempty"`

	// The job type.
	JobType       *SparkJobType           `json:"jobType,omitempty"`
	LivyInfo      *SparkSessionState      `json:"livyInfo,omitempty"`
	LogLines      *[]string               `json:"log,omitempty"`
	Name          *string                 `json:"name,omitempty"`
	Plugin        *SparkServicePlugin     `json:"pluginInfo,omitempty"`
	Result        *SparkSessionResultType `json:"result,omitempty"`
	Scheduler     *SparkScheduler         `json:"schedulerInfo,omitempty"`
	SparkPoolName *string                 `json:"sparkPoolName,omitempty"`
	State         *string                 `json:"state,omitempty"`
	SubmitterID   *string                 `json:"submitterId,omitempty"`
	SubmitterName *string                 `json:"submitterName,omitempty"`

	// Dictionary of
	Tags          *map[string]string `json:"tags,omitempty"`
	WorkspaceName *string            `json:"workspaceName,omitempty"`
}

// SparkSessionCancelSparkSessionOptions contains the optional parameters for the SparkSession.CancelSparkSession method.
type SparkSessionCancelSparkSessionOptions struct {
	// placeholder for future optional parameters
}

// SparkSessionCancelSparkStatementOptions contains the optional parameters for the SparkSession.CancelSparkStatement method.
type SparkSessionCancelSparkStatementOptions struct {
	// placeholder for future optional parameters
}

type SparkSessionCollection struct {
	From     *int32          `json:"from,omitempty"`
	Sessions *[]SparkSession `json:"sessions,omitempty"`
	Total    *int32          `json:"total,omitempty"`
}

// SparkSessionCollectionResponse is the response envelope for operations that return a SparkSessionCollection type.
type SparkSessionCollectionResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse            *http.Response
	SparkSessionCollection *SparkSessionCollection
}

// SparkSessionCreateSparkSessionOptions contains the optional parameters for the SparkSession.CreateSparkSession method.
type SparkSessionCreateSparkSessionOptions struct {
	// Optional query param specifying whether detailed response is returned beyond plain livy.
	Detailed *bool
}

// SparkSessionCreateSparkStatementOptions contains the optional parameters for the SparkSession.CreateSparkStatement method.
type SparkSessionCreateSparkStatementOptions struct {
	// placeholder for future optional parameters
}

// SparkSessionGetSparkSessionOptions contains the optional parameters for the SparkSession.GetSparkSession method.
type SparkSessionGetSparkSessionOptions struct {
	// Optional query param specifying whether detailed response is returned beyond plain livy.
	Detailed *bool
}

// SparkSessionGetSparkSessionsOptions contains the optional parameters for the SparkSession.GetSparkSessions method.
type SparkSessionGetSparkSessionsOptions struct {
	// Optional query param specifying whether detailed response is returned beyond plain livy.
	Detailed *bool
	// Optional param specifying which index the list should begin from.
	From *int32
	// Optional param specifying the size of the returned list.
	// By default it is 20 and that is the maximum.
	Size *int32
}

// SparkSessionGetSparkStatementOptions contains the optional parameters for the SparkSession.GetSparkStatement method.
type SparkSessionGetSparkStatementOptions struct {
	// placeholder for future optional parameters
}

// SparkSessionGetSparkStatementsOptions contains the optional parameters for the SparkSession.GetSparkStatements method.
type SparkSessionGetSparkStatementsOptions struct {
	// placeholder for future optional parameters
}

type SparkSessionOptions struct {
	Archives   *[]string `json:"archives,omitempty"`
	Arguments  *[]string `json:"args,omitempty"`
	ArtifactID *string   `json:"artifactId,omitempty"`
	ClassName  *string   `json:"className,omitempty"`

	// Dictionary of
	Configuration  *map[string]string `json:"conf,omitempty"`
	DriverCores    *int32             `json:"driverCores,omitempty"`
	DriverMemory   *string            `json:"driverMemory,omitempty"`
	ExecutorCores  *int32             `json:"executorCores,omitempty"`
	ExecutorCount  *int32             `json:"numExecutors,omitempty"`
	ExecutorMemory *string            `json:"executorMemory,omitempty"`
	File           *string            `json:"file,omitempty"`
	Files          *[]string          `json:"files,omitempty"`
	Jars           *[]string          `json:"jars,omitempty"`
	Name           *string            `json:"name,omitempty"`
	PythonFiles    *[]string          `json:"pyFiles,omitempty"`

	// Dictionary of
	Tags *map[string]string `json:"tags,omitempty"`
}

// SparkSessionResetSparkSessionTimeoutOptions contains the optional parameters for the SparkSession.ResetSparkSessionTimeout method.
type SparkSessionResetSparkSessionTimeoutOptions struct {
	// placeholder for future optional parameters
}

// SparkSessionResponse is the response envelope for operations that return a SparkSession type.
type SparkSessionResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse  *http.Response
	SparkSession *SparkSession
}

type SparkSessionState struct {
	BusyAt             *time.Time    `json:"busyAt,omitempty"`
	CurrentState       *string       `json:"currentState,omitempty"`
	DeadAt             *time.Time    `json:"deadAt,omitempty"`
	ErrorAt            *time.Time    `json:"errorAt,omitempty"`
	IDleAt             *time.Time    `json:"idleAt,omitempty"`
	JobCreationRequest *SparkRequest `json:"jobCreationRequest,omitempty"`
	NotStartedAt       *time.Time    `json:"notStartedAt,omitempty"`
	RecoveringAt       *time.Time    `json:"recoveringAt,omitempty"`
	ShuttingDownAt     *time.Time    `json:"shuttingDownAt,omitempty"`
	StartingAt         *time.Time    `json:"startingAt,omitempty"`
	TerminatedAt       *time.Time    `json:"killedAt,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type SparkSessionState.
func (s SparkSessionState) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if s.BusyAt != nil {
		objectMap["busyAt"] = (*timeRFC3339)(s.BusyAt)
	}
	if s.CurrentState != nil {
		objectMap["currentState"] = s.CurrentState
	}
	if s.DeadAt != nil {
		objectMap["deadAt"] = (*timeRFC3339)(s.DeadAt)
	}
	if s.ErrorAt != nil {
		objectMap["errorAt"] = (*timeRFC3339)(s.ErrorAt)
	}
	if s.IDleAt != nil {
		objectMap["idleAt"] = (*timeRFC3339)(s.IDleAt)
	}
	if s.JobCreationRequest != nil {
		objectMap["jobCreationRequest"] = s.JobCreationRequest
	}
	if s.NotStartedAt != nil {
		objectMap["notStartedAt"] = (*timeRFC3339)(s.NotStartedAt)
	}
	if s.RecoveringAt != nil {
		objectMap["recoveringAt"] = (*timeRFC3339)(s.RecoveringAt)
	}
	if s.ShuttingDownAt != nil {
		objectMap["shuttingDownAt"] = (*timeRFC3339)(s.ShuttingDownAt)
	}
	if s.StartingAt != nil {
		objectMap["startingAt"] = (*timeRFC3339)(s.StartingAt)
	}
	if s.TerminatedAt != nil {
		objectMap["killedAt"] = (*timeRFC3339)(s.TerminatedAt)
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SparkSessionState.
func (s *SparkSessionState) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "busyAt":
			if val != nil {
				var aux timeRFC3339
				err = json.Unmarshal(*val, &aux)
				s.BusyAt = (*time.Time)(&aux)
			}
			delete(rawMsg, key)
		case "currentState":
			if val != nil {
				err = json.Unmarshal(*val, &s.CurrentState)
			}
			delete(rawMsg, key)
		case "deadAt":
			if val != nil {
				var aux timeRFC3339
				err = json.Unmarshal(*val, &aux)
				s.DeadAt = (*time.Time)(&aux)
			}
			delete(rawMsg, key)
		case "errorAt":
			if val != nil {
				var aux timeRFC3339
				err = json.Unmarshal(*val, &aux)
				s.ErrorAt = (*time.Time)(&aux)
			}
			delete(rawMsg, key)
		case "idleAt":
			if val != nil {
				var aux timeRFC3339
				err = json.Unmarshal(*val, &aux)
				s.IDleAt = (*time.Time)(&aux)
			}
			delete(rawMsg, key)
		case "jobCreationRequest":
			if val != nil {
				err = json.Unmarshal(*val, &s.JobCreationRequest)
			}
			delete(rawMsg, key)
		case "notStartedAt":
			if val != nil {
				var aux timeRFC3339
				err = json.Unmarshal(*val, &aux)
				s.NotStartedAt = (*time.Time)(&aux)
			}
			delete(rawMsg, key)
		case "recoveringAt":
			if val != nil {
				var aux timeRFC3339
				err = json.Unmarshal(*val, &aux)
				s.RecoveringAt = (*time.Time)(&aux)
			}
			delete(rawMsg, key)
		case "shuttingDownAt":
			if val != nil {
				var aux timeRFC3339
				err = json.Unmarshal(*val, &aux)
				s.ShuttingDownAt = (*time.Time)(&aux)
			}
			delete(rawMsg, key)
		case "startingAt":
			if val != nil {
				var aux timeRFC3339
				err = json.Unmarshal(*val, &aux)
				s.StartingAt = (*time.Time)(&aux)
			}
			delete(rawMsg, key)
		case "killedAt":
			if val != nil {
				var aux timeRFC3339
				err = json.Unmarshal(*val, &aux)
				s.TerminatedAt = (*time.Time)(&aux)
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

type SparkStatement struct {
	Code   *string               `json:"code,omitempty"`
	ID     *int32                `json:"id,omitempty"`
	Output *SparkStatementOutput `json:"output,omitempty"`
	State  *string               `json:"state,omitempty"`
}

type SparkStatementCancellationResult struct {
	// The msg property from the Livy API. The value is always "canceled".
	Message *string `json:"msg,omitempty"`
}

// SparkStatementCancellationResultResponse is the response envelope for operations that return a SparkStatementCancellationResult type.
type SparkStatementCancellationResultResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse                      *http.Response
	SparkStatementCancellationResult *SparkStatementCancellationResult
}

type SparkStatementCollection struct {
	Statements *[]SparkStatement `json:"statements,omitempty"`
	Total      *int32            `json:"total_statements,omitempty"`
}

// SparkStatementCollectionResponse is the response envelope for operations that return a SparkStatementCollection type.
type SparkStatementCollectionResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse              *http.Response
	SparkStatementCollection *SparkStatementCollection
}

type SparkStatementOptions struct {
	Code *string                     `json:"code,omitempty"`
	Kind *SparkStatementLanguageType `json:"kind,omitempty"`
}

type SparkStatementOutput struct {
	// Any object
	Data           interface{} `json:"data,omitempty"`
	ErrorName      *string     `json:"ename,omitempty"`
	ErrorValue     *string     `json:"evalue,omitempty"`
	ExecutionCount *int32      `json:"execution_count,omitempty"`
	Status         *string     `json:"status,omitempty"`
	Traceback      *[]string   `json:"traceback,omitempty"`
}

// SparkStatementResponse is the response envelope for operations that return a SparkStatement type.
type SparkStatementResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse    *http.Response
	SparkStatement *SparkStatement
}
