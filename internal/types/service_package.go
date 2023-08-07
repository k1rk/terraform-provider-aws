// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package types

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ListTagsFunc represents a resource's tag listing functionality.
// On success (err == nil), tags are set in Context.
type ListTagsFunc func(context.Context, any, string) error

// UpdateTagsFunc represents a resource's tag updating functionality.
type UpdateTagsFunc func(context.Context, any, string, any, any) error

// ServicePackageResourceTags represents resource-level tagging information.
type ServicePackageResourceTags struct {
	IdentifierAttribute string // The attribute for the identifier for UpdateTags etc.
	ListTags            ListTagsFunc
	UpdateTags          UpdateTagsFunc
}

// ServicePackageFrameworkDataSource represents a Terraform Plugin Framework data source
// implemented by a service package.
type ServicePackageFrameworkDataSource struct {
	Factory func(context.Context) (datasource.DataSourceWithConfigure, error)
	Name    string
	Tags    *ServicePackageResourceTags
}

// ServicePackageFrameworkResource represents a Terraform Plugin Framework resource
// implemented by a service package.
type ServicePackageFrameworkResource struct {
	Factory func(context.Context) (resource.ResourceWithConfigure, error)
	Name    string
	Tags    *ServicePackageResourceTags
}

// ServicePackageSDKDataSource represents a Terraform Plugin SDK data source
// implemented by a service package.
type ServicePackageSDKDataSource struct {
	Factory  func() *schema.Resource
	TypeName string
	Name     string
	Tags     *ServicePackageResourceTags
}

// ServicePackageSDKResource represents a Terraform Plugin SDK resource
// implemented by a service package.
type ServicePackageSDKResource struct {
	Factory  func() *schema.Resource
	TypeName string
	Name     string
	Tags     *ServicePackageResourceTags
}
