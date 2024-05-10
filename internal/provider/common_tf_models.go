package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type NetlifySiteModel struct {
	ID            types.String   `tfsdk:"id"`
	AccountSlug   types.String   `tfsdk:"account_slug"`
	Name          types.String   `tfsdk:"name"`
	CustomDomain  types.String   `tfsdk:"custom_domain"`
	DomainAliases []types.String `tfsdk:"domain_aliases"`
}
