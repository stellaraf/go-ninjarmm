package ninjarmm

import "encoding/json"

type ninjaRMMBaseError struct {
	Error string `json:"error"`
}

type ninaRMMAPIError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	ErrorCode        int    `json:"error_code,omitempty"`
}

type ninjaRMMRequestError struct {
	ResultCode   string `json:"resultCode"`
	ErrorMessage string `json:"errorMessage"`
	IncidentID   string `json:"incidentId"`
}

type TestData struct {
	OrgID    int `json:"orgId"`
	DeviceID int `json:"deviceId"`
}

type Environment struct {
	ClientID             string
	ClientSecret         string
	BaseURL              string
	EncryptionPassphrase string
	TestData             string
}

const (
	Approval_APPROVED string = "APPROVED"
	Approval_PENDING  string = "PENDING"
)

const (
	NodeClass_CLOUD_MONITOR_TARGET         string = "CLOUD_MONITOR_TARGET"
	NodeClass_LINUX_SERVER                 string = "LINUX_SERVER"
	NodeClass_LINUX_WORKSTATION            string = "LINUX_WORKSTATION"
	NodeClass_MAC                          string = "MAC"
	NodeClass_MAC_SERVER                   string = "MAC_SERVER"
	NodeClass_NMS_APPLIANCE                string = "NMS_APPLIANCE"
	NodeClass_NMS_COMPUTER                 string = "NMS_COMPUTER"
	NodeClass_NMS_DIAL_MANAGER             string = "NMS_DIAL_MANAGER"
	NodeClass_NMS_FIREWALL                 string = "NMS_FIREWALL"
	NodeClass_NMS_IPSLA                    string = "NMS_IPSLA"
	NodeClass_NMS_NETWORK_MANAGEMENT_AGENT string = "NMS_NETWORK_MANAGEMENT_AGENT"
	NodeClass_NMS_OTHER                    string = "NMS_OTHER"
	NodeClass_NMS_PHONE                    string = "NMS_PHONE"
	NodeClass_NMS_PRINTER                  string = "NMS_PRINTER"
	NodeClass_NMS_PRIVATE_NETWORK_GATEWAY  string = "NMS_PRIVATE_NETWORK_GATEWAY"
	NodeClass_NMS_ROUTER                   string = "NMS_ROUTER"
	NodeClass_NMS_SCANNER                  string = "NMS_SCANNER"
	NodeClass_NMS_SERVER                   string = "NMS_SERVER"
	NodeClass_NMS_SWITCH                   string = "NMS_SWITCH"
	NodeClass_NMS_VIRTUAL_MACHINE          string = "NMS_VIRTUAL_MACHINE"
	NodeClass_NMS_VM_HOST                  string = "NMS_VM_HOST"
	NodeClass_NMS_WAP                      string = "NMS_WAP"
	NodeClass_VMWARE_VM_GUEST              string = "VMWARE_VM_GUEST"
	NodeClass_VMWARE_VM_HOST               string = "VMWARE_VM_HOST"
	NodeClass_WINDOWS_SERVER               string = "WINDOWS_SERVER"
	NodeClass_WINDOWS_WORKSTATION          string = "WINDOWS_WORKSTATION"
)

const (
	Severity_NONE     string = "NONE"
	Severity_MINOR    string = "MINOR"
	Severity_MODERATE string = "MODERATE"
	Severity_MAJOR    string = "MAJOR"
	Severity_CRITICAL string = "CRITICAL"
)

const (
	Priority_NONE   string = "NONE"
	Priority_LOW    string = "LOW"
	Priority_MEDIUM string = "MEDIUM"
	Priority_HIGH   string = "HIGH"
)

const (
	ActivityType_ACTION                    string = "ACTION"
	ActivityType_ACTIONSET                 string = "ACTIONSET"
	ActivityType_ANTIVIRUS                 string = "ANTIVIRUS"
	ActivityType_CLOUDBERRY                string = "CLOUDBERRY"
	ActivityType_CLOUDBERRY_BACKUP         string = "CLOUDBERRY_BACKUP"
	ActivityType_COMMENT                   string = "COMMENT"
	ActivityType_CONDITION                 string = "CONDITION"
	ActivityType_CONDITION_ACTION          string = "CONDITION_ACTION"
	ActivityType_CONDITION_ACTIONSET       string = "CONDITION_ACTIONSET"
	ActivityType_HELP_REQUEST              string = "HELP_REQUEST"
	ActivityType_IMAGEMANAGER              string = "IMAGEMANAGER"
	ActivityType_MDM                       string = "MDM"
	ActivityType_MONITOR                   string = "MONITOR"
	ActivityType_PATCH_MANAGEMENT          string = "PATCH_MANAGEMENT"
	ActivityType_PSA                       string = "PSA"
	ActivityType_RDP                       string = "RDP"
	ActivityType_REMOTE_TOOLS              string = "REMOTE_TOOLS"
	ActivityType_SCHEDULED_TASK            string = "SCHEDULED_TASK"
	ActivityType_SCRIPTING                 string = "SCRIPTING"
	ActivityType_SECURITY                  string = "SECURITY"
	ActivityType_SHADOWPROTECT             string = "SHADOWPROTECT"
	ActivityType_SOFTWARE_PATCH_MANAGEMENT string = "SOFTWARE_PATCH_MANAGEMENT"
	ActivityType_SPLASHTOP                 string = "SPLASHTOP"
	ActivityType_SYSTEM                    string = "SYSTEM"
	ActivityType_TEAMVIEWER                string = "TEAMVIEWER"
	ActivityType_VIRTUALIZATION            string = "VIRTUALIZATION"
)

const (
	StatusCode_ACKNOWLEDGED                                                         string = "ACKNOWLEDGED"
	StatusCode_ACTIVE_DIRECTORY_ADD_USER_TO_GROUP_AND_SET_PRIMARY_GROUP_FAILED      string = "ACTIVE_DIRECTORY_ADD_USER_TO_GROUP_AND_SET_PRIMARY_GROUP_FAILED"
	StatusCode_ACTIVE_DIRECTORY_ADD_USER_TO_GROUP_AND_SET_PRIMARY_GROUP_SUCCESS     string = "ACTIVE_DIRECTORY_ADD_USER_TO_GROUP_AND_SET_PRIMARY_GROUP_SUCCESS"
	StatusCode_ACTIVE_DIRECTORY_ADD_USER_TO_GROUP_FAILED                            string = "ACTIVE_DIRECTORY_ADD_USER_TO_GROUP_FAILED"
	StatusCode_ACTIVE_DIRECTORY_ADD_USER_TO_GROUP_SUCCESS                           string = "ACTIVE_DIRECTORY_ADD_USER_TO_GROUP_SUCCESS"
	StatusCode_ACTIVE_DIRECTORY_ALLOW_USER_PASSWORD_CHANGE_FAILED                   string = "ACTIVE_DIRECTORY_ALLOW_USER_PASSWORD_CHANGE_FAILED"
	StatusCode_ACTIVE_DIRECTORY_ALLOW_USER_PASSWORD_CHANGE_SUCCESS                  string = "ACTIVE_DIRECTORY_ALLOW_USER_PASSWORD_CHANGE_SUCCESS"
	StatusCode_ACTIVE_DIRECTORY_AUTO_DISCOVERY_CHANGE_CANDIDATE_NODE_STATUS_FAILED  string = "ACTIVE_DIRECTORY_AUTO_DISCOVERY_CHANGE_CANDIDATE_NODE_STATUS_FAILED"
	StatusCode_ACTIVE_DIRECTORY_AUTO_DISCOVERY_CHANGE_CANDIDATE_NODE_STATUS_SUCCESS string = "ACTIVE_DIRECTORY_AUTO_DISCOVERY_CHANGE_CANDIDATE_NODE_STATUS_SUCCESS"
	StatusCode_ACTIVE_DIRECTORY_AUTO_DISCOVERY_CLEAR_CANDIDATE_NODES_FAILED         string = "ACTIVE_DIRECTORY_AUTO_DISCOVERY_CLEAR_CANDIDATE_NODES_FAILED"
	StatusCode_ACTIVE_DIRECTORY_AUTO_DISCOVERY_CLEAR_CANDIDATE_NODES_SUCCESS        string = "ACTIVE_DIRECTORY_AUTO_DISCOVERY_CLEAR_CANDIDATE_NODES_SUCCESS"
	StatusCode_ACTIVE_DIRECTORY_AUTO_DISCOVERY_DEPLOYMENT_RESULT                    string = "ACTIVE_DIRECTORY_AUTO_DISCOVERY_DEPLOYMENT_RESULT"
	StatusCode_ACTIVE_DIRECTORY_AUTO_DISCOVERY_DISCOVERY_JOB_CREATED                string = "ACTIVE_DIRECTORY_AUTO_DISCOVERY_DISCOVERY_JOB_CREATED"
	StatusCode_ACTIVE_DIRECTORY_AUTO_DISCOVERY_DISCOVERY_JOB_DELETED                string = "ACTIVE_DIRECTORY_AUTO_DISCOVERY_DISCOVERY_JOB_DELETED"
	StatusCode_ACTIVE_DIRECTORY_AUTO_DISCOVERY_DISCOVERY_JOB_UPDATED                string = "ACTIVE_DIRECTORY_AUTO_DISCOVERY_DISCOVERY_JOB_UPDATED"
	StatusCode_ACTIVE_DIRECTORY_AUTO_DISCOVERY_JOB_RESULT_FAILED                    string = "ACTIVE_DIRECTORY_AUTO_DISCOVERY_JOB_RESULT_FAILED"
	StatusCode_ACTIVE_DIRECTORY_AUTO_DISCOVERY_JOB_RESULT_SUCCESS                   string = "ACTIVE_DIRECTORY_AUTO_DISCOVERY_JOB_RESULT_SUCCESS"
	StatusCode_ACTIVE_DIRECTORY_DISABLE_USER_FAILED                                 string = "ACTIVE_DIRECTORY_DISABLE_USER_FAILED"
	StatusCode_ACTIVE_DIRECTORY_DISABLE_USER_SUCCESS                                string = "ACTIVE_DIRECTORY_DISABLE_USER_SUCCESS"
	StatusCode_ACTIVE_DIRECTORY_DISABLED_ACCOUNT_EXPIRATION_FAILED                  string = "ACTIVE_DIRECTORY_DISABLED_ACCOUNT_EXPIRATION_FAILED"
	StatusCode_ACTIVE_DIRECTORY_DISABLED_ACCOUNT_EXPIRATION_SUCCESS                 string = "ACTIVE_DIRECTORY_DISABLED_ACCOUNT_EXPIRATION_SUCCESS"
	StatusCode_ACTIVE_DIRECTORY_DISABLED_PASSWORD_EXPIRATION_FAILED                 string = "ACTIVE_DIRECTORY_DISABLED_PASSWORD_EXPIRATION_FAILED"
	StatusCode_ACTIVE_DIRECTORY_DISABLED_PASSWORD_EXPIRATION_SUCCESS                string = "ACTIVE_DIRECTORY_DISABLED_PASSWORD_EXPIRATION_SUCCESS"
	StatusCode_ACTIVE_DIRECTORY_DISALLOW_USER_PASSWORD_CHANGE_FAILED                string = "ACTIVE_DIRECTORY_DISALLOW_USER_PASSWORD_CHANGE_FAILED"
	StatusCode_ACTIVE_DIRECTORY_DISALLOW_USER_PASSWORD_CHANGE_SUCCESS               string = "ACTIVE_DIRECTORY_DISALLOW_USER_PASSWORD_CHANGE_SUCCESS"
	StatusCode_ACTIVE_DIRECTORY_ENABLE_USER_FAILED                                  string = "ACTIVE_DIRECTORY_ENABLE_USER_FAILED"
	StatusCode_ACTIVE_DIRECTORY_ENABLE_USER_SUCCESS                                 string = "ACTIVE_DIRECTORY_ENABLE_USER_SUCCESS"
	StatusCode_ACTIVE_DIRECTORY_ENABLED_PASSWORD_EXPIRATION_FAILED                  string = "ACTIVE_DIRECTORY_ENABLED_PASSWORD_EXPIRATION_FAILED"
	StatusCode_ACTIVE_DIRECTORY_ENABLED_PASSWORD_EXPIRATION_SUCCESS                 string = "ACTIVE_DIRECTORY_ENABLED_PASSWORD_EXPIRATION_SUCCESS"
	StatusCode_ACTIVE_DIRECTORY_NOT_REQUIRE_PASSWORD_CHANGE_FAILED                  string = "ACTIVE_DIRECTORY_NOT_REQUIRE_PASSWORD_CHANGE_FAILED"
	StatusCode_ACTIVE_DIRECTORY_NOT_REQUIRE_PASSWORD_CHANGE_SUCCESS                 string = "ACTIVE_DIRECTORY_NOT_REQUIRE_PASSWORD_CHANGE_SUCCESS"
	StatusCode_ACTIVE_DIRECTORY_REMOVE_USER_FROM_GROUP_FAILED                       string = "ACTIVE_DIRECTORY_REMOVE_USER_FROM_GROUP_FAILED"
	StatusCode_ACTIVE_DIRECTORY_REMOVE_USER_FROM_GROUP_SUCCESS                      string = "ACTIVE_DIRECTORY_REMOVE_USER_FROM_GROUP_SUCCESS"
	StatusCode_ACTIVE_DIRECTORY_REQUIRE_PASSWORD_CHANGE_FAILED                      string = "ACTIVE_DIRECTORY_REQUIRE_PASSWORD_CHANGE_FAILED"
	StatusCode_ACTIVE_DIRECTORY_REQUIRE_PASSWORD_CHANGE_SUCCESS                     string = "ACTIVE_DIRECTORY_REQUIRE_PASSWORD_CHANGE_SUCCESS"
	StatusCode_ACTIVE_DIRECTORY_RESET_PASSWORD_FAILED                               string = "ACTIVE_DIRECTORY_RESET_PASSWORD_FAILED"
	StatusCode_ACTIVE_DIRECTORY_RESET_PASSWORD_SUCCESS                              string = "ACTIVE_DIRECTORY_RESET_PASSWORD_SUCCESS"
	StatusCode_ACTIVE_DIRECTORY_SET_ACCOUNT_EXPIRATION_FAILED                       string = "ACTIVE_DIRECTORY_SET_ACCOUNT_EXPIRATION_FAILED"
	StatusCode_ACTIVE_DIRECTORY_SET_ACCOUNT_EXPIRATION_SUCCESS                      string = "ACTIVE_DIRECTORY_SET_ACCOUNT_EXPIRATION_SUCCESS"
	StatusCode_ACTIVE_DIRECTORY_SET_PRIMARY_GROUP_FAILED                            string = "ACTIVE_DIRECTORY_SET_PRIMARY_GROUP_FAILED"
	StatusCode_ACTIVE_DIRECTORY_SET_PRIMARY_GROUP_SUCCESS                           string = "ACTIVE_DIRECTORY_SET_PRIMARY_GROUP_SUCCESS"
	StatusCode_ACTIVE_DIRECTORY_UNLOCK_USER_FAILED                                  string = "ACTIVE_DIRECTORY_UNLOCK_USER_FAILED"
	StatusCode_ACTIVE_DIRECTORY_UNLOCK_USER_SUCCESS                                 string = "ACTIVE_DIRECTORY_UNLOCK_USER_SUCCESS"
	StatusCode_ADAPTER_ADDED                                                        string = "ADAPTER_ADDED"
	StatusCode_ADAPTER_CONFIG_CHANGED                                               string = "ADAPTER_CONFIG_CHANGED"
	StatusCode_ADAPTER_REMOVED                                                      string = "ADAPTER_REMOVED"
	StatusCode_ADAPTER_STATUS_CHANGED                                               string = "ADAPTER_STATUS_CHANGED"
	StatusCode_AGENT_MESSAGE                                                        string = "AGENT_MESSAGE"
	StatusCode_API_ACCESS_DENIED                                                    string = "API_ACCESS_DENIED"
	StatusCode_API_ACCESS_GRANTED                                                   string = "API_ACCESS_GRANTED"
	StatusCode_API_ACCESS_REVOKED                                                   string = "API_ACCESS_REVOKED"
	StatusCode_API_WEBHOOK_CONFIGURATION_UPDATED                                    string = "API_WEBHOOK_CONFIGURATION_UPDATED"
	StatusCode_APP_USER_AUDIT_FAILED_LOGIN                                          string = "APP_USER_AUDIT_FAILED_LOGIN"
	StatusCode_APP_USER_CREATED                                                     string = "APP_USER_CREATED"
	StatusCode_APP_USER_CRITICAL_ACTION                                             string = "APP_USER_CRITICAL_ACTION"
	StatusCode_APP_USER_DELETED                                                     string = "APP_USER_DELETED"
	StatusCode_APP_USER_LOGGED_IN                                                   string = "APP_USER_LOGGED_IN"
	StatusCode_APP_USER_LOGGED_OUT                                                  string = "APP_USER_LOGGED_OUT"
	StatusCode_APP_USER_MFA_DELETED                                                 string = "APP_USER_MFA_DELETED"
	StatusCode_APP_USER_MFA_SETUP                                                   string = "APP_USER_MFA_SETUP"
	StatusCode_APP_USER_UPDATED                                                     string = "APP_USER_UPDATED"
	StatusCode_AUTOTASK_NODE_SYNC_COMPLETED                                         string = "AUTOTASK_NODE_SYNC_COMPLETED"
	StatusCode_AUTOTASK_NODE_SYNC_NODE_CREATED                                      string = "AUTOTASK_NODE_SYNC_NODE_CREATED"
	StatusCode_AUTOTASK_NODE_SYNC_NODE_DELETED                                      string = "AUTOTASK_NODE_SYNC_NODE_DELETED"
	StatusCode_AUTOTASK_NODE_SYNC_NODE_UPDATED                                      string = "AUTOTASK_NODE_SYNC_NODE_UPDATED"
	StatusCode_AUTOTASK_NODE_SYNC_STARTED                                           string = "AUTOTASK_NODE_SYNC_STARTED"
	StatusCode_AUTOTASK_UPDATED                                                     string = "AUTOTASK_UPDATED"
	StatusCode_BDAS_BITDEFENDER_PURGE_QUARANTINE_FAILED                             string = "BDAS_BITDEFENDER_PURGE_QUARANTINE_FAILED"
	StatusCode_BDAS_BITDEFENDER_PURGE_QUARANTINE_SUCCESS                            string = "BDAS_BITDEFENDER_PURGE_QUARANTINE_SUCCESS"
	StatusCode_BDAS_BITDEFENDER_THREAT_BLOCKED                                      string = "BDAS_BITDEFENDER_THREAT_BLOCKED"
	StatusCode_BDAS_BITDEFENDER_THREAT_CLEANED                                      string = "BDAS_BITDEFENDER_THREAT_CLEANED"
	StatusCode_BDAS_BITDEFENDER_THREAT_DELETED                                      string = "BDAS_BITDEFENDER_THREAT_DELETED"
	StatusCode_BDAS_BITDEFENDER_THREAT_IGNORED                                      string = "BDAS_BITDEFENDER_THREAT_IGNORED"
	StatusCode_BDAS_BITDEFENDER_THREAT_PRESENT                                      string = "BDAS_BITDEFENDER_THREAT_PRESENT"
	StatusCode_BDAS_BITDEFENDER_THREAT_QUARANTINE_DELETED                           string = "BDAS_BITDEFENDER_THREAT_QUARANTINE_DELETED"
	StatusCode_BDAS_BITDEFENDER_THREAT_QUARANTINE_RESTORED                          string = "BDAS_BITDEFENDER_THREAT_QUARANTINE_RESTORED"
	StatusCode_BDAS_BITDEFENDER_THREAT_QUARANTINE_RESTORED_CUSTOMPATH               string = "BDAS_BITDEFENDER_THREAT_QUARANTINE_RESTORED_CUSTOMPATH"
	StatusCode_BDAS_BITDEFENDER_THREAT_QUARANTINED                                  string = "BDAS_BITDEFENDER_THREAT_QUARANTINED"
	StatusCode_BITDEFENDER_DISABLED                                                 string = "BITDEFENDER_DISABLED"
	StatusCode_BITDEFENDER_DOWNLOAD_FAILED                                          string = "BITDEFENDER_DOWNLOAD_FAILED"
	StatusCode_BITDEFENDER_DOWNLOAD_STARTED                                         string = "BITDEFENDER_DOWNLOAD_STARTED"
	StatusCode_BITDEFENDER_DOWNLOAD_SUCCEEDED                                       string = "BITDEFENDER_DOWNLOAD_SUCCEEDED"
	StatusCode_BITDEFENDER_EXISTING_PRODUCT_UNINSTALL                               string = "BITDEFENDER_EXISTING_PRODUCT_UNINSTALL"
	StatusCode_BITDEFENDER_INSTALLATION_FAILED                                      string = "BITDEFENDER_INSTALLATION_FAILED"
	StatusCode_BITDEFENDER_INSTALLATION_STARTED                                     string = "BITDEFENDER_INSTALLATION_STARTED"
	StatusCode_BITDEFENDER_INSTALLATION_SUCCEEDED                                   string = "BITDEFENDER_INSTALLATION_SUCCEEDED"
	StatusCode_BITDEFENDER_RETRY_INSTALL_COMPLETED                                  string = "BITDEFENDER_RETRY_INSTALL_COMPLETED"
	StatusCode_BITDEFENDER_SCAN_COMPLETED                                           string = "BITDEFENDER_SCAN_COMPLETED"
	StatusCode_BITDEFENDER_SCAN_FAILED                                              string = "BITDEFENDER_SCAN_FAILED"
	StatusCode_BITDEFENDER_SCAN_STARTED                                             string = "BITDEFENDER_SCAN_STARTED"
	StatusCode_BITDEFENDER_THREAT_DELETE_FROM_QUARANTINE                            string = "BITDEFENDER_THREAT_DELETE_FROM_QUARANTINE"
	StatusCode_BITDEFENDER_THREAT_DELETE_FROM_QUARANTINE_FAILED                     string = "BITDEFENDER_THREAT_DELETE_FROM_QUARANTINE_FAILED"
	StatusCode_BITDEFENDER_THREAT_RESTORE_FROM_QUARANTINE                           string = "BITDEFENDER_THREAT_RESTORE_FROM_QUARANTINE"
	StatusCode_BITDEFENDER_THREAT_RESTORE_FROM_QUARANTINE_FAILED                    string = "BITDEFENDER_THREAT_RESTORE_FROM_QUARANTINE_FAILED"
	StatusCode_BITDEFENDER_UNINSTALLATION_FAILED                                    string = "BITDEFENDER_UNINSTALLATION_FAILED"
	StatusCode_BITDEFENDER_UNINSTALLATION_STARTED                                   string = "BITDEFENDER_UNINSTALLATION_STARTED"
	StatusCode_BITDEFENDER_UNINSTALLATION_SUCCEEDED                                 string = "BITDEFENDER_UNINSTALLATION_SUCCEEDED"
	StatusCode_BITDEFENDER_UNPACKING_FAILED                                         string = "BITDEFENDER_UNPACKING_FAILED"
	StatusCode_BITLOCKER_DISABLED                                                   string = "BITLOCKER_DISABLED"
	StatusCode_BITLOCKER_ENABLED                                                    string = "BITLOCKER_ENABLED"
	StatusCode_BLOCKED                                                              string = "BLOCKED"
	StatusCode_CANCEL_REQUESTED                                                     string = "CANCEL_REQUESTED"
	StatusCode_CANCELLED                                                            string = "CANCELLED"
	StatusCode_CLIENT_CREATED                                                       string = "CLIENT_CREATED"
	StatusCode_CLIENT_DELETED                                                       string = "CLIENT_DELETED"
	StatusCode_CLIENT_DOCUMENT_ATTRIBUTE_VALUE_DECRYPTED                            string = "CLIENT_DOCUMENT_ATTRIBUTE_VALUE_DECRYPTED"
	StatusCode_CLIENT_DOCUMENT_CREATED                                              string = "CLIENT_DOCUMENT_CREATED"
	StatusCode_CLIENT_DOCUMENT_DELETED                                              string = "CLIENT_DOCUMENT_DELETED"
	StatusCode_CLIENT_DOCUMENT_UPDATED                                              string = "CLIENT_DOCUMENT_UPDATED"
	StatusCode_CLIENT_UPDATED                                                       string = "CLIENT_UPDATED"
	StatusCode_CLOUDBERRY_BACKUPJOB_COMPLETED_WITH_WARNING                          string = "CLOUDBERRY_BACKUPJOB_COMPLETED_WITH_WARNING"
	StatusCode_CLOUDBERRY_BACKUPJOB_FAILED                                          string = "CLOUDBERRY_BACKUPJOB_FAILED"
	StatusCode_CLOUDBERRY_BACKUPJOB_STARTED                                         string = "CLOUDBERRY_BACKUPJOB_STARTED"
	StatusCode_CLOUDBERRY_BACKUPJOB_SUCCEEDED                                       string = "CLOUDBERRY_BACKUPJOB_SUCCEEDED"
	StatusCode_CLOUDBERRY_BACKUPPLAN_CREATED                                        string = "CLOUDBERRY_BACKUPPLAN_CREATED"
	StatusCode_CLOUDBERRY_BACKUPPLAN_CREATION_FAILED                                string = "CLOUDBERRY_BACKUPPLAN_CREATION_FAILED"
	StatusCode_CLOUDBERRY_BACKUPPLAN_DELETED                                        string = "CLOUDBERRY_BACKUPPLAN_DELETED"
	StatusCode_CLOUDBERRY_BACKUPPLAN_EDITED                                         string = "CLOUDBERRY_BACKUPPLAN_EDITED"
	StatusCode_CLOUDBERRY_INSTALL_FAILED                                            string = "CLOUDBERRY_INSTALL_FAILED"
	StatusCode_CLOUDBERRY_INSTALLED                                                 string = "CLOUDBERRY_INSTALLED"
	StatusCode_CLOUDBERRY_NETWORK_CREDENTIAL_CREATED                                string = "CLOUDBERRY_NETWORK_CREDENTIAL_CREATED"
	StatusCode_CLOUDBERRY_NETWORK_CREDENTIAL_CREATION_FAILED                        string = "CLOUDBERRY_NETWORK_CREDENTIAL_CREATION_FAILED"
	StatusCode_CLOUDBERRY_UNINSTALL_FAILED                                          string = "CLOUDBERRY_UNINSTALL_FAILED"
	StatusCode_CLOUDBERRY_UNINSTALLED                                               string = "CLOUDBERRY_UNINSTALLED"
	StatusCode_CLOUDBERRY_USER_CREATED                                              string = "CLOUDBERRY_USER_CREATED"
	StatusCode_COMMENT                                                              string = "COMMENT"
	StatusCode_COMPETITOR_EXISTING_PRODUCT_UNINSTALL                                string = "COMPETITOR_EXISTING_PRODUCT_UNINSTALL"
	StatusCode_COMPLETED                                                            string = "COMPLETED"
	StatusCode_CONNECTWISE_AGREEMENTS_SYNC_COMPLETED                                string = "CONNECTWISE_AGREEMENTS_SYNC_COMPLETED"
	StatusCode_CONNECTWISE_AGREEMENTS_SYNC_STARTED                                  string = "CONNECTWISE_AGREEMENTS_SYNC_STARTED"
	StatusCode_CONNECTWISE_NODE_SYNC_COMPLETED                                      string = "CONNECTWISE_NODE_SYNC_COMPLETED"
	StatusCode_CONNECTWISE_NODE_SYNC_NODE_CREATED                                   string = "CONNECTWISE_NODE_SYNC_NODE_CREATED"
	StatusCode_CONNECTWISE_NODE_SYNC_NODE_DELETED                                   string = "CONNECTWISE_NODE_SYNC_NODE_DELETED"
	StatusCode_CONNECTWISE_NODE_SYNC_NODE_UPDATED                                   string = "CONNECTWISE_NODE_SYNC_NODE_UPDATED"
	StatusCode_CONNECTWISE_NODE_SYNC_STARTED                                        string = "CONNECTWISE_NODE_SYNC_STARTED"
	StatusCode_CONNECTWISE_UPDATED                                                  string = "CONNECTWISE_UPDATED"
	StatusCode_CONNECTWISECONTROL_ATTEMPT                                           string = "CONNECTWISECONTROL_ATTEMPT"
	StatusCode_CONTACT_CREATED                                                      string = "CONTACT_CREATED"
	StatusCode_CONTACT_DELETED                                                      string = "CONTACT_DELETED"
	StatusCode_CONTACT_UPDATED                                                      string = "CONTACT_UPDATED"
	StatusCode_CPU_ADDED                                                            string = "CPU_ADDED"
	StatusCode_CPU_REMOVED                                                          string = "CPU_REMOVED"
	StatusCode_CREDENTIAL_CREATED                                                   string = "CREDENTIAL_CREATED"
	StatusCode_CREDENTIAL_DELETED                                                   string = "CREDENTIAL_DELETED"
	StatusCode_CREDENTIAL_UPDATED                                                   string = "CREDENTIAL_UPDATED"
	StatusCode_CREDENTIALS_CHANGED                                                  string = "CREDENTIALS_CHANGED"
	StatusCode_CUSTOM_HEALTH_STATUS_CHANGED                                         string = "CUSTOM_HEALTH_STATUS_CHANGED"
	StatusCode_CUSTOM_HEALTH_STATUS_RESET                                           string = "CUSTOM_HEALTH_STATUS_RESET"
	StatusCode_DEVICE_GROUP_CREATED                                                 string = "DEVICE_GROUP_CREATED"
	StatusCode_DEVICE_GROUP_DELETED                                                 string = "DEVICE_GROUP_DELETED"
	StatusCode_DEVICE_GROUP_UPDATED                                                 string = "DEVICE_GROUP_UPDATED"
	StatusCode_DISABLED                                                             string = "DISABLED"
	StatusCode_DISK_DRIVE_ADDED                                                     string = "DISK_DRIVE_ADDED"
	StatusCode_DISK_DRIVE_REMOVED                                                   string = "DISK_DRIVE_REMOVED"
	StatusCode_DISK_PARTITION_ADDED                                                 string = "DISK_PARTITION_ADDED"
	StatusCode_DISK_PARTITION_REMOVED                                               string = "DISK_PARTITION_REMOVED"
	StatusCode_DISK_VOLUME_ADDED                                                    string = "DISK_VOLUME_ADDED"
	StatusCode_DISK_VOLUME_REMOVED                                                  string = "DISK_VOLUME_REMOVED"
	StatusCode_DIVISION_FEATURE_DISABLED                                            string = "DIVISION_FEATURE_DISABLED"
	StatusCode_DIVISION_FEATURE_ENABLED                                             string = "DIVISION_FEATURE_ENABLED"
	StatusCode_DIVISION_STATUS_CHANGED                                              string = "DIVISION_STATUS_CHANGED"
	StatusCode_DOCUMENT_TEMPLATE_CREATED                                            string = "DOCUMENT_TEMPLATE_CREATED"
	StatusCode_DOCUMENT_TEMPLATE_DELETED                                            string = "DOCUMENT_TEMPLATE_DELETED"
	StatusCode_DOCUMENT_TEMPLATE_UPDATED                                            string = "DOCUMENT_TEMPLATE_UPDATED"
	StatusCode_END_USER_AUDIT_FAILED_LOGIN                                          string = "END_USER_AUDIT_FAILED_LOGIN"
	StatusCode_END_USER_CREATED                                                     string = "END_USER_CREATED"
	StatusCode_END_USER_DELETED                                                     string = "END_USER_DELETED"
	StatusCode_END_USER_LOGGED_IN                                                   string = "END_USER_LOGGED_IN"
	StatusCode_END_USER_LOGGED_OUT                                                  string = "END_USER_LOGGED_OUT"
	StatusCode_END_USER_MFA_DELETED                                                 string = "END_USER_MFA_DELETED"
	StatusCode_END_USER_MFA_SETUP                                                   string = "END_USER_MFA_SETUP"
	StatusCode_END_USER_UPDATED                                                     string = "END_USER_UPDATED"
	StatusCode_EVALUATION_FAILURE                                                   string = "EVALUATION_FAILURE"
	StatusCode_FILEVAULT_DISABLED                                                   string = "FILEVAULT_DISABLED"
	StatusCode_FILEVAULT_ENABLED                                                    string = "FILEVAULT_ENABLED"
	StatusCode_HELP_REQUEST_SUBMITTED                                               string = "HELP_REQUEST_SUBMITTED"
	StatusCode_IMAGEMANAGER_CONSOLIDATION_FAILED                                    string = "IMAGEMANAGER_CONSOLIDATION_FAILED"
	StatusCode_IMAGEMANAGER_INSTALL_FAILED                                          string = "IMAGEMANAGER_INSTALL_FAILED"
	StatusCode_IMAGEMANAGER_INSTALLED                                               string = "IMAGEMANAGER_INSTALLED"
	StatusCode_IMAGEMANAGER_LICENSE_ACTIVATED                                       string = "IMAGEMANAGER_LICENSE_ACTIVATED"
	StatusCode_IMAGEMANAGER_LICENSE_ACTIVATION_FAILED                               string = "IMAGEMANAGER_LICENSE_ACTIVATION_FAILED"
	StatusCode_IMAGEMANAGER_LICENSE_DEACTIVATED                                     string = "IMAGEMANAGER_LICENSE_DEACTIVATED"
	StatusCode_IMAGEMANAGER_LICENSE_DEACTIVATION_FAILED                             string = "IMAGEMANAGER_LICENSE_DEACTIVATION_FAILED"
	StatusCode_IMAGEMANAGER_LICENSE_PROVISION_FAILED                                string = "IMAGEMANAGER_LICENSE_PROVISION_FAILED"
	StatusCode_IMAGEMANAGER_LICENSE_PROVISIONED                                     string = "IMAGEMANAGER_LICENSE_PROVISIONED"
	StatusCode_IMAGEMANAGER_UNINSTALL_FAILED                                        string = "IMAGEMANAGER_UNINSTALL_FAILED"
	StatusCode_IMAGEMANAGER_UNINSTALLED                                             string = "IMAGEMANAGER_UNINSTALLED"
	StatusCode_IMAGEMANAGER_VERIFICATION_FAILED                                     string = "IMAGEMANAGER_VERIFICATION_FAILED"
	StatusCode_IN_PROCESS                                                           string = "IN_PROCESS"
	StatusCode_LANGUAGE_TAG_UPDATED                                                 string = "LANGUAGE_TAG_UPDATED"
	StatusCode_LOCATION_CREATED                                                     string = "LOCATION_CREATED"
	StatusCode_LOCATION_DELETED                                                     string = "LOCATION_DELETED"
	StatusCode_LOCATION_UPDATED                                                     string = "LOCATION_UPDATED"
	StatusCode_LOCKHART_BACKUP_CONFIGURE_FAILED                                     string = "LOCKHART_BACKUP_CONFIGURE_FAILED"
	StatusCode_LOCKHART_BACKUP_DOWNLOADED                                           string = "LOCKHART_BACKUP_DOWNLOADED"
	StatusCode_LOCKHART_BACKUP_NAS_ACCESS_FAILED                                    string = "LOCKHART_BACKUP_NAS_ACCESS_FAILED"
	StatusCode_LOCKHART_BACKUPJOB_CANCELLED                                         string = "LOCKHART_BACKUPJOB_CANCELLED"
	StatusCode_LOCKHART_BACKUPJOB_COMPLETED                                         string = "LOCKHART_BACKUPJOB_COMPLETED"
	StatusCode_LOCKHART_BACKUPJOB_COMPLETED_WITH_WARNING                            string = "LOCKHART_BACKUPJOB_COMPLETED_WITH_WARNING"
	StatusCode_LOCKHART_BACKUPJOB_FAILED                                            string = "LOCKHART_BACKUPJOB_FAILED"
	StatusCode_LOCKHART_BACKUPJOB_IN_PROCESS                                        string = "LOCKHART_BACKUPJOB_IN_PROCESS"
	StatusCode_LOCKHART_BACKUPJOB_PROCESSING                                        string = "LOCKHART_BACKUPJOB_PROCESSING"
	StatusCode_LOCKHART_BACKUPJOB_START_REQUESTED                                   string = "LOCKHART_BACKUPJOB_START_REQUESTED"
	StatusCode_LOCKHART_BACKUPJOB_STARTED                                           string = "LOCKHART_BACKUPJOB_STARTED"
	StatusCode_LOCKHART_BACKUPPLAN_ADDED                                            string = "LOCKHART_BACKUPPLAN_ADDED"
	StatusCode_LOCKHART_BACKUPPLAN_CREATION_FAILED                                  string = "LOCKHART_BACKUPPLAN_CREATION_FAILED"
	StatusCode_LOCKHART_BACKUPPLAN_DELETED                                          string = "LOCKHART_BACKUPPLAN_DELETED"
	StatusCode_LOCKHART_BACKUPPLAN_EDITED                                           string = "LOCKHART_BACKUPPLAN_EDITED"
	StatusCode_LOCKHART_FILE_DOWNLOAD                                               string = "LOCKHART_FILE_DOWNLOAD"
	StatusCode_LOCKHART_FILES_AND_FOLDERS_DELETED                                   string = "LOCKHART_FILES_AND_FOLDERS_DELETED"
	StatusCode_LOCKHART_FILES_DELETED                                               string = "LOCKHART_FILES_DELETED"
	StatusCode_LOCKHART_FOLDER_DOWNLOAD                                             string = "LOCKHART_FOLDER_DOWNLOAD"
	StatusCode_LOCKHART_FOLDERS_DELETED                                             string = "LOCKHART_FOLDERS_DELETED"
	StatusCode_LOCKHART_IMAGE_DOWNLOAD_COMPLETED                                    string = "LOCKHART_IMAGE_DOWNLOAD_COMPLETED"
	StatusCode_LOCKHART_IMAGE_DOWNLOAD_FAILED                                       string = "LOCKHART_IMAGE_DOWNLOAD_FAILED"
	StatusCode_LOCKHART_IMAGE_DOWNLOAD_FILE_COMPLETED                               string = "LOCKHART_IMAGE_DOWNLOAD_FILE_COMPLETED"
	StatusCode_LOCKHART_IMAGE_DOWNLOAD_FILE_FAILED                                  string = "LOCKHART_IMAGE_DOWNLOAD_FILE_FAILED"
	StatusCode_LOCKHART_IMAGE_DOWNLOAD_FILE_STARTED                                 string = "LOCKHART_IMAGE_DOWNLOAD_FILE_STARTED"
	StatusCode_LOCKHART_IMAGE_DOWNLOAD_FOLDER_COMPLETED                             string = "LOCKHART_IMAGE_DOWNLOAD_FOLDER_COMPLETED"
	StatusCode_LOCKHART_IMAGE_DOWNLOAD_FOLDER_FAILED                                string = "LOCKHART_IMAGE_DOWNLOAD_FOLDER_FAILED"
	StatusCode_LOCKHART_IMAGE_DOWNLOAD_FOLDER_STARTED                               string = "LOCKHART_IMAGE_DOWNLOAD_FOLDER_STARTED"
	StatusCode_LOCKHART_IMAGE_DOWNLOAD_STARTED                                      string = "LOCKHART_IMAGE_DOWNLOAD_STARTED"
	StatusCode_LOCKHART_IMAGE_MOUNT_COMPLETED                                       string = "LOCKHART_IMAGE_MOUNT_COMPLETED"
	StatusCode_LOCKHART_IMAGE_MOUNT_FAILED                                          string = "LOCKHART_IMAGE_MOUNT_FAILED"
	StatusCode_LOCKHART_IMAGE_RESTORE_COMPLETED                                     string = "LOCKHART_IMAGE_RESTORE_COMPLETED"
	StatusCode_LOCKHART_IMAGE_RESTORE_FAILED                                        string = "LOCKHART_IMAGE_RESTORE_FAILED"
	StatusCode_LOCKHART_IMAGE_RESTORE_STARTED                                       string = "LOCKHART_IMAGE_RESTORE_STARTED"
	StatusCode_LOCKHART_INSTALL_FAILED                                              string = "LOCKHART_INSTALL_FAILED"
	StatusCode_LOCKHART_INSTALLED                                                   string = "LOCKHART_INSTALLED"
	StatusCode_LOCKHART_RESTOREJOB_CANCELLED                                        string = "LOCKHART_RESTOREJOB_CANCELLED"
	StatusCode_LOCKHART_RESTOREJOB_COMPLETED                                        string = "LOCKHART_RESTOREJOB_COMPLETED"
	StatusCode_LOCKHART_RESTOREJOB_FAILED                                           string = "LOCKHART_RESTOREJOB_FAILED"
	StatusCode_LOCKHART_RESTOREJOB_IN_PROCESS                                       string = "LOCKHART_RESTOREJOB_IN_PROCESS"
	StatusCode_LOCKHART_RESTOREJOB_START_REQUESTED                                  string = "LOCKHART_RESTOREJOB_START_REQUESTED"
	StatusCode_LOCKHART_RESTOREJOB_STARTED                                          string = "LOCKHART_RESTOREJOB_STARTED"
	StatusCode_LOCKHART_REVISIONS_DELETE_FAILED                                     string = "LOCKHART_REVISIONS_DELETE_FAILED"
	StatusCode_LOCKHART_REVISIONS_DELETE_STARTED                                    string = "LOCKHART_REVISIONS_DELETE_STARTED"
	StatusCode_LOCKHART_REVISIONS_DELETED                                           string = "LOCKHART_REVISIONS_DELETED"
	StatusCode_LOCKHART_UNINSTALL_FAILED                                            string = "LOCKHART_UNINSTALL_FAILED"
	StatusCode_LOCKHART_UNINSTALLED                                                 string = "LOCKHART_UNINSTALLED"
	StatusCode_LOCKHART_UPSYNCJOB_CANCELLED                                         string = "LOCKHART_UPSYNCJOB_CANCELLED"
	StatusCode_LOCKHART_UPSYNCJOB_COMPLETED                                         string = "LOCKHART_UPSYNCJOB_COMPLETED"
	StatusCode_LOCKHART_UPSYNCJOB_FAILED                                            string = "LOCKHART_UPSYNCJOB_FAILED"
	StatusCode_LOCKHART_UPSYNCJOB_IN_PROCESS                                        string = "LOCKHART_UPSYNCJOB_IN_PROCESS"
	StatusCode_LOCKHART_UPSYNCJOB_PROCESSING                                        string = "LOCKHART_UPSYNCJOB_PROCESSING"
	StatusCode_LOCKHART_UPSYNCJOB_STARTED                                           string = "LOCKHART_UPSYNCJOB_STARTED"
	StatusCode_MAC_DAEMON_STARTED                                                   string = "MAC_DAEMON_STARTED"
	StatusCode_MAC_DAEMON_STOPPED                                                   string = "MAC_DAEMON_STOPPED"
	StatusCode_MAINTENANCE_MODE_CONFIGURED                                          string = "MAINTENANCE_MODE_CONFIGURED"
	StatusCode_MAINTENANCE_MODE_DELETED                                             string = "MAINTENANCE_MODE_DELETED"
	StatusCode_MAINTENANCE_MODE_FAILED                                              string = "MAINTENANCE_MODE_FAILED"
	StatusCode_MAINTENANCE_MODE_MODIFIED                                            string = "MAINTENANCE_MODE_MODIFIED"
	StatusCode_MDM_CLEAR_PASSCODE_STATUS_CREATED                                    string = "MDM_CLEAR_PASSCODE_STATUS_CREATED"
	StatusCode_MDM_CLEAR_PASSCODE_STATUS_UPDATED                                    string = "MDM_CLEAR_PASSCODE_STATUS_UPDATED"
	StatusCode_MDM_ERASE_STATUS_CREATED                                             string = "MDM_ERASE_STATUS_CREATED"
	StatusCode_MDM_ERASE_STATUS_UPDATED                                             string = "MDM_ERASE_STATUS_UPDATED"
	StatusCode_MDM_LOCK_DEVICE_STATUS_CREATED                                       string = "MDM_LOCK_DEVICE_STATUS_CREATED"
	StatusCode_MDM_LOCK_DEVICE_STATUS_UPDATED                                       string = "MDM_LOCK_DEVICE_STATUS_UPDATED"
	StatusCode_MDM_REBOOT_DEVICE_STATUS_CREATED                                     string = "MDM_REBOOT_DEVICE_STATUS_CREATED"
	StatusCode_MDM_REBOOT_DEVICE_STATUS_UPDATED                                     string = "MDM_REBOOT_DEVICE_STATUS_UPDATED"
	StatusCode_MDM_RELINQUISH_OWNERSHIP_DEVICE_STATUS_CREATED                       string = "MDM_RELINQUISH_OWNERSHIP_DEVICE_STATUS_CREATED"
	StatusCode_MDM_RELINQUISH_OWNERSHIP_DEVICE_STATUS_UPDATED                       string = "MDM_RELINQUISH_OWNERSHIP_DEVICE_STATUS_UPDATED"
	StatusCode_MDM_RESET_PASSCODE_DEVICE_STATUS_CREATED                             string = "MDM_RESET_PASSCODE_DEVICE_STATUS_CREATED"
	StatusCode_MDM_RESET_PASSCODE_DEVICE_STATUS_UPDATED                             string = "MDM_RESET_PASSCODE_DEVICE_STATUS_UPDATED"
	StatusCode_MEMORY_ADDED                                                         string = "MEMORY_ADDED"
	StatusCode_MEMORY_REMOVED                                                       string = "MEMORY_REMOVED"
	StatusCode_MOBILE_DEVICE_UNREGISTERED                                           string = "MOBILE_DEVICE_UNREGISTERED"
	StatusCode_NC_SESSION_FAILED_TO_START                                           string = "NC_SESSION_FAILED_TO_START"
	StatusCode_NC_SESSION_REQUESTED                                                 string = "NC_SESSION_REQUESTED"
	StatusCode_NC_SESSION_SESSION_TERMINATION_REQUESTED                             string = "NC_SESSION_SESSION_TERMINATION_REQUESTED"
	StatusCode_NC_SESSION_STARTED                                                   string = "NC_SESSION_STARTED"
	StatusCode_NC_SESSION_TERMINATED                                                string = "NC_SESSION_TERMINATED"
	StatusCode_NINJA_TICKETING_ATTRIBUTE_ACTIVED                                    string = "NINJA_TICKETING_ATTRIBUTE_ACTIVED"
	StatusCode_NINJA_TICKETING_ATTRIBUTE_CREATED                                    string = "NINJA_TICKETING_ATTRIBUTE_CREATED"
	StatusCode_NINJA_TICKETING_ATTRIBUTE_DEACTIVATED                                string = "NINJA_TICKETING_ATTRIBUTE_DEACTIVATED"
	StatusCode_NINJA_TICKETING_ATTRIBUTE_UPDATED                                    string = "NINJA_TICKETING_ATTRIBUTE_UPDATED"
	StatusCode_NINJA_TICKETING_CONDITION_RULE_MAKE_DEFAULT                          string = "NINJA_TICKETING_CONDITION_RULE_MAKE_DEFAULT"
	StatusCode_NINJA_TICKETING_CREATION_FAILED                                      string = "NINJA_TICKETING_CREATION_FAILED"
	StatusCode_NINJA_TICKETING_DISABLED                                             string = "NINJA_TICKETING_DISABLED"
	StatusCode_NINJA_TICKETING_EMAIL_ADDRESS_CONFIG_CREATED                         string = "NINJA_TICKETING_EMAIL_ADDRESS_CONFIG_CREATED"
	StatusCode_NINJA_TICKETING_EMAIL_ADDRESS_CONFIG_DELETED                         string = "NINJA_TICKETING_EMAIL_ADDRESS_CONFIG_DELETED"
	StatusCode_NINJA_TICKETING_EMAIL_ADDRESS_CONFIG_MAKE_DEFAULT                    string = "NINJA_TICKETING_EMAIL_ADDRESS_CONFIG_MAKE_DEFAULT"
	StatusCode_NINJA_TICKETING_EMAIL_ADDRESS_CONFIG_UPDATED                         string = "NINJA_TICKETING_EMAIL_ADDRESS_CONFIG_UPDATED"
	StatusCode_NINJA_TICKETING_ENABLED                                              string = "NINJA_TICKETING_ENABLED"
	StatusCode_NINJA_TICKETING_FORM_ACTIVED                                         string = "NINJA_TICKETING_FORM_ACTIVED"
	StatusCode_NINJA_TICKETING_FORM_CREATED                                         string = "NINJA_TICKETING_FORM_CREATED"
	StatusCode_NINJA_TICKETING_FORM_DEACTIVED                                       string = "NINJA_TICKETING_FORM_DEACTIVED"
	StatusCode_NINJA_TICKETING_FORM_MAKE_DEFAULT                                    string = "NINJA_TICKETING_FORM_MAKE_DEFAULT"
	StatusCode_NINJA_TICKETING_FORM_UPDATED                                         string = "NINJA_TICKETING_FORM_UPDATED"
	StatusCode_NINJA_TICKETING_PENDING_EMAIL_CREATION_APPROVED                      string = "NINJA_TICKETING_PENDING_EMAIL_CREATION_APPROVED"
	StatusCode_NINJA_TICKETING_PENDING_EMAIL_RECEIVED                               string = "NINJA_TICKETING_PENDING_EMAIL_RECEIVED"
	StatusCode_NINJA_TICKETING_PENDING_EMAIL_REJECTED                               string = "NINJA_TICKETING_PENDING_EMAIL_REJECTED"
	StatusCode_NINJA_TICKETING_PENDING_EMAIL_UPDATE_APPROVED                        string = "NINJA_TICKETING_PENDING_EMAIL_UPDATE_APPROVED"
	StatusCode_NINJA_TICKETING_RULESET_CREATED                                      string = "NINJA_TICKETING_RULESET_CREATED"
	StatusCode_NINJA_TICKETING_RULESET_DELETED                                      string = "NINJA_TICKETING_RULESET_DELETED"
	StatusCode_NINJA_TICKETING_RULESET_UPDATED                                      string = "NINJA_TICKETING_RULESET_UPDATED"
	StatusCode_NINJA_TICKETING_SCRIPT_RULE_MAKE_DEFAULT                             string = "NINJA_TICKETING_SCRIPT_RULE_MAKE_DEFAULT"
	StatusCode_NINJA_TICKETING_TICKET_CREATED                                       string = "NINJA_TICKETING_TICKET_CREATED"
	StatusCode_NINJA_TICKETING_TICKET_DELETED                                       string = "NINJA_TICKETING_TICKET_DELETED"
	StatusCode_NINJA_TICKETING_TRIGGER_CREATED                                      string = "NINJA_TICKETING_TRIGGER_CREATED"
	StatusCode_NINJA_TICKETING_TRIGGER_DELETED                                      string = "NINJA_TICKETING_TRIGGER_DELETED"
	StatusCode_NINJA_TICKETING_TRIGGER_UPDATED                                      string = "NINJA_TICKETING_TRIGGER_UPDATED"
	StatusCode_NINJA_TICKETING_UPDATED                                              string = "NINJA_TICKETING_UPDATED"
	StatusCode_NODE_ACCESS_DENIED                                                   string = "NODE_ACCESS_DENIED"
	StatusCode_NODE_ACCESS_GRANTED                                                  string = "NODE_ACCESS_GRANTED"
	StatusCode_NODE_ATTRIBUTE_CREATED                                               string = "NODE_ATTRIBUTE_CREATED"
	StatusCode_NODE_ATTRIBUTE_DELETED                                               string = "NODE_ATTRIBUTE_DELETED"
	StatusCode_NODE_ATTRIBUTE_UPDATED                                               string = "NODE_ATTRIBUTE_UPDATED"
	StatusCode_NODE_ATTRIBUTE_VALUE_DECRYPTED                                       string = "NODE_ATTRIBUTE_VALUE_DECRYPTED"
	StatusCode_NODE_ATTRIBUTE_VALUE_UPDATED                                         string = "NODE_ATTRIBUTE_VALUE_UPDATED"
	StatusCode_NODE_AUTOMATICALLY_APPROVED                                          string = "NODE_AUTOMATICALLY_APPROVED"
	StatusCode_NODE_AUTOMATICALLY_REJECTED                                          string = "NODE_AUTOMATICALLY_REJECTED"
	StatusCode_NODE_CLONE_ADVISED_TO_REGISTER                                       string = "NODE_CLONE_ADVISED_TO_REGISTER"
	StatusCode_NODE_CLONING_DETECTED                                                string = "NODE_CLONING_DETECTED"
	StatusCode_NODE_CREATED                                                         string = "NODE_CREATED"
	StatusCode_NODE_DELETED                                                         string = "NODE_DELETED"
	StatusCode_NODE_IDENTIFICATION_UPDATED                                          string = "NODE_IDENTIFICATION_UPDATED"
	StatusCode_NODE_JOBS_CANCELLED                                                  string = "NODE_JOBS_CANCELLED"
	StatusCode_NODE_MANUALLY_APPROVED                                               string = "NODE_MANUALLY_APPROVED"
	StatusCode_NODE_MANUALLY_REJECTED                                               string = "NODE_MANUALLY_REJECTED"
	StatusCode_NODE_REGISTRATION_REJECTED                                           string = "NODE_REGISTRATION_REJECTED"
	StatusCode_NODE_ROLE_CREATED                                                    string = "NODE_ROLE_CREATED"
	StatusCode_NODE_ROLE_DELETED                                                    string = "NODE_ROLE_DELETED"
	StatusCode_NODE_ROLE_UPDATED                                                    string = "NODE_ROLE_UPDATED"
	StatusCode_NODE_UPDATED                                                         string = "NODE_UPDATED"
	StatusCode_PATCH_MANAGEMENT_APPLY_PATCH_COMPLETED                               string = "PATCH_MANAGEMENT_APPLY_PATCH_COMPLETED"
	StatusCode_PATCH_MANAGEMENT_APPLY_PATCH_STARTED                                 string = "PATCH_MANAGEMENT_APPLY_PATCH_STARTED"
	StatusCode_PATCH_MANAGEMENT_FAILURE                                             string = "PATCH_MANAGEMENT_FAILURE"
	StatusCode_PATCH_MANAGEMENT_INSTALL_FAILED                                      string = "PATCH_MANAGEMENT_INSTALL_FAILED"
	StatusCode_PATCH_MANAGEMENT_INSTALLED                                           string = "PATCH_MANAGEMENT_INSTALLED"
	StatusCode_PATCH_MANAGEMENT_MESSAGE                                             string = "PATCH_MANAGEMENT_MESSAGE"
	StatusCode_PATCH_MANAGEMENT_SCAN_COMPLETED                                      string = "PATCH_MANAGEMENT_SCAN_COMPLETED"
	StatusCode_PATCH_MANAGEMENT_SCAN_STARTED                                        string = "PATCH_MANAGEMENT_SCAN_STARTED"
	StatusCode_POLICY_CREATED                                                       string = "POLICY_CREATED"
	StatusCode_POLICY_DELETED                                                       string = "POLICY_DELETED"
	StatusCode_POLICY_UPDATED                                                       string = "POLICY_UPDATED"
	StatusCode_PORT_CLOSED                                                          string = "PORT_CLOSED"
	StatusCode_PORT_OPENED                                                          string = "PORT_OPENED"
	StatusCode_PROCESS_STARTED                                                      string = "PROCESS_STARTED"
	StatusCode_PROCESS_STOPPED                                                      string = "PROCESS_STOPPED"
	StatusCode_PSA_CREDENTIALS_FAILED                                               string = "PSA_CREDENTIALS_FAILED"
	StatusCode_PSA_DISABLED                                                         string = "PSA_DISABLED"
	StatusCode_PSA_ENABLED                                                          string = "PSA_ENABLED"
	StatusCode_PSA_TICKET_CREATION_FAILED                                           string = "PSA_TICKET_CREATION_FAILED"
	StatusCode_PSA_TICKET_CREATION_SUCCEEDED                                        string = "PSA_TICKET_CREATION_SUCCEEDED"
	StatusCode_PSA_TICKET_CREATION_TEST                                             string = "PSA_TICKET_CREATION_TEST"
	StatusCode_RAID_CONTROLLER_ADDED                                                string = "RAID_CONTROLLER_ADDED"
	StatusCode_RAID_CONTROLLER_REMOVED                                              string = "RAID_CONTROLLER_REMOVED"
	StatusCode_RAID_LOGICAL_DISK_ADDED                                              string = "RAID_LOGICAL_DISK_ADDED"
	StatusCode_RAID_LOGICAL_DISK_REMOVED                                            string = "RAID_LOGICAL_DISK_REMOVED"
	StatusCode_RAID_PHYSICAL_DRIVE_ADDED                                            string = "RAID_PHYSICAL_DRIVE_ADDED"
	StatusCode_RAID_PHYSICAL_DRIVE_REMOVED                                          string = "RAID_PHYSICAL_DRIVE_REMOVED"
	StatusCode_RDP_AUTO_PROVISION                                                   string = "RDP_AUTO_PROVISION"
	StatusCode_RDP_CONNECTION_ESTABLISHED                                           string = "RDP_CONNECTION_ESTABLISHED"
	StatusCode_RDP_CONNECTION_INITIATED                                             string = "RDP_CONNECTION_INITIATED"
	StatusCode_RDP_CONNECTION_TERMINATED                                            string = "RDP_CONNECTION_TERMINATED"
	StatusCode_REJECTED_NODE_CLEARED                                                string = "REJECTED_NODE_CLEARED"
	StatusCode_REJECTED_NODE_DELETED                                                string = "REJECTED_NODE_DELETED"
	StatusCode_REMOTE_SUPPORT_CREATED                                               string = "REMOTE_SUPPORT_CREATED"
	StatusCode_REMOTE_SUPPORT_DELETED                                               string = "REMOTE_SUPPORT_DELETED"
	StatusCode_REMOTE_SUPPORT_UPDATED                                               string = "REMOTE_SUPPORT_UPDATED"
	StatusCode_REMOTE_TOOLS_ACTIVE_DIRECTORY_INITIATED                              string = "REMOTE_TOOLS_ACTIVE_DIRECTORY_INITIATED"
	StatusCode_REMOTE_TOOLS_COPY_OBJECT_FAILED                                      string = "REMOTE_TOOLS_COPY_OBJECT_FAILED"
	StatusCode_REMOTE_TOOLS_COPY_OBJECT_SUCCESS                                     string = "REMOTE_TOOLS_COPY_OBJECT_SUCCESS"
	StatusCode_REMOTE_TOOLS_CREATE_DIRECTORY_FAILED                                 string = "REMOTE_TOOLS_CREATE_DIRECTORY_FAILED"
	StatusCode_REMOTE_TOOLS_CREATE_DIRECTORY_INITIATED                              string = "REMOTE_TOOLS_CREATE_DIRECTORY_INITIATED"
	StatusCode_REMOTE_TOOLS_CREATE_DIRECTORY_SUCCESS                                string = "REMOTE_TOOLS_CREATE_DIRECTORY_SUCCESS"
	StatusCode_REMOTE_TOOLS_CREATE_KEY_FAILED                                       string = "REMOTE_TOOLS_CREATE_KEY_FAILED"
	StatusCode_REMOTE_TOOLS_CREATE_KEY_SUCCESS                                      string = "REMOTE_TOOLS_CREATE_KEY_SUCCESS"
	StatusCode_REMOTE_TOOLS_CREATE_PARAMETER_FAILED                                 string = "REMOTE_TOOLS_CREATE_PARAMETER_FAILED"
	StatusCode_REMOTE_TOOLS_CREATE_PARAMETER_SUCCESS                                string = "REMOTE_TOOLS_CREATE_PARAMETER_SUCCESS"
	StatusCode_REMOTE_TOOLS_DELETE_FILE_INITIATED                                   string = "REMOTE_TOOLS_DELETE_FILE_INITIATED"
	StatusCode_REMOTE_TOOLS_DELETE_KEY_FAILED                                       string = "REMOTE_TOOLS_DELETE_KEY_FAILED"
	StatusCode_REMOTE_TOOLS_DELETE_KEY_SUCCESS                                      string = "REMOTE_TOOLS_DELETE_KEY_SUCCESS"
	StatusCode_REMOTE_TOOLS_DELETE_OBJECT_FAILED                                    string = "REMOTE_TOOLS_DELETE_OBJECT_FAILED"
	StatusCode_REMOTE_TOOLS_DELETE_OBJECT_SUCCESS                                   string = "REMOTE_TOOLS_DELETE_OBJECT_SUCCESS"
	StatusCode_REMOTE_TOOLS_DELETE_PARAMETER_FAILED                                 string = "REMOTE_TOOLS_DELETE_PARAMETER_FAILED"
	StatusCode_REMOTE_TOOLS_DELETE_PARAMETER_SUCCESS                                string = "REMOTE_TOOLS_DELETE_PARAMETER_SUCCESS"
	StatusCode_REMOTE_TOOLS_DOWNLOAD_FILE_INITIATED                                 string = "REMOTE_TOOLS_DOWNLOAD_FILE_INITIATED"
	StatusCode_REMOTE_TOOLS_FILE_TRANSFER_FAILED                                    string = "REMOTE_TOOLS_FILE_TRANSFER_FAILED"
	StatusCode_REMOTE_TOOLS_FILE_TRANSFER_SUCCESS                                   string = "REMOTE_TOOLS_FILE_TRANSFER_SUCCESS"
	StatusCode_REMOTE_TOOLS_MODIFY_OBJECT_FAILED                                    string = "REMOTE_TOOLS_MODIFY_OBJECT_FAILED"
	StatusCode_REMOTE_TOOLS_MODIFY_OBJECT_SUCCESS                                   string = "REMOTE_TOOLS_MODIFY_OBJECT_SUCCESS"
	StatusCode_REMOTE_TOOLS_MODIFY_PARAMETER_FAILED                                 string = "REMOTE_TOOLS_MODIFY_PARAMETER_FAILED"
	StatusCode_REMOTE_TOOLS_MODIFY_PARAMETER_SUCCESS                                string = "REMOTE_TOOLS_MODIFY_PARAMETER_SUCCESS"
	StatusCode_REMOTE_TOOLS_MOVE_OBJECT_FAILED                                      string = "REMOTE_TOOLS_MOVE_OBJECT_FAILED"
	StatusCode_REMOTE_TOOLS_MOVE_OBJECT_SUCCESS                                     string = "REMOTE_TOOLS_MOVE_OBJECT_SUCCESS"
	StatusCode_REMOTE_TOOLS_PROCESS_CONTROL_INITIATED                               string = "REMOTE_TOOLS_PROCESS_CONTROL_INITIATED"
	StatusCode_REMOTE_TOOLS_REGISTRY_CONTROL_INITIATED                              string = "REMOTE_TOOLS_REGISTRY_CONTROL_INITIATED"
	StatusCode_REMOTE_TOOLS_RENAME_FILE_INITIATED                                   string = "REMOTE_TOOLS_RENAME_FILE_INITIATED"
	StatusCode_REMOTE_TOOLS_RENAME_KEY_FAILED                                       string = "REMOTE_TOOLS_RENAME_KEY_FAILED"
	StatusCode_REMOTE_TOOLS_RENAME_KEY_SUCCESS                                      string = "REMOTE_TOOLS_RENAME_KEY_SUCCESS"
	StatusCode_REMOTE_TOOLS_RENAME_PARAMETER_FAILED                                 string = "REMOTE_TOOLS_RENAME_PARAMETER_FAILED"
	StatusCode_REMOTE_TOOLS_RENAME_PARAMETER_SUCCESS                                string = "REMOTE_TOOLS_RENAME_PARAMETER_SUCCESS"
	StatusCode_REMOTE_TOOLS_RESTART_SERVICE_FAILED                                  string = "REMOTE_TOOLS_RESTART_SERVICE_FAILED"
	StatusCode_REMOTE_TOOLS_RESTART_SERVICE_SUCCESS                                 string = "REMOTE_TOOLS_RESTART_SERVICE_SUCCESS"
	StatusCode_REMOTE_TOOLS_SERVICE_CONTROL_INITIATED                               string = "REMOTE_TOOLS_SERVICE_CONTROL_INITIATED"
	StatusCode_REMOTE_TOOLS_SET_PROCESS_PRIORITY_FAILED                             string = "REMOTE_TOOLS_SET_PROCESS_PRIORITY_FAILED"
	StatusCode_REMOTE_TOOLS_SET_PROCESS_PRIORITY_SUCESS                             string = "REMOTE_TOOLS_SET_PROCESS_PRIORITY_SUCESS"
	StatusCode_REMOTE_TOOLS_START_SERVICE_FAILED                                    string = "REMOTE_TOOLS_START_SERVICE_FAILED"
	StatusCode_REMOTE_TOOLS_START_SERVICE_SUCCESS                                   string = "REMOTE_TOOLS_START_SERVICE_SUCCESS"
	StatusCode_REMOTE_TOOLS_START_TYPE_CHANGE_FAILED                                string = "REMOTE_TOOLS_START_TYPE_CHANGE_FAILED"
	StatusCode_REMOTE_TOOLS_START_TYPE_CHANGE_SUCCESS                               string = "REMOTE_TOOLS_START_TYPE_CHANGE_SUCCESS"
	StatusCode_REMOTE_TOOLS_STOP_SERVICE_FAILED                                     string = "REMOTE_TOOLS_STOP_SERVICE_FAILED"
	StatusCode_REMOTE_TOOLS_STOP_SERVICE_SUCCESS                                    string = "REMOTE_TOOLS_STOP_SERVICE_SUCCESS"
	StatusCode_REMOTE_TOOLS_TERMINATE_PROCESS_FAILED                                string = "REMOTE_TOOLS_TERMINATE_PROCESS_FAILED"
	StatusCode_REMOTE_TOOLS_TERMINATE_PROCESS_SUCCESS                               string = "REMOTE_TOOLS_TERMINATE_PROCESS_SUCCESS"
	StatusCode_REMOTE_TOOLS_TERMINATE_PROCESS_TREE_FAILED                           string = "REMOTE_TOOLS_TERMINATE_PROCESS_TREE_FAILED"
	StatusCode_REMOTE_TOOLS_TERMINATE_PROCESS_TREE_SUCCESS                          string = "REMOTE_TOOLS_TERMINATE_PROCESS_TREE_SUCCESS"
	StatusCode_REMOTE_TOOLS_UPLOAD_FILE_INITIATED                                   string = "REMOTE_TOOLS_UPLOAD_FILE_INITIATED"
	StatusCode_REPORT_CREATED                                                       string = "REPORT_CREATED"
	StatusCode_REPORT_DELETED                                                       string = "REPORT_DELETED"
	StatusCode_REPORT_UPDATED                                                       string = "REPORT_UPDATED"
	StatusCode_RESET                                                                string = "RESET"
	StatusCode_RESET_BY_PSA_TICKET_CALLBACK                                         string = "RESET_BY_PSA_TICKET_CALLBACK"
	StatusCode_SCHEDULE_INSTALL_OPTION_CHANGED                                      string = "SCHEDULE_INSTALL_OPTION_CHANGED"
	StatusCode_SCHEDULED_TASK_CREATED                                               string = "SCHEDULED_TASK_CREATED"
	StatusCode_SCHEDULED_TASK_DELETED                                               string = "SCHEDULED_TASK_DELETED"
	StatusCode_SCHEDULED_TASK_UPDATED                                               string = "SCHEDULED_TASK_UPDATED"
	StatusCode_SCRIPT_CREATED                                                       string = "SCRIPT_CREATED"
	StatusCode_SCRIPT_DELETED                                                       string = "SCRIPT_DELETED"
	StatusCode_SCRIPT_UPDATED                                                       string = "SCRIPT_UPDATED"
	StatusCode_SECURITY_CREDENTIAL_ACCESS_DENIED                                    string = "SECURITY_CREDENTIAL_ACCESS_DENIED"
	StatusCode_SECURITY_CREDENTIAL_ACCESS_GRANTED                                   string = "SECURITY_CREDENTIAL_ACCESS_GRANTED"
	StatusCode_SERVER_MESSAGE                                                       string = "SERVER_MESSAGE"
	StatusCode_SHADOWPROTECT_BACKUPJOB_ABORTED                                      string = "SHADOWPROTECT_BACKUPJOB_ABORTED"
	StatusCode_SHADOWPROTECT_BACKUPJOB_FAILED                                       string = "SHADOWPROTECT_BACKUPJOB_FAILED"
	StatusCode_SHADOWPROTECT_INSTALL_FAILED                                         string = "SHADOWPROTECT_INSTALL_FAILED"
	StatusCode_SHADOWPROTECT_INSTALLED                                              string = "SHADOWPROTECT_INSTALLED"
	StatusCode_SHADOWPROTECT_LICENSE_ACTIVATED                                      string = "SHADOWPROTECT_LICENSE_ACTIVATED"
	StatusCode_SHADOWPROTECT_LICENSE_ACTIVATION_FAILED                              string = "SHADOWPROTECT_LICENSE_ACTIVATION_FAILED"
	StatusCode_SHADOWPROTECT_LICENSE_DEACTIVATED                                    string = "SHADOWPROTECT_LICENSE_DEACTIVATED"
	StatusCode_SHADOWPROTECT_LICENSE_DEACTIVATION_FAILED                            string = "SHADOWPROTECT_LICENSE_DEACTIVATION_FAILED"
	StatusCode_SHADOWPROTECT_LICENSE_PROVISION_FAILED                               string = "SHADOWPROTECT_LICENSE_PROVISION_FAILED"
	StatusCode_SHADOWPROTECT_LICENSE_PROVISIONED                                    string = "SHADOWPROTECT_LICENSE_PROVISIONED"
	StatusCode_SHADOWPROTECT_UNINSTALL_FAILED                                       string = "SHADOWPROTECT_UNINSTALL_FAILED"
	StatusCode_SHADOWPROTECT_UNINSTALLED                                            string = "SHADOWPROTECT_UNINSTALLED"
	StatusCode_SOFTWARE_ADDED                                                       string = "SOFTWARE_ADDED"
	StatusCode_SOFTWARE_PATCH_MANAGEMENT_APPLY_PATCH_COMPLETED                      string = "SOFTWARE_PATCH_MANAGEMENT_APPLY_PATCH_COMPLETED"
	StatusCode_SOFTWARE_PATCH_MANAGEMENT_APPLY_PATCH_STARTED                        string = "SOFTWARE_PATCH_MANAGEMENT_APPLY_PATCH_STARTED"
	StatusCode_SOFTWARE_PATCH_MANAGEMENT_INSTALL_FAILED                             string = "SOFTWARE_PATCH_MANAGEMENT_INSTALL_FAILED"
	StatusCode_SOFTWARE_PATCH_MANAGEMENT_INSTALLED                                  string = "SOFTWARE_PATCH_MANAGEMENT_INSTALLED"
	StatusCode_SOFTWARE_PATCH_MANAGEMENT_MESSAGE                                    string = "SOFTWARE_PATCH_MANAGEMENT_MESSAGE"
	StatusCode_SOFTWARE_PATCH_MANAGEMENT_SCAN_COMPLETED                             string = "SOFTWARE_PATCH_MANAGEMENT_SCAN_COMPLETED"
	StatusCode_SOFTWARE_PATCH_MANAGEMENT_SCAN_STARTED                               string = "SOFTWARE_PATCH_MANAGEMENT_SCAN_STARTED"
	StatusCode_SOFTWARE_REMOVED                                                     string = "SOFTWARE_REMOVED"
	StatusCode_SPLASHTOP_CONNECTION_ESTABLISHED                                     string = "SPLASHTOP_CONNECTION_ESTABLISHED"
	StatusCode_SPLASHTOP_CONNECTION_INITIATED                                       string = "SPLASHTOP_CONNECTION_INITIATED"
	StatusCode_SPLASHTOP_CONNECTION_TERMINATED                                      string = "SPLASHTOP_CONNECTION_TERMINATED"
	StatusCode_START_REQUESTED                                                      string = "START_REQUESTED"
	StatusCode_STARTED                                                              string = "STARTED"
	StatusCode_SYSTEM_REBOOTED                                                      string = "SYSTEM_REBOOTED"
	StatusCode_TEAMVIEWER_ACCOUNTNAME_ADDED                                         string = "TEAMVIEWER_ACCOUNTNAME_ADDED"
	StatusCode_TEAMVIEWER_ACCOUNTNAME_CHANGED                                       string = "TEAMVIEWER_ACCOUNTNAME_CHANGED"
	StatusCode_TEAMVIEWER_ACCOUNTNAME_REMOVED                                       string = "TEAMVIEWER_ACCOUNTNAME_REMOVED"
	StatusCode_TEAMVIEWER_CONFIG_CHANGED                                            string = "TEAMVIEWER_CONFIG_CHANGED"
	StatusCode_TEAMVIEWER_CONNECTION_CANCELLED                                      string = "TEAMVIEWER_CONNECTION_CANCELLED"
	StatusCode_TEAMVIEWER_CONNECTION_ESTABLISHED                                    string = "TEAMVIEWER_CONNECTION_ESTABLISHED"
	StatusCode_TEAMVIEWER_CONNECTION_TERMINATED                                     string = "TEAMVIEWER_CONNECTION_TERMINATED"
	StatusCode_TEAMVIEWER_INSTALL_FAILED                                            string = "TEAMVIEWER_INSTALL_FAILED"
	StatusCode_TEAMVIEWER_INSTALLED                                                 string = "TEAMVIEWER_INSTALLED"
	StatusCode_TEAMVIEWER_PERMANENT_PASSWORD_CHANGED                                string = "TEAMVIEWER_PERMANENT_PASSWORD_CHANGED"
	StatusCode_TEAMVIEWER_UNINSTALL_FAILED                                          string = "TEAMVIEWER_UNINSTALL_FAILED"
	StatusCode_TEAMVIEWER_UNINSTALLED                                               string = "TEAMVIEWER_UNINSTALLED"
	StatusCode_TICKET_TEMPLATE_CREATED                                              string = "TICKET_TEMPLATE_CREATED"
	StatusCode_TICKET_TEMPLATE_DELETED                                              string = "TICKET_TEMPLATE_DELETED"
	StatusCode_TICKET_TEMPLATE_UPDATED                                              string = "TICKET_TEMPLATE_UPDATED"
	StatusCode_TIME_ZONE_UPDATED                                                    string = "TIME_ZONE_UPDATED"
	StatusCode_TRIGGERED                                                            string = "TRIGGERED"
	StatusCode_TRUSTED_PLATFORM_MODULE_DISABLED                                     string = "TRUSTED_PLATFORM_MODULE_DISABLED"
	StatusCode_TRUSTED_PLATFORM_MODULE_ENABLED                                      string = "TRUSTED_PLATFORM_MODULE_ENABLED"
	StatusCode_TRUSTED_PLATFORM_MODULE_INSTALLED                                    string = "TRUSTED_PLATFORM_MODULE_INSTALLED"
	StatusCode_TRUSTED_PLATFORM_MODULE_UNINSTALLED                                  string = "TRUSTED_PLATFORM_MODULE_UNINSTALLED"
	StatusCode_UNKNOWN                                                              string = "UNKNOWN"
	StatusCode_USER_ACCOUNT_ADDED                                                   string = "USER_ACCOUNT_ADDED"
	StatusCode_USER_ACCOUNT_REMOVED                                                 string = "USER_ACCOUNT_REMOVED"
	StatusCode_USER_LOGGED_IN                                                       string = "USER_LOGGED_IN"
	StatusCode_USER_LOGGED_OUT                                                      string = "USER_LOGGED_OUT"
	StatusCode_VIPREAV_ACTIVEPROTECTION_THREAT_QUARANTINED                          string = "VIPREAV_ACTIVEPROTECTION_THREAT_QUARANTINED"
	StatusCode_VIPREAV_DISABLED                                                     string = "VIPREAV_DISABLED"
	StatusCode_VIPREAV_INSTALL_FAILED                                               string = "VIPREAV_INSTALL_FAILED"
	StatusCode_VIPREAV_INSTALLED                                                    string = "VIPREAV_INSTALLED"
	StatusCode_VIPREAV_QUARANTINED_THREAT_REMOVED                                   string = "VIPREAV_QUARANTINED_THREAT_REMOVED"
	StatusCode_VIPREAV_REBOOT_REQUIRED                                              string = "VIPREAV_REBOOT_REQUIRED"
	StatusCode_VIPREAV_SCAN_ABORTED                                                 string = "VIPREAV_SCAN_ABORTED"
	StatusCode_VIPREAV_SCAN_COMPLETED                                               string = "VIPREAV_SCAN_COMPLETED"
	StatusCode_VIPREAV_SCAN_FAILED                                                  string = "VIPREAV_SCAN_FAILED"
	StatusCode_VIPREAV_SCAN_PAUSED                                                  string = "VIPREAV_SCAN_PAUSED"
	StatusCode_VIPREAV_SCAN_STARTED                                                 string = "VIPREAV_SCAN_STARTED"
	StatusCode_VIPREAV_SCAN_THREAT_QUARANTINED                                      string = "VIPREAV_SCAN_THREAT_QUARANTINED"
	StatusCode_VIPREAV_UNINSTALL_FAILED                                             string = "VIPREAV_UNINSTALL_FAILED"
	StatusCode_VIPREAV_UNINSTALLED                                                  string = "VIPREAV_UNINSTALLED"
	StatusCode_VIPREAV_USERINITIATED_THREAT_QUARANTINED                             string = "VIPREAV_USERINITIATED_THREAT_QUARANTINED"
	StatusCode_WEBAPP_MESSAGE                                                       string = "WEBAPP_MESSAGE"
	StatusCode_WEBROOT_COMMAND_SUBMITTED                                            string = "WEBROOT_COMMAND_SUBMITTED"
	StatusCode_WEBROOT_INSTALL_FAILED                                               string = "WEBROOT_INSTALL_FAILED"
	StatusCode_WEBROOT_THREAT_DETECTED                                              string = "WEBROOT_THREAT_DETECTED"
	StatusCode_WINDOWS_SERVICE_STARTED                                              string = "WINDOWS_SERVICE_STARTED"
	StatusCode_WINDOWS_SERVICE_STOPPED                                              string = "WINDOWS_SERVICE_STOPPED"
)

const (
	ActivityResult_SUCCESS     string = "SUCCESS"
	ActivityResult_FAILURE     string = "FAILURE"
	ActivityResult_UNSUPPORTED string = "UNSUPPORTED"
	ActivityResult_UNCOMPLETED string = "UNCOMPLETED"
)

const (
	DeviceApprovalMode_AUTOMATIC string = "AUTOMATIC"
	DeviceApprovalMode_MANUAL    string = "MANUAL"
	DeviceApprovalMode_REJECT    string = "REJECT"
)

type BaseObject struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PolicyRef struct {
	NodeRoleID int `json:"nodeRoleId"`
	PolicyID   int `json:"policyId"`
}

type SettingBase struct {
	Enabled bool `json:"enabled"`
}

type TeamviewerSetting struct {
	SettingBase
	Product string   `json:"product"`
	Targets []string `json:"targets"`
}

type OrganizationSettings struct {
	TrayIcon   SettingBase       `json:"trayicon"`
	Splashtop  SettingBase       `json:"splashtop"`
	Teamviewer TeamviewerSetting `json:"teamviewer"`
	PSA        SettingBase       `json:"psa"`
}

type Location struct {
	BaseObject
	Address     string      `json:"address,omitempty"`
	Description string      `json:"description,omitempty"`
	Tags        []string    `json:"tags"`
	Fields      interface{} `json:"fields,omitempty"`
}

type Organization struct {
	BaseObject
	Locations []Location           `json:"locations"`
	Policies  []PolicyRef          `json:"policies"`
	Settings  OrganizationSettings `json:"settings"`
}

type OrganizationSummary struct {
	BaseObject
	Tags             []string    `json:"tags"`
	Description      string      `json:"description,omitempty"`
	Fields           interface{} `json:"fields,omitempty"`
	UserData         interface{} `json:"userData,omitempty"`
	NodeApprovalMode string      `json:"nodeApprovalMode,omitempty"`
}

type Activity struct {
	ID              json.Number `json:"id"`
	ActivityTime    string      `json:"activityTime"`
	DeviceID        json.Number `json:"deviceId"`
	Severity        string      `json:"severity"`
	Priority        string      `json:"priority"`
	ActivityType    string      `json:"activityType"`
	StatusCode      string      `json:"statusCode"`
	Status          string      `json:"status"`
	ActivityResult  string      `json:"activityResult"`
	SourceConfigUID string      `json:"sourceConfigUid"`
	SourceName      string      `json:"sourceName"`
	Subject         string      `json:"subject"`
	UserID          json.Number `json:"userId"`
	Message         string      `json:"message"`
	Type            string      `json:"type"`
	Data            interface{} `json:"data"`
}

type DeviceOS struct {
	Manufacturer            string      `json:"manufacturer"`
	Name                    string      `json:"name"`
	Architecture            string      `json:"architecture"`
	LastBootTime            json.Number `json:"lastBootTime"`
	BuildNumber             string      `json:"buildNumber"`
	ReleaseID               string      `json:"releaseId"`
	ServicePackMajorVersion json.Number `json:"servicePackMajorVersion"`
	ServicePackMinorVersion json.Number `json:"servicePackMinorVersion"`
	Locale                  string      `json:"locale"`
	Language                string      `json:"language"`
	NeedsReboot             bool        `json:"needsReboot"`
}

type DeviceSystem struct {
	Name                string      `json:"name"`
	Manufacturer        string      `json:"manufacturer"`
	Model               string      `json:"model"`
	BiosSerialNumber    string      `json:"biosSerialNumber"`
	SerialNumber        string      `json:"serialNumber"`
	Domain              string      `json:"domain"`
	DomainRole          string      `json:"domainRole"`
	NumberOfProcessors  json.Number `json:"numberOfProcessors"`
	TotalPhysicalMemory json.Number `json:"totalPhysicalMemory"`
	VirtualMachine      bool        `json:"virtualMachine"`
	ChassisType         string      `json:"chassisType"`
}

type DeviceMemory struct {
	Capacity json.Number `json:"capacity"`
}

type DeviceProcessor struct {
	Architecture    string      `json:"architecture"`
	MaxClockSpeed   json.Number `json:"maxClockSpeed"`
	ClockSpeed      json.Number `json:"clockSpeed"`
	Name            string      `json:"name"`
	NumCores        json.Number `json:"numCores"`
	NumLogicalCores json.Number `json:"numLogicalCores"`
}

type DeviceVolume struct {
	Name         string      `json:"name"`
	Label        string      `json:"label"`
	DeviceType   string      `json:"deviceType"`
	FileSystem   string      `json:"fileSystem"`
	AutoMount    bool        `json:"autoMount"`
	Compressed   bool        `json:"compressed"`
	Capacity     json.Number `json:"capacity"`
	SerialNumber string      `json:"serialNumber"`
}

type Device struct {
	ID             int         `json:"id"`
	OrganizationID int         `json:"organizationId"`
	LocationID     int         `json:"locationId"`
	NodeClass      string      `json:"nodeClass"`
	NodeRoleId     int         `json:"nodeRoleId"`
	RolePolicyID   int         `json:"rolePolicyId"`
	PolicyID       int         `json:"policyId"`
	ApprovalStatus string      `json:"approvalStatus"`
	Offline        bool        `json:"offline"`
	DisplayName    string      `json:"displayName"`
	SystemName     string      `json:"systemName"`
	DNSName        string      `json:"dnsName"`
	Created        json.Number `json:"created"`
	LastContact    json.Number `json:"lastContact"`
	LastUpdate     json.Number `json:"lastUpdate"`
}

type DeviceDetails struct {
	Device
	IPAddresses      []string          `json:"ipAddresses"`
	MACAddresses     []string          `json:"macAddresses"`
	PublicIP         string            `json:"publicIP"`
	OS               DeviceOS          `json:"os"`
	System           DeviceSystem      `json:"system"`
	Memory           DeviceMemory      `json:"memory"`
	Processors       []DeviceProcessor `json:"processors"`
	Volumes          []DeviceVolume    `json:"volumes"`
	LastLoggedInUser string            `json:"lastLoggedInUser"`
	DeviceType       string            `json:"deviceType"`
}

type OSPatchReport struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Severity  string      `json:"severity"`
	Status    string      `json:"status"`
	Type      string      `json:"type"`
	KBNumber  string      `json:"kbNumber"`
	DeviceID  int         `json:"deviceId"`
	Timestamp json.Number `json:"timestamp"`
}

type OSPatchReportQuery struct {
	Results []OSPatchReport `json:"results"`
}

type OSPatchReportDetail struct {
	ID        string        `json:"id"`
	Name      string        `json:"name"`
	Severity  string        `json:"severity"`
	Status    string        `json:"status"`
	Type      string        `json:"type"`
	KBNumber  string        `json:"kbNumber"`
	Timestamp json.Number   `json:"timestamp"`
	Device    DeviceDetails `json:"device"`
}

type WebhookBase struct {
	ID              int         `json:"id"`
	ActivityTime    json.Number `json:"activityTime"`
	DeviceID        int         `json:"deviceId"`
	Severity        string      `json:"severity"`
	Priority        string      `json:"priority"`
	SeriesUID       string      `json:"seriesUid"`
	ActivityType    string      `json:"activityType"`
	ActivityResult  string      `json:"activityResult,omitempty"`
	StatusCode      string      `json:"statusCode"`
	Status          string      `json:"status"`
	SourceConfigUID string      `json:"sourceConfigUid"`
	SourceName      string      `json:"sourceName"`
	Message         string      `json:"message"`
	Type            string      `json:"type"`
	Data            interface{} `json:"data"`
}

type Webhook struct {
	WebhookBase
	Device Device `json:"device"`
}
