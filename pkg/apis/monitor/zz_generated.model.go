// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by model-api-gen. DO NOT EDIT.

package monitor

import (
	time "time"

	jsonutils "yunion.io/x/jsonutils"

	"yunion.io/x/onecloud/pkg/apis"
)

// SAlert is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SAlert.
type SAlert struct {
	// db.SVirtualResourceBase
	apis.SEnabledResourceBase
	apis.SStatusStandaloneResourceBase
	SMonitorScopedResource
	// Frequency is evaluate period
	Frequency int64         `json:"frequency"`
	Settings  *AlertSetting `json:"settings"`
	Level     string        `json:"level"`
	Message   string        `json:"message"`
	UsedBy    string        `json:"used_by"`
	// Silenced       bool
	ExecutionError string `json:"execution_error"`
	// If an alert rule has a configured `For` and the query violates the configured threshold
	// it will first go from `OK` to `Pending`. Going from `OK` to `Pending` will not send any
	// notifications. Once the alert rule has been firing for more than `For` duration, it will
	// change to `Alerting` and send alert notifications.
	For                 int64                `json:"for"`
	EvalData            jsonutils.JSONObject `json:"eval_data"`
	State               string               `json:"state"`
	NoDataState         string               `json:"no_data_state"`
	ExecutionErrorState string               `json:"execution_error_state"`
	LastStateChange     time.Time            `json:"last_state_change"`
	StateChanges        int                  `json:"state_changes"`
	CustomizeConfig     jsonutils.JSONObject `json:"customize_config"`
	ResType             string               `json:"res_type"`
}

// SAlertDashBoard is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SAlertDashBoard.
type SAlertDashBoard struct {
	// db.SVirtualResourceBase
	apis.SEnabledResourceBase
	apis.SStatusStandaloneResourceBase
	apis.SScopedResourceBase
	Refresh string `json:"refresh"`
}

// SAlertDashboardPanel is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SAlertDashboardPanel.
type SAlertDashboardPanel struct {
	apis.SVirtualJointResourceBase
	DashboardId string `json:"dashboard_id"`
	PanelId     string `json:"panel_id"`
	Index       int    `json:"index"`
}

// SAlertPanel is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SAlertPanel.
type SAlertPanel struct {
	apis.SStatusStandaloneResourceBase
	apis.SScopedResourceBase
	Settings *AlertSetting `json:"settings"`
	Message  string        `json:"message"`
}

// SAlertRecord is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SAlertRecord.
type SAlertRecord struct {
	// db.SVirtualResourceBase
	apis.SEnabledResourceBase
	apis.SStandaloneAnonResourceBase
	SMonitorScopedResource
	AlertId   string               `json:"alert_id"`
	Level     string               `json:"level"`
	State     string               `json:"state"`
	SendState string               `json:"send_state"`
	EvalData  jsonutils.JSONObject `json:"eval_data"`
	AlertRule jsonutils.JSONObject `json:"alert_rule"`
	ResType   string               `json:"res_type"`
	ResIds    string               `json:"res_ids"`
}

// SAlertRecordShield is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SAlertRecordShield.
type SAlertRecordShield struct {
	// db.SVirtualResourceBase
	apis.SEnabledResourceBase
	apis.SStatusStandaloneResourceBase
	SMonitorScopedResource
	AlertId   string    `json:"alert_id"`
	ResId     string    `json:"res_id"`
	ResType   string    `json:"res_type"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

// SAlertResource is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SAlertResource.
type SAlertResource struct {
	apis.SStandaloneResourceBase
	Type string `json:"type"`
}

// SCommonAlert is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SCommonAlert.
type SCommonAlert struct {
	SAlert
}

// SDataSource is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SDataSource.
type SDataSource struct {
	apis.SStandaloneResourceBase
	Type      string `json:"type"`
	Url       string `json:"url"`
	User      string `json:"user"`
	Password  string `json:"password"`
	Database  string `json:"database"`
	IsDefault *bool  `json:"is_default,omitempty"`
}

// SMeterAlert is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SMeterAlert.
type SMeterAlert struct {
	SV1Alert
}

// SMetric is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SMetric.
type SMetric struct {
	apis.SVirtualJointResourceBase
	MeasurementId string `json:"measurement_id"`
	FieldId       string `json:"field_id"`
}

// SMetricField is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SMetricField.
type SMetricField struct {
	// db.SVirtualResourceBase
	apis.SEnabledResourceBase
	apis.SStatusStandaloneResourceBase
	apis.SScopedResourceBase
	DisplayName string `json:"display_name"`
	Unit        string `json:"unit"`
	ValueType   string `json:"value_type"`
	Score       int    `json:"score"`
}

// SMetricMeasurement is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SMetricMeasurement.
type SMetricMeasurement struct {
	// db.SVirtualResourceBase
	apis.SEnabledResourceBase
	apis.SStatusStandaloneResourceBase
	apis.SScopedResourceBase
	ResType     string `json:"res_type"`
	Database    string `json:"database"`
	DisplayName string `json:"display_name"`
	Score       int    `json:"score"`
}

// SMigrationAlert is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SMigrationAlert.
type SMigrationAlert struct {
	SAlert
	MetricType   string               `json:"metric_type"`
	MigrateNotes jsonutils.JSONObject `json:"migrate_notes"`
}

// SMonitorResource is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SMonitorResource.
type SMonitorResource struct {
	apis.SVirtualResourceBase
	apis.SEnabledResourceBase
	AlertState string `json:"alert_state"`
	ResId      string `json:"res_id"`
	ResType    string `json:"res_type"`
}

// SMonitorResourceAlert is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SMonitorResourceAlert.
type SMonitorResourceAlert struct {
	apis.SJointResourceBase
	MonitorResourceId string               `json:"monitor_resource_id"`
	AlertId           string               `json:"alert_id"`
	AlertRecordId     string               `json:"alert_record_id"`
	ResType           string               `json:"res_type"`
	Metric            string               `json:"metric"`
	AlertState        string               `json:"alert_state"`
	SendState         string               `json:"send_state"`
	TriggerTime       time.Time            `json:"trigger_time"`
	Data              jsonutils.JSONObject `json:"data"`
}

// SMonitorScopedResource is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SMonitorScopedResource.
type SMonitorScopedResource struct {
	apis.SScopedResourceBase
}

// SNodeAlert is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SNodeAlert.
type SNodeAlert struct {
	SCommonAlert
}

// SNotification is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SNotification.
type SNotification struct {
	apis.SVirtualResourceBase
	Type                  string `json:"type"`
	IsDefault             bool   `json:"is_default"`
	SendReminder          bool   `json:"send_reminder"`
	DisableResolveMessage bool   `json:"disable_resolve_message"`
	// unit is second
	Frequency            int64                `json:"frequency"`
	Settings             jsonutils.JSONObject `json:"settings"`
	LastSendNotification time.Time            `json:"last_send_notification"`
}

// SV1Alert is an autogenerated struct via yunion.io/x/onecloud/pkg/monitor/models.SV1Alert.
type SV1Alert struct {
	SAlert
}
