// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package types

import "encoding/json"

type BillingGetOResponse struct {
	Items json.RawMessage `json:"items"`
}

type DeviceGetResponse struct {
	Benchmarks json.RawMessage `json:"benchmarks"`
}

type DeviceScriptAssignmentsListResponse struct {
	Items json.RawMessage `json:"items"`
}

type DevicesCreateResponse struct {
	Explanation string          `json:"explanation"`
	Filters     json.RawMessage `json:"filters"`
}

type FactsCreateCustom2Response struct {
	Items    json.RawMessage `json:"items"`
	Metadata json.RawMessage `json:"metadata"`
}

type FeatureBetasListResponse struct {
	Items    json.RawMessage `json:"items"`
	Metadata json.RawMessage `json:"metadata"`
}

type IntegrationsGetResponse struct {
	Items json.RawMessage `json:"items"`
}

type MaintenanceCreateQueryResponse struct {
	Items    json.RawMessage `json:"items"`
	Metadata json.RawMessage `json:"metadata"`
}

type MaintenanceCreateStaged3Response struct {
	Items    json.RawMessage `json:"items"`
	Metadata json.RawMessage `json:"metadata"`
}

type ManagedAppConfigurationsCreateResponse struct {
}

type ManagedAppConfigurationsListResponse struct {
}

type MdmCreateConfigurations3Response struct {
	Payloads json.RawMessage `json:"payloads"`
}

type MdmCreateConfigurationsResponse struct {
	Payloads json.RawMessage `json:"payloads"`
}

type MdmUpdateResponse struct {
	Payloads json.RawMessage `json:"payloads"`
}

type MonitoringCreateQueryResponse struct {
	Items    json.RawMessage `json:"items"`
	Metadata json.RawMessage `json:"metadata"`
}

type OaCreateCompliancerulesResponse struct {
	Items    json.RawMessage `json:"items"`
	Metadata json.RawMessage `json:"metadata"`
}

type OaCreateDevicesResponse struct {
	Benchmarks json.RawMessage `json:"benchmarks"`
}

type OaCreateFilesResponse struct {
	Items json.RawMessage `json:"items"`
}

type OaCreateIdentityResponse struct {
	Items    json.RawMessage `json:"items"`
	Metadata json.RawMessage `json:"metadata"`
}

type OaCreateIntegrationsResponse struct {
	Items json.RawMessage `json:"items"`
}

type OaCreateReportsResponse struct {
	Items    json.RawMessage `json:"items"`
	Metadata json.RawMessage `json:"metadata"`
}

type OaCreateResponse struct {
	Items    json.RawMessage `json:"items"`
	Metadata json.RawMessage `json:"metadata"`
}

type OaCreateVariablesResponse struct {
	Items    json.RawMessage `json:"items"`
	Metadata json.RawMessage `json:"metadata"`
}

type OaCreateWebhooks3Response struct {
	Items    json.RawMessage `json:"items"`
	Metadata json.RawMessage `json:"metadata"`
}

type OaCreateWebhooksResponse struct {
	Items    json.RawMessage `json:"items"`
	Metadata json.RawMessage `json:"metadata"`
}

type OaListBenchmarksResponse struct {
	Benchmarks json.RawMessage `json:"benchmarks"`
}

type OaListCompliancerules2Response struct {
	Benchmarks json.RawMessage `json:"benchmarks"`
}

type OaListCompliancerulesResponse struct {
	Benchmarks json.RawMessage `json:"benchmarks"`
}

type OaListIntegrationsResponse struct {
	Items json.RawMessage `json:"items"`
}

type OaListReports2Response struct {
	Items json.RawMessage `json:"items"`
}

type StaticFieldsListResponse struct {
	Items    json.RawMessage `json:"items"`
	Metadata json.RawMessage `json:"metadata"`
}

type StaticFieldsListStaticfieldsResponse struct {
	Items    json.RawMessage `json:"items"`
	Metadata json.RawMessage `json:"metadata"`
}

type VariablesGetItem struct {
	OrganizationId string          `json:"organization_id"`
	PolicyValues   json.RawMessage `json:"policy_values"`
	UpdatedDate    string          `json:"updated_date"`
	VariableKey    string          `json:"variable_key"`
}

type VariablesGetO2Response struct {
	Value string `json:"value"`
}

type VariablesGetO3Response struct {
	Value string `json:"value"`
}

type addigy_errors_response struct {
	ErrorChain      json.RawMessage `json:"error_chain"`
	InternalMessage string          `json:"internal_message"`
	Message         string          `json:"message"`
	Origin          string          `json:"origin"`
}

type ade_account struct {
	AdminId    string          `json:"admin_id"`
	OrgAddress string          `json:"org_address"`
	OrgEmail   string          `json:"org_email"`
	OrgName    string          `json:"org_name"`
	OrgPhone   string          `json:"org_phone"`
	ServerName string          `json:"server_name"`
	ServerUuid string          `json:"server_uuid"`
	Urls       json.RawMessage `json:"urls"`
}

type alert_entities_Monitoring struct {
	Category           string          `json:"category"`
	Emails             json.RawMessage `json:"emails"`
	Fact               string          `json:"fact"`
	FactIdentifier     string          `json:"fact_identifier"`
	HasScript          bool            `json:"has_script"`
	Id                 string          `json:"id"`
	Instructions       json.RawMessage `json:"instructions"`
	IsInBlueprint      bool            `json:"is_in_blueprint"`
	Level              string          `json:"level"`
	MaxValue           float64         `json:"max_value"`
	MinValue           float64         `json:"min_value"`
	Name               string          `json:"name"`
	PolicyRestricted   bool            `json:"policy_restricted"`
	Provider           string          `json:"provider"`
	RemediationEnabled bool            `json:"remediation_enabled"`
	RemediationTime    int             `json:"remediation_time"`
	Script             json.RawMessage `json:"script"`
	ScriptId           string          `json:"script_id"`
	Selector           string          `json:"selector"`
	SendTicket         bool            `json:"send_ticket"`
	Source             string          `json:"source"`
	SourceId           string          `json:"source_id"`
	Value              string          `json:"value"`
	ValueType          json.RawMessage `json:"value_type"`
	Version            int             `json:"version"`
}

type alert_entities_MonitoringQueryRequest struct {
	ExcludedIds   json.RawMessage `json:"excluded_ids"`
	Ids           json.RawMessage `json:"ids"`
	Limit         int             `json:"limit"`
	NameContains  string          `json:"name_contains"`
	Skip          int             `json:"skip"`
	SortDirection string          `json:"sort_direction"`
	SortField     string          `json:"sort_field"`
}

type alert_entities_MonitoringQueryResponse struct {
	AlertPolicies json.RawMessage `json:"alert_policies"`
	Alerts        json.RawMessage `json:"alerts"`
	StagedAlerts  json.RawMessage `json:"staged_alerts"`
}

type alert_entities_ScheduledAlert struct {
	Category           string          `json:"category"`
	Emails             json.RawMessage `json:"emails"`
	Fact               string          `json:"fact"`
	FactIdentifier     string          `json:"fact_identifier"`
	HasScript          bool            `json:"has_script"`
	Instructions       json.RawMessage `json:"instructions"`
	Level              string          `json:"level"`
	MaxValue           float64         `json:"max_value"`
	MinValue           float64         `json:"min_value"`
	Name               string          `json:"name"`
	Orgid              string          `json:"orgid"`
	PolicyRestricted   bool            `json:"policy_restricted"`
	Provider           string          `json:"provider"`
	RemediationEnabled bool            `json:"remediation_enabled"`
	RemediationTime    int             `json:"remediation_time"`
	Script             json.RawMessage `json:"script"`
	ScriptId           string          `json:"script_id"`
	Selector           string          `json:"selector"`
	SendTicket         bool            `json:"send_ticket"`
	Value              string          `json:"value"`
	ValueType          json.RawMessage `json:"value_type"`
}

type alert_entities_StagedAlertRequest struct {
	Id string `json:"id"`
}

type apps_and_books_token_client_config struct {
	CountryISO2ACode            string          `json:"countryISO2ACode"`
	DefaultPlatform             string          `json:"defaultPlatform"`
	LocationName                bool            `json:"locationName"`
	MdmInfo                     json.RawMessage `json:"mdmInfo"`
	SubscribedNotificationTypes json.RawMessage `json:"subscribedNotificationTypes"`
	UId                         string          `json:"uId"`
	WebsiteURL                  bool            `json:"websiteURL"`
}

type apps_and_books_token_client_config_mdm_info struct {
	Id       string `json:"id"`
	Metadata string `json:"metadata"`
	Name     string `json:"name"`
}

type available_report struct {
	Fields     json.RawMessage `json:"fields"`
	Properties string          `json:"properties"`
}

type azure_ca_account_metadata struct {
	AzurePartnerComplianceGroupId string `json:"azure_partner_compliance_group_id"`
	CreationDate                  string `json:"creation_date"`
	CustomerEnrollmentUrl         string `json:"customer_enrollment_url"`
	CustomerTenantId              string `json:"customer_tenant_id"`
	Enabled                       bool   `json:"enabled"`
	LastProvisionDate             string `json:"last_provision_date"`
	OrganizationId                string `json:"organization_id"`
	PolicyId                      string `json:"policy_id"`
}

type benchmark struct {
	ComplianceRules    json.RawMessage `json:"compliance_rules"`
	ComplianceRulesIds json.RawMessage `json:"compliance_rules_ids"`
	CreatedDate        string          `json:"created_date"`
	Id                 string          `json:"id"`
	OrganizationId     string          `json:"organization_id"`
	UpdatedDate        string          `json:"updated_date"`
}

type benchmark_clone_request struct {
	BenchmarkId        string          `json:"benchmark_id"`
	ComplianceRulesIds json.RawMessage `json:"compliance_rules_ids"`
	Name               string          `json:"name"`
}

type benchmark_create_request struct {
	ComplianceRulesIds json.RawMessage `json:"compliance_rules_ids"`
	MaximumOsVersion   string          `json:"maximum_os_version"`
	MinimumOsVersion   string          `json:"minimum_os_version"`
	Name               string          `json:"name"`
	TargetOs           string          `json:"target_os"`
}

type benchmark_update_request struct {
	ComplianceRulesIds json.RawMessage `json:"compliance_rules_ids"`
	Id                 string          `json:"id"`
	MaximumOsVersion   string          `json:"maximum_os_version"`
	MinimumOsVersion   string          `json:"minimum_os_version"`
	Name               string          `json:"name"`
	TargetOs           string          `json:"target_os"`
}

type certificate_entities_CertificatesRequest struct {
	Page    int             `json:"page"`
	PerPage int             `json:"per_page"`
	Query   json.RawMessage `json:"query"`
}

type certificate_entities_certificates_request struct {
	Page    int             `json:"page"`
	PerPage int             `json:"per_page"`
	Query   json.RawMessage `json:"query"`
}

type compliance_failed_rule struct {
	Error                string `json:"error"`
	RemediationCommandId string `json:"remediation_command_id"`
	RuleId               string `json:"rule_id"`
}

type compliance_rule struct {
	AgentRemediationScriptId string          `json:"agent_remediation_script_id"`
	CreatedDate              string          `json:"created_date"`
	FilterSets               json.RawMessage `json:"filter_sets"`
	Id                       string          `json:"id"`
	Name                     string          `json:"name"`
	OrganizationId           string          `json:"organization_id"`
	RemediationEnabled       bool            `json:"remediation_enabled"`
	UpdatedDate              string          `json:"updated_date"`
}

type compliance_rule_create_request struct {
	AgentRemediationScriptId string          `json:"agent_remediation_script_id"`
	FilterSets               json.RawMessage `json:"filter_sets"`
	Name                     string          `json:"name"`
	RemediationEnabled       bool            `json:"remediation_enabled"`
}

type compliance_rule_update_request struct {
	AgentRemediationScriptId string          `json:"agent_remediation_script_id"`
	FilterSets               json.RawMessage `json:"filter_sets"`
	Id                       string          `json:"id"`
	Name                     string          `json:"name"`
	RemediationEnabled       bool            `json:"remediation_enabled"`
}

type configuration_profile_definition struct {
	AddigyPayloadType  string `json:"addigy_payload_type"`
	PayloadDescription string `json:"payload_description"`
	PayloadName        string `json:"payload_name"`
	PayloadObject      string `json:"payload_object"`
	PayloadType        string `json:"payload_type"`
}

type configuration_profile_entities_Definition struct {
	AddigyPayloadType  string `json:"addigy_payload_type"`
	PayloadDescription string `json:"payload_description"`
	PayloadName        string `json:"payload_name"`
	PayloadObject      string `json:"payload_object"`
	PayloadType        string `json:"payload_type"`
}

type create_response struct {
	Id string `json:"id"`
}

type ddm_system_updates_statuses_response struct {
	BuildVersion  string          `json:"build_version"`
	DeclarationId string          `json:"declaration_id"`
	FailureReason string          `json:"failure_reason"`
	InstallState  string          `json:"install_state"`
	LastUpdated   string          `json:"last_updated"`
	OrgId         string          `json:"org_id"`
	OsVersion     string          `json:"os_version"`
	Reason        json.RawMessage `json:"reason"`
	Udid          string          `json:"udid"`
}

type default_asset_alert struct {
	Category           string          `json:"category"`
	Fact               string          `json:"fact"`
	FactIdentifier     string          `json:"fact_identifier"`
	HasScript          bool            `json:"has_script"`
	Id                 string          `json:"id"`
	Instructions       json.RawMessage `json:"instructions"`
	Level              string          `json:"level"`
	MaxValue           float64         `json:"max_value"`
	MinValue           float64         `json:"min_value"`
	Name               string          `json:"name"`
	Provider           string          `json:"provider"`
	RemediationEnabled bool            `json:"remediation_enabled"`
	RemediationTime    int             `json:"remediation_time"`
	Script             json.RawMessage `json:"script"`
	ScriptId           string          `json:"script_id"`
	Selector           string          `json:"selector"`
	SendTicket         bool            `json:"send_ticket"`
	ShortDescription   string          `json:"short_description"`
	Value              string          `json:"value"`
	ValueType          json.RawMessage `json:"value_type"`
}

type default_asset_alert_query struct {
	Ids          json.RawMessage `json:"ids"`
	NameContains string          `json:"name_contains"`
}

type default_asset_maintenance struct {
	Day              string          `json:"day"`
	Enabled          bool            `json:"enabled"`
	Frequency        string          `json:"frequency"`
	Id               string          `json:"id"`
	Instructions     json.RawMessage `json:"instructions"`
	JobName          string          `json:"job_name"`
	MaxTryCount      int             `json:"max_try_count"`
	PromptUser       bool            `json:"prompt_user"`
	ScheduledTime    string          `json:"scheduled_time"`
	ShortDescription string          `json:"short_description"`
}

type default_asset_maintenance_query struct {
	Ids          json.RawMessage `json:"ids"`
	NameContains string          `json:"name_contains"`
}

type default_asset_mdm_configuration struct {
	AddigyPayloadType          string `json:"addigy_payload_type"`
	AddigyPayloadVersion       int    `json:"addigy_payload_version"`
	PayloadDisplayName         string `json:"payload_display_name"`
	PayloadEnabled             bool   `json:"payload_enabled"`
	PayloadGroupId             string `json:"payload_group_id"`
	PayloadIdentifier          string `json:"payload_identifier"`
	PayloadType                string `json:"payload_type"`
	PayloadUuid                string `json:"payload_uuid"`
	PayloadVersion             int    `json:"payload_version"`
	PolicyRestricted           bool   `json:"policy_restricted"`
	RequiresDeviceSupervision  bool   `json:"requires_device_supervision"`
	RequiresMdmProfileApproved bool   `json:"requires_mdm_profile_approved"`
}

type default_asset_mdm_configuration_query struct {
	Ids          json.RawMessage `json:"ids"`
	NameContains string          `json:"name_contains"`
}

type default_asset_self_service_configuration struct {
	AppLogo                    json.RawMessage `json:"app_logo"`
	DockIcon                   json.RawMessage `json:"dock_icon"`
	FilevaultPromptText        string          `json:"filevault_prompt_text"`
	HideChat                   string          `json:"hide_chat"`
	HomeScreenAddress          string          `json:"home_screen_address"`
	HomeScreenCompanyName      string          `json:"home_screen_company_name"`
	HomeScreenConfigureDetails string          `json:"home_screen_configure_details"`
	HomeScreenDescription      string          `json:"home_screen_description"`
	HomeScreenEmail            string          `json:"home_screen_email"`
	HomeScreenPhone            string          `json:"home_screen_phone"`
	HomeScreenShowAddress      string          `json:"home_screen_show_address"`
	HomeScreenShowDescription  string          `json:"home_screen_show_description"`
	HomeScreenShowEmail        string          `json:"home_screen_show_email"`
	HomeScreenShowPhone        string          `json:"home_screen_show_phone"`
	InstructionId              string          `json:"instruction_id"`
	IntegrationIntuneEnabled   string          `json:"integration_intune_enabled"`
	MaintenancePromptText      string          `json:"maintenance_prompt_text"`
	MenubarIcon                json.RawMessage `json:"menubar_icon"`
	MsOfficeUpdatesPromptText  string          `json:"ms_office_updates_prompt_text"`
	Name                       string          `json:"name"`
	ScreenviewPromptText       string          `json:"screenview_prompt_text"`
	ShowDockIcon               string          `json:"show_dock_icon"`
	ShowInApplications         string          `json:"show_in_applications"`
	ShowMenubarIcon            string          `json:"show_menubar_icon"`
	ShowSupport                string          `json:"show_support"`
	UserSentimentPromptText    string          `json:"user_sentiment_prompt_text"`
}

type default_asset_self_service_query struct {
	Ids          json.RawMessage `json:"ids"`
	NameContains string          `json:"name_contains"`
}

type default_maintenance_instruction struct {
	Instruction string `json:"instruction"`
	JobName     string `json:"job_name"`
}

type device_active_ddm_update_declaration struct {
	Active          bool   `json:"active"`
	DeclarationType string `json:"declaration_type"`
	Identifier      string `json:"identifier"`
	ServerToken     string `json:"server-token"`
	Valid           string `json:"valid"`
}

type device_compliance_benchmark_status struct {
	BenchmarkId    string          `json:"benchmark_id"`
	FailedRules    json.RawMessage `json:"failed_rules"`
	Id             string          `json:"id"`
	IsCompliant    bool            `json:"is_compliant"`
	LastUpdated    string          `json:"last_updated"`
	OrganizationId string          `json:"organization_id"`
}

type device_compliance_status struct {
	AgentId     string `json:"agent_id"`
	IsCompliant bool   `json:"is_compliant"`
}

type device_entities_DeviceAudit struct {
	AgentAuditDate string          `json:"agent_audit_date"`
	Agentid        string          `json:"agentid"`
	AuditDate      string          `json:"audit_date"`
	Facts          json.RawMessage `json:"facts"`
	Orgid          string          `json:"orgid"`
}

type device_entities_DeviceFact struct {
	ErrorMsg string `json:"error_msg"`
	Value    string `json:"value"`
}

type device_entities_DeviceResponse struct {
	Items    json.RawMessage `json:"items"`
	Metadata json.RawMessage `json:"metadata"`
}

type device_entities_device_audit struct {
	AgentAuditDate string          `json:"agent_audit_date"`
	Agentid        string          `json:"agentid"`
	AuditDate      string          `json:"audit_date"`
	Facts          json.RawMessage `json:"facts"`
	Orgid          string          `json:"orgid"`
}

type device_entities_device_fact struct {
	ErrorMsg string `json:"error_msg"`
	Value    string `json:"value"`
}

type device_entities_device_response struct {
	Items    json.RawMessage `json:"items"`
	Metadata json.RawMessage `json:"metadata"`
}

type device_filter struct {
	AuditField string          `json:"audit_field"`
	Operation  string          `json:"operation"`
	RangeValue float64         `json:"range_value"`
	Value      json.RawMessage `json:"value"`
}

type device_script_assignments_Device struct {
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

type device_script_assignments_DeviceScriptAssignment struct {
	AgentId  string          `json:"agent_id"`
	Device   json.RawMessage `json:"device"`
	Script   json.RawMessage `json:"script"`
	ScriptId string          `json:"script_id"`
}

type device_script_assignments_Script struct {
	Name string `json:"name"`
}

type device_script_assignments_device struct {
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

type device_script_assignments_device_script_assignment struct {
	AgentId  string          `json:"agent_id"`
	Device   json.RawMessage `json:"device"`
	Script   json.RawMessage `json:"script"`
	ScriptId string          `json:"script_id"`
}

type device_script_assignments_script struct {
	Name string `json:"name"`
}

type fact_entities_Fact struct {
	Id              string          `json:"id"`
	Name            string          `json:"name"`
	Notes           string          `json:"notes"`
	OsArchitectures json.RawMessage `json:"os_architectures"`
	ReturnType      string          `json:"return_type"`
	Version         int             `json:"version"`
}

type fact_entities_FactOSArchitecture struct {
	IsSupported bool   `json:"is_supported"`
	Language    string `json:"language"`
	Script      string `json:"script"`
	Shebang     string `json:"shebang"`
}

type fact_entities_FactResponse struct {
	Items    json.RawMessage `json:"items"`
	Metadata json.RawMessage `json:"metadata"`
}

type fact_entities_OSArchitectures struct {
	Linux json.RawMessage `json:"linux"`
	MacOS json.RawMessage `json:"macOS"`
}

type fact_entities_fact struct {
	Id              string          `json:"id"`
	Name            string          `json:"name"`
	Notes           string          `json:"notes"`
	OsArchitectures json.RawMessage `json:"os_architectures"`
	ReturnType      string          `json:"return_type"`
	Version         int             `json:"version"`
}

type fact_entities_fact_instruction_response struct {
	Fact        json.RawMessage `json:"fact"`
	Instruction json.RawMessage `json:"instruction"`
}

type fact_entities_fact_os_architecture struct {
	IsSupported bool   `json:"is_supported"`
	Language    string `json:"language"`
	Script      string `json:"script"`
	Shebang     string `json:"shebang"`
}

type fact_entities_fact_response struct {
	Items    json.RawMessage `json:"items"`
	Metadata json.RawMessage `json:"metadata"`
}

type fact_entities_os_architectures struct {
	Linux json.RawMessage `json:"linux"`
	MacOS json.RawMessage `json:"macOS"`
}

type feature_betas_entities_FeatureBeta struct {
	FeatureFlagKey string `json:"feature_flag_key"`
	KbLink         string `json:"kb_link"`
	Name           string `json:"name"`
	Public         bool   `json:"public"`
}

type feature_betas_entities_featurebeta struct {
	FeatureFlagKey string `json:"feature_flag_key"`
	KbLink         string `json:"kb_link"`
	Name           string `json:"name"`
	Public         bool   `json:"public"`
}

type file struct {
	ContentType    string `json:"content_type"`
	Created        string `json:"created"`
	Filename       string `json:"filename"`
	Id             string `json:"id"`
	Md5Hash        string `json:"md5_hash"`
	OrganizationId string `json:"organization_id"`
	Provider       string `json:"provider"`
	Size           int    `json:"size"`
	UserEmail      string `json:"user_email"`
}

type files_tracked_entity struct {
	FeatureName string `json:"feature_name"`
	FileId      string `json:"file_id"`
	ItemId      string `json:"item_id"`
	ItemName    string `json:"item_name"`
}

type files_tracked_request struct {
	FileIds json.RawMessage `json:"file_ids"`
}

type filter_set struct {
	AuditField   string          `json:"audit_field"`
	BooleanValue bool            `json:"boolean_value"`
	DateValue    string          `json:"date_value"`
	ListValue    json.RawMessage `json:"list_value"`
	NumberValue  float64         `json:"number_value"`
	Operation    string          `json:"operation"`
	StringValue  string          `json:"string_value"`
}

type homescreen_icon_item struct {
	BundleId    string          `json:"bundle_id"`
	DisplayName string          `json:"display_name"`
	Pages       json.RawMessage `json:"pages"`
	Url         string          `json:"url"`
}

type homescreen_layout_IconItem struct {
	BundleId    string          `json:"bundle_id"`
	DisplayName string          `json:"display_name"`
	Pages       json.RawMessage `json:"pages"`
	Url         string          `json:"url"`
}

type homescreen_request struct {
	Assigned   bool            `json:"assigned"`
	DeviceType string          `json:"device_type"`
	Dock       json.RawMessage `json:"dock"`
	Pages      json.RawMessage `json:"pages"`
	PolicyId   string          `json:"policy_id"`
}

type homescreen_response struct {
	AddigyPayloadType    string          `json:"addigy_payload_type"`
	AddigyPayloadVersion int             `json:"addigy_payload_version"`
	Assigned             bool            `json:"assigned"`
	DeviceType           string          `json:"device_type"`
	Dock                 json.RawMessage `json:"dock"`
	HasManifest          bool            `json:"has_manifest"`
	Orgid                string          `json:"orgid"`
	Pages                json.RawMessage `json:"pages"`
	PayloadDisplayName   string          `json:"payload_display_name"`
	PayloadEnabled       bool            `json:"payload_enabled"`
	PayloadGroupId       string          `json:"payload_group_id"`
	PayloadIdentifier    string          `json:"payload_identifier"`
	PayloadPriority      float64         `json:"payload_priority"`
	PayloadType          string          `json:"payload_type"`
	PayloadUuid          float64         `json:"payload_uuid"`
	PayloadVersion       int             `json:"payload_version"`
	PolicyId             string          `json:"policy_id"`
	PolicyRestricted     bool            `json:"policy_restricted"`
}

type identity_addigy_sync struct {
	ActiveService      string          `json:"active_service"`
	AwaitConfiguration bool            `json:"await_configuration"`
	ConfigVersion      int             `json:"config_version"`
	Services           json.RawMessage `json:"services"`
}

type identity_addigy_sync_store struct {
	ActiveService      string          `json:"active_service"`
	AwaitConfiguration bool            `json:"await_configuration"`
	Services           json.RawMessage `json:"services"`
}

type identity_image struct {
	ContentType string `json:"content_type"`
	Created     string `json:"created"`
	Filename    string `json:"filename"`
	Id          string `json:"id"`
	Md5Hash     string `json:"md5_hash"`
	Provider    string `json:"provider"`
	Size        int    `json:"size"`
	UserEmail   string `json:"user_email"`
}

type identity_provider struct {
	AllowLocalLogin       bool            `json:"allow_local_login"`
	AllowRevertLogin      bool            `json:"allow_revert_login"`
	AllowSyncUsers        bool            `json:"allow_sync_users"`
	BackgroundImage       json.RawMessage `json:"background_image"`
	ClientId              string          `json:"client_id"`
	CollectUserAttributes bool            `json:"collect_user_attributes"`
	IsAdmin               bool            `json:"is_admin"`
	LoginLogo             json.RawMessage `json:"login_logo"`
	RedirectUri           string          `json:"redirect_uri"`
}

type identity_service_azure struct {
	AllowLocalLogin       bool            `json:"allow_local_login"`
	AllowRevertLogin      bool            `json:"allow_revert_login"`
	AllowSyncUsers        bool            `json:"allow_sync_users"`
	BackgroundImage       json.RawMessage `json:"background_image"`
	ClientId              string          `json:"client_id"`
	ClientSecret          string          `json:"client_secret"`
	CollectUserAttributes bool            `json:"collect_user_attributes"`
	IsAdmin               bool            `json:"is_admin"`
	LoginLogo             json.RawMessage `json:"login_logo"`
	RedirectUri           string          `json:"redirect_uri"`
	TenantId              string          `json:"tenant_id"`
}

type identity_service_google struct {
	AllowLocalLogin       bool            `json:"allow_local_login"`
	AllowRevertLogin      bool            `json:"allow_revert_login"`
	AllowSyncUsers        bool            `json:"allow_sync_users"`
	BackgroundImage       json.RawMessage `json:"background_image"`
	ClientId              string          `json:"client_id"`
	ClientSecret          string          `json:"client_secret"`
	CollectUserAttributes bool            `json:"collect_user_attributes"`
	IsAdmin               bool            `json:"is_admin"`
	LoginLogo             json.RawMessage `json:"login_logo"`
	RedirectUri           string          `json:"redirect_uri"`
}

type identity_service_okta struct {
	AllowLocalLogin       bool            `json:"allow_local_login"`
	AllowRevertLogin      bool            `json:"allow_revert_login"`
	AllowSyncUsers        bool            `json:"allow_sync_users"`
	ApiToken              string          `json:"api_token"`
	BackgroundImage       json.RawMessage `json:"background_image"`
	ClientId              string          `json:"client_id"`
	CollectUserAttributes bool            `json:"collect_user_attributes"`
	Domain                string          `json:"domain"`
	HasApiManagement      bool            `json:"has_api_management"`
	IsAdmin               bool            `json:"is_admin"`
	LoginLogo             json.RawMessage `json:"login_logo"`
	RedirectUri           string          `json:"redirect_uri"`
}

type identity_services struct {
	Azure  json.RawMessage `json:"azure"`
	Google json.RawMessage `json:"google"`
	Okta   json.RawMessage `json:"okta"`
}

type installed_apps_mdm_entities_Application struct {
	AgentId            string `json:"agent_id"`
	HasUpdateAvailable bool   `json:"has_update_available"`
	Identifier         string `json:"identifier"`
	LastUpdated        string `json:"last_updated"`
	Name               string `json:"name"`
	Orgid              string `json:"orgid"`
	ShortVersion       string `json:"short_version"`
	Udid               string `json:"udid"`
	Version            string `json:"version"`
}

type installed_apps_mdm_entities_Request struct {
	AgentIds      json.RawMessage `json:"agent_ids"`
	Limit         int             `json:"limit"`
	Skip          int             `json:"skip"`
	SortDirection string          `json:"sort_direction"`
	SortField     string          `json:"sort_field"`
}

type installed_apps_mdm_entities_Response struct {
	Items    json.RawMessage `json:"items"`
	Metadata json.RawMessage `json:"metadata"`
}

type installed_apps_mdm_entities_application struct {
	AgentId            string `json:"agent_id"`
	HasUpdateAvailable bool   `json:"has_update_available"`
	Identifier         string `json:"identifier"`
	LastUpdated        string `json:"last_updated"`
	Name               string `json:"name"`
	Orgid              string `json:"orgid"`
	ShortVersion       string `json:"short_version"`
	Udid               string `json:"udid"`
	Version            string `json:"version"`
}

type installed_apps_mdm_entities_request struct {
	AgentIds      json.RawMessage `json:"agent_ids"`
	Limit         int             `json:"limit"`
	Skip          int             `json:"skip"`
	SortDirection string          `json:"sort_direction"`
	SortField     string          `json:"sort_field"`
}

type installed_apps_mdm_entities_response struct {
	Items    json.RawMessage `json:"items"`
	Metadata json.RawMessage `json:"metadata"`
}

type installed_system_updates_Request struct {
	Email    string `json:"email"`
	PolicyId string `json:"policy_id"`
}

type installed_system_updates_entities_Response struct {
	HumanReadableName string `json:"human_readable_name"`
	LastUpdated       string `json:"last_updated"`
	Version           string `json:"version"`
}

type installed_system_updates_request struct {
	Email    string `json:"email"`
	PolicyId string `json:"policy_id"`
}

type installed_system_updates_response struct {
	HumanReadableName string `json:"human_readable_name"`
	LastUpdated       string `json:"last_updated"`
	Version           string `json:"version"`
}

type instruction struct {
	AssetVersion     int    `json:"asset_version"`
	Condition        string `json:"condition"`
	Identifier       string `json:"identifier"`
	InstructionId    string `json:"instruction_id"`
	IsInBlueprint    bool   `json:"is_in_blueprint"`
	Label            string `json:"label"`
	Name             string `json:"name"`
	OsType           string `json:"os_type"`
	PolicyRestricted bool   `json:"policy_restricted"`
	Provider         string `json:"provider"`
	Public           bool   `json:"public"`
	RemoveScript     string `json:"remove_script"`
	RunOnSuccess     bool   `json:"run_on_success"`
	StatusOnSkipped  string `json:"status_on_skipped"`
	UserEmail        string `json:"user_email"`
}

type maintenance struct {
	Day                     string          `json:"day"`
	Enabled                 bool            `json:"enabled"`
	ExpectedRemediationTime int             `json:"expected_remediation_time"`
	Frequency               string          `json:"frequency"`
	Id                      string          `json:"id"`
	IsInBlueprint           bool            `json:"is_in_blueprint"`
	LocalTime               bool            `json:"local_time"`
	MaxTryCount             int             `json:"max_try_count"`
	Name                    string          `json:"name"`
	OrganizationId          string          `json:"organization_id"`
	Policies                json.RawMessage `json:"policies"`
	PolicyRestricted        bool            `json:"policy_restricted"`
	PromptUser              bool            `json:"prompt_user"`
	ScheduledTime           string          `json:"scheduled_time"`
	Scripts                 json.RawMessage `json:"scripts"`
	Source                  string          `json:"source"`
	SourceId                string          `json:"source_id"`
	Version                 int             `json:"version"`
}

type maintenance_instruction struct {
	Name   string `json:"name"`
	Script string `json:"script"`
}

type maintenance_post_request struct {
	Day                     string          `json:"day"`
	Enabled                 bool            `json:"enabled"`
	ExpectedRemediationTime int             `json:"expected_remediation_time"`
	Frequency               string          `json:"frequency"`
	IsInBlueprint           bool            `json:"is_in_blueprint"`
	LocalTime               bool            `json:"local_time"`
	MaxTryCount             int             `json:"max_try_count"`
	Name                    string          `json:"name"`
	PromptUser              bool            `json:"prompt_user"`
	ScheduledTime           string          `json:"scheduled_time"`
	Scripts                 json.RawMessage `json:"scripts"`
	TimeoutSeconds          int             `json:"timeout_seconds"`
	Version                 int             `json:"version"`
}

type maintenance_put_request struct {
	Day                     string          `json:"day"`
	Enabled                 bool            `json:"enabled"`
	ExpectedRemediationTime int             `json:"expected_remediation_time"`
	Frequency               string          `json:"frequency"`
	Id                      string          `json:"id"`
	IsInBlueprint           bool            `json:"is_in_blueprint"`
	LocalTime               bool            `json:"local_time"`
	MaxTryCount             int             `json:"max_try_count"`
	Name                    string          `json:"name"`
	PromptUser              bool            `json:"prompt_user"`
	ScheduledTime           string          `json:"scheduled_time"`
	Scripts                 json.RawMessage `json:"scripts"`
	TimeoutSeconds          int             `json:"timeout_seconds"`
	Version                 int             `json:"version"`
}

type maintenance_service_Maintenance struct {
	Day                     string          `json:"day"`
	Enabled                 bool            `json:"enabled"`
	ExpectedRemediationTime int             `json:"expected_remediation_time"`
	Frequency               string          `json:"frequency"`
	Id                      string          `json:"id"`
	IsInBlueprint           bool            `json:"is_in_blueprint"`
	LocalTime               bool            `json:"local_time"`
	MaxTryCount             int             `json:"max_try_count"`
	Name                    string          `json:"name"`
	OrganizationId          string          `json:"organization_id"`
	Policies                json.RawMessage `json:"policies"`
	PolicyRestricted        bool            `json:"policy_restricted"`
	PromptUser              bool            `json:"prompt_user"`
	ScheduledTime           string          `json:"scheduled_time"`
	Scripts                 json.RawMessage `json:"scripts"`
	Source                  string          `json:"source"`
	SourceId                string          `json:"source_id"`
	Version                 int             `json:"version"`
}

type maintenance_staged_request struct {
	Id string `json:"id"`
}

type managed_app_configuration struct {
	AssetId       string          `json:"asset_id"`
	BundleId      string          `json:"bundle_id"`
	Configuration json.RawMessage `json:"configuration"`
	LocationId    string          `json:"location_id"`
}

type managed_app_configuration_request struct {
	BundleId      string          `json:"bundle_id"`
	Configuration json.RawMessage `json:"configuration"`
	LocationId    string          `json:"location_id"`
}

type managed_apps_Configuration struct {
	AssetId       string          `json:"asset_id"`
	BundleId      string          `json:"bundle_id"`
	Configuration json.RawMessage `json:"configuration"`
	LocationId    string          `json:"location_id"`
}

type managed_apps_ConfigurationRequest struct {
	BundleId      string          `json:"bundle_id"`
	Configuration json.RawMessage `json:"configuration"`
	LocationId    string          `json:"location_id"`
}

type manifest_based_mdm_profiles_response struct {
	PoliciesMdmPayloads json.RawMessage `json:"policies_mdm_payloads"`
	Profiles            json.RawMessage `json:"profiles"`
	StagedPayloads      json.RawMessage `json:"staged_payloads"`
}

type mbov_account_status struct {
	Account    json.RawMessage `json:"account"`
	Enabled    bool            `json:"enabled"`
	MbovActive bool            `json:"mbov_active"`
}

type mbov_catalog_usage struct {
	CatalogCode string `json:"catalog_code"`
	Name        string `json:"name"`
	Usage       int    `json:"usage"`
}

type mbov_enable_request struct {
	BillingCity    string `json:"billing_city"`
	BillingCountry string `json:"billing_country"`
	BillingState   string `json:"billing_state"`
	BillingStreet  string `json:"billing_street"`
	BillingZipCode string `json:"billing_zip_code"`
	Email          string `json:"email"`
}

type mbov_site struct {
	CompanyName     string `json:"company_name"`
	Email           string `json:"email"`
	Firstname       string `json:"firstname"`
	Id              string `json:"id"`
	Lastname        string `json:"lastname"`
	NebulaAccountId string `json:"nebula_account_id"`
	SiteEndDate     string `json:"site_end_date"`
}

type mbov_sync_request struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type mdm_certificates_MDMInstalledCertificate struct {
	CertOrgName      string `json:"cert_org_name"`
	CommonName       string `json:"common_name"`
	Data             string `json:"data"`
	DeviceUuid       string `json:"device_uuid"`
	IsIdentity       bool   `json:"is_identity"`
	IssuerCommonName string `json:"issuer_common_name"`
	NotAfter         string `json:"not_after"`
	NotBefore        string `json:"not_before"`
	Version          int    `json:"version"`
}

type mdm_certificates_MDMInstalledCertificateResponse struct {
	Items    json.RawMessage `json:"items"`
	Metadata json.RawMessage `json:"metadata"`
}

type mdm_certificates_mdm_installed_certificate struct {
	CertOrgName      string `json:"cert_org_name"`
	CommonName       string `json:"common_name"`
	Data             string `json:"data"`
	DeviceUuid       string `json:"device_uuid"`
	IsIdentity       bool   `json:"is_identity"`
	IssuerCommonName string `json:"issuer_common_name"`
	NotAfter         string `json:"not_after"`
	NotBefore        string `json:"not_before"`
	Version          int    `json:"version"`
}

type mdm_certificates_mdm_installed_certificate_response struct {
	Items    json.RawMessage `json:"items"`
	Metadata json.RawMessage `json:"metadata"`
}

type mdm_command_command struct {
	RequestRequiresNetworkTether bool   `json:"request_requires_network_tether"`
	RequestType                  string `json:"request_type"`
}

type mdm_command_error_chain struct {
	ErrorCode            float64 `json:"error_code"`
	ErrorDomain          string  `json:"error_domain"`
	LocalizedDescription string  `json:"localized_description"`
	UsEnglishDescription string  `json:"us_english_description"`
}

type mdm_command_response struct {
	Command           json.RawMessage `json:"command"`
	CommandUuid       string          `json:"command_uuid"`
	CreatedDate       string          `json:"created_date"`
	DeviceUdid        string          `json:"device_udid"`
	ErrorChain        json.RawMessage `json:"error_chain"`
	ExecutionDeadline string          `json:"execution_deadline"`
	ExpirationDate    string          `json:"expiration_date"`
	LastUpdated       string          `json:"last_updated"`
	ManagedUserId     string          `json:"managed_user_id"`
	Name              string          `json:"name"`
	StartingDeadline  string          `json:"starting_deadline"`
	Status            string          `json:"status"`
	Weight            float64         `json:"weight"`
}

type mdm_configurations_MDMProfiles struct {
	AddigyPayloadType          string `json:"addigy_payload_type"`
	AddigyPayloadVersion       int    `json:"addigy_payload_version"`
	PayloadDisplayName         string `json:"payload_display_name"`
	PayloadEnabled             bool   `json:"payload_enabled"`
	PayloadGroupId             string `json:"payload_group_id"`
	PayloadIdentifier          string `json:"payload_identifier"`
	PayloadType                string `json:"payload_type"`
	PayloadUuid                string `json:"payload_uuid"`
	PayloadVersion             int    `json:"payload_version"`
	PolicyRestricted           bool   `json:"policy_restricted"`
	RequiresDeviceSupervision  bool   `json:"requires_device_supervision"`
	RequiresMdmProfileApproved bool   `json:"requires_mdm_profile_approved"`
}

type mdm_configurations_mdm_profiles struct {
	AddigyPayloadType          string `json:"addigy_payload_type"`
	AddigyPayloadVersion       int    `json:"addigy_payload_version"`
	PayloadDisplayName         string `json:"payload_display_name"`
	PayloadEnabled             bool   `json:"payload_enabled"`
	PayloadGroupId             string `json:"payload_group_id"`
	PayloadIdentifier          string `json:"payload_identifier"`
	PayloadType                string `json:"payload_type"`
	PayloadUuid                string `json:"payload_uuid"`
	PayloadVersion             int    `json:"payload_version"`
	PolicyRestricted           bool   `json:"policy_restricted"`
	RequiresDeviceSupervision  bool   `json:"requires_device_supervision"`
	RequiresMdmProfileApproved bool   `json:"requires_mdm_profile_approved"`
}

type mdm_device_details struct {
	ApnCertificate    json.RawMessage `json:"apn_certificate"`
	EnrollmentProfile json.RawMessage `json:"enrollment profile"`
	LastResponse      json.RawMessage `json:"last_response"`
}

type mdm_device_users_response struct {
	DataQuota     int    `json:"data_quota"`
	DataUsed      int    `json:"data_used"`
	FullName      string `json:"full_name"`
	HasDataToSync bool   `json:"has_data_to_sync"`
	MobileAccount bool   `json:"mobile_account"`
	Orgid         string `json:"orgid"`
	Udid          string `json:"udid"`
	UserName      string `json:"user_name"`
	UserUid       int    `json:"user_uid"`
}

type mdm_profile_installed_payload struct {
	DisplayName  string  `json:"display_name"`
	Identifier   string  `json:"identifier"`
	Organization string  `json:"organization"`
	Priority     float64 `json:"priority"`
	Uuid         string  `json:"uuid"`
	Version      int     `json:"version"`
}

type mdm_profiles_MDMListResponse struct {
	Count   int             `json:"count"`
	Results json.RawMessage `json:"results"`
	Total   int             `json:"total"`
}

type mdm_profiles_MDMProfileResponse struct {
	Profiles       json.RawMessage `json:"profiles"`
	StagedPayloads json.RawMessage `json:"staged_payloads"`
}

type mdm_profiles_mdm_list_response struct {
	Count   int             `json:"count"`
	Results json.RawMessage `json:"results"`
	Total   int             `json:"total"`
}

type mdm_profiles_mdm_profile_response struct {
	Profiles       json.RawMessage `json:"profiles"`
	StagedPayloads json.RawMessage `json:"staged_payloads"`
}

type mdm_profiles_policies struct {
	ApnCertificateId string          `json:"apn_certificate_id"`
	Certificate      json.RawMessage `json:"certificate"`
	Name             string          `json:"name"`
	Policy           json.RawMessage `json:"policy"`
}

type mdm_profiles_policies_response struct {
	Items    json.RawMessage `json:"items"`
	Metadata json.RawMessage `json:"metadata"`
}

type mdm_restart_device_response struct {
	Command           json.RawMessage `json:"Command"`
	CommandUUID       string          `json:"CommandUUID"`
	CreatedDate       string          `json:"CreatedDate"`
	DeviceUDID        string          `json:"DeviceUDID"`
	ErrorChain        json.RawMessage `json:"ErrorChain"`
	ExecutionDeadline string          `json:"ExecutionDeadline"`
	ExpirationDate    string          `json:"ExpirationDate"`
	LastUpdated       string          `json:"LastUpdated"`
	ManagedUserID     string          `json:"ManagedUserID"`
	Name              string          `json:"Name"`
	StartingDeadline  string          `json:"StartingDeadline"`
	Status            string          `json:"Status"`
	Weight            int             `json:"Weight"`
}

type metadata_entities_Metadata struct {
	Page        int `json:"page"`
	PageCount   int `json:"page_count"`
	PerPage     int `json:"per_page"`
	ResultCount int `json:"result_count"`
	Total       int `json:"total"`
}

type metadata_entities_metadata struct {
	Page        int `json:"page"`
	PageCount   int `json:"page_count"`
	PerPage     int `json:"per_page"`
	ResultCount int `json:"result_count"`
	Total       int `json:"total"`
}

type monitoring struct {
	Category           string          `json:"category"`
	Emails             json.RawMessage `json:"emails"`
	Fact               string          `json:"fact"`
	FactIdentifier     string          `json:"fact_identifier"`
	HasScript          bool            `json:"has_script"`
	Id                 string          `json:"id"`
	Instructions       json.RawMessage `json:"instructions"`
	IsInBlueprint      bool            `json:"is_in_blueprint"`
	Level              string          `json:"level"`
	MaxValue           float64         `json:"max_value"`
	MinValue           float64         `json:"min_value"`
	Name               string          `json:"name"`
	PolicyRestricted   bool            `json:"policy_restricted"`
	Provider           string          `json:"provider"`
	RemediationEnabled bool            `json:"remediation_enabled"`
	RemediationTime    int             `json:"remediation_time"`
	Script             json.RawMessage `json:"script"`
	ScriptId           string          `json:"script_id"`
	Selector           string          `json:"selector"`
	SendTicket         bool            `json:"send_ticket"`
	Source             string          `json:"source"`
	SourceId           string          `json:"source_id"`
	Value              string          `json:"value"`
	ValueType          json.RawMessage `json:"value_type"`
	Version            int             `json:"version"`
}

type monitoring_post_request struct {
	Category           string          `json:"category"`
	Emails             json.RawMessage `json:"emails"`
	FactIdentifier     string          `json:"fact_identifier"`
	HasScript          bool            `json:"has_script"`
	Instructions       json.RawMessage `json:"instructions"`
	IsInBlueprint      bool            `json:"is_in_blueprint"`
	Level              string          `json:"level"`
	MaxValue           float64         `json:"max_value"`
	MinValue           float64         `json:"min_value"`
	Name               string          `json:"name"`
	RemediationEnabled bool            `json:"remediation_enabled"`
	RemediationTime    int             `json:"remediation_time"`
	ScriptId           string          `json:"script_id"`
	Selector           string          `json:"selector"`
	SendTicket         bool            `json:"send_ticket"`
	Value              json.RawMessage `json:"value"`
	ValueType          string          `json:"value_type"`
	Version            int             `json:"version"`
}

type monitoring_put_request struct {
	Category           string          `json:"category"`
	Emails             json.RawMessage `json:"emails"`
	Fact               string          `json:"fact"`
	FactIdentifier     string          `json:"fact_identifier"`
	HasScript          bool            `json:"has_script"`
	Id                 string          `json:"id"`
	Instructions       json.RawMessage `json:"instructions"`
	IsInBlueprint      bool            `json:"is_in_blueprint"`
	Level              string          `json:"level"`
	MaxValue           float64         `json:"max_value"`
	MinValue           float64         `json:"min_value"`
	Name               string          `json:"name"`
	RemediationEnabled bool            `json:"remediation_enabled"`
	RemediationTime    int             `json:"remediation_time"`
	ScriptId           string          `json:"script_id"`
	Selector           string          `json:"selector"`
	SendTicket         bool            `json:"send_ticket"`
	Value              json.RawMessage `json:"value"`
	ValueType          string          `json:"value_type"`
	Version            int             `json:"version"`
}

type monitoring_query_request struct {
	ExcludedIds   json.RawMessage `json:"excluded_ids"`
	Ids           json.RawMessage `json:"ids"`
	Limit         int             `json:"limit"`
	NameContains  string          `json:"name_contains"`
	Skip          int             `json:"skip"`
	SortDirection string          `json:"sort_direction"`
	SortField     string          `json:"sort_field"`
}

type monitoring_query_response struct {
	AlertPolicies json.RawMessage `json:"alert_policies"`
	Alerts        json.RawMessage `json:"alerts"`
	StagedAlerts  json.RawMessage `json:"staged_alerts"`
}

type monitoring_received_alerts struct {
	AgentId            string          `json:"agent_id"`
	Category           string          `json:"category"`
	CreatedDate        string          `json:"created_date"`
	Emails             json.RawMessage `json:"emails"`
	FactIdentifier     string          `json:"fact_identifier"`
	FactName           string          `json:"fact_name"`
	Id                 string          `json:"id"`
	Level              string          `json:"level"`
	Name               string          `json:"name"`
	Orgid              string          `json:"orgid"`
	RemediationEnabled bool            `json:"remediation_enabled"`
	RemediationTime    int             `json:"remediation_time"`
	ResolvedDate       string          `json:"resolved_date"`
	ResolvedUserEmail  string          `json:"resolved_user_email"`
	SecondsToResolved  int             `json:"seconds_to_resolved"`
	Selector           string          `json:"selector"`
	Status             string          `json:"status"`
	Value              string          `json:"value"`
	ValueType          json.RawMessage `json:"value_type"`
}

type monitoring_received_alerts_response struct {
	Items    json.RawMessage `json:"items"`
	Metadata json.RawMessage `json:"metadata"`
}

type monitoring_scheduled_alert struct {
	Category           string          `json:"category"`
	Emails             json.RawMessage `json:"emails"`
	Fact               string          `json:"fact"`
	FactIdentifier     string          `json:"fact_identifier"`
	HasScript          bool            `json:"has_script"`
	Instructions       json.RawMessage `json:"instructions"`
	Level              string          `json:"level"`
	MaxValue           float64         `json:"max_value"`
	MinValue           float64         `json:"min_value"`
	Name               string          `json:"name"`
	Orgid              string          `json:"orgid"`
	PolicyRestricted   bool            `json:"policy_restricted"`
	Provider           string          `json:"provider"`
	RemediationEnabled bool            `json:"remediation_enabled"`
	RemediationTime    int             `json:"remediation_time"`
	Script             json.RawMessage `json:"script"`
	ScriptId           string          `json:"script_id"`
	Selector           string          `json:"selector"`
	SendTicket         bool            `json:"send_ticket"`
	Value              string          `json:"value"`
	ValueType          json.RawMessage `json:"value_type"`
}

type monitoring_stage_request struct {
	Id string `json:"id"`
}

type organization_entities_children_response struct {
	Items    json.RawMessage `json:"items"`
	Metadata json.RawMessage `json:"metadata"`
}

type organization_entities_organization struct {
	BillingPlatform         string  `json:"billing_platform"`
	CompanyName             string  `json:"company_name"`
	DeviceCount             int     `json:"device_count"`
	Email                   string  `json:"email"`
	Enabled                 bool    `json:"enabled"`
	FirstName               string  `json:"first_name"`
	GoLiveSsEnabled         bool    `json:"go_live_ss_enabled"`
	LastName                string  `json:"last_name"`
	LastVisit               float64 `json:"last_visit"`
	OrganizationId          string  `json:"organization_id"`
	Parent                  string  `json:"parent"`
	SoftwareMeteringEnabled bool    `json:"software_metering_enabled"`
	SoftwareSharingAllowed  bool    `json:"software_sharing_allowed"`
	SoftwareSharingEnabled  bool    `json:"software_sharing_enabled"`
	Subdomain               string  `json:"subdomain"`
	SupportAccessEnabled    bool    `json:"support_access_enabled"`
	SupportAccessExpires    float64 `json:"support_access_expires"`
	TrialEnd                int     `json:"trial_end"`
	UlaDate                 string  `json:"ula_date"`
	UlaIp                   string  `json:"ula_ip"`
}

type organization_files struct {
	ContentType string `json:"content_type"`
	Created     string `json:"created"`
	Filename    string `json:"filename"`
	Id          string `json:"id"`
	Md5Hash     string `json:"md5_hash"`
	Orgid       string `json:"orgid"`
	Provider    string `json:"provider"`
	Size        int    `json:"size"`
	UserEmail   string `json:"user_email"`
}

type organization_users_response struct {
	Items    json.RawMessage `json:"items"`
	Metadata json.RawMessage `json:"metadata"`
}

type policies_mdm_payloads struct {
	ConfigurationId string `json:"configuration_id"`
	Orgid           string `json:"orgid"`
	PolicyId        string `json:"policy_id"`
}

type policy_ade_token struct {
	Account              json.RawMessage `json:"account"`
	DevicesSyncCompleted bool            `json:"devices_sync_completed"`
	LastScanTime         string          `json:"last_scan_time"`
	PolicyId             string          `json:"policy_id"`
	SyncingError         string          `json:"syncing_error"`
}

type policy_apps_and_books_token struct {
	Active            bool            `json:"active"`
	ClientConfig      json.RawMessage `json:"clientConfig"`
	ErrMsg            string          `json:"errMsg"`
	LastSync          string          `json:"lastSync"`
	LastSyncingStatus string          `json:"lastSyncingStatus"`
	LocationId        string          `json:"locationId"`
	OrgName           string          `json:"orgName"`
	PolicyId          string          `json:"policyId"`
	Syncing           bool            `json:"syncing"`
	Version           string          `json:"version"`
}

type policy_assignment_rule struct {
	Created  string `json:"created"`
	Disabled bool   `json:"disabled"`
	RuleId   string `json:"rule_id"`
	Script   string `json:"script"`
}

type policy_assignment_rule_create_request struct {
	AutoRemove bool            `json:"auto_remove"`
	Disabled   bool            `json:"disabled"`
	Filters    json.RawMessage `json:"filters"`
	PolicyId   string          `json:"policy_id"`
}

type policy_create_request struct {
	Color          string `json:"color"`
	Icon           string `json:"icon"`
	Name           string `json:"name"`
	ParentPolicyId string `json:"parent_policy_id"`
}

type policy_entities_policy_response struct {
	AddigySync                json.RawMessage `json:"addigy_sync"`
	AgentPath                 string          `json:"agent_path"`
	AgentVersion              string          `json:"agent_version"`
	AutotaskAccountId         int             `json:"autotask_account_id"`
	CollectorSettings         json.RawMessage `json:"collector_settings"`
	Color                     string          `json:"color"`
	ConnectwiseAccountId      int             `json:"connectwise_account_id"`
	CreationTime              string          `json:"creation_time"`
	DownloadPath              string          `json:"download_path"`
	Icon                      string          `json:"icon"`
	IgnoreUpdates             bool            `json:"ignore_updates"`
	Instructions              json.RawMessage `json:"instructions"`
	ItglueAccountId           string          `json:"itglue_account_id"`
	LastDeployed              string          `json:"last_deployed"`
	Name                      string          `json:"name"`
	Orgid                     string          `json:"orgid"`
	Parent                    string          `json:"parent"`
	PolicyId                  string          `json:"policyId"`
	Schedules                 json.RawMessage `json:"schedules"`
	SelfServiceInstructionIds json.RawMessage `json:"self_service_instruction_ids"`
	SplashtopSettings         json.RawMessage `json:"splashtop_settings"`
	SshSettings               json.RawMessage `json:"ssh_settings"`
	VncSettings               json.RawMessage `json:"vnc_settings"`
}

type policy_identity struct {
	AddigySync json.RawMessage `json:"addigy_sync"`
	Id         string          `json:"id"`
	PolicyId   string          `json:"policy_id"`
}

type policy_identity_store struct {
	AddigySync json.RawMessage `json:"addigy_sync"`
	Id         string          `json:"id"`
	PolicyId   string          `json:"policy_id"`
}

type policy_service_Artwork struct {
	Height int    `json:"height"`
	Url    string `json:"url"`
	Width  int    `json:"width"`
}

type policy_service_OaSelfServiceAssetsRequests struct {
	DeviceFamily string          `json:"device_family"`
	LocationIds  json.RawMessage `json:"location_ids"`
	PolicyIds    json.RawMessage `json:"policy_ids"`
}

type policy_service_Offers struct {
	Assets  json.RawMessage `json:"assets"`
	Version json.RawMessage `json:"version"`
}

type policy_service_OffersAssets struct {
	Flavor string `json:"flavor"`
	Size   string `json:"size"`
}

type policy_service_SelfServiceLocationAsset struct {
	AssetDetails json.RawMessage `json:"asset_details"`
}

type policy_update_request struct {
	Color    string `json:"color"`
	Icon     string `json:"icon"`
	Name     string `json:"name"`
	PolicyId string `json:"policy_id"`
}

type prebuilt_apps_app struct {
	Id         string          `json:"id"`
	LatestId   string          `json:"latest_id"`
	Name       string          `json:"name"`
	PngIcon    string          `json:"png_icon"`
	SearchPath string          `json:"search_path"`
	SvgIcon    string          `json:"svg_icon"`
	VersionIds json.RawMessage `json:"version_ids"`
}

type prebuilt_apps_app_query_response struct {
	Items    json.RawMessage `json:"items"`
	Metadata json.RawMessage `json:"metadata"`
}

type prebuilt_apps_configuration struct {
	Id                   string          `json:"id"`
	PrebuiltAppId        string          `json:"prebuilt_app_id"`
	PrebuiltAppVersionId string          `json:"prebuilt_app_version_id"`
	RunRemovalScript     bool            `json:"run_removal_script"`
	UserDeferral         int             `json:"user_deferral"`
	Variables            json.RawMessage `json:"variables"`
}

type prebuilt_apps_configurations_assignment_request struct {
	PolicyIds json.RawMessage `json:"policy_ids"`
}

type prebuilt_apps_configurations_create_request struct {
	AutoUpdate           bool            `json:"auto_update"`
	PolicyIds            json.RawMessage `json:"policy_ids"`
	PrebuiltAppId        string          `json:"prebuilt_app_id"`
	PrebuiltAppVersionId string          `json:"prebuilt_app_version_id"`
	RunRemovalScript     bool            `json:"run_removal_script"`
	UserDeferral         int             `json:"user_deferral"`
}

type prebuilt_apps_configurations_put_request struct {
	AutoUpdate       bool `json:"auto_update"`
	RunRemovalScript bool `json:"run_removal_script"`
	UserDeferral     int  `json:"user_deferral"`
}

type prebuilt_apps_configurations_query_request struct {
	Ids                   json.RawMessage `json:"ids"`
	Page                  int             `json:"page"`
	PerPage               int             `json:"per_page"`
	PrebuiltAppIds        json.RawMessage `json:"prebuilt_app_ids"`
	PrebuiltAppVersionIds json.RawMessage `json:"prebuilt_app_version_ids"`
	SortDirection         string          `json:"sort_direction"`
	SortField             string          `json:"sort_field"`
}

type prebuilt_apps_configurations_query_response struct {
	Items    json.RawMessage `json:"items"`
	Metadata json.RawMessage `json:"metadata"`
}

type prebuilt_apps_create_app_request struct {
	Name       string `json:"name"`
	PngIcon    string `json:"png_icon"`
	SearchPath string `json:"search_path"`
	SvgIcon    string `json:"svg_icon"`
}

type prebuilt_apps_query_apps_request struct {
	Ids           json.RawMessage `json:"ids"`
	Name          string          `json:"name"`
	Page          int             `json:"page"`
	PerPage       int             `json:"per_page"`
	SortDirection string          `json:"sort_direction"`
	SortField     string          `json:"sort_field"`
	VersionId     string          `json:"version_id"`
}

type prebuilt_apps_update_app_request struct {
	LatestId   string          `json:"latest_id"`
	PngIcon    string          `json:"png_icon"`
	SearchPath string          `json:"search_path"`
	SvgIcon    string          `json:"svg_icon"`
	VersionIds json.RawMessage `json:"version_ids"`
}

type prebuilt_apps_versions_create_update struct {
	AppId           string          `json:"app_id"`
	ConditionScript string          `json:"condition_script"`
	Files           json.RawMessage `json:"files"`
	InstallScript   string          `json:"install_script"`
	Notification    string          `json:"notification"`
	Profiles        json.RawMessage `json:"profiles"`
	PublishedDate   string          `json:"published_date"`
	ReleaseNotesUrl string          `json:"release_notes_url"`
	RemoveScript    string          `json:"remove_script"`
	Variables       json.RawMessage `json:"variables"`
	Version         string          `json:"version"`
}

type prebuilt_apps_versions_profiles struct {
	AddigyPayloadType            string          `json:"addigy_payload_type"`
	AllowUserOverrides           bool            `json:"allow_user_overrides"`
	AllowedSystemExtensions      json.RawMessage `json:"allowed_system_extensions"`
	AllowedSystemExtensionsTypes json.RawMessage `json:"allowed_system_extensions_types"`
	AllowedTeamIdentifiers       json.RawMessage `json:"allowed_team_identifiers"`
	Bundle                       string          `json:"bundle"`
	Custom                       bool            `json:"custom"`
	Events                       json.RawMessage `json:"events"`
	FileId                       string          `json:"file_id"`
	Identifier                   string          `json:"identifier"`
	IdentifierType               string          `json:"identifier_type"`
	Name                         string          `json:"name"`
	Path                         string          `json:"path"`
	Permissions                  json.RawMessage `json:"permissions"`
	RemovableSystemExtensions    json.RawMessage `json:"removable_system_extensions"`
	Requirements                 string          `json:"requirements"`
	Rules                        json.RawMessage `json:"rules"`
	Signature                    string          `json:"signature"`
}

type prebuilt_apps_versions_query_request struct {
	AppIds        json.RawMessage `json:"app_ids"`
	Ids           json.RawMessage `json:"ids"`
	Page          int             `json:"page"`
	PerPage       int             `json:"per_page"`
	SortDirection string          `json:"sort_direction"`
	SortField     string          `json:"sort_field"`
	Version       string          `json:"version"`
}

type prebuilt_apps_versions_query_response struct {
	Items json.RawMessage `json:"items"`
	Total int             `json:"total"`
}

type prebuilt_apps_versions_version struct {
	AppId           string          `json:"app_id"`
	ConditionScript string          `json:"condition_script"`
	Deprecated      bool            `json:"deprecated"`
	DeprecatedDate  string          `json:"deprecated_date"`
	Files           json.RawMessage `json:"files"`
	Id              string          `json:"id"`
	InstallScript   string          `json:"install_script"`
	Notification    string          `json:"notification"`
	Profiles        json.RawMessage `json:"profiles"`
	PublishedDate   string          `json:"published_date"`
	ReleaseNotesUrl string          `json:"release_notes_url"`
	RemoveScript    string          `json:"remove_script"`
	Variables       json.RawMessage `json:"variables"`
	Version         string          `json:"version"`
}

type report_create_request struct {
	EmailSettings json.RawMessage `json:"email_settings"`
	Payload       json.RawMessage `json:"payload"`
}

type report_status struct {
	Id             string `json:"id"`
	LastUpdated    string `json:"last_updated"`
	OrganizationId string `json:"organization_id"`
	Status         string `json:"status"`
}

type response_Error struct {
	Code       int             `json:"code"`
	ErrorChain json.RawMessage `json:"error_chain"`
	Message    string          `json:"message"`
}

type response_error struct {
	Code       int             `json:"code"`
	ErrorChain json.RawMessage `json:"error_chain"`
	Message    string          `json:"message"`
}

type schedule_starting_time struct {
	Hour   string `json:"hour"`
	Minute string `json:"minute"`
}

type self_service_asset_artwork struct {
	Height int    `json:"height"`
	Url    string `json:"url"`
	Width  int    `json:"width"`
}

type self_service_asset_offers struct {
	Assets  json.RawMessage `json:"assets"`
	Version json.RawMessage `json:"version"`
}

type self_service_asset_offers_assets struct {
	Flavor string `json:"flavor"`
	Size   string `json:"size"`
}

type self_service_configuration_Request struct {
	AppLogo                    json.RawMessage `json:"app_logo"`
	DockIcon                   json.RawMessage `json:"dock_icon"`
	FilevaultPromptText        string          `json:"filevault_prompt_text"`
	HideChat                   bool            `json:"hide_chat"`
	HomeScreenAddress          string          `json:"home_screen_address"`
	HomeScreenCompanyName      string          `json:"home_screen_company_name"`
	HomeScreenConfigureDetails bool            `json:"home_screen_configure_details"`
	HomeScreenDescription      string          `json:"home_screen_description"`
	HomeScreenEmail            string          `json:"home_screen_email"`
	HomeScreenPhone            string          `json:"home_screen_phone"`
	HomeScreenShowAddress      bool            `json:"home_screen_show_address"`
	HomeScreenShowDescription  bool            `json:"home_screen_show_description"`
	HomeScreenShowEmail        bool            `json:"home_screen_show_email"`
	HomeScreenShowPhone        bool            `json:"home_screen_show_phone"`
	IntegrationIntuneEnabled   bool            `json:"integration_intune_enabled"`
	IsInBlueprint              bool            `json:"is_in_blueprint"`
	IsOnboardingConfig         bool            `json:"is_onboarding_config"`
	MaintenancePromptText      string          `json:"maintenance_prompt_text"`
	MenubarIcon                json.RawMessage `json:"menubar_icon"`
	MsOfficeUpdatesPromptText  string          `json:"ms_office_updates_prompt_text"`
	Name                       string          `json:"name"`
	OsType                     string          `json:"os_type"`
	ScreenviewPromptText       string          `json:"screenview_prompt_text"`
	ShowDockIcon               bool            `json:"show_dock_icon"`
	ShowInApplications         bool            `json:"show_in_applications"`
	ShowMenubarIcon            bool            `json:"show_menubar_icon"`
	ShowSupport                bool            `json:"show_support"`
	UserSentimentPromptText    string          `json:"user_sentiment_prompt_text"`
	Version                    int             `json:"version"`
}

type self_service_configuration_post_request struct {
	AppLogo                    json.RawMessage `json:"app_logo"`
	DockIcon                   json.RawMessage `json:"dock_icon"`
	FilevaultPromptText        string          `json:"filevault_prompt_text"`
	HideChat                   bool            `json:"hide_chat"`
	HomeScreenAddress          string          `json:"home_screen_address"`
	HomeScreenCompanyName      string          `json:"home_screen_company_name"`
	HomeScreenConfigureDetails bool            `json:"home_screen_configure_details"`
	HomeScreenDescription      string          `json:"home_screen_description"`
	HomeScreenEmail            string          `json:"home_screen_email"`
	HomeScreenPhone            string          `json:"home_screen_phone"`
	HomeScreenShowAddress      bool            `json:"home_screen_show_address"`
	HomeScreenShowDescription  bool            `json:"home_screen_show_description"`
	HomeScreenShowEmail        bool            `json:"home_screen_show_email"`
	HomeScreenShowPhone        bool            `json:"home_screen_show_phone"`
	IntegrationIntuneEnabled   bool            `json:"integration_intune_enabled"`
	IsInBlueprint              bool            `json:"is_in_blueprint"`
	IsOnboardingConfig         bool            `json:"is_onboarding_config"`
	MaintenancePromptText      string          `json:"maintenance_prompt_text"`
	MenubarIcon                json.RawMessage `json:"menubar_icon"`
	MsOfficeUpdatesPromptText  string          `json:"ms_office_updates_prompt_text"`
	Name                       string          `json:"name"`
	OsType                     string          `json:"os_type"`
	ScreenviewPromptText       string          `json:"screenview_prompt_text"`
	ShowDockIcon               bool            `json:"show_dock_icon"`
	ShowInApplications         bool            `json:"show_in_applications"`
	ShowMenubarIcon            bool            `json:"show_menubar_icon"`
	ShowSupport                bool            `json:"show_support"`
	UserSentimentPromptText    string          `json:"user_sentiment_prompt_text"`
	Version                    int             `json:"version"`
}

type self_service_configuration_post_response struct {
	AppLogo                    json.RawMessage `json:"app_logo"`
	DockIcon                   json.RawMessage `json:"dock_icon"`
	FilevaultPromptText        string          `json:"filevault_prompt_text"`
	HideChat                   bool            `json:"hide_chat"`
	HomeScreenAddress          string          `json:"home_screen_address"`
	HomeScreenCompanyName      string          `json:"home_screen_company_name"`
	HomeScreenConfigureDetails bool            `json:"home_screen_configure_details"`
	HomeScreenDescription      string          `json:"home_screen_description"`
	HomeScreenEmail            string          `json:"home_screen_email"`
	HomeScreenPhone            string          `json:"home_screen_phone"`
	HomeScreenShowAddress      bool            `json:"home_screen_show_address"`
	HomeScreenShowDescription  bool            `json:"home_screen_show_description"`
	HomeScreenShowEmail        bool            `json:"home_screen_show_email"`
	HomeScreenShowPhone        bool            `json:"home_screen_show_phone"`
	InstructionId              string          `json:"instruction_id"`
	IntegrationIntuneEnabled   bool            `json:"integration_intune_enabled"`
	IsInBlueprint              bool            `json:"is_in_blueprint"`
	IsOnboardingConfig         bool            `json:"is_onboarding_config"`
	MaintenancePromptText      string          `json:"maintenance_prompt_text"`
	MenubarIcon                json.RawMessage `json:"menubar_icon"`
	MsOfficeUpdatesPromptText  string          `json:"ms_office_updates_prompt_text"`
	Name                       string          `json:"name"`
	OsType                     string          `json:"os_type"`
	ScreenviewPromptText       string          `json:"screenview_prompt_text"`
	ShowDockIcon               bool            `json:"show_dock_icon"`
	ShowInApplications         bool            `json:"show_in_applications"`
	ShowMenubarIcon            bool            `json:"show_menubar_icon"`
	ShowSupport                bool            `json:"show_support"`
	UserSentimentPromptText    string          `json:"user_sentiment_prompt_text"`
	Version                    int             `json:"version"`
}

type self_service_location_asset struct {
	Artwork          json.RawMessage `json:"artwork"`
	AssetId          string          `json:"asset_id"`
	Assigned         bool            `json:"assigned"`
	AssignedCount    int             `json:"assigned_count"`
	AssignedPolicyId string          `json:"assigned_policy_id"`
	AutoUpdate       bool            `json:"auto_update"`
	AvailableCount   int             `json:"available_count"`
	LocationId       string          `json:"location_id"`
	NameRaw          string          `json:"name_raw"`
	Offers           json.RawMessage `json:"offers"`
	OrganizationId   string          `json:"organization_id"`
	PolicyId         string          `json:"policy_id"`
	ShortUrl         string          `json:"short_url"`
	SupportedDevices json.RawMessage `json:"supported_devices"`
	TotalCount       int             `json:"total_count"`
}

type self_service_location_assets_request struct {
	DeviceFamily string          `json:"device_family"`
	LocationIds  json.RawMessage `json:"location_ids"`
	PolicyIds    json.RawMessage `json:"policy_ids"`
}

type self_service_location_assets_response struct {
	AssetDetails json.RawMessage `json:"asset_details"`
}

type self_service_policy_assigned_instructions struct {
	OrganizationId                   string          `json:"organization_id"`
	PolicyIdForIos                   string          `json:"policy_id_for_ios"`
	PolicyIdForMacOs                 string          `json:"policy_id_for_mac_os"`
	PolicyNameForIos                 string          `json:"policy_name_for_ios"`
	PolicyNameForMacOs               string          `json:"policy_name_for_mac_os"`
	SelfServiceConfigurationForIos   json.RawMessage `json:"self_service_configuration_for_ios"`
	SelfServiceConfigurationForMacOs json.RawMessage `json:"self_service_configuration_for_mac_os"`
}

type static_fields_entities_StaticField struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type static_fields_entities_StaticFieldValue struct {
	Agentid       string `json:"agentid"`
	StaticFieldId string `json:"static_field_id"`
	Value         string `json:"value"`
}

type static_fields_entities_StaticFieldValueResponse struct {
	Failed    json.RawMessage `json:"failed"`
	Succeeded json.RawMessage `json:"succeeded"`
}

type static_fields_entities_static_field struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type static_fields_entities_static_field_value struct {
	Agentid       string `json:"agentid"`
	StaticFieldId string `json:"static_field_id"`
	Value         string `json:"value"`
}

type static_fields_entities_static_field_value_response struct {
	Failed    json.RawMessage `json:"failed"`
	Succeeded json.RawMessage `json:"succeeded"`
}

type system_events_entities_Response struct {
	Items    json.RawMessage `json:"items"`
	Metadata json.RawMessage `json:"metadata"`
}

type system_events_entities_SystemEvent struct {
	Action         json.RawMessage `json:"action"`
	ActionReceiver json.RawMessage `json:"action_receiver"`
	ActionSender   json.RawMessage `json:"action_sender"`
	Date           string          `json:"date"`
	EventId        string          `json:"event_id"`
	Level          string          `json:"level"`
	Orgid          string          `json:"orgid"`
	Result         json.RawMessage `json:"result"`
}

type system_events_entities_response struct {
	Items    json.RawMessage `json:"items"`
	Metadata json.RawMessage `json:"metadata"`
}

type system_events_entities_system_event struct {
	Action         json.RawMessage `json:"action"`
	ActionReceiver json.RawMessage `json:"action_receiver"`
	ActionSender   json.RawMessage `json:"action_sender"`
	Date           string          `json:"date"`
	EventId        string          `json:"event_id"`
	Level          string          `json:"level"`
	Orgid          string          `json:"orgid"`
	Result         json.RawMessage `json:"result"`
}

type system_events_query struct {
	Fields json.RawMessage `json:"fields"`
	Query  string          `json:"query"`
}

type system_events_search_request struct {
	FromDateTime  string          `json:"from_date_time"`
	Options       json.RawMessage `json:"options"`
	Page          int             `json:"page"`
	PerPage       int             `json:"per_page"`
	Queries       json.RawMessage `json:"queries"`
	SortDirection string          `json:"sort_direction"`
	ToDateTime    string          `json:"to_date_time"`
}

type system_events_search_response struct {
	Aggregations json.RawMessage `json:"aggregations"`
	Items        json.RawMessage `json:"items"`
	Keywords     json.RawMessage `json:"keywords"`
	Metadata     json.RawMessage `json:"metadata"`
	Took         int             `json:"took"`
}

type system_updates_available_request struct {
	DeviceUdid string `json:"device_udid"`
}

type system_updates_available_response struct {
	AllowsInstallLater         bool            `json:"allows_install_later"`
	AppIdentifiersToClose      json.RawMessage `json:"app_identifiers_to_close"`
	Build                      string          `json:"build"`
	DeferredUntil              string          `json:"deferred_until"`
	DownloadSize               float64         `json:"download_size"`
	HumanReadableName          string          `json:"human_readable_name"`
	HumanReadableNameLocale    string          `json:"human_readable_name_locale"`
	InstallSize                float64         `json:"install_size"`
	IsConfigDataUpdate         bool            `json:"is_config_data_update"`
	IsCritical                 bool            `json:"is_critical"`
	IsFirmwareUpdate           bool            `json:"is_firmware_update"`
	IsMajorOsUpdate            bool            `json:"is_major_os_update"`
	IsSecurityResponse         bool            `json:"is_security_response"`
	LastUpdated                string          `json:"last_updated"`
	MetadataUrl                string          `json:"metadata_url"`
	Orgid                      string          `json:"orgid"`
	ProductKey                 string          `json:"product_key"`
	ProductName                string          `json:"product_name"`
	RequiresBootstrapToken     bool            `json:"requires_bootstrap_token"`
	RestartRequired            bool            `json:"restart_required"`
	SupplementalBuildVersion   string          `json:"supplemental_build_version"`
	SupplementalOsVersionExtra string          `json:"supplemental_os_version_extra"`
	Udid                       string          `json:"udid"`
	Version                    string          `json:"version"`
}

type system_updates_available_with_statuses_response struct {
	Status json.RawMessage `json:"status"`
	Update json.RawMessage `json:"update"`
}

type system_updates_enqueue_schedule_request struct {
	DeviceUdid     string `json:"device_udid"`
	InstallAction  string `json:"install_action"`
	ProductKey     string `json:"product_key"`
	ProductVersion string `json:"product_version"`
}

type system_updates_installed_response struct {
	AgentId           string `json:"agent_id"`
	CreatedDate       string `json:"created_date"`
	DeviceModel       string `json:"device_model"`
	DeviceUuid        string `json:"device_uuid"`
	HumanReadableName string `json:"human_readable_name"`
	LastUpdated       string `json:"last_updated"`
	OrganizationId    string `json:"organization_id"`
	Version           string `json:"version"`
}

type system_updates_latest_update struct {
	Build               string `json:"build"`
	ExpirationDate      string `json:"expiration_date"`
	PostingDate         string `json:"posting_date"`
	PrerequisiteBuild   string `json:"prerequisite_build"`
	ProductVersionExtra string `json:"product_version_extra"`
	Version             string `json:"version"`
}

type system_updates_latest_updates struct {
	Ios         json.RawMessage `json:"ios"`
	Ipados      json.RawMessage `json:"ipados"`
	LastUpdated string          `json:"last_updated"`
	Macos       json.RawMessage `json:"macos"`
	Tvos        json.RawMessage `json:"tvos"`
	Watchos     json.RawMessage `json:"watchos"`
}

type system_updates_on_demand_devices struct {
	DeviceUuids json.RawMessage `json:"device_uuids"`
}

type system_updates_on_demand_entities_DevicesRequest struct {
	DeviceUuids json.RawMessage `json:"device_uuids"`
}

type system_updates_on_demand_entities_PolicyRequest struct {
	PolicyId string `json:"policy_id"`
}

type system_updates_on_demand_policy struct {
	PolicyId string `json:"policy_id"`
}

type system_updates_scan_request struct {
	DeviceUdid string `json:"device_udid"`
	Force      bool   `json:"force"`
}

type system_updates_schedule struct {
	CutOffTime        string          `json:"cut_off_time"`
	Enabled           bool            `json:"enabled"`
	MaintenanceWindow string          `json:"maintenance_window"`
	StartingTime      json.RawMessage `json:"starting_time"`
	WeekDays          json.RawMessage `json:"week_days"`
}

type system_updates_settings struct {
	AllowBetaUpdatesInDdm   bool            `json:"allow_beta_updates_in_ddm"`
	AllowedDays             json.RawMessage `json:"allowed_days"`
	DaysAfterRelease        int             `json:"days_after_release"`
	DaysAfterReleaseRsr     int             `json:"days_after_release_rsr"`
	DeniedPeriod            json.RawMessage `json:"denied_period"`
	Enabled                 bool            `json:"enabled"`
	HoursAfterRelease       int             `json:"hours_after_release"`
	InstallAction           string          `json:"install_action"`
	KeepOsUpdated           bool            `json:"keep_os_updated"`
	MaxOsVersionAllowed     string          `json:"max_os_version_allowed"`
	MaxUserDeferrals        int             `json:"max_user_deferrals"`
	MinutesAfterRelease     int             `json:"minutes_after_release"`
	ResendUpdateCommandHour int             `json:"resend_update_command_hour"`
}

type system_updates_settings_entities_Request struct {
	IosSettings    json.RawMessage `json:"ios_settings"`
	IpadosSettings json.RawMessage `json:"ipados_settings"`
	MacosSettings  json.RawMessage `json:"macos_settings"`
	PolicyId       string          `json:"policy_id"`
	Schedule       json.RawMessage `json:"schedule"`
	TvosSettings   json.RawMessage `json:"tvos_settings"`
}

type system_updates_settings_entities_Response struct {
	IosSettings    json.RawMessage `json:"ios_settings"`
	IpadosSettings json.RawMessage `json:"ipados_settings"`
	LastUpdated    string          `json:"last_updated"`
	MacosSettings  json.RawMessage `json:"macos_settings"`
	OrgId          string          `json:"org_id"`
	PolicyId       string          `json:"policy_id"`
	Schedule       json.RawMessage `json:"schedule"`
	TvosSettings   json.RawMessage `json:"tvos_settings"`
}

type system_updates_settings_request struct {
	IosSettings    json.RawMessage `json:"ios_settings"`
	IpadosSettings json.RawMessage `json:"ipados_settings"`
	MacosSettings  json.RawMessage `json:"macos_settings"`
	PolicyId       string          `json:"policy_id"`
	Schedule       json.RawMessage `json:"schedule"`
	TvosSettings   json.RawMessage `json:"tvos_settings"`
}

type system_updates_settings_response struct {
	IosSettings    json.RawMessage `json:"ios_settings"`
	IpadosSettings json.RawMessage `json:"ipados_settings"`
	LastUpdated    string          `json:"last_updated"`
	MacosSettings  json.RawMessage `json:"macos_settings"`
	OrgId          string          `json:"org_id"`
	PolicyId       string          `json:"policy_id"`
	Schedule       json.RawMessage `json:"schedule"`
	TvosSettings   json.RawMessage `json:"tvos_settings"`
}

type system_updates_status_list_request struct {
	DeviceUdid string `json:"device_udid"`
}

type system_updates_statuses_response struct {
	CheckedForCompletion    bool            `json:"checked_for_completion"`
	DeferralsRemaining      float64         `json:"deferrals_remaining"`
	DownloadPercentComplete float64         `json:"download_percent_complete"`
	ErrorChain              json.RawMessage `json:"error_chain"`
	ErrorMessage            string          `json:"error_message"`
	ExecutionDeadline       string          `json:"execution_deadline"`
	ExpirationDate          string          `json:"expiration_date"`
	HumanReadableName       string          `json:"human_readable_name"`
	IsDownloaded            bool            `json:"is_downloaded"`
	LastUpdated             string          `json:"last_updated"`
	MaxDeferrals            float64         `json:"max_deferrals"`
	NextScheduledInstall    string          `json:"next_scheduled_install"`
	Orgid                   string          `json:"orgid"`
	PastNotifications       json.RawMessage `json:"past_notifications"`
	ProductKey              string          `json:"product_key"`
	StartingDeadline        string          `json:"starting_deadline"`
	Status                  string          `json:"status"`
	Udid                    string          `json:"udid"`
	UpdateCommandReceived   bool            `json:"update_command_received"`
	UpdateCommandSent       bool            `json:"update_command_sent"`
	Version                 string          `json:"version"`
}

type template struct {
	Color                              string          `json:"color"`
	DefaultComplianceAssignments       json.RawMessage `json:"default_compliance_assignments"`
	DefaultMdmConfigurationAssignments json.RawMessage `json:"default_mdm_configuration_assignments"`
	DefaultMonitoringAssignments       json.RawMessage `json:"default_monitoring_assignments"`
	DefaultSelfServiceAssignments      json.RawMessage `json:"default_self_service_assignments"`
	DefaultSoftwareAssignments         json.RawMessage `json:"default_software_assignments"`
	Icon                               string          `json:"icon"`
	Id                                 string          `json:"id"`
	Name                               string          `json:"name"`
	Order                              int             `json:"order"`
	ShortDescription                   string          `json:"short_description"`
	Version                            int             `json:"version"`
}

type template_asset struct {
	Id string `json:"id"`
}

type template_assignment struct {
	AssetId string `json:"asset_id"`
	Id      string `json:"id"`
	Name    string `json:"name"`
	Status  string `json:"status"`
}

type template_create_policy_request struct {
	Assets     json.RawMessage `json:"assets"`
	Color      string          `json:"color"`
	Icon       string          `json:"icon"`
	Id         string          `json:"id"`
	ParentId   string          `json:"parent_id"`
	PolicyName string          `json:"policy_name"`
}

type template_query_request struct {
	Ids json.RawMessage `json:"ids"`
}

type template_response struct {
	Items json.RawMessage `json:"items"`
}

type ticketing_policy_device_sync_request struct {
	PolicyId string `json:"policy_id"`
}

type user struct {
	AddigyRole string          `json:"addigy_role"`
	Email      string          `json:"email"`
	Name       string          `json:"name"`
	Orgid      string          `json:"orgid"`
	Phone      string          `json:"phone"`
	Policies   json.RawMessage `json:"policies"`
}

type variable struct {
	CreatedDate    string          `json:"created_date"`
	DefaultValue   json.RawMessage `json:"default_value"`
	Key            string          `json:"key"`
	OrganizationId string          `json:"organization_id"`
	UpdatedDate    string          `json:"updated_date"`
}

type variable_create_request struct {
	DefaultValue json.RawMessage `json:"default_value"`
	Key          string          `json:"key"`
}

type variable_update_request struct {
	DefaultValue json.RawMessage `json:"default_value"`
	Key          string          `json:"key"`
}

type variable_usage struct {
	AssetIds       json.RawMessage `json:"asset_ids"`
	AssetType      string          `json:"asset_type"`
	OrganizationId string          `json:"organization_id"`
	VariableKey    string          `json:"variable_key"`
}

type webhook struct {
	Action         json.RawMessage `json:"action"`
	Disabled       bool            `json:"disabled"`
	Id             string          `json:"id"`
	Name           string          `json:"name"`
	OrganizationId string          `json:"organization_id"`
	SecretToken    string          `json:"secret_token"`
	Trigger        json.RawMessage `json:"trigger"`
}

type webhook_action struct {
	Url string `json:"url"`
}

type webhook_create_request struct {
	Action  json.RawMessage `json:"action"`
	Name    string          `json:"name"`
	Trigger json.RawMessage `json:"trigger"`
}

type webhook_status struct {
	Count string `json:"count"`
}

type webhook_trigger struct {
	ActionEntityIdentifier string `json:"action_entity_identifier"`
	ActionEntityName       string `json:"action_entity_name"`
	ActionEntityType       string `json:"action_entity_type"`
	ActionName             string `json:"action_name"`
	ReceiverIdentifier     string `json:"receiver_identifier"`
	ReceiverName           string `json:"receiver_name"`
	ReceiverType           string `json:"receiver_type"`
	SenderIdentifier       string `json:"sender_identifier"`
	SenderName             string `json:"sender_name"`
	SenderType             string `json:"sender_type"`
}

type webhook_update_request struct {
	Action  json.RawMessage `json:"action"`
	Id      string          `json:"id"`
	Name    string          `json:"name"`
	Trigger json.RawMessage `json:"trigger"`
}

type webhooks_schedule_count struct {
	ErrorCounter   string `json:"error_counter"`
	LastError      string `json:"last_error"`
	LastExecute    string `json:"last_execute"`
	LastProcess    string `json:"last_process"`
	LastSchedule   string `json:"last_schedule"`
	OrganizationId string `json:"organization_id"`
	RunAfter       string `json:"run_after"`
	Status         string `json:"status"`
	WebhookId      string `json:"webhook_id"`
}
