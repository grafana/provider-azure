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

type AccessControlObservation struct {
}

type AccessControlParameters struct {

	// One or more akamai_signature_header_authentication_key blocks as defined below.
	// +kubebuilder:validation:Optional
	AkamaiSignatureHeaderAuthenticationKey []AkamaiSignatureHeaderAuthenticationKeyParameters `json:"akamaiSignatureHeaderAuthenticationKey,omitempty" tf:"akamai_signature_header_authentication_key,omitempty"`

	// A ip block as defined below.
	// +kubebuilder:validation:Optional
	IPAllow []IPAllowParameters `json:"ipAllow,omitempty" tf:"ip_allow,omitempty"`
}

type AkamaiSignatureHeaderAuthenticationKeyObservation struct {
}

type AkamaiSignatureHeaderAuthenticationKeyParameters struct {

	// Authentication key.
	// +kubebuilder:validation:Optional
	Base64Key *string `json:"base64Key,omitempty" tf:"base64_key,omitempty"`

	// The expiration time of the authentication key.
	// +kubebuilder:validation:Optional
	Expiration *string `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// Identifier of the key.
	// +kubebuilder:validation:Optional
	Identifier *string `json:"identifier,omitempty" tf:"identifier,omitempty"`
}

type IPAllowObservation struct {
}

type IPAllowParameters struct {

	// The IP address to allow.
	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The name which should be used for this Streaming Endpoint maximum length is 24. Changing this forces a new Streaming Endpoint to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The subnet mask prefix length (see CIDR notation).
	// +kubebuilder:validation:Optional
	SubnetPrefixLength *float64 `json:"subnetPrefixLength,omitempty" tf:"subnet_prefix_length,omitempty"`
}

type StreamingEndpointCrossSiteAccessPolicyObservation struct {
}

type StreamingEndpointCrossSiteAccessPolicyParameters struct {

	// The content of clientaccesspolicy.xml used by Silverlight.
	// +kubebuilder:validation:Optional
	ClientAccessPolicy *string `json:"clientAccessPolicy,omitempty" tf:"client_access_policy,omitempty"`

	// The content of crossdomain.xml used by Silverlight.
	// +kubebuilder:validation:Optional
	CrossDomainPolicy *string `json:"crossDomainPolicy,omitempty" tf:"cross_domain_policy,omitempty"`
}

type StreamingEndpointObservation struct {

	// The host name of the Streaming Endpoint.
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// The ID of the Streaming Endpoint.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type StreamingEndpointParameters struct {

	// A access_control block as defined below.
	// +kubebuilder:validation:Optional
	AccessControl []AccessControlParameters `json:"accessControl,omitempty" tf:"access_control,omitempty"`

	// The flag indicates if the resource should be automatically started on creation.
	// +kubebuilder:validation:Optional
	AutoStartEnabled *bool `json:"autoStartEnabled,omitempty" tf:"auto_start_enabled,omitempty"`

	// The CDN enabled flag.
	// +kubebuilder:validation:Optional
	CdnEnabled *bool `json:"cdnEnabled,omitempty" tf:"cdn_enabled,omitempty"`

	// The CDN profile name.
	// +kubebuilder:validation:Optional
	CdnProfile *string `json:"cdnProfile,omitempty" tf:"cdn_profile,omitempty"`

	// The CDN provider name. Supported value are StandardVerizon,PremiumVerizon and StandardAkamai
	// +kubebuilder:validation:Optional
	CdnProvider *string `json:"cdnProvider,omitempty" tf:"cdn_provider,omitempty"`

	// A cross_site_access_policy block as defined below.
	// +kubebuilder:validation:Optional
	CrossSiteAccessPolicy []StreamingEndpointCrossSiteAccessPolicyParameters `json:"crossSiteAccessPolicy,omitempty" tf:"cross_site_access_policy,omitempty"`

	// The custom host names of the streaming endpoint.
	// +kubebuilder:validation:Optional
	CustomHostNames []*string `json:"customHostNames,omitempty" tf:"custom_host_names,omitempty"`

	// The streaming endpoint description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The Azure Region where the Streaming Endpoint should exist. Changing this forces a new Streaming Endpoint to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Max cache age in seconds.
	// +kubebuilder:validation:Optional
	MaxCacheAgeSeconds *float64 `json:"maxCacheAgeSeconds,omitempty" tf:"max_cache_age_seconds,omitempty"`

	// The Media Services account name. Changing this forces a new Streaming Endpoint to be created.
	// +crossplane:generate:reference:type=ServicesAccount
	// +kubebuilder:validation:Optional
	MediaServicesAccountName *string `json:"mediaServicesAccountName,omitempty" tf:"media_services_account_name,omitempty"`

	// Reference to a ServicesAccount to populate mediaServicesAccountName.
	// +kubebuilder:validation:Optional
	MediaServicesAccountNameRef *v1.Reference `json:"mediaServicesAccountNameRef,omitempty" tf:"-"`

	// Selector for a ServicesAccount to populate mediaServicesAccountName.
	// +kubebuilder:validation:Optional
	MediaServicesAccountNameSelector *v1.Selector `json:"mediaServicesAccountNameSelector,omitempty" tf:"-"`

	// The name of the Resource Group where the Streaming Endpoint should exist. Changing this forces a new Streaming Endpoint to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The number of scale units. To create a Standard Streaming Endpoint set 0. For Premium Streaming Endpoint valid values are between 1 and 10.
	// +kubebuilder:validation:Required
	ScaleUnits *float64 `json:"scaleUnits" tf:"scale_units,omitempty"`

	// A mapping of tags which should be assigned to the Streaming Endpoint.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// StreamingEndpointSpec defines the desired state of StreamingEndpoint
type StreamingEndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StreamingEndpointParameters `json:"forProvider"`
}

// StreamingEndpointStatus defines the observed state of StreamingEndpoint.
type StreamingEndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StreamingEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StreamingEndpoint is the Schema for the StreamingEndpoints API. Manages a Streaming Endpoint.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type StreamingEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StreamingEndpointSpec   `json:"spec"`
	Status            StreamingEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StreamingEndpointList contains a list of StreamingEndpoints
type StreamingEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StreamingEndpoint `json:"items"`
}

// Repository type metadata.
var (
	StreamingEndpoint_Kind             = "StreamingEndpoint"
	StreamingEndpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StreamingEndpoint_Kind}.String()
	StreamingEndpoint_KindAPIVersion   = StreamingEndpoint_Kind + "." + CRDGroupVersion.String()
	StreamingEndpoint_GroupVersionKind = CRDGroupVersion.WithKind(StreamingEndpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&StreamingEndpoint{}, &StreamingEndpointList{})
}