// Copyright (c) Illumio, Inc.
// SPDX-License-Identifier: MPL-2.0
// Code generated by providergen. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	resource_schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	configv1 "github.com/illumio/terraform-provider-illumio-cloudsecure/api/illumio/cloud/config/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// ProviderData is the state of the provider, which is passed to resources and data sources at runtime as their ProviderData.
type ProviderData interface {
	// Client returns the CloudSecure Config API client.
	Client() configv1.ConfigServiceClient

	// RequestTimeout returns the maximum duration of each API request.
	RequestTimeout() time.Duration
}

func (p *Provider) Resources(ctx context.Context) []func() resource.Resource {
	resources := p.schema.Resources()
	resp := make([]func() resource.Resource, 0, len(resources))
	for _, r := range resources {
		switch r.TypeName {
		case "aws_account":
			resp = append(resp, func() resource.Resource { return NewAwsAccountResource(r.Schema) })
		case "aws_organization":
			resp = append(resp, func() resource.Resource { return NewAwsOrganizationResource(r.Schema) })
		case "k8s_cluster_onboarding_credential":
			resp = append(resp, func() resource.Resource { return NewK8SClusterOnboardingCredentialResource(r.Schema) })
		}
	}
	return resp
}

func (p *Provider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		// TODO: Add support for data sources.
	}
}

// AwsAccountResource implements the aws_account resource.
type AwsAccountResource struct {
	// schema is the schema of the aws_account resource.
	schema resource_schema.Schema

	// providerData is the provider configuration.
	config ProviderData
}

var _ resource.ResourceWithConfigure = &AwsAccountResource{}
var _ resource.ResourceWithImportState = &AwsAccountResource{}

// NewAwsAccountResource returns a new aws_account resource.
func NewAwsAccountResource(schema resource_schema.Schema) resource.Resource {
	return &AwsAccountResource{
		schema: schema,
	}
}

func (r *AwsAccountResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aws_account"
}

func (r *AwsAccountResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = r.schema
}

func (r *AwsAccountResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	providerData, ok := req.ProviderData.(ProviderData)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected ProviderData, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.config = providerData
}

func (r *AwsAccountResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AwsAccountResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	protoReq := NewCreateAwsAccountRequest(&data)

	tflog.Trace(ctx, "creating a resource", map[string]any{"type": "aws_account"})

	rpcCtx, rpcCancel := context.WithTimeout(ctx, r.config.RequestTimeout())
	protoResp, err := r.config.Client().CreateAwsAccount(rpcCtx, protoReq)
	rpcCancel()
	if err != nil {
		resp.Diagnostics.AddError("Config API Error", fmt.Sprintf("Unable to create aws_account, got error: %s", err))
		return
	}

	CopyCreateAwsAccountResponse(&data, protoResp)

	tflog.Trace(ctx, "created a resource", map[string]any{"type": "aws_account", "id": protoResp.Id})

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AwsAccountResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AwsAccountResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	protoReq := NewReadAwsAccountRequest(&data)

	tflog.Trace(ctx, "reading a resource", map[string]any{"type": "aws_account", "id": protoReq.Id})

	rpcCtx, rpcCancel := context.WithTimeout(ctx, r.config.RequestTimeout())
	protoResp, err := r.config.Client().ReadAwsAccount(rpcCtx, protoReq)
	rpcCancel()
	if err != nil {
		switch status.Code(err) {
		case codes.NotFound:
			resp.Diagnostics.AddWarning("Resource Not Found", fmt.Sprintf("No aws_account found with id %s", protoReq.Id))
			resp.State.RemoveResource(ctx)
			return
		default:
			resp.Diagnostics.AddError("Config API Error", fmt.Sprintf("Unable to read aws_account, got error: %s", err))
			return
		}
	}

	CopyReadAwsAccountResponse(&data, protoResp)

	tflog.Trace(ctx, "read a resource", map[string]any{"type": "aws_account", "id": protoResp.Id})

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AwsAccountResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var beforeData AwsAccountResourceModel
	var afterData AwsAccountResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &beforeData)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(req.Plan.Get(ctx, &afterData)...)
	if resp.Diagnostics.HasError() {
		return
	}

	protoReq := NewUpdateAwsAccountRequest(&beforeData, &afterData)

	tflog.Trace(ctx, "updating a resource", map[string]any{"type": "aws_account", "id": protoReq.Id, "update_mask": protoReq.UpdateMask.Paths})

	rpcCtx, rpcCancel := context.WithTimeout(ctx, r.config.RequestTimeout())
	protoResp, err := r.config.Client().UpdateAwsAccount(rpcCtx, protoReq)
	rpcCancel()
	if err != nil {
		switch status.Code(err) {
		case codes.NotFound:
			resp.Diagnostics.AddError("Resource Not Found", fmt.Sprintf("No aws_account found with id %s", protoReq.Id))
			resp.State.RemoveResource(ctx)
			return
		default:
			resp.Diagnostics.AddError("Config API Error", fmt.Sprintf("Unable to update aws_account, got error: %s", err))
			return
		}
	}

	CopyUpdateAwsAccountResponse(&afterData, protoResp)

	tflog.Trace(ctx, "updated a resource", map[string]any{"type": "aws_account", "id": protoResp.Id})

	resp.Diagnostics.Append(resp.State.Set(ctx, &afterData)...)
}

func (r *AwsAccountResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AwsAccountResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	protoReq := NewDeleteAwsAccountRequest(&data)

	tflog.Trace(ctx, "deleting a resource", map[string]any{"type": "aws_account", "id": protoReq.Id})

	rpcCtx, rpcCancel := context.WithTimeout(ctx, r.config.RequestTimeout())
	_, err := r.config.Client().DeleteAwsAccount(rpcCtx, protoReq)
	rpcCancel()
	if err != nil {
		switch status.Code(err) {
		case codes.NotFound:
			tflog.Trace(ctx, "resource was already deleted", map[string]any{"type": "aws_account", "id": protoReq.Id})
		default:
			resp.Diagnostics.AddError("Config API Error", fmt.Sprintf("Unable to delete aws_account, got error: %s", err))
			return
		}
	}

	tflog.Trace(ctx, "deleted a resource", map[string]any{"type": "aws_account", "id": protoReq.Id})
}

func (r *AwsAccountResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// TODO
}

// AwsOrganizationResource implements the aws_organization resource.
type AwsOrganizationResource struct {
	// schema is the schema of the aws_organization resource.
	schema resource_schema.Schema

	// providerData is the provider configuration.
	config ProviderData
}

var _ resource.ResourceWithConfigure = &AwsOrganizationResource{}
var _ resource.ResourceWithImportState = &AwsOrganizationResource{}

// NewAwsOrganizationResource returns a new aws_organization resource.
func NewAwsOrganizationResource(schema resource_schema.Schema) resource.Resource {
	return &AwsOrganizationResource{
		schema: schema,
	}
}

func (r *AwsOrganizationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aws_organization"
}

func (r *AwsOrganizationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = r.schema
}

func (r *AwsOrganizationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	providerData, ok := req.ProviderData.(ProviderData)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected ProviderData, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.config = providerData
}

func (r *AwsOrganizationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AwsOrganizationResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	protoReq := NewCreateAwsOrganizationRequest(&data)

	tflog.Trace(ctx, "creating a resource", map[string]any{"type": "aws_organization"})

	rpcCtx, rpcCancel := context.WithTimeout(ctx, r.config.RequestTimeout())
	protoResp, err := r.config.Client().CreateAwsOrganization(rpcCtx, protoReq)
	rpcCancel()
	if err != nil {
		resp.Diagnostics.AddError("Config API Error", fmt.Sprintf("Unable to create aws_organization, got error: %s", err))
		return
	}

	CopyCreateAwsOrganizationResponse(&data, protoResp)

	tflog.Trace(ctx, "created a resource", map[string]any{"type": "aws_organization", "id": protoResp.Id})

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AwsOrganizationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AwsOrganizationResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	protoReq := NewReadAwsOrganizationRequest(&data)

	tflog.Trace(ctx, "reading a resource", map[string]any{"type": "aws_organization", "id": protoReq.Id})

	rpcCtx, rpcCancel := context.WithTimeout(ctx, r.config.RequestTimeout())
	protoResp, err := r.config.Client().ReadAwsOrganization(rpcCtx, protoReq)
	rpcCancel()
	if err != nil {
		switch status.Code(err) {
		case codes.NotFound:
			resp.Diagnostics.AddWarning("Resource Not Found", fmt.Sprintf("No aws_organization found with id %s", protoReq.Id))
			resp.State.RemoveResource(ctx)
			return
		default:
			resp.Diagnostics.AddError("Config API Error", fmt.Sprintf("Unable to read aws_organization, got error: %s", err))
			return
		}
	}

	CopyReadAwsOrganizationResponse(&data, protoResp)

	tflog.Trace(ctx, "read a resource", map[string]any{"type": "aws_organization", "id": protoResp.Id})

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AwsOrganizationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var beforeData AwsOrganizationResourceModel
	var afterData AwsOrganizationResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &beforeData)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(req.Plan.Get(ctx, &afterData)...)
	if resp.Diagnostics.HasError() {
		return
	}

	protoReq := NewUpdateAwsOrganizationRequest(&beforeData, &afterData)

	tflog.Trace(ctx, "updating a resource", map[string]any{"type": "aws_organization", "id": protoReq.Id, "update_mask": protoReq.UpdateMask.Paths})

	rpcCtx, rpcCancel := context.WithTimeout(ctx, r.config.RequestTimeout())
	protoResp, err := r.config.Client().UpdateAwsOrganization(rpcCtx, protoReq)
	rpcCancel()
	if err != nil {
		switch status.Code(err) {
		case codes.NotFound:
			resp.Diagnostics.AddError("Resource Not Found", fmt.Sprintf("No aws_organization found with id %s", protoReq.Id))
			resp.State.RemoveResource(ctx)
			return
		default:
			resp.Diagnostics.AddError("Config API Error", fmt.Sprintf("Unable to update aws_organization, got error: %s", err))
			return
		}
	}

	CopyUpdateAwsOrganizationResponse(&afterData, protoResp)

	tflog.Trace(ctx, "updated a resource", map[string]any{"type": "aws_organization", "id": protoResp.Id})

	resp.Diagnostics.Append(resp.State.Set(ctx, &afterData)...)
}

func (r *AwsOrganizationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AwsOrganizationResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	protoReq := NewDeleteAwsOrganizationRequest(&data)

	tflog.Trace(ctx, "deleting a resource", map[string]any{"type": "aws_organization", "id": protoReq.Id})

	rpcCtx, rpcCancel := context.WithTimeout(ctx, r.config.RequestTimeout())
	_, err := r.config.Client().DeleteAwsOrganization(rpcCtx, protoReq)
	rpcCancel()
	if err != nil {
		switch status.Code(err) {
		case codes.NotFound:
			tflog.Trace(ctx, "resource was already deleted", map[string]any{"type": "aws_organization", "id": protoReq.Id})
		default:
			resp.Diagnostics.AddError("Config API Error", fmt.Sprintf("Unable to delete aws_organization, got error: %s", err))
			return
		}
	}

	tflog.Trace(ctx, "deleted a resource", map[string]any{"type": "aws_organization", "id": protoReq.Id})
}

func (r *AwsOrganizationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// TODO
}

// K8SClusterOnboardingCredentialResource implements the k8s_cluster_onboarding_credential resource.
type K8SClusterOnboardingCredentialResource struct {
	// schema is the schema of the k8s_cluster_onboarding_credential resource.
	schema resource_schema.Schema

	// providerData is the provider configuration.
	config ProviderData
}

var _ resource.ResourceWithConfigure = &K8SClusterOnboardingCredentialResource{}
var _ resource.ResourceWithImportState = &K8SClusterOnboardingCredentialResource{}

// NewK8SClusterOnboardingCredentialResource returns a new k8s_cluster_onboarding_credential resource.
func NewK8SClusterOnboardingCredentialResource(schema resource_schema.Schema) resource.Resource {
	return &K8SClusterOnboardingCredentialResource{
		schema: schema,
	}
}

func (r *K8SClusterOnboardingCredentialResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_k8s_cluster_onboarding_credential"
}

func (r *K8SClusterOnboardingCredentialResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = r.schema
}

func (r *K8SClusterOnboardingCredentialResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	providerData, ok := req.ProviderData.(ProviderData)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected ProviderData, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.config = providerData
}

func (r *K8SClusterOnboardingCredentialResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data K8SClusterOnboardingCredentialResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	protoReq := NewCreateK8SClusterOnboardingCredentialRequest(&data)

	tflog.Trace(ctx, "creating a resource", map[string]any{"type": "k8s_cluster_onboarding_credential"})

	rpcCtx, rpcCancel := context.WithTimeout(ctx, r.config.RequestTimeout())
	protoResp, err := r.config.Client().CreateK8SClusterOnboardingCredential(rpcCtx, protoReq)
	rpcCancel()
	if err != nil {
		resp.Diagnostics.AddError("Config API Error", fmt.Sprintf("Unable to create k8s_cluster_onboarding_credential, got error: %s", err))
		return
	}

	CopyCreateK8SClusterOnboardingCredentialResponse(&data, protoResp)

	tflog.Trace(ctx, "created a resource", map[string]any{"type": "k8s_cluster_onboarding_credential", "id": protoResp.Id})

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *K8SClusterOnboardingCredentialResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data K8SClusterOnboardingCredentialResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	protoReq := NewReadK8SClusterOnboardingCredentialRequest(&data)

	tflog.Trace(ctx, "reading a resource", map[string]any{"type": "k8s_cluster_onboarding_credential", "id": protoReq.Id})

	rpcCtx, rpcCancel := context.WithTimeout(ctx, r.config.RequestTimeout())
	protoResp, err := r.config.Client().ReadK8SClusterOnboardingCredential(rpcCtx, protoReq)
	rpcCancel()
	if err != nil {
		switch status.Code(err) {
		case codes.NotFound:
			resp.Diagnostics.AddWarning("Resource Not Found", fmt.Sprintf("No k8s_cluster_onboarding_credential found with id %s", protoReq.Id))
			resp.State.RemoveResource(ctx)
			return
		default:
			resp.Diagnostics.AddError("Config API Error", fmt.Sprintf("Unable to read k8s_cluster_onboarding_credential, got error: %s", err))
			return
		}
	}

	CopyReadK8SClusterOnboardingCredentialResponse(&data, protoResp)

	tflog.Trace(ctx, "read a resource", map[string]any{"type": "k8s_cluster_onboarding_credential", "id": protoResp.Id})

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *K8SClusterOnboardingCredentialResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var beforeData K8SClusterOnboardingCredentialResourceModel
	var afterData K8SClusterOnboardingCredentialResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &beforeData)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(req.Plan.Get(ctx, &afterData)...)
	if resp.Diagnostics.HasError() {
		return
	}

	protoReq := NewUpdateK8SClusterOnboardingCredentialRequest(&beforeData, &afterData)

	tflog.Trace(ctx, "updating a resource", map[string]any{"type": "k8s_cluster_onboarding_credential", "id": protoReq.Id, "update_mask": protoReq.UpdateMask.Paths})

	rpcCtx, rpcCancel := context.WithTimeout(ctx, r.config.RequestTimeout())
	protoResp, err := r.config.Client().UpdateK8SClusterOnboardingCredential(rpcCtx, protoReq)
	rpcCancel()
	if err != nil {
		switch status.Code(err) {
		case codes.NotFound:
			resp.Diagnostics.AddError("Resource Not Found", fmt.Sprintf("No k8s_cluster_onboarding_credential found with id %s", protoReq.Id))
			resp.State.RemoveResource(ctx)
			return
		default:
			resp.Diagnostics.AddError("Config API Error", fmt.Sprintf("Unable to update k8s_cluster_onboarding_credential, got error: %s", err))
			return
		}
	}

	CopyUpdateK8SClusterOnboardingCredentialResponse(&afterData, protoResp)

	tflog.Trace(ctx, "updated a resource", map[string]any{"type": "k8s_cluster_onboarding_credential", "id": protoResp.Id})

	resp.Diagnostics.Append(resp.State.Set(ctx, &afterData)...)
}

func (r *K8SClusterOnboardingCredentialResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data K8SClusterOnboardingCredentialResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	protoReq := NewDeleteK8SClusterOnboardingCredentialRequest(&data)

	tflog.Trace(ctx, "deleting a resource", map[string]any{"type": "k8s_cluster_onboarding_credential", "id": protoReq.Id})

	rpcCtx, rpcCancel := context.WithTimeout(ctx, r.config.RequestTimeout())
	_, err := r.config.Client().DeleteK8SClusterOnboardingCredential(rpcCtx, protoReq)
	rpcCancel()
	if err != nil {
		switch status.Code(err) {
		case codes.NotFound:
			tflog.Trace(ctx, "resource was already deleted", map[string]any{"type": "k8s_cluster_onboarding_credential", "id": protoReq.Id})
		default:
			resp.Diagnostics.AddError("Config API Error", fmt.Sprintf("Unable to delete k8s_cluster_onboarding_credential, got error: %s", err))
			return
		}
	}

	tflog.Trace(ctx, "deleted a resource", map[string]any{"type": "k8s_cluster_onboarding_credential", "id": protoReq.Id})
}

func (r *K8SClusterOnboardingCredentialResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// TODO
}

type AwsAccountResourceModel struct {
	Id                          types.String `tfsdk:"id"`
	AccountId                   types.String `tfsdk:"account_id"`
	Mode                        types.String `tfsdk:"mode"`
	Name                        types.String `tfsdk:"name"`
	OrganizationMasterAccountId types.String `tfsdk:"organization_master_account_id"`
	RoleArn                     types.String `tfsdk:"role_arn"`
	RoleExternalId              types.String `tfsdk:"role_external_id"`
}

type AwsOrganizationResourceModel struct {
	Id              types.String `tfsdk:"id"`
	MasterAccountId types.String `tfsdk:"master_account_id"`
	Mode            types.String `tfsdk:"mode"`
	Name            types.String `tfsdk:"name"`
	OrganizationId  types.String `tfsdk:"organization_id"`
	RoleArn         types.String `tfsdk:"role_arn"`
	RoleExternalId  types.String `tfsdk:"role_external_id"`
}

type K8SClusterOnboardingCredentialResourceModel struct {
	Id            types.String `tfsdk:"id"`
	ClientId      types.String `tfsdk:"client_id"`
	ClientSecret  types.String `tfsdk:"client_secret"`
	CreatedAt     types.String `tfsdk:"created_at"`
	Description   types.String `tfsdk:"description"`
	IllumioRegion types.String `tfsdk:"illumio_region"`
	Name          types.String `tfsdk:"name"`
}

func NewCreateAwsAccountRequest(data *AwsAccountResourceModel) *configv1.CreateAwsAccountRequest {
	proto := &configv1.CreateAwsAccountRequest{}
	if !data.AccountId.IsUnknown() && !data.AccountId.IsNull() {
		var dataValue attr.Value = data.AccountId
		var protoValue string
		protoValue = dataValue.(types.String).ValueString()
		proto.AccountId = protoValue
	}
	if !data.Mode.IsUnknown() && !data.Mode.IsNull() {
		var dataValue attr.Value = data.Mode
		var protoValue string
		protoValue = dataValue.(types.String).ValueString()
		proto.Mode = protoValue
	}
	if !data.Name.IsUnknown() && !data.Name.IsNull() {
		var dataValue attr.Value = data.Name
		var protoValue string
		protoValue = dataValue.(types.String).ValueString()
		proto.Name = protoValue
	}
	if !data.OrganizationMasterAccountId.IsUnknown() && !data.OrganizationMasterAccountId.IsNull() {
		var dataValue attr.Value = data.OrganizationMasterAccountId
		var protoValue string
		protoValue = dataValue.(types.String).ValueString()
		proto.OrganizationMasterAccountId = &protoValue
	}
	if !data.RoleArn.IsUnknown() && !data.RoleArn.IsNull() {
		var dataValue attr.Value = data.RoleArn
		var protoValue string
		protoValue = dataValue.(types.String).ValueString()
		proto.RoleArn = protoValue
	}
	if !data.RoleExternalId.IsUnknown() && !data.RoleExternalId.IsNull() {
		var dataValue attr.Value = data.RoleExternalId
		var protoValue string
		protoValue = dataValue.(types.String).ValueString()
		proto.RoleExternalId = protoValue
	}
	return proto
}

func NewReadAwsAccountRequest(data *AwsAccountResourceModel) *configv1.ReadAwsAccountRequest {
	proto := &configv1.ReadAwsAccountRequest{}
	if !data.Id.IsUnknown() && !data.Id.IsNull() {
		var dataValue attr.Value = data.Id
		var protoValue string
		protoValue = dataValue.(types.String).ValueString()
		proto.Id = protoValue
	}
	return proto
}

func NewDeleteAwsAccountRequest(data *AwsAccountResourceModel) *configv1.DeleteAwsAccountRequest {
	proto := &configv1.DeleteAwsAccountRequest{}
	if !data.Id.IsUnknown() && !data.Id.IsNull() {
		var dataValue attr.Value = data.Id
		var protoValue string
		protoValue = dataValue.(types.String).ValueString()
		proto.Id = protoValue
	}
	return proto
}

func NewCreateAwsOrganizationRequest(data *AwsOrganizationResourceModel) *configv1.CreateAwsOrganizationRequest {
	proto := &configv1.CreateAwsOrganizationRequest{}
	if !data.MasterAccountId.IsUnknown() && !data.MasterAccountId.IsNull() {
		var dataValue attr.Value = data.MasterAccountId
		var protoValue string
		protoValue = dataValue.(types.String).ValueString()
		proto.MasterAccountId = protoValue
	}
	if !data.Mode.IsUnknown() && !data.Mode.IsNull() {
		var dataValue attr.Value = data.Mode
		var protoValue string
		protoValue = dataValue.(types.String).ValueString()
		proto.Mode = protoValue
	}
	if !data.Name.IsUnknown() && !data.Name.IsNull() {
		var dataValue attr.Value = data.Name
		var protoValue string
		protoValue = dataValue.(types.String).ValueString()
		proto.Name = protoValue
	}
	if !data.OrganizationId.IsUnknown() && !data.OrganizationId.IsNull() {
		var dataValue attr.Value = data.OrganizationId
		var protoValue string
		protoValue = dataValue.(types.String).ValueString()
		proto.OrganizationId = protoValue
	}
	if !data.RoleArn.IsUnknown() && !data.RoleArn.IsNull() {
		var dataValue attr.Value = data.RoleArn
		var protoValue string
		protoValue = dataValue.(types.String).ValueString()
		proto.RoleArn = protoValue
	}
	if !data.RoleExternalId.IsUnknown() && !data.RoleExternalId.IsNull() {
		var dataValue attr.Value = data.RoleExternalId
		var protoValue string
		protoValue = dataValue.(types.String).ValueString()
		proto.RoleExternalId = protoValue
	}
	return proto
}

func NewReadAwsOrganizationRequest(data *AwsOrganizationResourceModel) *configv1.ReadAwsOrganizationRequest {
	proto := &configv1.ReadAwsOrganizationRequest{}
	if !data.Id.IsUnknown() && !data.Id.IsNull() {
		var dataValue attr.Value = data.Id
		var protoValue string
		protoValue = dataValue.(types.String).ValueString()
		proto.Id = protoValue
	}
	return proto
}

func NewDeleteAwsOrganizationRequest(data *AwsOrganizationResourceModel) *configv1.DeleteAwsOrganizationRequest {
	proto := &configv1.DeleteAwsOrganizationRequest{}
	if !data.Id.IsUnknown() && !data.Id.IsNull() {
		var dataValue attr.Value = data.Id
		var protoValue string
		protoValue = dataValue.(types.String).ValueString()
		proto.Id = protoValue
	}
	return proto
}

func NewCreateK8SClusterOnboardingCredentialRequest(data *K8SClusterOnboardingCredentialResourceModel) *configv1.CreateK8SClusterOnboardingCredentialRequest {
	proto := &configv1.CreateK8SClusterOnboardingCredentialRequest{}
	if !data.Description.IsUnknown() && !data.Description.IsNull() {
		var dataValue attr.Value = data.Description
		var protoValue string
		protoValue = dataValue.(types.String).ValueString()
		proto.Description = &protoValue
	}
	if !data.IllumioRegion.IsUnknown() && !data.IllumioRegion.IsNull() {
		var dataValue attr.Value = data.IllumioRegion
		var protoValue string
		protoValue = dataValue.(types.String).ValueString()
		proto.IllumioRegion = protoValue
	}
	if !data.Name.IsUnknown() && !data.Name.IsNull() {
		var dataValue attr.Value = data.Name
		var protoValue string
		protoValue = dataValue.(types.String).ValueString()
		proto.Name = protoValue
	}
	return proto
}

func NewReadK8SClusterOnboardingCredentialRequest(data *K8SClusterOnboardingCredentialResourceModel) *configv1.ReadK8SClusterOnboardingCredentialRequest {
	proto := &configv1.ReadK8SClusterOnboardingCredentialRequest{}
	if !data.Id.IsUnknown() && !data.Id.IsNull() {
		var dataValue attr.Value = data.Id
		var protoValue string
		protoValue = dataValue.(types.String).ValueString()
		proto.Id = protoValue
	}
	return proto
}

func NewDeleteK8SClusterOnboardingCredentialRequest(data *K8SClusterOnboardingCredentialResourceModel) *configv1.DeleteK8SClusterOnboardingCredentialRequest {
	proto := &configv1.DeleteK8SClusterOnboardingCredentialRequest{}
	if !data.Id.IsUnknown() && !data.Id.IsNull() {
		var dataValue attr.Value = data.Id
		var protoValue string
		protoValue = dataValue.(types.String).ValueString()
		proto.Id = protoValue
	}
	return proto
}

func NewUpdateAwsAccountRequest(beforeData, afterData *AwsAccountResourceModel) *configv1.UpdateAwsAccountRequest {
	proto := &configv1.UpdateAwsAccountRequest{}
	proto.UpdateMask, _ = fieldmaskpb.New(proto)
	proto.Id = beforeData.Id.ValueString()
	if !afterData.Name.Equal(beforeData.Name) {
		proto.UpdateMask.Append(proto, "name")
		if !afterData.Name.IsUnknown() && !afterData.Name.IsNull() {
			var dataValue attr.Value = afterData.Name
			var protoValue string
			protoValue = dataValue.(types.String).ValueString()
			proto.Name = protoValue
		}
	}
	return proto
}

func NewUpdateAwsOrganizationRequest(beforeData, afterData *AwsOrganizationResourceModel) *configv1.UpdateAwsOrganizationRequest {
	proto := &configv1.UpdateAwsOrganizationRequest{}
	proto.UpdateMask, _ = fieldmaskpb.New(proto)
	proto.Id = beforeData.Id.ValueString()
	if !afterData.Name.Equal(beforeData.Name) {
		proto.UpdateMask.Append(proto, "name")
		if !afterData.Name.IsUnknown() && !afterData.Name.IsNull() {
			var dataValue attr.Value = afterData.Name
			var protoValue string
			protoValue = dataValue.(types.String).ValueString()
			proto.Name = protoValue
		}
	}
	return proto
}

func NewUpdateK8SClusterOnboardingCredentialRequest(beforeData, afterData *K8SClusterOnboardingCredentialResourceModel) *configv1.UpdateK8SClusterOnboardingCredentialRequest {
	proto := &configv1.UpdateK8SClusterOnboardingCredentialRequest{}
	proto.UpdateMask, _ = fieldmaskpb.New(proto)
	proto.Id = beforeData.Id.ValueString()
	if !afterData.Description.Equal(beforeData.Description) {
		proto.UpdateMask.Append(proto, "description")
		if !afterData.Description.IsUnknown() && !afterData.Description.IsNull() {
			var dataValue attr.Value = afterData.Description
			var protoValue string
			protoValue = dataValue.(types.String).ValueString()
			proto.Description = &protoValue
		}
	}
	if !afterData.Name.Equal(beforeData.Name) {
		proto.UpdateMask.Append(proto, "name")
		if !afterData.Name.IsUnknown() && !afterData.Name.IsNull() {
			var dataValue attr.Value = afterData.Name
			var protoValue string
			protoValue = dataValue.(types.String).ValueString()
			proto.Name = protoValue
		}
	}
	return proto
}
func CopyCreateAwsAccountResponse(dst *AwsAccountResourceModel, src *configv1.CreateAwsAccountResponse) {
	dst.Id = types.StringValue(src.Id)
	dst.AccountId = types.StringValue(src.AccountId)
	dst.Mode = types.StringValue(src.Mode)
	dst.Name = types.StringValue(src.Name)
	dst.OrganizationMasterAccountId = types.StringPointerValue(src.OrganizationMasterAccountId)
	dst.RoleArn = types.StringValue(src.RoleArn)
	dst.RoleExternalId = types.StringValue(src.RoleExternalId)
}
func CopyReadAwsAccountResponse(dst *AwsAccountResourceModel, src *configv1.ReadAwsAccountResponse) {
	dst.Id = types.StringValue(src.Id)
	dst.AccountId = types.StringValue(src.AccountId)
	dst.Mode = types.StringValue(src.Mode)
	dst.Name = types.StringValue(src.Name)
	dst.OrganizationMasterAccountId = types.StringPointerValue(src.OrganizationMasterAccountId)
	dst.RoleArn = types.StringValue(src.RoleArn)
	dst.RoleExternalId = types.StringValue(src.RoleExternalId)
}
func CopyUpdateAwsAccountResponse(dst *AwsAccountResourceModel, src *configv1.UpdateAwsAccountResponse) {
	dst.Id = types.StringValue(src.Id)
	dst.AccountId = types.StringValue(src.AccountId)
	dst.Mode = types.StringValue(src.Mode)
	dst.Name = types.StringValue(src.Name)
	dst.OrganizationMasterAccountId = types.StringPointerValue(src.OrganizationMasterAccountId)
	dst.RoleArn = types.StringValue(src.RoleArn)
	dst.RoleExternalId = types.StringValue(src.RoleExternalId)
}
func CopyCreateAwsOrganizationResponse(dst *AwsOrganizationResourceModel, src *configv1.CreateAwsOrganizationResponse) {
	dst.Id = types.StringValue(src.Id)
	dst.MasterAccountId = types.StringValue(src.MasterAccountId)
	dst.Mode = types.StringValue(src.Mode)
	dst.Name = types.StringValue(src.Name)
	dst.OrganizationId = types.StringValue(src.OrganizationId)
	dst.RoleArn = types.StringValue(src.RoleArn)
	dst.RoleExternalId = types.StringValue(src.RoleExternalId)
}
func CopyReadAwsOrganizationResponse(dst *AwsOrganizationResourceModel, src *configv1.ReadAwsOrganizationResponse) {
	dst.Id = types.StringValue(src.Id)
	dst.MasterAccountId = types.StringValue(src.MasterAccountId)
	dst.Mode = types.StringValue(src.Mode)
	dst.Name = types.StringValue(src.Name)
	dst.OrganizationId = types.StringValue(src.OrganizationId)
	dst.RoleArn = types.StringValue(src.RoleArn)
	dst.RoleExternalId = types.StringValue(src.RoleExternalId)
}
func CopyUpdateAwsOrganizationResponse(dst *AwsOrganizationResourceModel, src *configv1.UpdateAwsOrganizationResponse) {
	dst.Id = types.StringValue(src.Id)
	dst.MasterAccountId = types.StringValue(src.MasterAccountId)
	dst.Mode = types.StringValue(src.Mode)
	dst.Name = types.StringValue(src.Name)
	dst.OrganizationId = types.StringValue(src.OrganizationId)
	dst.RoleArn = types.StringValue(src.RoleArn)
	dst.RoleExternalId = types.StringValue(src.RoleExternalId)
}
func CopyCreateK8SClusterOnboardingCredentialResponse(dst *K8SClusterOnboardingCredentialResourceModel, src *configv1.CreateK8SClusterOnboardingCredentialResponse) {
	dst.Id = types.StringValue(src.Id)
	dst.ClientId = types.StringValue(src.ClientId)
	dst.ClientSecret = types.StringValue(src.ClientSecret)
	dst.CreatedAt = types.StringValue(src.CreatedAt)
	dst.Description = types.StringPointerValue(src.Description)
	dst.IllumioRegion = types.StringValue(src.IllumioRegion)
	dst.Name = types.StringValue(src.Name)
}
func CopyReadK8SClusterOnboardingCredentialResponse(dst *K8SClusterOnboardingCredentialResourceModel, src *configv1.ReadK8SClusterOnboardingCredentialResponse) {
	dst.Id = types.StringValue(src.Id)
	dst.ClientId = types.StringValue(src.ClientId)
	dst.CreatedAt = types.StringValue(src.CreatedAt)
	dst.Description = types.StringPointerValue(src.Description)
	dst.IllumioRegion = types.StringValue(src.IllumioRegion)
	dst.Name = types.StringValue(src.Name)
}
func CopyUpdateK8SClusterOnboardingCredentialResponse(dst *K8SClusterOnboardingCredentialResourceModel, src *configv1.UpdateK8SClusterOnboardingCredentialResponse) {
	dst.Id = types.StringValue(src.Id)
	dst.ClientId = types.StringValue(src.ClientId)
	dst.CreatedAt = types.StringValue(src.CreatedAt)
	dst.Description = types.StringPointerValue(src.Description)
	dst.IllumioRegion = types.StringValue(src.IllumioRegion)
	dst.Name = types.StringValue(src.Name)
}
