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
	_ validator.String = ForbiddenIfEqualsValidator{}
)

type ForbiddenIfEqualsValidator struct {
	PredicateValue  string
	PathExpressions path.Expressions
}

type ForbiddenIfEqualsValidatorRequest struct {
	Config         tfsdk.Config
	ConfigValue    attr.Value
	Path           path.Path
	PathExpression path.Expression
}

type ForbiddenIfEqualsValidatorResponse struct {
	Diagnostics diag.Diagnostics
}

func (av ForbiddenIfEqualsValidator) Description(ctx context.Context) string {
	return av.MarkdownDescription(ctx)
}

func (av ForbiddenIfEqualsValidator) MarkdownDescription(_ context.Context) string {
	return fmt.Sprintf("Ensure that if an attribute is set to %q, these are not set: %q", av.PredicateValue, av.PathExpressions)
}

func (av ForbiddenIfEqualsValidator) Validate(ctx context.Context, req ForbiddenIfEqualsValidatorRequest, res *ForbiddenIfEqualsValidatorResponse) {
	if req.ConfigValue.IsNull() || !req.ConfigValue.Equal(types.StringValue(av.PredicateValue)) {
		return
	}

	expressions := req.PathExpression.MergeExpressions(av.PathExpressions...)

	for _, expression := range expressions {
		matchedPaths, diags := req.Config.PathMatches(ctx, expression)

		res.Diagnostics.Append(diags...)

		// Collect all errors
		if diags.HasError() {
			continue
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

			if !mpVal.IsNull() {
				res.Diagnostics.Append(validatordiag.InvalidAttributeCombinationDiagnostic(
					req.Path,
					fmt.Sprintf("Attribute %q must not be specified when %q is set to %q", mp, req.Path, av.PredicateValue),
				))
			}
		}
	}
}

func (av ForbiddenIfEqualsValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	validateReq := ForbiddenIfEqualsValidatorRequest{
		Config:         req.Config,
		ConfigValue:    req.ConfigValue,
		Path:           req.Path,
		PathExpression: req.PathExpression,
	}
	validateResp := &ForbiddenIfEqualsValidatorResponse{}

	av.Validate(ctx, validateReq, validateResp)

	resp.Diagnostics.Append(validateResp.Diagnostics...)
}

func ForbiddenIfEquals(predicateValue string, pathExpressions ...path.Expression) validator.String {
	return ForbiddenIfEqualsValidator{
		PredicateValue:  predicateValue,
		PathExpressions: path.Expressions(pathExpressions),
	}
}
