// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armlargeinstance

const host = "https://management.azure.com"

const (
	moduleName    = "armlargeinstance"
	moduleVersion = "v0.1.0"
)

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	// ActionTypeInternal - Actions are for internal-only APIs.
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// AzureLargeInstanceForcePowerState - Enum of two possible values to determine if the ALI instance restart operation should
// forcefully terminate and halt any existing processes that may be running on the server or not.
type AzureLargeInstanceForcePowerState string

const (
	// AzureLargeInstanceForcePowerStateActive - Active means that the restart operation will terminate and halt existing processes
	// that may be running on the server
	AzureLargeInstanceForcePowerStateActive AzureLargeInstanceForcePowerState = "active"
	// AzureLargeInstanceForcePowerStateInactive - Inactive means that the restart operation will not terminate and halt existing
	// processes that may be running on the server
	AzureLargeInstanceForcePowerStateInactive AzureLargeInstanceForcePowerState = "inactive"
)

// PossibleAzureLargeInstanceForcePowerStateValues returns the possible values for the AzureLargeInstanceForcePowerState const type.
func PossibleAzureLargeInstanceForcePowerStateValues() []AzureLargeInstanceForcePowerState {
	return []AzureLargeInstanceForcePowerState{
		AzureLargeInstanceForcePowerStateActive,
		AzureLargeInstanceForcePowerStateInactive,
	}
}

// AzureLargeInstanceHardwareTypeNamesEnum - Enum of the hardware options (vendor and/or their product name) for an Azure
// Large Instance
type AzureLargeInstanceHardwareTypeNamesEnum string

const (
	// AzureLargeInstanceHardwareTypeNamesEnumCiscoUCS - Hardware type of UCS from vendor Cisco
	AzureLargeInstanceHardwareTypeNamesEnumCiscoUCS AzureLargeInstanceHardwareTypeNamesEnum = "Cisco_UCS"
	// AzureLargeInstanceHardwareTypeNamesEnumHPE - Hardware type of HPE from vendor Hewlett Packard Enterprise
	AzureLargeInstanceHardwareTypeNamesEnumHPE AzureLargeInstanceHardwareTypeNamesEnum = "HPE"
	// AzureLargeInstanceHardwareTypeNamesEnumSDFLEX - Hardware type of SDFLEX
	AzureLargeInstanceHardwareTypeNamesEnumSDFLEX AzureLargeInstanceHardwareTypeNamesEnum = "SDFLEX"
)

// PossibleAzureLargeInstanceHardwareTypeNamesEnumValues returns the possible values for the AzureLargeInstanceHardwareTypeNamesEnum const type.
func PossibleAzureLargeInstanceHardwareTypeNamesEnumValues() []AzureLargeInstanceHardwareTypeNamesEnum {
	return []AzureLargeInstanceHardwareTypeNamesEnum{
		AzureLargeInstanceHardwareTypeNamesEnumCiscoUCS,
		AzureLargeInstanceHardwareTypeNamesEnumHPE,
		AzureLargeInstanceHardwareTypeNamesEnumSDFLEX,
	}
}

// AzureLargeInstancePowerStateEnum - Power states that an Azure Large Instance can be in
type AzureLargeInstancePowerStateEnum string

const (
	// AzureLargeInstancePowerStateEnumRestarting - Restarting means that the Azure Large Instance resource is restarting.
	AzureLargeInstancePowerStateEnumRestarting AzureLargeInstancePowerStateEnum = "restarting"
	// AzureLargeInstancePowerStateEnumStarted - Started means that the Azure Large Instance resource has been powered on.
	AzureLargeInstancePowerStateEnumStarted AzureLargeInstancePowerStateEnum = "started"
	// AzureLargeInstancePowerStateEnumStarting - Starting means that the Azure Large Instance resource is turning on.
	AzureLargeInstancePowerStateEnumStarting AzureLargeInstancePowerStateEnum = "starting"
	// AzureLargeInstancePowerStateEnumStopped - Stopped means that the Azure Large Instance resource has shut down.
	AzureLargeInstancePowerStateEnumStopped AzureLargeInstancePowerStateEnum = "stopped"
	// AzureLargeInstancePowerStateEnumStopping - Stopping means that the Azure Large Instance resource is shutting down.
	AzureLargeInstancePowerStateEnumStopping AzureLargeInstancePowerStateEnum = "stopping"
	// AzureLargeInstancePowerStateEnumUnknown - Unknown means that the state of the Azure Large Instance is unknown.
	AzureLargeInstancePowerStateEnumUnknown AzureLargeInstancePowerStateEnum = "unknown"
)

// PossibleAzureLargeInstancePowerStateEnumValues returns the possible values for the AzureLargeInstancePowerStateEnum const type.
func PossibleAzureLargeInstancePowerStateEnumValues() []AzureLargeInstancePowerStateEnum {
	return []AzureLargeInstancePowerStateEnum{
		AzureLargeInstancePowerStateEnumRestarting,
		AzureLargeInstancePowerStateEnumStarted,
		AzureLargeInstancePowerStateEnumStarting,
		AzureLargeInstancePowerStateEnumStopped,
		AzureLargeInstancePowerStateEnumStopping,
		AzureLargeInstancePowerStateEnumUnknown,
	}
}

// AzureLargeInstanceProvisioningStatesEnum - Provisioning states that an Azure Large Instance can be in
type AzureLargeInstanceProvisioningStatesEnum string

const (
	// AzureLargeInstanceProvisioningStatesEnumAccepted - Accepted means Azure Large Instance resource provisioning has been accepted.
	AzureLargeInstanceProvisioningStatesEnumAccepted AzureLargeInstanceProvisioningStatesEnum = "Accepted"
	// AzureLargeInstanceProvisioningStatesEnumCanceled - Cancelled Azure Large Instance resource operation has been cancelled
	AzureLargeInstanceProvisioningStatesEnumCanceled AzureLargeInstanceProvisioningStatesEnum = "Canceled"
	// AzureLargeInstanceProvisioningStatesEnumCreating - Creating means Azure Large Instance resource is being created.
	AzureLargeInstanceProvisioningStatesEnumCreating AzureLargeInstanceProvisioningStatesEnum = "Creating"
	// AzureLargeInstanceProvisioningStatesEnumDeleting - Deleting means Azure Large Instance resource is in the process of being
	// deleted
	AzureLargeInstanceProvisioningStatesEnumDeleting AzureLargeInstanceProvisioningStatesEnum = "Deleting"
	// AzureLargeInstanceProvisioningStatesEnumFailed - Failed means Azure Large Instance resource is in failed state
	AzureLargeInstanceProvisioningStatesEnumFailed AzureLargeInstanceProvisioningStatesEnum = "Failed"
	// AzureLargeInstanceProvisioningStatesEnumMigrating - Migrating means Azure Large Instance resource is being migrated from
	// one subscription or resource group to another
	AzureLargeInstanceProvisioningStatesEnumMigrating AzureLargeInstanceProvisioningStatesEnum = "Migrating"
	// AzureLargeInstanceProvisioningStatesEnumSucceeded - Succeeded means Azure Large Instance resource creation succeeded during
	// last create/update
	AzureLargeInstanceProvisioningStatesEnumSucceeded AzureLargeInstanceProvisioningStatesEnum = "Succeeded"
	// AzureLargeInstanceProvisioningStatesEnumUpdating - Updating means an existing Azure Large Instance resource is being updated
	AzureLargeInstanceProvisioningStatesEnumUpdating AzureLargeInstanceProvisioningStatesEnum = "Updating"
)

// PossibleAzureLargeInstanceProvisioningStatesEnumValues returns the possible values for the AzureLargeInstanceProvisioningStatesEnum const type.
func PossibleAzureLargeInstanceProvisioningStatesEnumValues() []AzureLargeInstanceProvisioningStatesEnum {
	return []AzureLargeInstanceProvisioningStatesEnum{
		AzureLargeInstanceProvisioningStatesEnumAccepted,
		AzureLargeInstanceProvisioningStatesEnumCanceled,
		AzureLargeInstanceProvisioningStatesEnumCreating,
		AzureLargeInstanceProvisioningStatesEnumDeleting,
		AzureLargeInstanceProvisioningStatesEnumFailed,
		AzureLargeInstanceProvisioningStatesEnumMigrating,
		AzureLargeInstanceProvisioningStatesEnumSucceeded,
		AzureLargeInstanceProvisioningStatesEnumUpdating,
	}
}

// AzureLargeInstanceSizeNamesEnum - Enum of available model types (each of which have their own storage / memory sizes) for
// an Azure Large Instance type. See https://docs.microsoft.com/azure/sap/large-instances/hana-available-skus
type AzureLargeInstanceSizeNamesEnum string

const (
	// AzureLargeInstanceSizeNamesEnumS112 - No longer offered or used.
	AzureLargeInstanceSizeNamesEnumS112 AzureLargeInstanceSizeNamesEnum = "S112"
	// AzureLargeInstanceSizeNamesEnumS144 - Type I class SKU that can't be purchased anymore
	AzureLargeInstanceSizeNamesEnumS144 AzureLargeInstanceSizeNamesEnum = "S144"
	// AzureLargeInstanceSizeNamesEnumS144M - Type I class SKU that can't be purchased anymore
	AzureLargeInstanceSizeNamesEnumS144M AzureLargeInstanceSizeNamesEnum = "S144m"
	// AzureLargeInstanceSizeNamesEnumS192 - Type I class SKU that can't be purchased anymore
	AzureLargeInstanceSizeNamesEnumS192 AzureLargeInstanceSizeNamesEnum = "S192"
	// AzureLargeInstanceSizeNamesEnumS192M - Type I class SKU that can't be purchased anymore
	AzureLargeInstanceSizeNamesEnumS192M AzureLargeInstanceSizeNamesEnum = "S192m"
	// AzureLargeInstanceSizeNamesEnumS192Xm - Type I class SKU that can't be purchased anymore
	AzureLargeInstanceSizeNamesEnumS192Xm AzureLargeInstanceSizeNamesEnum = "S192xm"
	// AzureLargeInstanceSizeNamesEnumS224 - 4 sockets, 224 CPU threads, 112 CPU cores, 3 TB total memory, 3 TB DRAM, 6.3 TB storage,
	// Cisco_UCS hardware type
	AzureLargeInstanceSizeNamesEnumS224 AzureLargeInstanceSizeNamesEnum = "S224"
	// AzureLargeInstanceSizeNamesEnumS224M - 4 sockets, 224 CPU threads, 112 CPU cores, 6 TB total memory, 6 TB DRAM, 10.5 TB
	// storage, Cisco_UCS hardware type
	AzureLargeInstanceSizeNamesEnumS224M AzureLargeInstanceSizeNamesEnum = "S224m"
	// AzureLargeInstanceSizeNamesEnumS224Om - 4 sockets, 224 CPU threads, 112 CPU cores, 6 TB total memory, 3 TB DRAM, 3 TB memory
	// optane, 10.5 TB storage, Cisco_UCS hardware type
	AzureLargeInstanceSizeNamesEnumS224Om AzureLargeInstanceSizeNamesEnum = "S224om"
	// AzureLargeInstanceSizeNamesEnumS224Oo - 4 sockets, 224 CPU threads, 112 CPU cores, 4.5 TB total memory, 1.5 TB DRAM, 3
	// TB memory optane, 8.4 TB storage, Cisco_UCS hardware type
	AzureLargeInstanceSizeNamesEnumS224Oo AzureLargeInstanceSizeNamesEnum = "S224oo"
	// AzureLargeInstanceSizeNamesEnumS224Oom - 4 sockets, 224 CPU threads, 112 CPU cores, 9 TB total memory, 3 TB DRAM, 6 TB
	// memory optane, 14.8 TB storage, Cisco_UCS hardware type
	AzureLargeInstanceSizeNamesEnumS224Oom AzureLargeInstanceSizeNamesEnum = "S224oom"
	// AzureLargeInstanceSizeNamesEnumS224Ooo - 4 sockets, 224 CPU threads, 112 CPU cores, 7.5TB total memory, 1.5 TB DRAM, 6
	// TB memory optane, 12.7 TB storage, Cisco_UCS hardware type
	AzureLargeInstanceSizeNamesEnumS224Ooo AzureLargeInstanceSizeNamesEnum = "S224ooo"
	// AzureLargeInstanceSizeNamesEnumS224Se - 4 sockets, 448 CPU threads, 6 TB total memory, SDFLEX hardware type
	AzureLargeInstanceSizeNamesEnumS224Se AzureLargeInstanceSizeNamesEnum = "S224se"
	// AzureLargeInstanceSizeNamesEnumS384 - 8 sockets, 384 CPU threads, 192 CPU cores, 4 TB total memory, 4 TB DRAM, 16 TB storage,
	// HPEMc990x hardware type
	AzureLargeInstanceSizeNamesEnumS384 AzureLargeInstanceSizeNamesEnum = "S384"
	// AzureLargeInstanceSizeNamesEnumS384M - 8 sockets, 384 CPU threads, 192 CPU cores, 6 TB total memory, 6 TB DRAM, 18 TB storage,
	// HPEMc990x hardware type
	AzureLargeInstanceSizeNamesEnumS384M AzureLargeInstanceSizeNamesEnum = "S384m"
	// AzureLargeInstanceSizeNamesEnumS384Xm - 8 sockets, 384 CPU threads, 192 CPU cores, 8 TB total memory, 8 TB DRAM, 22 TB
	// storage, HPEMc990x hardware type
	AzureLargeInstanceSizeNamesEnumS384Xm AzureLargeInstanceSizeNamesEnum = "S384xm"
	// AzureLargeInstanceSizeNamesEnumS384Xxm - 8 sockets, 384 CPU threads, 12 TB total memory, HPEMc990x hardware type
	AzureLargeInstanceSizeNamesEnumS384Xxm AzureLargeInstanceSizeNamesEnum = "S384xxm"
	// AzureLargeInstanceSizeNamesEnumS448 - 8 sockets, 448 CPU threads, 224 CPU cores, 6 TB total memory, 6 TB DRAM, 10.5 TB
	// storage, SDFLEX hardware type
	AzureLargeInstanceSizeNamesEnumS448 AzureLargeInstanceSizeNamesEnum = "S448"
	// AzureLargeInstanceSizeNamesEnumS448M - 8 sockets, 448 CPU threads, 224 CPU cores, 12 TB total memory, 12 TB DRAM, 18.9
	// TB storage, SDFLEX hardware type
	AzureLargeInstanceSizeNamesEnumS448M AzureLargeInstanceSizeNamesEnum = "S448m"
	// AzureLargeInstanceSizeNamesEnumS448Om - 8 sockets, 448 CPU threads, 224 CPU cores, 12 TB total memory, 6 TB DRAM, 6 TB
	// memory optane, 18.9 TB storage, SDFLEX hardware type
	AzureLargeInstanceSizeNamesEnumS448Om AzureLargeInstanceSizeNamesEnum = "S448om"
	// AzureLargeInstanceSizeNamesEnumS448Oo - 8 sockets, 448 CPU threads, 224 CPU cores, 9 TB total memory, 3 TB DRAM, 6 TB memory
	// optane, 14.8 TB storage, SDFLEX hardware type
	AzureLargeInstanceSizeNamesEnumS448Oo AzureLargeInstanceSizeNamesEnum = "S448oo"
	// AzureLargeInstanceSizeNamesEnumS448Oom - 8 sockets, 448 CPU threads, 224 CPU cores, 18 TB total memory, 6 TB DRAM, 12 memory
	// optane, 27.4 TB storage, SDFLEX hardware type
	AzureLargeInstanceSizeNamesEnumS448Oom AzureLargeInstanceSizeNamesEnum = "S448oom"
	// AzureLargeInstanceSizeNamesEnumS448Ooo - 8 sockets, 448 CPU threads, 224 CPU cores, 15 TB total memory, 3 TB DRAM, 12 memory
	// optane, 23.2 TB storage, SDFLEX hardware type
	AzureLargeInstanceSizeNamesEnumS448Ooo AzureLargeInstanceSizeNamesEnum = "S448ooo"
	// AzureLargeInstanceSizeNamesEnumS448Se - 8 sockets, 448 CPU threads, 12 TB total memory, SDFLEX hardware type
	AzureLargeInstanceSizeNamesEnumS448Se AzureLargeInstanceSizeNamesEnum = "S448se"
	// AzureLargeInstanceSizeNamesEnumS576M - 12 sockets, 576 CPU threads, 288 CPU cores, 12 TB total memory, 12 TB DRAM, 28 TB
	// storage, HPEMc990x hardware type
	AzureLargeInstanceSizeNamesEnumS576M AzureLargeInstanceSizeNamesEnum = "S576m"
	// AzureLargeInstanceSizeNamesEnumS576Xm - 12 sockets, 576 CPU threads, 288 CPU cores, 18 TB total memory, HPEMc990x hardware
	// type
	AzureLargeInstanceSizeNamesEnumS576Xm AzureLargeInstanceSizeNamesEnum = "S576xm"
	// AzureLargeInstanceSizeNamesEnumS672 - 12 sockets, 672 CPU threads, 336 CPU cores, 9 TB total memory, 9 TB DRAM, 14.7 TB
	// storage, SDFLEX hardware type
	AzureLargeInstanceSizeNamesEnumS672 AzureLargeInstanceSizeNamesEnum = "S672"
	// AzureLargeInstanceSizeNamesEnumS672M - 12 sockets, 672 CPU threads, 336 CPU cores, 18 TB total memory, 18 TB DRAM, 27.4
	// TB storage, SDFLEX hardware type
	AzureLargeInstanceSizeNamesEnumS672M AzureLargeInstanceSizeNamesEnum = "S672m"
	// AzureLargeInstanceSizeNamesEnumS672Om - 12 sockets, 672 CPU threads, 336 CPU cores, 18 TB total memory, 9 TB DRAM, 9 TB
	// memory optane, 27.4 TB storage, SDFLEX hardware type
	AzureLargeInstanceSizeNamesEnumS672Om AzureLargeInstanceSizeNamesEnum = "S672om"
	// AzureLargeInstanceSizeNamesEnumS672Oo - 12 sockets, 672 CPU threads, 336 CPU cores, 13.5 TB total memory, 4.5 TB DRAM,
	// 9 TB memory optane, 21.1 TB storage, SDFLEX hardware type
	AzureLargeInstanceSizeNamesEnumS672Oo AzureLargeInstanceSizeNamesEnum = "S672oo"
	// AzureLargeInstanceSizeNamesEnumS672Oom - 12 sockets, 672 CPU threads, 336 CPU cores, 27 TB total memory, 9 TB DRAM, 18
	// TB memory optane, 40 TB storage, SDFLEX hardware type
	AzureLargeInstanceSizeNamesEnumS672Oom AzureLargeInstanceSizeNamesEnum = "S672oom"
	// AzureLargeInstanceSizeNamesEnumS672Ooo - 12 sockets, 672 CPU threads, 336 CPU cores, 22.5 TB total memory, 4.5 TB DRAM,
	// 18 TB memory optane, 33.7 TB storage, SDFLEX hardware type
	AzureLargeInstanceSizeNamesEnumS672Ooo AzureLargeInstanceSizeNamesEnum = "S672ooo"
	// AzureLargeInstanceSizeNamesEnumS72 - Type I class SKU that can't be purchased anymore
	AzureLargeInstanceSizeNamesEnumS72 AzureLargeInstanceSizeNamesEnum = "S72"
	// AzureLargeInstanceSizeNamesEnumS72M - Type I class SKU that can't be purchased anymore
	AzureLargeInstanceSizeNamesEnumS72M AzureLargeInstanceSizeNamesEnum = "S72m"
	// AzureLargeInstanceSizeNamesEnumS768 - No longer offered or used.
	AzureLargeInstanceSizeNamesEnumS768 AzureLargeInstanceSizeNamesEnum = "S768"
	// AzureLargeInstanceSizeNamesEnumS768M - 16 sockets, 768 CPU threads, 384 CPU cores, 16 TB total memory, 16 TB DRAM, 36 TB
	// storage, HPEMc990x hardware type
	AzureLargeInstanceSizeNamesEnumS768M AzureLargeInstanceSizeNamesEnum = "S768m"
	// AzureLargeInstanceSizeNamesEnumS768Xm - 16 sockets, 768 CPU threads, 384 CPU cores, 24 TB total memory, 24 TB DRAM, 56
	// TB storage, HPEMc990x hardware type
	AzureLargeInstanceSizeNamesEnumS768Xm AzureLargeInstanceSizeNamesEnum = "S768xm"
	// AzureLargeInstanceSizeNamesEnumS896 - 16 sockets, 896 CPU threads, 448 CPU cores, 12 TB total memory, 12 TB DRAM, 18.9
	// TB storage, SDFLEX hardware type
	AzureLargeInstanceSizeNamesEnumS896 AzureLargeInstanceSizeNamesEnum = "S896"
	// AzureLargeInstanceSizeNamesEnumS896M - 16 sockets, 896 CPU threads, 448 CPU cores, 24 TB total memory, 24 TB DRAM, 35.8
	// TB storage, SDFLEX hardware type
	AzureLargeInstanceSizeNamesEnumS896M AzureLargeInstanceSizeNamesEnum = "S896m"
	// AzureLargeInstanceSizeNamesEnumS896Om - 16 sockets, 896 CPU threads, 448 CPU cores, 24 TB total memory, 12 TB DRAM, 12
	// TB memory optane, 35.8 TB storage, SDFLEX hardware type
	AzureLargeInstanceSizeNamesEnumS896Om AzureLargeInstanceSizeNamesEnum = "S896om"
	// AzureLargeInstanceSizeNamesEnumS896Oo - 16 sockets, 896 CPU threads, 448 CPU cores, 18 TB total memory, 6 TB DRAM, 12 TB
	// memory optane, 27.4 TB storage, SDFLEX hardware type
	AzureLargeInstanceSizeNamesEnumS896Oo AzureLargeInstanceSizeNamesEnum = "S896oo"
	// AzureLargeInstanceSizeNamesEnumS896Oom - 16 sockets, 896 CPU threads, 448 CPU cores, 36 TB total memory, 12 TB DRAM, 24
	// TB memory optane, 52.7 TB storage, SDFLEX hardware type
	AzureLargeInstanceSizeNamesEnumS896Oom AzureLargeInstanceSizeNamesEnum = "S896oom"
	// AzureLargeInstanceSizeNamesEnumS896Ooo - 16 sockets, 896 CPU threads, 448 CPU cores, 30 TB total memory, 6 TB DRAM, 24
	// TB memory optane, 44.3 TB storage, SDFLEX hardware type
	AzureLargeInstanceSizeNamesEnumS896Ooo AzureLargeInstanceSizeNamesEnum = "S896ooo"
	// AzureLargeInstanceSizeNamesEnumS96 - 2 sockets, 96 CPU threads, 48 CPU cores, 768 GB total memory, 768 GB DRAM, 3.0 TB
	// storage, Cisco_UCS hardware type
	AzureLargeInstanceSizeNamesEnumS96 AzureLargeInstanceSizeNamesEnum = "S96"
	// AzureLargeInstanceSizeNamesEnumS960M - 20 sockets, 960 CPU threads, 480 CPU cores, 20 TB total memory, 20 TB DRAM, 46 TB
	// storage, HPEMc990x hardware type
	AzureLargeInstanceSizeNamesEnumS960M AzureLargeInstanceSizeNamesEnum = "S960m"
)

// PossibleAzureLargeInstanceSizeNamesEnumValues returns the possible values for the AzureLargeInstanceSizeNamesEnum const type.
func PossibleAzureLargeInstanceSizeNamesEnumValues() []AzureLargeInstanceSizeNamesEnum {
	return []AzureLargeInstanceSizeNamesEnum{
		AzureLargeInstanceSizeNamesEnumS112,
		AzureLargeInstanceSizeNamesEnumS144,
		AzureLargeInstanceSizeNamesEnumS144M,
		AzureLargeInstanceSizeNamesEnumS192,
		AzureLargeInstanceSizeNamesEnumS192M,
		AzureLargeInstanceSizeNamesEnumS192Xm,
		AzureLargeInstanceSizeNamesEnumS224,
		AzureLargeInstanceSizeNamesEnumS224M,
		AzureLargeInstanceSizeNamesEnumS224Om,
		AzureLargeInstanceSizeNamesEnumS224Oo,
		AzureLargeInstanceSizeNamesEnumS224Oom,
		AzureLargeInstanceSizeNamesEnumS224Ooo,
		AzureLargeInstanceSizeNamesEnumS224Se,
		AzureLargeInstanceSizeNamesEnumS384,
		AzureLargeInstanceSizeNamesEnumS384M,
		AzureLargeInstanceSizeNamesEnumS384Xm,
		AzureLargeInstanceSizeNamesEnumS384Xxm,
		AzureLargeInstanceSizeNamesEnumS448,
		AzureLargeInstanceSizeNamesEnumS448M,
		AzureLargeInstanceSizeNamesEnumS448Om,
		AzureLargeInstanceSizeNamesEnumS448Oo,
		AzureLargeInstanceSizeNamesEnumS448Oom,
		AzureLargeInstanceSizeNamesEnumS448Ooo,
		AzureLargeInstanceSizeNamesEnumS448Se,
		AzureLargeInstanceSizeNamesEnumS576M,
		AzureLargeInstanceSizeNamesEnumS576Xm,
		AzureLargeInstanceSizeNamesEnumS672,
		AzureLargeInstanceSizeNamesEnumS672M,
		AzureLargeInstanceSizeNamesEnumS672Om,
		AzureLargeInstanceSizeNamesEnumS672Oo,
		AzureLargeInstanceSizeNamesEnumS672Oom,
		AzureLargeInstanceSizeNamesEnumS672Ooo,
		AzureLargeInstanceSizeNamesEnumS72,
		AzureLargeInstanceSizeNamesEnumS72M,
		AzureLargeInstanceSizeNamesEnumS768,
		AzureLargeInstanceSizeNamesEnumS768M,
		AzureLargeInstanceSizeNamesEnumS768Xm,
		AzureLargeInstanceSizeNamesEnumS896,
		AzureLargeInstanceSizeNamesEnumS896M,
		AzureLargeInstanceSizeNamesEnumS896Om,
		AzureLargeInstanceSizeNamesEnumS896Oo,
		AzureLargeInstanceSizeNamesEnumS896Oom,
		AzureLargeInstanceSizeNamesEnumS896Ooo,
		AzureLargeInstanceSizeNamesEnumS96,
		AzureLargeInstanceSizeNamesEnumS960M,
	}
}

// CreatedByType - The kind of entity that created the resource.
type CreatedByType string

const (
	// CreatedByTypeApplication - The entity was created by an application.
	CreatedByTypeApplication CreatedByType = "Application"
	// CreatedByTypeKey - The entity was created by a key.
	CreatedByTypeKey CreatedByType = "Key"
	// CreatedByTypeManagedIdentity - The entity was created by a managed identity.
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	// CreatedByTypeUser - The entity was created by a user.
	CreatedByTypeUser CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	// OriginSystem - Indicates the operation is initiated by a system.
	OriginSystem Origin = "system"
	// OriginUser - Indicates the operation is initiated by a user.
	OriginUser Origin = "user"
	// OriginUserSystem - Indicates the operation is initiated by a user or system.
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// ProvisioningState - An enum of possible operation states for an AzureLargeStorageInstances
type ProvisioningState string

const (
	// ProvisioningStateAccepted - Accepted means ARM resource has been accepted.
	ProvisioningStateAccepted ProvisioningState = "Accepted"
	// ProvisioningStateCanceled - Cancelled means resource operation has been cancelled
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateCreating - Creating means ARM resource is being created.
	ProvisioningStateCreating ProvisioningState = "Creating"
	// ProvisioningStateDeleting - Deleting means resource is in the process of being deleted
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed - Failed means resource is in failed state
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateMigrating - Migrating means resource is being migrated from one subscription or resource group to another
	ProvisioningStateMigrating ProvisioningState = "Migrating"
	// ProvisioningStateSucceeded - Succeeded means resource creation succeeded during last create/update
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUpdating - Updating means an existing ARM resource is being updated
	ProvisioningStateUpdating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateMigrating,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// ResourceProvisioningState - The provisioning state of a resource type.
type ResourceProvisioningState string

const (
	// ResourceProvisioningStateCanceled - Resource creation was canceled.
	ResourceProvisioningStateCanceled ResourceProvisioningState = "Canceled"
	// ResourceProvisioningStateFailed - Resource creation failed.
	ResourceProvisioningStateFailed ResourceProvisioningState = "Failed"
	// ResourceProvisioningStateSucceeded - Resource has been created.
	ResourceProvisioningStateSucceeded ResourceProvisioningState = "Succeeded"
)

// PossibleResourceProvisioningStateValues returns the possible values for the ResourceProvisioningState const type.
func PossibleResourceProvisioningStateValues() []ResourceProvisioningState {
	return []ResourceProvisioningState{
		ResourceProvisioningStateCanceled,
		ResourceProvisioningStateFailed,
		ResourceProvisioningStateSucceeded,
	}
}
