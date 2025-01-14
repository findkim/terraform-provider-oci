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

// AutonomousDatabase An Oracle Autonomous Database.
type AutonomousDatabase struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Autonomous Database.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of the database.
	LifecycleState AutonomousDatabaseLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

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
	LicenseModel AutonomousDatabaseLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`

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
	DbWorkload AutonomousDatabaseDbWorkloadEnum `mandatory:"false" json:"dbWorkload,omitempty"`

	// The client IP access control list (ACL). This feature is available for serverless deployments (https://docs.cloud.oracle.com/Content/Database/Concepts/adboverview.htm#AEI) only.
	// Only clients connecting from an IP address included in the ACL may access the Autonomous Database instance. This is an array of CIDR (Classless Inter-Domain Routing) notations for a subnet.
	WhitelistedIps []string `mandatory:"false" json:"whitelistedIps"`

	// Indicates if auto scaling is enabled for the Autonomous Database CPU core count. Note that auto scaling is available for serverless deployments (https://docs.cloud.oracle.com/Content/Database/Concepts/adboverview.htm#AEI) only.
	IsAutoScalingEnabled *bool `mandatory:"false" json:"isAutoScalingEnabled"`

	// Status of the Data Safe registration for this Autonomous Database.
	DataSafeStatus AutonomousDatabaseDataSafeStatusEnum `mandatory:"false" json:"dataSafeStatus,omitempty"`
}

func (m AutonomousDatabase) String() string {
	return common.PointerString(m)
}

// AutonomousDatabaseLifecycleStateEnum Enum with underlying type: string
type AutonomousDatabaseLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousDatabaseLifecycleStateEnum
const (
	AutonomousDatabaseLifecycleStateProvisioning            AutonomousDatabaseLifecycleStateEnum = "PROVISIONING"
	AutonomousDatabaseLifecycleStateAvailable               AutonomousDatabaseLifecycleStateEnum = "AVAILABLE"
	AutonomousDatabaseLifecycleStateStopping                AutonomousDatabaseLifecycleStateEnum = "STOPPING"
	AutonomousDatabaseLifecycleStateStopped                 AutonomousDatabaseLifecycleStateEnum = "STOPPED"
	AutonomousDatabaseLifecycleStateStarting                AutonomousDatabaseLifecycleStateEnum = "STARTING"
	AutonomousDatabaseLifecycleStateTerminating             AutonomousDatabaseLifecycleStateEnum = "TERMINATING"
	AutonomousDatabaseLifecycleStateTerminated              AutonomousDatabaseLifecycleStateEnum = "TERMINATED"
	AutonomousDatabaseLifecycleStateUnavailable             AutonomousDatabaseLifecycleStateEnum = "UNAVAILABLE"
	AutonomousDatabaseLifecycleStateRestoreInProgress       AutonomousDatabaseLifecycleStateEnum = "RESTORE_IN_PROGRESS"
	AutonomousDatabaseLifecycleStateRestoreFailed           AutonomousDatabaseLifecycleStateEnum = "RESTORE_FAILED"
	AutonomousDatabaseLifecycleStateBackupInProgress        AutonomousDatabaseLifecycleStateEnum = "BACKUP_IN_PROGRESS"
	AutonomousDatabaseLifecycleStateScaleInProgress         AutonomousDatabaseLifecycleStateEnum = "SCALE_IN_PROGRESS"
	AutonomousDatabaseLifecycleStateAvailableNeedsAttention AutonomousDatabaseLifecycleStateEnum = "AVAILABLE_NEEDS_ATTENTION"
	AutonomousDatabaseLifecycleStateUpdating                AutonomousDatabaseLifecycleStateEnum = "UPDATING"
	AutonomousDatabaseLifecycleStateMaintenanceInProgress   AutonomousDatabaseLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
)

var mappingAutonomousDatabaseLifecycleState = map[string]AutonomousDatabaseLifecycleStateEnum{
	"PROVISIONING":              AutonomousDatabaseLifecycleStateProvisioning,
	"AVAILABLE":                 AutonomousDatabaseLifecycleStateAvailable,
	"STOPPING":                  AutonomousDatabaseLifecycleStateStopping,
	"STOPPED":                   AutonomousDatabaseLifecycleStateStopped,
	"STARTING":                  AutonomousDatabaseLifecycleStateStarting,
	"TERMINATING":               AutonomousDatabaseLifecycleStateTerminating,
	"TERMINATED":                AutonomousDatabaseLifecycleStateTerminated,
	"UNAVAILABLE":               AutonomousDatabaseLifecycleStateUnavailable,
	"RESTORE_IN_PROGRESS":       AutonomousDatabaseLifecycleStateRestoreInProgress,
	"RESTORE_FAILED":            AutonomousDatabaseLifecycleStateRestoreFailed,
	"BACKUP_IN_PROGRESS":        AutonomousDatabaseLifecycleStateBackupInProgress,
	"SCALE_IN_PROGRESS":         AutonomousDatabaseLifecycleStateScaleInProgress,
	"AVAILABLE_NEEDS_ATTENTION": AutonomousDatabaseLifecycleStateAvailableNeedsAttention,
	"UPDATING":                  AutonomousDatabaseLifecycleStateUpdating,
	"MAINTENANCE_IN_PROGRESS":   AutonomousDatabaseLifecycleStateMaintenanceInProgress,
}

// GetAutonomousDatabaseLifecycleStateEnumValues Enumerates the set of values for AutonomousDatabaseLifecycleStateEnum
func GetAutonomousDatabaseLifecycleStateEnumValues() []AutonomousDatabaseLifecycleStateEnum {
	values := make([]AutonomousDatabaseLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousDatabaseLifecycleState {
		values = append(values, v)
	}
	return values
}

// AutonomousDatabaseLicenseModelEnum Enum with underlying type: string
type AutonomousDatabaseLicenseModelEnum string

// Set of constants representing the allowable values for AutonomousDatabaseLicenseModelEnum
const (
	AutonomousDatabaseLicenseModelLicenseIncluded     AutonomousDatabaseLicenseModelEnum = "LICENSE_INCLUDED"
	AutonomousDatabaseLicenseModelBringYourOwnLicense AutonomousDatabaseLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingAutonomousDatabaseLicenseModel = map[string]AutonomousDatabaseLicenseModelEnum{
	"LICENSE_INCLUDED":       AutonomousDatabaseLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": AutonomousDatabaseLicenseModelBringYourOwnLicense,
}

// GetAutonomousDatabaseLicenseModelEnumValues Enumerates the set of values for AutonomousDatabaseLicenseModelEnum
func GetAutonomousDatabaseLicenseModelEnumValues() []AutonomousDatabaseLicenseModelEnum {
	values := make([]AutonomousDatabaseLicenseModelEnum, 0)
	for _, v := range mappingAutonomousDatabaseLicenseModel {
		values = append(values, v)
	}
	return values
}

// AutonomousDatabaseDbWorkloadEnum Enum with underlying type: string
type AutonomousDatabaseDbWorkloadEnum string

// Set of constants representing the allowable values for AutonomousDatabaseDbWorkloadEnum
const (
	AutonomousDatabaseDbWorkloadOltp AutonomousDatabaseDbWorkloadEnum = "OLTP"
	AutonomousDatabaseDbWorkloadDw   AutonomousDatabaseDbWorkloadEnum = "DW"
)

var mappingAutonomousDatabaseDbWorkload = map[string]AutonomousDatabaseDbWorkloadEnum{
	"OLTP": AutonomousDatabaseDbWorkloadOltp,
	"DW":   AutonomousDatabaseDbWorkloadDw,
}

// GetAutonomousDatabaseDbWorkloadEnumValues Enumerates the set of values for AutonomousDatabaseDbWorkloadEnum
func GetAutonomousDatabaseDbWorkloadEnumValues() []AutonomousDatabaseDbWorkloadEnum {
	values := make([]AutonomousDatabaseDbWorkloadEnum, 0)
	for _, v := range mappingAutonomousDatabaseDbWorkload {
		values = append(values, v)
	}
	return values
}

// AutonomousDatabaseDataSafeStatusEnum Enum with underlying type: string
type AutonomousDatabaseDataSafeStatusEnum string

// Set of constants representing the allowable values for AutonomousDatabaseDataSafeStatusEnum
const (
	AutonomousDatabaseDataSafeStatusRegistering   AutonomousDatabaseDataSafeStatusEnum = "REGISTERING"
	AutonomousDatabaseDataSafeStatusRegistered    AutonomousDatabaseDataSafeStatusEnum = "REGISTERED"
	AutonomousDatabaseDataSafeStatusDeregistering AutonomousDatabaseDataSafeStatusEnum = "DEREGISTERING"
	AutonomousDatabaseDataSafeStatusNotRegistered AutonomousDatabaseDataSafeStatusEnum = "NOT_REGISTERED"
	AutonomousDatabaseDataSafeStatusFailed        AutonomousDatabaseDataSafeStatusEnum = "FAILED"
)

var mappingAutonomousDatabaseDataSafeStatus = map[string]AutonomousDatabaseDataSafeStatusEnum{
	"REGISTERING":    AutonomousDatabaseDataSafeStatusRegistering,
	"REGISTERED":     AutonomousDatabaseDataSafeStatusRegistered,
	"DEREGISTERING":  AutonomousDatabaseDataSafeStatusDeregistering,
	"NOT_REGISTERED": AutonomousDatabaseDataSafeStatusNotRegistered,
	"FAILED":         AutonomousDatabaseDataSafeStatusFailed,
}

// GetAutonomousDatabaseDataSafeStatusEnumValues Enumerates the set of values for AutonomousDatabaseDataSafeStatusEnum
func GetAutonomousDatabaseDataSafeStatusEnumValues() []AutonomousDatabaseDataSafeStatusEnum {
	values := make([]AutonomousDatabaseDataSafeStatusEnum, 0)
	for _, v := range mappingAutonomousDatabaseDataSafeStatus {
		values = append(values, v)
	}
	return values
}
