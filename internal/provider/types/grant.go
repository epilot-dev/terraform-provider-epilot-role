// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Grant struct {
	Action       types.String     `tfsdk:"action"`
	Conditions   []GrantCondition `tfsdk:"conditions"`
	Dependencies types.String     `tfsdk:"dependencies"`
	Effect       types.String     `tfsdk:"effect"`
	Resource     types.String     `tfsdk:"resource"`
}
