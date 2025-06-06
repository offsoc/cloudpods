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

package notify

import (
	time "time"

	jsonutils "yunion.io/x/jsonutils"

	"yunion.io/x/onecloud/pkg/apis"
)

// SConfig is an autogenerated struct via yunion.io/x/onecloud/pkg/notify/models.SConfig.
type SConfig struct {
	apis.SDomainLevelResourceBase
	Type        string                `json:"type"`
	Content     *SNotifyConfigContent `json:"content"`
	Attribution string                `json:"attribution"`
}

// SEmailQueue is an autogenerated struct via yunion.io/x/onecloud/pkg/notify/models.SEmailQueue.
type SEmailQueue struct {
	apis.SLogBase
	RecvAt          time.Time            `json:"recv_at"`
	Dest            string               `json:"dest"`
	Subject         string               `json:"subject"`
	DestCc          string               `json:"dest_cc"`
	DestBcc         string               `json:"dest_bcc"`
	SessionId       string               `json:"session_id"`
	Content         jsonutils.JSONObject `json:"content"`
	MoreDetails     jsonutils.JSONObject `json:"more_details"`
	ProjectId       string               `json:"project_id"`
	Project         string               `json:"project"`
	ProjectDomainId string               `json:"project_domain_id"`
	ProjectDomain   string               `json:"project_domain"`
	UserId          string               `json:"user_id"`
	User            string               `json:"user"`
	DomainId        string               `json:"domain_id"`
	Domain          string               `json:"domain"`
	Roles           string               `json:"roles"`
}

// SEmailQueueStatus is an autogenerated struct via yunion.io/x/onecloud/pkg/notify/models.SEmailQueueStatus.
type SEmailQueueStatus struct {
	Id      int64     `json:"id"`
	SentAt  time.Time `json:"sent_at"`
	Status  string    `json:"status"`
	Results string    `json:"results"`
}

// SEvent is an autogenerated struct via yunion.io/x/onecloud/pkg/notify/models.SEvent.
type SEvent struct {
	apis.SLogBase
	// 资源创建时间
	CreatedAt    time.Time `json:"created_at"`
	Message      string    `json:"message"`
	Event        string    `json:"event"`
	ResourceType string    `json:"resource_type"`
	Action       string    `json:"action"`
	AdvanceDays  int       `json:"advance_days"`
	TopicId      string    `json:"topic_id"`
}

// SNotification is an autogenerated struct via yunion.io/x/onecloud/pkg/notify/models.SNotification.
type SNotification struct {
	apis.SStatusStandaloneResourceBase
	ContactType string `json:"contact_type"`
	// swagger:ignore
	Topic    string `json:"topic"`
	Priority string `json:"priority"`
	// swagger:ignore
	Message string `json:"message"`
	// swagger:ignore
	TopicType string `json:"topic_type"`
	// swagger:ignore
	TopicId    string    `json:"topic_id"`
	ReceivedAt time.Time `json:"received_at"`
	EventId    string    `json:"event_id"`
	SendTimes  int       `json:"send_times"`
}

// SNotificationGroup is an autogenerated struct via yunion.io/x/onecloud/pkg/notify/models.SNotificationGroup.
type SNotificationGroup struct {
	Id       string `json:"id"`
	GroupKey string `json:"group_key"`
	Title    string `json:"title"`
	// swagger:ignore
	Message     string               `json:"message"`
	ReceiverId  string               `json:"receiver_id"`
	Body        jsonutils.JSONObject `json:"body"`
	Header      jsonutils.JSONObject `json:"header"`
	MsgKey      string               `json:"msg_key"`
	ContactType string               `json:"contact_type"`
	Contact     string               `json:"contact"`
	DomainId    string               `json:"domain_id"`
}

// SNotificationLog is an autogenerated struct via yunion.io/x/onecloud/pkg/notify/models.SNotificationLog.
type SNotificationLog struct {
	apis.SLogBase
	ContactType string `json:"contact_type"`
	// swagger:ignore
	Topic    string `json:"topic"`
	Priority string `json:"priority"`
	// swagger:ignore
	Message string `json:"message"`
	// swagger:ignore
	TopicType  string    `json:"topic_type"`
	ReceivedAt time.Time `json:"received_at"`
	EventId    string    `json:"event_id"`
	SendTimes  int       `json:"send_times"`
}

// SReceiver is an autogenerated struct via yunion.io/x/onecloud/pkg/notify/models.SReceiver.
type SReceiver struct {
	apis.SVirtualResourceBase
	apis.SEnabledResourceBase
	Email string `json:"email"`
	// swagger:ignore
	Mobile string `json:"mobile"`
	Lang   string `json:"lang"`
	// swagger:ignore
	EnabledEmail *bool `json:"enabled_email,omitempty"`
	// swagger:ignore
	VerifiedEmail *bool `json:"verified_email,omitempty"`
	// swagger:ignore
	EnabledMobile *bool `json:"enabled_mobile,omitempty"`
	// swagger:ignore
	VerifiedMobile *bool `json:"verified_mobile,omitempty"`
}

// SRobot is an autogenerated struct via yunion.io/x/onecloud/pkg/notify/models.SRobot.
type SRobot struct {
	apis.SSharableVirtualResourceBase
	apis.SEnabledResourceBase
	Type        string               `json:"type"`
	Address     string               `json:"address"`
	Lang        string               `json:"lang"`
	Header      jsonutils.JSONObject `json:"header"`
	Body        jsonutils.JSONObject `json:"body"`
	MsgKey      string               `json:"msg_key"`
	UseTemplate *bool                `json:"use_template,omitempty"`
}

// SSubscriber is an autogenerated struct via yunion.io/x/onecloud/pkg/notify/models.SSubscriber.
type SSubscriber struct {
	apis.SStandaloneAnonResourceBase
	apis.SEnabledResourceBase
	TopicId                 string `json:"topic_id"`
	Type                    string `json:"type"`
	Identification          string `json:"identification"`
	RoleScope               string `json:"role_scope"`
	ResourceScope           string `json:"resource_scope"`
	ResourceAttributionId   string `json:"resource_attribution_id"`
	ResourceAttributionName string `json:"resource_attribution_name"`
	Scope                   string `json:"scope"`
	DomainId                string `json:"domain_id"`
	// minutes
	GroupTimes uint32 `json:"group_times"`
}

// SSubscriberReceiver is an autogenerated struct via yunion.io/x/onecloud/pkg/notify/models.SSubscriberReceiver.
type SSubscriberReceiver struct {
	apis.SJointResourceBase
	SubscriberId string `json:"subscriber_id"`
	ReceiverId   string `json:"receiver_id"`
}

// STemplate is an autogenerated struct via yunion.io/x/onecloud/pkg/notify/models.STemplate.
type STemplate struct {
	apis.SStandaloneAnonResourceBase
	ContactType string `json:"contact_type"`
	Topic       string `json:"topic"`
	// title | content | remote
	TemplateType string `json:"template_type"`
	Content      string `json:"content"`
	Lang         string `json:"lang"`
	Example      string `json:"example"`
}

// STopic is an autogenerated struct via yunion.io/x/onecloud/pkg/notify/models.STopic.
type STopic struct {
	apis.SEnabledStatusStandaloneResourceBase
	Type              string           `json:"type"`
	Results           *bool            `json:"results,omitempty"`
	TitleCn           string           `json:"title_cn"`
	TitleEn           string           `json:"title_en"`
	ContentCn         string           `json:"content_cn"`
	ContentEn         string           `json:"content_en"`
	GroupKeys         *STopicGroupKeys `json:"group_keys"`
	AdvanceDays       []int            `json:"advance_days"`
	WebconsoleDisable *bool            `json:"webconsole_disable,omitempty"`
}

// STopicAction is an autogenerated struct via yunion.io/x/onecloud/pkg/notify/models.STopicAction.
type STopicAction struct {
	apis.SResourceBase
	ActionId string `json:"action_id"`
	TopicId  string `json:"topic_id"`
}

// STopicResource is an autogenerated struct via yunion.io/x/onecloud/pkg/notify/models.STopicResource.
type STopicResource struct {
	apis.SResourceBase
	ResourceId string `json:"resource_id"`
	TopicId    string `json:"topic_id"`
}
