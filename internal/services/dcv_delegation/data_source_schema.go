// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dcv_delegation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*DCVDelegationDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"uuid": schema.StringAttribute{
				Description: "The DCV Delegation unique identifier.",
				Computed:    true,
			},
		},
	}
}

func (d *DCVDelegationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *DCVDelegationDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
