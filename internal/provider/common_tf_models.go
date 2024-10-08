package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type netlifyDomainModel struct {
	ID           types.String `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	RegisteredAt types.String `tfsdk:"registered_at"`
	ExpiresAt    types.String `tfsdk:"expires_at"`
	RenewalPrice types.String `tfsdk:"renewal_price"`
	AutoRenew    types.Bool   `tfsdk:"auto_renew"`
	AutoRenewAt  types.String `tfsdk:"auto_renew_at"`
}

func (m netlifyDomainModel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":            types.StringType,
		"name":          types.StringType,
		"registered_at": types.StringType,
		"expires_at":    types.StringType,
		"renewal_price": types.StringType,
		"auto_renew":    types.BoolType,
		"auto_renew_at": types.StringType,
	}
}
