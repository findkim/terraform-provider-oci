// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service.
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AutonomousDatabaseSummary An Oracle Autonomous Database.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type AutonomousDatabaseSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Autonomous Database.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of the database.
	LifecycleState AutonomousDatabaseSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The database name.
	DbName *string `mandatory:"true" json:"dbName"`

	// The number of CPU cores to be made available to the database.
	CpuCoreCount *int `mandatory:"true" json:"cpuCoreCount"`

	// The quantity of data in the database, in terabytes.
	DataStorageSizeInTBs *int `mandatory:"true" json:"dataStorageSizeInTBs"`

	// Information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Indicates if this is an Always Free resource. The default value is false. Note that Always Free Autonomous Databases have 1 CPU and 20GB memory. For Always Free databases, memory and CPU cannot be scaled.
	IsFreeTier *bool `mandatory:"false" json:"isFreeTier"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The date and time the Always Free database will be stopped because of inactivity. If this time is reached without any database activity, the database will automatically be put into the STOPPED state.
	TimeReclamationOfFreeAutonomousDatabase *common.SDKTime `mandatory:"false" json:"timeReclamationOfFreeAutonomousDatabase"`

	// The date and time the Always Free database will be automatically deleted because of inactivity. If the database is in the STOPPED state and without activity until this time, it will be deleted.
	TimeDeletionOfFreeAutonomousDatabase *common.SDKTime `mandatory:"false" json:"timeDeletionOfFreeAutonomousDatabase"`

	// True if the database uses the dedicated deployment (https://docs.cloud.oracle.com/Content/Database/Concepts/adbddoverview.htm) option.
	IsDedicated *bool `mandatory:"false" json:"isDedicated"`

	// The Autonomous Container Database OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	AutonomousContainerDatabaseId *string `mandatory:"false" json:"autonomousContainerDatabaseId"`

	// The date and time the database was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The user-friendly name for the Autonomous Database. The name does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The URL of the Service Console for the Autonomous Database.
	ServiceConsoleUrl *string `mandatory:"false" json:"serviceConsoleUrl"`

	// The connection string used to connect to the Autonomous Database. The username for the Service Console is ADMIN. Use the password you entered when creating the Autonomous Database for the password value.
	ConnectionStrings *AutonomousDatabaseConnectionStrings `mandatory:"false" json:"connectionStrings"`

	ConnectionUrls *AutonomousDatabaseConnectionUrls `mandatory:"false" json:"connectionUrls"`

	// The Oracle license model that applies to the Oracle Autonomous Database. The default for Autonomous Database using the shared deployment is BRING_YOUR_OWN_LICENSE. Note that when provisioning an Autonomous Database using the dedicated deployment (https://docs.cloud.oracle.com/Content/Database/Concepts/adbddoverview.htm) option, this attribute must be null because the attribute is already set on Autonomous Exadata Infrastructure level.
	LicenseModel AutonomousDatabaseSummaryLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

	// The amount of storage that has been used, in terabytes.
	UsedDataStorageSizeInTBs *int `mandatory:"false" json:"usedDataStorageSizeInTBs"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A valid Oracle Database version for Autonomous Database.
	DbVersion *string `mandatory:"false" json:"dbVersion"`

	// Indicates if the Autonomous Database version is a preview version.
	IsPreview *bool `mandatory:"false" json:"isPreview"`

	// The Autonomous Database workload type. OLTP indicates an Autonomous Transaction Processing database and DW indicates an Autonomous Data Warehouse database.
	DbWorkload AutonomousDatabaseSummaryDbWorkloadEnum `mandatory:"false" json:"dbWorkload,omitempty"`

	// The client IP access control list (ACL). This feature is available for serverless deployments (https://docs.cloud.oracle.com/Content/Database/Concepts/adboverview.htm#AEI) only.
	// Only clients connecting from an IP address included in the ACL may access the Autonomous Database instance. This is an array of CIDR (Classless Inter-Domain Routing) notations for a subnet.
	WhitelistedIps []string `mandatory:"false" json:"whitelistedIps"`

	// Indicates if auto scaling is enabled for the Autonomous Database CPU core count. Note that auto scaling is available for serverless deployments (https://docs.cloud.oracle.com/Content/Database/Concepts/adboverview.htm#AEI) only.
	IsAutoScalingEnabled *bool `mandatory:"false" json:"isAutoScalingEnabled"`

	// Status of the Data Safe registration for this Autonomous Database.
	DataSafeStatus AutonomousDatabaseSummaryDataSafeStatusEnum `mandatory:"false" json:"dataSafeStatus,omitempty"`
}

func (m AutonomousDatabaseSummary) String() string {
	return common.PointerString(m)
}

// AutonomousDatabaseSummaryLifecycleStateEnum Enum with underlying type: string
type AutonomousDatabaseSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryLifecycleStateEnum
const (
	AutonomousDatabaseSummaryLifecycleStateProvisioning            AutonomousDatabaseSummaryLifecycleStateEnum = "PROVISIONING"
	AutonomousDatabaseSummaryLifecycleStateAvailable               AutonomousDatabaseSummaryLifecycleStateEnum = "AVAILABLE"
	AutonomousDatabaseSummaryLifecycleStateStopping                AutonomousDatabaseSummaryLifecycleStateEnum = "STOPPING"
	AutonomousDatabaseSummaryLifecycleStateStopped                 AutonomousDatabaseSummaryLifecycleStateEnum = "STOPPED"
	AutonomousDatabaseSummaryLifecycleStateStarting                AutonomousDatabaseSummaryLifecycleStateEnum = "STARTING"
	AutonomousDatabaseSummaryLifecycleStateTerminating             AutonomousDatabaseSummaryLifecycleStateEnum = "TERMINATING"
	AutonomousDatabaseSummaryLifecycleStateTerminated              AutonomousDatabaseSummaryLifecycleStateEnum = "TERMINATED"
	AutonomousDatabaseSummaryLifecycleStateUnavailable             AutonomousDatabaseSummaryLifecycleStateEnum = "UNAVAILABLE"
	AutonomousDatabaseSummaryLifecycleStateRestoreInProgress       AutonomousDatabaseSummaryLifecycleStateEnum = "RESTORE_IN_PROGRESS"
	AutonomousDatabaseSummaryLifecycleStateRestoreFailed           AutonomousDatabaseSummaryLifecycleStateEnum = "RESTORE_FAILED"
	AutonomousDatabaseSummaryLifecycleStateBackupInProgress        AutonomousDatabaseSummaryLifecycleStateEnum = "BACKUP_IN_PROGRESS"
	AutonomousDatabaseSummaryLifecycleStateScaleInProgress         AutonomousDatabaseSummaryLifecycleStateEnum = "SCALE_IN_PROGRESS"
	AutonomousDatabaseSummaryLifecycleStateAvailableNeedsAttention AutonomousDatabaseSummaryLifecycleStateEnum = "AVAILABLE_NEEDS_ATTENTION"
	AutonomousDatabaseSummaryLifecycleStateUpdating                AutonomousDatabaseSummaryLifecycleStateEnum = "UPDATING"
	AutonomousDatabaseSummaryLifecycleStateMaintenanceInProgress   AutonomousDatabaseSummaryLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
)

var mappingAutonomousDatabaseSummaryLifecycleState = map[string]AutonomousDatabaseSummaryLifecycleStateEnum{
	"PROVISIONING":              AutonomousDatabaseSummaryLifecycleStateProvisioning,
	"AVAILABLE":                 AutonomousDatabaseSummaryLifecycleStateAvailable,
	"STOPPING":                  AutonomousDatabaseSummaryLifecycleStateStopping,
	"STOPPED":                   AutonomousDatabaseSummaryLifecycleStateStopped,
	"STARTING":                  AutonomousDatabaseSummaryLifecycleStateStarting,
	"TERMINATING":               AutonomousDatabaseSummaryLifecycleStateTerminating,
	"TERMINATED":                AutonomousDatabaseSummaryLifecycleStateTerminated,
	"UNAVAILABLE":               AutonomousDatabaseSummaryLifecycleStateUnavailable,
	"RESTORE_IN_PROGRESS":       AutonomousDatabaseSummaryLifecycleStateRestoreInProgress,
	"RESTORE_FAILED":            AutonomousDatabaseSummaryLifecycleStateRestoreFailed,
	"BACKUP_IN_PROGRESS":        AutonomousDatabaseSummaryLifecycleStateBackupInProgress,
	"SCALE_IN_PROGRESS":         AutonomousDatabaseSummaryLifecycleStateScaleInProgress,
	"AVAILABLE_NEEDS_ATTENTION": AutonomousDatabaseSummaryLifecycleStateAvailableNeedsAttention,
	"UPDATING":                  AutonomousDatabaseSummaryLifecycleStateUpdating,
	"MAINTENANCE_IN_PROGRESS":   AutonomousDatabaseSummaryLifecycleStateMaintenanceInProgress,
}

// GetAutonomousDatabaseSummaryLifecycleStateEnumValues Enumerates the set of values for AutonomousDatabaseSummaryLifecycleStateEnum
func GetAutonomousDatabaseSummaryLifecycleStateEnumValues() []AutonomousDatabaseSummaryLifecycleStateEnum {
	values := make([]AutonomousDatabaseSummaryLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}

// AutonomousDatabaseSummaryLicenseModelEnum Enum with underlying type: string
type AutonomousDatabaseSummaryLicenseModelEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryLicenseModelEnum
const (
	AutonomousDatabaseSummaryLicenseModelLicenseIncluded     AutonomousDatabaseSummaryLicenseModelEnum = "LICENSE_INCLUDED"
	AutonomousDatabaseSummaryLicenseModelBringYourOwnLicense AutonomousDatabaseSummaryLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingAutonomousDatabaseSummaryLicenseModel = map[string]AutonomousDatabaseSummaryLicenseModelEnum{
	"LICENSE_INCLUDED":       AutonomousDatabaseSummaryLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": AutonomousDatabaseSummaryLicenseModelBringYourOwnLicense,
}

// GetAutonomousDatabaseSummaryLicenseModelEnumValues Enumerates the set of values for AutonomousDatabaseSummaryLicenseModelEnum
func GetAutonomousDatabaseSummaryLicenseModelEnumValues() []AutonomousDatabaseSummaryLicenseModelEnum {
	values := make([]AutonomousDatabaseSummaryLicenseModelEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryLicenseModel {
		values = append(values, v)
	}
	return values
}

// AutonomousDatabaseSummaryDbWorkloadEnum Enum with underlying type: string
type AutonomousDatabaseSummaryDbWorkloadEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryDbWorkloadEnum
const (
	AutonomousDatabaseSummaryDbWorkloadOltp AutonomousDatabaseSummaryDbWorkloadEnum = "OLTP"
	AutonomousDatabaseSummaryDbWorkloadDw   AutonomousDatabaseSummaryDbWorkloadEnum = "DW"
)

var mappingAutonomousDatabaseSummaryDbWorkload = map[string]AutonomousDatabaseSummaryDbWorkloadEnum{
	"OLTP": AutonomousDatabaseSummaryDbWorkloadOltp,
	"DW":   AutonomousDatabaseSummaryDbWorkloadDw,
}

// GetAutonomousDatabaseSummaryDbWorkloadEnumValues Enumerates the set of values for AutonomousDatabaseSummaryDbWorkloadEnum
func GetAutonomousDatabaseSummaryDbWorkloadEnumValues() []AutonomousDatabaseSummaryDbWorkloadEnum {
	values := make([]AutonomousDatabaseSummaryDbWorkloadEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryDbWorkload {
		values = append(values, v)
	}
	return values
}

// AutonomousDatabaseSummaryDataSafeStatusEnum Enum with underlying type: string
type AutonomousDatabaseSummaryDataSafeStatusEnum string

// Set of constants representing the allowable values for AutonomousDatabaseSummaryDataSafeStatusEnum
const (
	AutonomousDatabaseSummaryDataSafeStatusRegistering   AutonomousDatabaseSummaryDataSafeStatusEnum = "REGISTERING"
	AutonomousDatabaseSummaryDataSafeStatusRegistered    AutonomousDatabaseSummaryDataSafeStatusEnum = "REGISTERED"
	AutonomousDatabaseSummaryDataSafeStatusDeregistering AutonomousDatabaseSummaryDataSafeStatusEnum = "DEREGISTERING"
	AutonomousDatabaseSummaryDataSafeStatusNotRegistered AutonomousDatabaseSummaryDataSafeStatusEnum = "NOT_REGISTERED"
	AutonomousDatabaseSummaryDataSafeStatusFailed        AutonomousDatabaseSummaryDataSafeStatusEnum = "FAILED"
)

var mappingAutonomousDatabaseSummaryDataSafeStatus = map[string]AutonomousDatabaseSummaryDataSafeStatusEnum{
	"REGISTERING":    AutonomousDatabaseSummaryDataSafeStatusRegistering,
	"REGISTERED":     AutonomousDatabaseSummaryDataSafeStatusRegistered,
	"DEREGISTERING":  AutonomousDatabaseSummaryDataSafeStatusDeregistering,
	"NOT_REGISTERED": AutonomousDatabaseSummaryDataSafeStatusNotRegistered,
	"FAILED":         AutonomousDatabaseSummaryDataSafeStatusFailed,
}

// GetAutonomousDatabaseSummaryDataSafeStatusEnumValues Enumerates the set of values for AutonomousDatabaseSummaryDataSafeStatusEnum
func GetAutonomousDatabaseSummaryDataSafeStatusEnumValues() []AutonomousDatabaseSummaryDataSafeStatusEnum {
	values := make([]AutonomousDatabaseSummaryDataSafeStatusEnum, 0)
	for _, v := range mappingAutonomousDatabaseSummaryDataSafeStatus {
		values = append(values, v)
	}
	return values
}
