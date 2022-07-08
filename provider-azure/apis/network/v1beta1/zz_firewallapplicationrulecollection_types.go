/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type FirewallApplicationRuleCollectionObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FirewallApplicationRuleCollectionParameters struct {

	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// +crossplane:generate:reference:type=Firewall
	// +kubebuilder:validation:Optional
	AzureFirewallName *string `json:"azureFirewallName,omitempty" tf:"azure_firewall_name,omitempty"`

	// +kubebuilder:validation:Optional
	AzureFirewallNameRef *v1.Reference `json:"azureFirewallNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	AzureFirewallNameSelector *v1.Selector `json:"azureFirewallNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Priority *float64 `json:"priority" tf:"priority,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Rule []RuleParameters `json:"rule" tf:"rule,omitempty"`
}

type ProtocolObservation struct {
}

type ProtocolParameters struct {

	// +kubebuilder:validation:Required
	Port *float64 `json:"port" tf:"port,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type RuleObservation struct {
}

type RuleParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	FqdnTags []*string `json:"fqdnTags,omitempty" tf:"fqdn_tags,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Protocol []ProtocolParameters `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	SourceAddresses []*string `json:"sourceAddresses,omitempty" tf:"source_addresses,omitempty"`

	// +kubebuilder:validation:Optional
	SourceIPGroups []*string `json:"sourceIpGroups,omitempty" tf:"source_ip_groups,omitempty"`

	// +kubebuilder:validation:Optional
	TargetFqdns []*string `json:"targetFqdns,omitempty" tf:"target_fqdns,omitempty"`
}

// FirewallApplicationRuleCollectionSpec defines the desired state of FirewallApplicationRuleCollection
type FirewallApplicationRuleCollectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallApplicationRuleCollectionParameters `json:"forProvider"`
}

// FirewallApplicationRuleCollectionStatus defines the observed state of FirewallApplicationRuleCollection.
type FirewallApplicationRuleCollectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallApplicationRuleCollectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallApplicationRuleCollection is the Schema for the FirewallApplicationRuleCollections API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FirewallApplicationRuleCollection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallApplicationRuleCollectionSpec   `json:"spec"`
	Status            FirewallApplicationRuleCollectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallApplicationRuleCollectionList contains a list of FirewallApplicationRuleCollections
type FirewallApplicationRuleCollectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirewallApplicationRuleCollection `json:"items"`
}

// Repository type metadata.
var (
	FirewallApplicationRuleCollection_Kind             = "FirewallApplicationRuleCollection"
	FirewallApplicationRuleCollection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FirewallApplicationRuleCollection_Kind}.String()
	FirewallApplicationRuleCollection_KindAPIVersion   = FirewallApplicationRuleCollection_Kind + "." + CRDGroupVersion.String()
	FirewallApplicationRuleCollection_GroupVersionKind = CRDGroupVersion.WithKind(FirewallApplicationRuleCollection_Kind)
)

func init() {
	SchemeBuilder.Register(&FirewallApplicationRuleCollection{}, &FirewallApplicationRuleCollectionList{})
}