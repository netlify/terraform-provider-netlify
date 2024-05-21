package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type netlifyDomainModel struct {
	ID           types.String `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	RegisteredAt types.String `tfsdk:"registered_at"`
	ExpiresAt    types.String `tfsdk:"expires_at"`
	RenewalPrice types.String `tfsdk:"renewal_price"`
	AutoRenew    types.Bool   `tfsdk:"auto_renew"`
	AutoRenewAt  types.String `tfsdk:"auto_renew_at"`
}
