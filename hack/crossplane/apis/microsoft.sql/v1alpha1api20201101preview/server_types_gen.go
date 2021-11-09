// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101preview

import (
	"github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:rbac:groups=microsoft.sql.azure.com,resources=servers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=microsoft.sql.azure.com,resources={servers/status,servers/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
//Generated from: https://schema.management.azure.com/schemas/2020-11-01-preview/Microsoft.Sql.json#/resourceDefinitions/servers
type Server struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Servers_Spec  `json:"spec,omitempty"`
	Status            Server_Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
//Generated from: https://schema.management.azure.com/schemas/2020-11-01-preview/Microsoft.Sql.json#/resourceDefinitions/servers
type ServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Server `json:"items"`
}

type Server_Status struct {
	v1alpha1.ResourceStatus `json:",inline"`
	AtProvider              ServerObservation `json:"atProvider"`
}

// +kubebuilder:validation:Enum={"2020-11-01-preview"}
type ServersSpecAPIVersion string

const ServersSpecAPIVersion20201101Preview = ServersSpecAPIVersion("2020-11-01-preview")

type Servers_Spec struct {
	v1alpha1.ResourceSpec `json:",inline"`
	ForProvider           ServersParameters `json:"forProvider"`
}

type ServerObservation struct {
	//AdministratorLogin: Administrator username for the server. Once created it
	//cannot be changed.
	AdministratorLogin *string `json:"administratorLogin,omitempty"`

	//AdministratorLoginPassword: The administrator login password (required for
	//server creation).
	AdministratorLoginPassword *string `json:"administratorLoginPassword,omitempty"`

	//Administrators: The Azure Active Directory identity of the server.
	Administrators *ServerExternalAdministrator_Status `json:"administrators,omitempty"`

	//FullyQualifiedDomainName: The fully qualified domain name of the server.
	FullyQualifiedDomainName *string `json:"fullyQualifiedDomainName,omitempty"`

	//Id: Resource ID.
	Id *string `json:"id,omitempty"`

	//Identity: The Azure Active Directory identity of the server.
	Identity *ResourceIdentity_Status `json:"identity,omitempty"`

	//KeyId: A CMK URI of the key to use for encryption.
	KeyId *string `json:"keyId,omitempty"`

	//Kind: Kind of sql server. This is metadata used for the Azure portal experience.
	Kind *string `json:"kind,omitempty"`

	//Location: Resource location.
	Location *string `json:"location,omitempty"`

	//MinimalTlsVersion: Minimal TLS version. Allowed values: '1.0', '1.1', '1.2'
	MinimalTlsVersion *string `json:"minimalTlsVersion,omitempty"`

	//Name: Resource name.
	Name *string `json:"name,omitempty"`

	//PrimaryUserAssignedIdentityId: The resource id of a user assigned identity to be
	//used by default.
	PrimaryUserAssignedIdentityId *string `json:"primaryUserAssignedIdentityId,omitempty"`

	//PrivateEndpointConnections: List of private endpoint connections on a server
	PrivateEndpointConnections []ServerPrivateEndpointConnection_Status `json:"privateEndpointConnections,omitempty"`

	//PublicNetworkAccess: Whether or not public endpoint access is allowed for this
	//server.  Value is optional but if passed in, must be 'Enabled' or 'Disabled'
	PublicNetworkAccess *ServerPropertiesStatusPublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	//State: The state of the server.
	State *string `json:"state,omitempty"`

	//Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	//Type: Resource type.
	Type *string `json:"type,omitempty"`

	//Version: The version of the server.
	Version *string `json:"version,omitempty"`

	//WorkspaceFeature: Whether or not existing server has a workspace created and if
	//it allows connection from workspace
	WorkspaceFeature *ServerPropertiesStatusWorkspaceFeature `json:"workspaceFeature,omitempty"`
}

type ServersParameters struct {
	//AdministratorLogin: Administrator username for the server. Once created it
	//cannot be changed.
	AdministratorLogin *string `json:"administratorLogin,omitempty"`

	//AdministratorLoginPassword: The administrator login password (required for
	//server creation).
	AdministratorLoginPassword *string `json:"administratorLoginPassword,omitempty"`

	//Administrators: Properties of a active directory administrator.
	Administrators *ServerExternalAdministrator `json:"administrators,omitempty"`

	//Identity: Azure Active Directory identity configuration for a resource.
	Identity *ResourceIdentity `json:"identity,omitempty"`

	//KeyId: A CMK URI of the key to use for encryption.
	KeyId *string `json:"keyId,omitempty"`

	//Location: Location to deploy resource to
	Location string `json:"location,omitempty"`

	//MinimalTlsVersion: Minimal TLS version. Allowed values: '1.0', '1.1', '1.2'
	MinimalTlsVersion *string `json:"minimalTlsVersion,omitempty"`

	// +kubebuilder:validation:Required
	//Name: The name of the server.
	Name string `json:"name"`

	//PrimaryUserAssignedIdentityId: The resource id of a user assigned identity to be
	//used by default.
	PrimaryUserAssignedIdentityId *string `json:"primaryUserAssignedIdentityId,omitempty"`

	//PublicNetworkAccess: Whether or not public endpoint access is allowed for this
	//server.  Value is optional but if passed in, must be 'Enabled' or 'Disabled'.
	PublicNetworkAccess       *ServerPropertiesPublicNetworkAccess `json:"publicNetworkAccess,omitempty"`
	ResourceGroupName         string                               `json:"resourceGroupName"`
	ResourceGroupNameRef      *v1alpha1.Reference                  `json:"resourceGroupNameRef,omitempty"`
	ResourceGroupNameSelector *v1alpha1.Selector                   `json:"resourceGroupNameSelector,omitempty"`

	//Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`

	//Version: The version of the server.
	Version *string `json:"version,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-11-01-preview/Microsoft.Sql.json#/definitions/ResourceIdentity
type ResourceIdentity struct {
	//Type: The identity type. Set this to 'SystemAssigned' in order to automatically
	//create and assign an Azure Active Directory principal for the resource.
	Type *ResourceIdentityType `json:"type,omitempty"`
}

type ResourceIdentity_Status struct {
	//PrincipalId: The Azure Active Directory principal id.
	PrincipalId *string `json:"principalId,omitempty"`

	//TenantId: The Azure Active Directory tenant id.
	TenantId *string `json:"tenantId,omitempty"`

	//Type: The identity type. Set this to 'SystemAssigned' in order to automatically
	//create and assign an Azure Active Directory principal for the resource.
	Type *ResourceIdentityStatusType `json:"type,omitempty"`

	//UserAssignedIdentities: The resource ids of the user assigned identities to use
	UserAssignedIdentities map[string]UserIdentity_Status `json:"userAssignedIdentities,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2020-11-01-preview/Microsoft.Sql.json#/definitions/ServerExternalAdministrator
type ServerExternalAdministrator struct {
	//AdministratorType: Type of the sever administrator.
	AdministratorType *ServerExternalAdministratorAdministratorType `json:"administratorType,omitempty"`

	//AzureADOnlyAuthentication: Azure Active Directory only Authentication enabled.
	AzureADOnlyAuthentication *bool `json:"azureADOnlyAuthentication,omitempty"`

	//Login: Login name of the server administrator.
	Login *string `json:"login,omitempty"`

	//PrincipalType: Principal Type of the sever administrator.
	PrincipalType *ServerExternalAdministratorPrincipalType `json:"principalType,omitempty"`

	// +kubebuilder:validation:Pattern="^[0-9a-fA-F]{8}(-[0-9a-fA-F]{4}){3}-[0-9a-fA-F]{12}$"
	//Sid: SID (object ID) of the server administrator.
	Sid *string `json:"sid,omitempty"`

	// +kubebuilder:validation:Pattern="^[0-9a-fA-F]{8}(-[0-9a-fA-F]{4}){3}-[0-9a-fA-F]{12}$"
	//TenantId: Tenant ID of the administrator.
	TenantId *string `json:"tenantId,omitempty"`
}

type ServerExternalAdministrator_Status struct {
	//AdministratorType: Type of the sever administrator.
	AdministratorType *ServerExternalAdministratorStatusAdministratorType `json:"administratorType,omitempty"`

	//AzureADOnlyAuthentication: Azure Active Directory only Authentication enabled.
	AzureADOnlyAuthentication *bool `json:"azureADOnlyAuthentication,omitempty"`

	//Login: Login name of the server administrator.
	Login *string `json:"login,omitempty"`

	//PrincipalType: Principal Type of the sever administrator.
	PrincipalType *ServerExternalAdministratorStatusPrincipalType `json:"principalType,omitempty"`

	//Sid: SID (object ID) of the server administrator.
	Sid *string `json:"sid,omitempty"`

	//TenantId: Tenant ID of the administrator.
	TenantId *string `json:"tenantId,omitempty"`
}

type ServerPrivateEndpointConnection_Status struct {
	//Id: Resource ID.
	Id *string `json:"id,omitempty"`

	//Properties: Private endpoint connection properties
	Properties *PrivateEndpointConnectionProperties_Status `json:"properties,omitempty"`
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type ServerPropertiesPublicNetworkAccess string

const (
	ServerPropertiesPublicNetworkAccessDisabled = ServerPropertiesPublicNetworkAccess("Disabled")
	ServerPropertiesPublicNetworkAccessEnabled  = ServerPropertiesPublicNetworkAccess("Enabled")
)

type ServerPropertiesStatusPublicNetworkAccess string

const (
	ServerPropertiesStatusPublicNetworkAccessDisabled = ServerPropertiesStatusPublicNetworkAccess("Disabled")
	ServerPropertiesStatusPublicNetworkAccessEnabled  = ServerPropertiesStatusPublicNetworkAccess("Enabled")
)

type ServerPropertiesStatusWorkspaceFeature string

const (
	ServerPropertiesStatusWorkspaceFeatureConnected    = ServerPropertiesStatusWorkspaceFeature("Connected")
	ServerPropertiesStatusWorkspaceFeatureDisconnected = ServerPropertiesStatusWorkspaceFeature("Disconnected")
)

type PrivateEndpointConnectionProperties_Status struct {
	//PrivateEndpoint: Private endpoint which the connection belongs to.
	PrivateEndpoint *PrivateEndpointProperty_Status `json:"privateEndpoint,omitempty"`

	//PrivateLinkServiceConnectionState: Connection state of the private endpoint
	//connection.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateProperty_Status `json:"privateLinkServiceConnectionState,omitempty"`

	//ProvisioningState: State of the private endpoint connection.
	ProvisioningState *PrivateEndpointConnectionPropertiesStatusProvisioningState `json:"provisioningState,omitempty"`
}

type ResourceIdentityStatusType string

const (
	ResourceIdentityStatusTypeNone                       = ResourceIdentityStatusType("None")
	ResourceIdentityStatusTypeSystemAssigned             = ResourceIdentityStatusType("SystemAssigned")
	ResourceIdentityStatusTypeSystemAssignedUserAssigned = ResourceIdentityStatusType("SystemAssigned,UserAssigned")
	ResourceIdentityStatusTypeUserAssigned               = ResourceIdentityStatusType("UserAssigned")
)

// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned,UserAssigned","UserAssigned"}
type ResourceIdentityType string

const (
	ResourceIdentityTypeNone                       = ResourceIdentityType("None")
	ResourceIdentityTypeSystemAssigned             = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeSystemAssignedUserAssigned = ResourceIdentityType("SystemAssigned,UserAssigned")
	ResourceIdentityTypeUserAssigned               = ResourceIdentityType("UserAssigned")
)

// +kubebuilder:validation:Enum={"ActiveDirectory"}
type ServerExternalAdministratorAdministratorType string

const ServerExternalAdministratorAdministratorTypeActiveDirectory = ServerExternalAdministratorAdministratorType("ActiveDirectory")

// +kubebuilder:validation:Enum={"Application","Group","User"}
type ServerExternalAdministratorPrincipalType string

const (
	ServerExternalAdministratorPrincipalTypeApplication = ServerExternalAdministratorPrincipalType("Application")
	ServerExternalAdministratorPrincipalTypeGroup       = ServerExternalAdministratorPrincipalType("Group")
	ServerExternalAdministratorPrincipalTypeUser        = ServerExternalAdministratorPrincipalType("User")
)

type ServerExternalAdministratorStatusAdministratorType string

const ServerExternalAdministratorStatusAdministratorTypeActiveDirectory = ServerExternalAdministratorStatusAdministratorType("ActiveDirectory")

type ServerExternalAdministratorStatusPrincipalType string

const (
	ServerExternalAdministratorStatusPrincipalTypeApplication = ServerExternalAdministratorStatusPrincipalType("Application")
	ServerExternalAdministratorStatusPrincipalTypeGroup       = ServerExternalAdministratorStatusPrincipalType("Group")
	ServerExternalAdministratorStatusPrincipalTypeUser        = ServerExternalAdministratorStatusPrincipalType("User")
)

type UserIdentity_Status struct {
	//ClientId: The Azure Active Directory client id.
	ClientId *string `json:"clientId,omitempty"`

	//PrincipalId: The Azure Active Directory principal id.
	PrincipalId *string `json:"principalId,omitempty"`
}

type PrivateEndpointConnectionPropertiesStatusProvisioningState string

const (
	PrivateEndpointConnectionPropertiesStatusProvisioningStateApproving = PrivateEndpointConnectionPropertiesStatusProvisioningState("Approving")
	PrivateEndpointConnectionPropertiesStatusProvisioningStateDropping  = PrivateEndpointConnectionPropertiesStatusProvisioningState("Dropping")
	PrivateEndpointConnectionPropertiesStatusProvisioningStateFailed    = PrivateEndpointConnectionPropertiesStatusProvisioningState("Failed")
	PrivateEndpointConnectionPropertiesStatusProvisioningStateReady     = PrivateEndpointConnectionPropertiesStatusProvisioningState("Ready")
	PrivateEndpointConnectionPropertiesStatusProvisioningStateRejecting = PrivateEndpointConnectionPropertiesStatusProvisioningState("Rejecting")
)

type PrivateEndpointProperty_Status struct {
	//Id: Resource id of the private endpoint.
	Id *string `json:"id,omitempty"`
}

type PrivateLinkServiceConnectionStateProperty_Status struct {
	//ActionsRequired: The actions required for private link service connection.
	ActionsRequired *PrivateLinkServiceConnectionStatePropertyStatusActionsRequired `json:"actionsRequired,omitempty"`

	// +kubebuilder:validation:Required
	//Description: The private link service connection description.
	Description string `json:"description"`

	// +kubebuilder:validation:Required
	//Status: The private link service connection status.
	Status PrivateLinkServiceConnectionStatePropertyStatusStatus `json:"status"`
}

type PrivateLinkServiceConnectionStatePropertyStatusActionsRequired string

const PrivateLinkServiceConnectionStatePropertyStatusActionsRequiredNone = PrivateLinkServiceConnectionStatePropertyStatusActionsRequired("None")

type PrivateLinkServiceConnectionStatePropertyStatusStatus string

const (
	PrivateLinkServiceConnectionStatePropertyStatusStatusApproved     = PrivateLinkServiceConnectionStatePropertyStatusStatus("Approved")
	PrivateLinkServiceConnectionStatePropertyStatusStatusDisconnected = PrivateLinkServiceConnectionStatePropertyStatusStatus("Disconnected")
	PrivateLinkServiceConnectionStatePropertyStatusStatusPending      = PrivateLinkServiceConnectionStatePropertyStatusStatus("Pending")
	PrivateLinkServiceConnectionStatePropertyStatusStatusRejected     = PrivateLinkServiceConnectionStatePropertyStatusStatus("Rejected")
)

func init() {
	SchemeBuilder.Register(&Server{}, &ServerList{})
}
