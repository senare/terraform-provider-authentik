package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	api "goauthentik.io/api/v3"
)

func resourceStageAuthenticatorStatic() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceStageAuthenticatorStaticCreate,
		ReadContext:   resourceStageAuthenticatorStaticRead,
		UpdateContext: resourceStageAuthenticatorStaticUpdate,
		DeleteContext: resourceStageAuthenticatorStaticDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"configure_flow": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"token_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  6,
			},
		},
	}
}

func resourceStageAuthenticatorStaticSchemaToProvider(d *schema.ResourceData) *api.AuthenticatorStaticStageRequest {
	r := api.AuthenticatorStaticStageRequest{
		Name:       d.Get("name").(string),
		TokenCount: intToPointer(d.Get("token_count").(int)),
	}

	if h, hSet := d.GetOk("configure_flow"); hSet {
		r.ConfigureFlow.Set(stringToPointer(h.(string)))
	}
	return &r
}

func resourceStageAuthenticatorStaticCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*APIClient)

	r := resourceStageAuthenticatorStaticSchemaToProvider(d)

	res, hr, err := c.client.StagesApi.StagesAuthenticatorStaticCreate(ctx).AuthenticatorStaticStageRequest(*r).Execute()
	if err != nil {
		return httpToDiag(d, hr, err)
	}

	d.SetId(res.Pk)
	return resourceStageAuthenticatorStaticRead(ctx, d, m)
}

func resourceStageAuthenticatorStaticRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*APIClient)

	res, hr, err := c.client.StagesApi.StagesAuthenticatorStaticRetrieve(ctx, d.Id()).Execute()
	if err != nil {
		return httpToDiag(d, hr, err)
	}

	setWrapper(d, "name", res.Name)
	setWrapper(d, "token_count", res.TokenCount)
	if res.ConfigureFlow.IsSet() {
		setWrapper(d, "configure_flow", res.ConfigureFlow.Get())
	}
	return diags
}

func resourceStageAuthenticatorStaticUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*APIClient)

	app := resourceStageAuthenticatorStaticSchemaToProvider(d)

	res, hr, err := c.client.StagesApi.StagesAuthenticatorStaticUpdate(ctx, d.Id()).AuthenticatorStaticStageRequest(*app).Execute()
	if err != nil {
		return httpToDiag(d, hr, err)
	}

	d.SetId(res.Pk)
	return resourceStageAuthenticatorStaticRead(ctx, d, m)
}

func resourceStageAuthenticatorStaticDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*APIClient)
	hr, err := c.client.StagesApi.StagesAuthenticatorStaticDestroy(ctx, d.Id()).Execute()
	if err != nil {
		return httpToDiag(d, hr, err)
	}
	return diag.Diagnostics{}
}
