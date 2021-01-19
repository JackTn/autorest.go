// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azspark

type PluginCurrentState string

const (
	PluginCurrentStateCleanup             PluginCurrentState = "Cleanup"
	PluginCurrentStateEnded               PluginCurrentState = "Ended"
	PluginCurrentStateMonitoring          PluginCurrentState = "Monitoring"
	PluginCurrentStatePreparation         PluginCurrentState = "Preparation"
	PluginCurrentStateQueued              PluginCurrentState = "Queued"
	PluginCurrentStateResourceAcquisition PluginCurrentState = "ResourceAcquisition"
	PluginCurrentStateSubmission          PluginCurrentState = "Submission"
)

// PossiblePluginCurrentStateValues returns the possible values for the PluginCurrentState const type.
func PossiblePluginCurrentStateValues() []PluginCurrentState {
	return []PluginCurrentState{
		PluginCurrentStateCleanup,
		PluginCurrentStateEnded,
		PluginCurrentStateMonitoring,
		PluginCurrentStatePreparation,
		PluginCurrentStateQueued,
		PluginCurrentStateResourceAcquisition,
		PluginCurrentStateSubmission,
	}
}

func (c PluginCurrentState) ToPtr() *PluginCurrentState {
	return &c
}

type SchedulerCurrentState string

const (
	SchedulerCurrentStateEnded     SchedulerCurrentState = "Ended"
	SchedulerCurrentStateQueued    SchedulerCurrentState = "Queued"
	SchedulerCurrentStateScheduled SchedulerCurrentState = "Scheduled"
)

// PossibleSchedulerCurrentStateValues returns the possible values for the SchedulerCurrentState const type.
func PossibleSchedulerCurrentStateValues() []SchedulerCurrentState {
	return []SchedulerCurrentState{
		SchedulerCurrentStateEnded,
		SchedulerCurrentStateQueued,
		SchedulerCurrentStateScheduled,
	}
}

func (c SchedulerCurrentState) ToPtr() *SchedulerCurrentState {
	return &c
}

// SparkBatchJobResultType - The Spark batch job result.
type SparkBatchJobResultType string

const (
	SparkBatchJobResultTypeCancelled SparkBatchJobResultType = "Cancelled"
	SparkBatchJobResultTypeFailed    SparkBatchJobResultType = "Failed"
	SparkBatchJobResultTypeSucceeded SparkBatchJobResultType = "Succeeded"
	SparkBatchJobResultTypeUncertain SparkBatchJobResultType = "Uncertain"
)

// PossibleSparkBatchJobResultTypeValues returns the possible values for the SparkBatchJobResultType const type.
func PossibleSparkBatchJobResultTypeValues() []SparkBatchJobResultType {
	return []SparkBatchJobResultType{
		SparkBatchJobResultTypeCancelled,
		SparkBatchJobResultTypeFailed,
		SparkBatchJobResultTypeSucceeded,
		SparkBatchJobResultTypeUncertain,
	}
}

func (c SparkBatchJobResultType) ToPtr() *SparkBatchJobResultType {
	return &c
}

type SparkErrorSource string

const (
	SparkErrorSourceDependency SparkErrorSource = "Dependency"
	SparkErrorSourceSystem     SparkErrorSource = "System"
	SparkErrorSourceUnknown    SparkErrorSource = "Unknown"
	SparkErrorSourceUser       SparkErrorSource = "User"
)

// PossibleSparkErrorSourceValues returns the possible values for the SparkErrorSource const type.
func PossibleSparkErrorSourceValues() []SparkErrorSource {
	return []SparkErrorSource{
		SparkErrorSourceDependency,
		SparkErrorSourceSystem,
		SparkErrorSourceUnknown,
		SparkErrorSourceUser,
	}
}

func (c SparkErrorSource) ToPtr() *SparkErrorSource {
	return &c
}

// SparkJobType - The job type.
type SparkJobType string

const (
	SparkJobTypeSparkBatch   SparkJobType = "SparkBatch"
	SparkJobTypeSparkSession SparkJobType = "SparkSession"
)

// PossibleSparkJobTypeValues returns the possible values for the SparkJobType const type.
func PossibleSparkJobTypeValues() []SparkJobType {
	return []SparkJobType{
		SparkJobTypeSparkBatch,
		SparkJobTypeSparkSession,
	}
}

func (c SparkJobType) ToPtr() *SparkJobType {
	return &c
}

type SparkSessionResultType string

const (
	SparkSessionResultTypeCancelled SparkSessionResultType = "Cancelled"
	SparkSessionResultTypeFailed    SparkSessionResultType = "Failed"
	SparkSessionResultTypeSucceeded SparkSessionResultType = "Succeeded"
	SparkSessionResultTypeUncertain SparkSessionResultType = "Uncertain"
)

// PossibleSparkSessionResultTypeValues returns the possible values for the SparkSessionResultType const type.
func PossibleSparkSessionResultTypeValues() []SparkSessionResultType {
	return []SparkSessionResultType{
		SparkSessionResultTypeCancelled,
		SparkSessionResultTypeFailed,
		SparkSessionResultTypeSucceeded,
		SparkSessionResultTypeUncertain,
	}
}

func (c SparkSessionResultType) ToPtr() *SparkSessionResultType {
	return &c
}

type SparkStatementLanguageType string

const (
	SparkStatementLanguageTypeDotNetSpark SparkStatementLanguageType = "dotnetspark"
	SparkStatementLanguageTypePySpark     SparkStatementLanguageType = "pyspark"
	SparkStatementLanguageTypeSQL         SparkStatementLanguageType = "sql"
	SparkStatementLanguageTypeSpark       SparkStatementLanguageType = "spark"
)

// PossibleSparkStatementLanguageTypeValues returns the possible values for the SparkStatementLanguageType const type.
func PossibleSparkStatementLanguageTypeValues() []SparkStatementLanguageType {
	return []SparkStatementLanguageType{
		SparkStatementLanguageTypeDotNetSpark,
		SparkStatementLanguageTypePySpark,
		SparkStatementLanguageTypeSQL,
		SparkStatementLanguageTypeSpark,
	}
}

func (c SparkStatementLanguageType) ToPtr() *SparkStatementLanguageType {
	return &c
}
