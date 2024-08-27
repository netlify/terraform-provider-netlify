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
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ validator.String = EnvironmentVariableContextParameterValidator{}
)

type EnvironmentVariableContextParameterValidator struct {
	ContextPathExpression path.Expression
}

type EnvironmentVariableContextParameterValidatorRequest struct {
	Config         tfsdk.Config
	ConfigValue    attr.Value
	Path           path.Path
	PathExpression path.Expression
}

type EnvironmentVariableContextParameterValidatorResponse struct {
	Diagnostics diag.Diagnostics
}

func (av EnvironmentVariableContextParameterValidator) Description(ctx context.Context) string {
	return av.MarkdownDescription(ctx)
}

func (av EnvironmentVariableContextParameterValidator) MarkdownDescription(_ context.Context) string {
	return fmt.Sprintf("Ensure that an attribute is a non-empty string iff %q is set to \"branch\"", av.ContextPathExpression)
}

func (av EnvironmentVariableContextParameterValidator) Validate(ctx context.Context, req EnvironmentVariableContextParameterValidatorRequest, res *EnvironmentVariableContextParameterValidatorResponse) {
	// Delay validation until all involved attributes have a known value
	if req.ConfigValue.IsUnknown() {
		return
	}

	isNonEmpty := !req.ConfigValue.IsNull() && !req.ConfigValue.Equal(types.StringValue(""))

	matchedPaths, diags := req.Config.PathMatches(ctx, req.PathExpression.Merge(av.ContextPathExpression))
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

		var listValue basetypes.ListValue
		listValue, diags = types.ListValue(types.StringType, []attr.Value{types.StringValue("branch")})
		res.Diagnostics.Append(diags...)

		// Collect all errors
		if diags.HasError() {
			continue
		}

		var mpValList basetypes.ListValue
		mpValList, diags = types.ListValue(types.StringType, []attr.Value{mpVal})
		res.Diagnostics.Append(diags...)

		// Collect all errors
		if diags.HasError() {
			continue
		}
		isBranch := !mpVal.IsNull() && mpValList.Equal(listValue)

		if isNonEmpty != isBranch {
			res.Diagnostics.Append(validatordiag.InvalidAttributeCombinationDiagnostic(
				req.Path,
				fmt.Sprintf("Attribute %q must be a non-empty string iff %q is specified %q %q %q", req.Path, mp, listValue, mpVal, req.ConfigValue),
			))
		}
	}
}

func (av EnvironmentVariableContextParameterValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	validateReq := EnvironmentVariableContextParameterValidatorRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &EnvironmentVariableContextParameterValidatorResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}
