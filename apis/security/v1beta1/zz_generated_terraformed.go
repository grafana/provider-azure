/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this AdvancedThreatProtection
func (mg *AdvancedThreatProtection) GetTerraformResourceType() string {
	return "azurerm_advanced_threat_protection"
}

// GetConnectionDetailsMapping for this AdvancedThreatProtection
func (tr *AdvancedThreatProtection) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this AdvancedThreatProtection
func (tr *AdvancedThreatProtection) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this AdvancedThreatProtection
func (tr *AdvancedThreatProtection) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this AdvancedThreatProtection
func (tr *AdvancedThreatProtection) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this AdvancedThreatProtection
func (tr *AdvancedThreatProtection) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this AdvancedThreatProtection
func (tr *AdvancedThreatProtection) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this AdvancedThreatProtection using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *AdvancedThreatProtection) LateInitialize(attrs []byte) (bool, error) {
	params := &AdvancedThreatProtectionParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *AdvancedThreatProtection) GetTerraformSchemaVersion() int {
	return 1
}

// GetTerraformResourceType returns Terraform resource type for this IOTSecurityDeviceGroup
func (mg *IOTSecurityDeviceGroup) GetTerraformResourceType() string {
	return "azurerm_iot_security_device_group"
}

// GetConnectionDetailsMapping for this IOTSecurityDeviceGroup
func (tr *IOTSecurityDeviceGroup) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this IOTSecurityDeviceGroup
func (tr *IOTSecurityDeviceGroup) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this IOTSecurityDeviceGroup
func (tr *IOTSecurityDeviceGroup) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this IOTSecurityDeviceGroup
func (tr *IOTSecurityDeviceGroup) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this IOTSecurityDeviceGroup
func (tr *IOTSecurityDeviceGroup) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this IOTSecurityDeviceGroup
func (tr *IOTSecurityDeviceGroup) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this IOTSecurityDeviceGroup using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *IOTSecurityDeviceGroup) LateInitialize(attrs []byte) (bool, error) {
	params := &IOTSecurityDeviceGroupParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *IOTSecurityDeviceGroup) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this IOTSecuritySolution
func (mg *IOTSecuritySolution) GetTerraformResourceType() string {
	return "azurerm_iot_security_solution"
}

// GetConnectionDetailsMapping for this IOTSecuritySolution
func (tr *IOTSecuritySolution) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this IOTSecuritySolution
func (tr *IOTSecuritySolution) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this IOTSecuritySolution
func (tr *IOTSecuritySolution) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this IOTSecuritySolution
func (tr *IOTSecuritySolution) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this IOTSecuritySolution
func (tr *IOTSecuritySolution) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this IOTSecuritySolution
func (tr *IOTSecuritySolution) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this IOTSecuritySolution using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *IOTSecuritySolution) LateInitialize(attrs []byte) (bool, error) {
	params := &IOTSecuritySolutionParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *IOTSecuritySolution) GetTerraformSchemaVersion() int {
	return 0
}