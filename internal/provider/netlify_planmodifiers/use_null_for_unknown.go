package netlify_planmodifiers

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func UseNullForUnknown() planmodifier.String {
	return useNullForUnknownModifier{}
}

type useNullForUnknownModifier struct{}

func (m useNullForUnknownModifier) Description(_ context.Context) string {
	return "The value of this attribute in state will always be null."
}

func (m useNullForUnknownModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m useNullForUnknownModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if !req.PlanValue.IsUnknown() {
		return
	}

	resp.PlanValue = types.StringNull()
}
