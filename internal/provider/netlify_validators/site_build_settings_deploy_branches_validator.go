package netlify_validators

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ validator.List = SiteBuildSettingsDeployBranchesValidator{}
)

type SiteBuildSettingsDeployBranchesValidator struct {
	AllBranchesPathExpression path.Expression
}

type SiteBuildSettingsDeployBranchesValidatorRequest struct {
	Config         tfsdk.Config
	ConfigValue    attr.Value
	Path           path.Path
	PathExpression path.Expression
}

type SiteBuildSettingsDeployBranchesValidatorResponse struct {
	Diagnostics diag.Diagnostics
}

func (av SiteBuildSettingsDeployBranchesValidator) Description(ctx context.Context) string {
	return av.MarkdownDescription(ctx)
}

func (av SiteBuildSettingsDeployBranchesValidator) MarkdownDescription(_ context.Context) string {
	return fmt.Sprintf("Ensure that an attribute is a non-empty list only if %q is set to false", av.AllBranchesPathExpression)
}

func (av SiteBuildSettingsDeployBranchesValidator) Validate(ctx context.Context, req SiteBuildSettingsDeployBranchesValidatorRequest, res *SiteBuildSettingsDeployBranchesValidatorResponse) {
	// Delay validation until all involved attributes have a known value
	if req.ConfigValue.IsUnknown() {
		return
	}

	emptyList, diags := types.ListValue(types.StringType, []attr.Value{})
	res.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	isNonEmpty := !req.ConfigValue.IsNull() && !req.ConfigValue.Equal(emptyList)

	matchedPaths, diags := req.Config.PathMatches(ctx, req.PathExpression.Merge(av.AllBranchesPathExpression))
	res.Diagnostics.Append(diags...)
	if diags.HasError() {
		return
	}

	for _, mp := range matchedPaths {
		var mpVal attr.Value
		diags = req.Config.GetAttribute(ctx, mp, &mpVal)
		res.Diagnostics.Append(diags...)

		// Collect all errors
		if diags.HasError() {
			continue
		}

		// Delay validation until all involved attributes have a known value
		if mpVal.IsUnknown() {
			return
		}

		isAllBranches := mpVal.Equal(types.BoolValue(true))

		if isAllBranches && isNonEmpty {
			res.Diagnostics.Append(validatordiag.InvalidAttributeCombinationDiagnostic(
				req.Path,
				fmt.Sprintf("Attribute %q must be an empty list if %q is set to true", req.Path, mp),
			))
		}
	}
}

func (av SiteBuildSettingsDeployBranchesValidator) ValidateList(ctx context.Context, req validator.ListRequest, resp *validator.ListResponse) {
	validateReq := SiteBuildSettingsDeployBranchesValidatorRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &SiteBuildSettingsDeployBranchesValidatorResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}
