/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this ContactPoint
func (mg *ContactPoint) GetTerraformResourceType() string {
	return "grafana_contact_point"
}

// GetConnectionDetailsMapping for this ContactPoint
func (tr *ContactPoint) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"alertmanager[*].basic_auth_password": "spec.forProvider.alertmanager[*].basicAuthPasswordSecretRef", "alertmanager[*].settings": "spec.forProvider.alertmanager[*].settingsSecretRef", "dingding[*].settings": "spec.forProvider.dingding[*].settingsSecretRef", "discord[*].settings": "spec.forProvider.discord[*].settingsSecretRef", "discord[*].url": "spec.forProvider.discord[*].urlSecretRef", "email[*].settings": "spec.forProvider.email[*].settingsSecretRef", "googlechat[*].settings": "spec.forProvider.googlechat[*].settingsSecretRef", "googlechat[*].url": "spec.forProvider.googlechat[*].urlSecretRef", "kafka[*].rest_proxy_url": "spec.forProvider.kafka[*].restProxyUrlSecretRef", "kafka[*].settings": "spec.forProvider.kafka[*].settingsSecretRef", "opsgenie[*].api_key": "spec.forProvider.opsgenie[*].apiKeySecretRef", "opsgenie[*].settings": "spec.forProvider.opsgenie[*].settingsSecretRef", "pagerduty[*].integration_key": "spec.forProvider.pagerduty[*].integrationKeySecretRef", "pagerduty[*].settings": "spec.forProvider.pagerduty[*].settingsSecretRef", "pushover[*].api_token": "spec.forProvider.pushover[*].apiTokenSecretRef", "pushover[*].settings": "spec.forProvider.pushover[*].settingsSecretRef", "pushover[*].user_key": "spec.forProvider.pushover[*].userKeySecretRef", "sensugo[*].api_key": "spec.forProvider.sensugo[*].apiKeySecretRef", "sensugo[*].settings": "spec.forProvider.sensugo[*].settingsSecretRef", "slack[*].settings": "spec.forProvider.slack[*].settingsSecretRef", "slack[*].token": "spec.forProvider.slack[*].tokenSecretRef", "slack[*].url": "spec.forProvider.slack[*].urlSecretRef", "teams[*].settings": "spec.forProvider.teams[*].settingsSecretRef", "teams[*].url": "spec.forProvider.teams[*].urlSecretRef", "telegram[*].settings": "spec.forProvider.telegram[*].settingsSecretRef", "telegram[*].token": "spec.forProvider.telegram[*].tokenSecretRef", "threema[*].api_secret": "spec.forProvider.threema[*].apiSecretSecretRef", "threema[*].settings": "spec.forProvider.threema[*].settingsSecretRef", "victorops[*].settings": "spec.forProvider.victorops[*].settingsSecretRef", "webhook[*].authorization_credentials": "spec.forProvider.webhook[*].authorizationCredentialsSecretRef", "webhook[*].basic_auth_password": "spec.forProvider.webhook[*].basicAuthPasswordSecretRef", "webhook[*].settings": "spec.forProvider.webhook[*].settingsSecretRef", "wecom[*].settings": "spec.forProvider.wecom[*].settingsSecretRef", "wecom[*].url": "spec.forProvider.wecom[*].urlSecretRef"}
}

// GetObservation of this ContactPoint
func (tr *ContactPoint) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ContactPoint
func (tr *ContactPoint) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ContactPoint
func (tr *ContactPoint) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ContactPoint
func (tr *ContactPoint) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ContactPoint
func (tr *ContactPoint) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ContactPoint using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ContactPoint) LateInitialize(attrs []byte) (bool, error) {
	params := &ContactPointParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ContactPoint) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this MessageTemplate
func (mg *MessageTemplate) GetTerraformResourceType() string {
	return "grafana_message_template"
}

// GetConnectionDetailsMapping for this MessageTemplate
func (tr *MessageTemplate) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this MessageTemplate
func (tr *MessageTemplate) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this MessageTemplate
func (tr *MessageTemplate) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this MessageTemplate
func (tr *MessageTemplate) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this MessageTemplate
func (tr *MessageTemplate) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this MessageTemplate
func (tr *MessageTemplate) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this MessageTemplate using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *MessageTemplate) LateInitialize(attrs []byte) (bool, error) {
	params := &MessageTemplateParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *MessageTemplate) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this MuteTiming
func (mg *MuteTiming) GetTerraformResourceType() string {
	return "grafana_mute_timing"
}

// GetConnectionDetailsMapping for this MuteTiming
func (tr *MuteTiming) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this MuteTiming
func (tr *MuteTiming) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this MuteTiming
func (tr *MuteTiming) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this MuteTiming
func (tr *MuteTiming) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this MuteTiming
func (tr *MuteTiming) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this MuteTiming
func (tr *MuteTiming) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this MuteTiming using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *MuteTiming) LateInitialize(attrs []byte) (bool, error) {
	params := &MuteTimingParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *MuteTiming) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this NotificationPolicy
func (mg *NotificationPolicy) GetTerraformResourceType() string {
	return "grafana_notification_policy"
}

// GetConnectionDetailsMapping for this NotificationPolicy
func (tr *NotificationPolicy) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this NotificationPolicy
func (tr *NotificationPolicy) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this NotificationPolicy
func (tr *NotificationPolicy) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this NotificationPolicy
func (tr *NotificationPolicy) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this NotificationPolicy
func (tr *NotificationPolicy) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this NotificationPolicy
func (tr *NotificationPolicy) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this NotificationPolicy using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *NotificationPolicy) LateInitialize(attrs []byte) (bool, error) {
	params := &NotificationPolicyParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *NotificationPolicy) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this RuleGroup
func (mg *RuleGroup) GetTerraformResourceType() string {
	return "grafana_rule_group"
}

// GetConnectionDetailsMapping for this RuleGroup
func (tr *RuleGroup) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this RuleGroup
func (tr *RuleGroup) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this RuleGroup
func (tr *RuleGroup) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this RuleGroup
func (tr *RuleGroup) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this RuleGroup
func (tr *RuleGroup) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this RuleGroup
func (tr *RuleGroup) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this RuleGroup using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *RuleGroup) LateInitialize(attrs []byte) (bool, error) {
	params := &RuleGroupParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *RuleGroup) GetTerraformSchemaVersion() int {
	return 0
}