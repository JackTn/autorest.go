// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azspark

type LivyStatementStates string

const (
	LivyStatementStatesAvailable  LivyStatementStates = "available"
	LivyStatementStatesCancelled  LivyStatementStates = "cancelled"
	LivyStatementStatesCancelling LivyStatementStates = "cancelling"
	LivyStatementStatesError      LivyStatementStates = "error"
	LivyStatementStatesRunning    LivyStatementStates = "running"
	LivyStatementStatesWaiting    LivyStatementStates = "waiting"
)

// PossibleLivyStatementStatesValues returns the possible values for the LivyStatementStates const type.
func PossibleLivyStatementStatesValues() []LivyStatementStates {
	return []LivyStatementStates{
		LivyStatementStatesAvailable,
		LivyStatementStatesCancelled,
		LivyStatementStatesCancelling,
		LivyStatementStatesError,
		LivyStatementStatesRunning,
		LivyStatementStatesWaiting,
	}
}

// LivyStates - The batch state
type LivyStates string

const (
	LivyStatesBusy         LivyStates = "busy"
	LivyStatesDead         LivyStates = "dead"
	LivyStatesError        LivyStates = "error"
	LivyStatesIdle         LivyStates = "idle"
	LivyStatesKilled       LivyStates = "killed"
	LivyStatesNotStarted   LivyStates = "not_started"
	LivyStatesRecovering   LivyStates = "recovering"
	LivyStatesRunning      LivyStates = "running"
	LivyStatesShuttingDown LivyStates = "shutting_down"
	LivyStatesStarting     LivyStates = "starting"
	LivyStatesSuccess      LivyStates = "success"
)

// PossibleLivyStatesValues returns the possible values for the LivyStates const type.
func PossibleLivyStatesValues() []LivyStates {
	return []LivyStates{
		LivyStatesBusy,
		LivyStatesDead,
		LivyStatesError,
		LivyStatesIdle,
		LivyStatesKilled,
		LivyStatesNotStarted,
		LivyStatesRecovering,
		LivyStatesRunning,
		LivyStatesShuttingDown,
		LivyStatesStarting,
		LivyStatesSuccess,
	}
}

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
