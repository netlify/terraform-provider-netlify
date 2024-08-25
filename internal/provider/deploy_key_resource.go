package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &deployKeyResource{}
	_ resource.ResourceWithConfigure   = &deployKeyResource{}
	_ resource.ResourceWithImportState = &deployKeyResource{}
)

func NewDeployKeyResource() resource.Resource {
	return &deployKeyResource{}
}

type deployKeyResource struct {
	data NetlifyProviderData
}

type deployKeyResourceModel struct {
	ID          types.String `tfsdk:"id"`
	LastUpdated types.String `tfsdk:"last_updated"`
	PublicKey   types.String `tfsdk:"public_key"`
}

func (r *deployKeyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_deploy_key"
}

func (r *deployKeyResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	data, ok := req.ProviderData.(NetlifyProviderData)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected NetlifyProviderData, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.data = data
}

func (r *deployKeyResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "Deploy key for Git repositories. Avoid creating this resource directly if possible.",
		MarkdownDescription: "Deploy key for Git repositories. Avoid creating this resource directly if possible. [Read more](https://docs.netlify.com/git/repo-permissions-linking/#deploy-keys)",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"public_key": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *deployKeyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan deployKeyResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	key, _, err := r.data.client.DeployKeysAPI.CreateDeployKey(ctx).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Netlify deploy key",
			fmt.Sprintf("Could not create Netlify deploy key: %q", err.Error()),
		)
		return
	}

	plan.ID = types.StringValue(key.Id)
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC3339))
	plan.PublicKey = types.StringValue(key.PublicKey)

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *deployKeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state deployKeyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	key, _, err := r.data.client.DeployKeysAPI.GetDeployKey(ctx, state.ID.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading Netlify deploy key",
			fmt.Sprintf("Could not read Netlify deploy key %q: %q", state.ID.ValueString(), err.Error()),
		)
		return
	}

	state.PublicKey = types.StringValue(key.PublicKey)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *deployKeyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError(
		"Update not supported for Netlify deploy keys",
		"Update is not supported for Netlify deploy keys at this time.",
	)
}

func (r *deployKeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state deployKeyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.data.client.DeployKeysAPI.DeleteDeployKey(ctx, state.ID.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting Netlify deploy key",
			fmt.Sprintf("Could not delete Netlify deploy key %q: %q", state.ID.ValueString(), err.Error()),
		)
		return
	}
}

func (r *deployKeyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
