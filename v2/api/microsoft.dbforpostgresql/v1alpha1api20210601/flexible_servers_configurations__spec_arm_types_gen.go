// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type FlexibleServersConfigurations_SpecARM struct {
	//Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	//Name: The name of the server configuration.
	Name string `json:"name"`

	//Properties: The properties of a configuration.
	Properties ConfigurationPropertiesARM `json:"properties"`

	//Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &FlexibleServersConfigurations_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (flexibleServersConfigurationsSpecARM FlexibleServersConfigurations_SpecARM) GetAPIVersion() string {
	return "2021-06-01"
}

// GetName returns the Name of the resource
func (flexibleServersConfigurationsSpecARM FlexibleServersConfigurations_SpecARM) GetName() string {
	return flexibleServersConfigurationsSpecARM.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforPostgreSQL/flexibleServers/configurations"
func (flexibleServersConfigurationsSpecARM FlexibleServersConfigurations_SpecARM) GetType() string {
	return "Microsoft.DBforPostgreSQL/flexibleServers/configurations"
}

//Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/ConfigurationProperties
type ConfigurationPropertiesARM struct {
	//Source: Source of the configuration.
	Source *string `json:"source,omitempty"`

	//Value: Value of the configuration.
	Value *string `json:"value,omitempty"`
}
