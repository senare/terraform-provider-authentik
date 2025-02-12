package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGroups() *schema.Resource {
	groupSchema := map[string]*schema.Schema{}
	for k, v := range dataSourceGroup().Schema {
		groupSchema[k] = &schema.Schema{}
		*groupSchema[k] = *v
		groupSchema[k].Computed = true
		groupSchema[k].Optional = false
		groupSchema[k].Required = false
		groupSchema[k].ExactlyOneOf = []string{}
	}
	return &schema.Resource{
		ReadContext: dataSourceGroupsRead,
		Description: "Get groups list",
		Schema: map[string]*schema.Schema{
			"attributes": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_superuser": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"members_by_pk": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"members_by_username": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ordering": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"search": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: groupSchema,
				},
			},
		},
	}
}

func dataSourceGroupsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*APIClient)

	req := c.client.CoreApi.CoreGroupsList(ctx)

	for key := range dataSourceGroups().Schema {
		if v, ok := d.GetOk(key); ok {
			switch key {
			case "attributes":
				req = req.Attributes(v.(string))
			case "is_superuser":
				req = req.IsSuperuser(v.(bool))
			case "members_by_pk":
				members := make([]int32, len(v.([]int)))
				for i, pk := range v.([]int) {
					members[i] = int32(pk)
				}
				req = req.MembersByPk(members)
			case "members_by_username":
				req = req.MembersByUsername(v.([]string))
			case "name":
				req = req.Name(v.(string))
			case "ordering":
				req = req.Ordering(v.(string))
			case "search":
				req = req.Search(v.(string))
			}
		}
	}

	groups := make([]map[string]interface{}, 0)
	for page := int32(1); true; page++ {
		req = req.Page(page)
		res, hr, err := req.Execute()
		if err != nil {
			return httpToDiag(d, hr, err)
		}

		for _, groupRes := range res.Results {
			u, err := mapFromGroup(groupRes)
			if err != nil {
				return diag.FromErr(err)
			}
			groups = append(groups, u)
		}

		if res.Pagination.Next == 0 {
			break
		}
	}

	d.SetId("0")
	setWrapper(d, "groups", groups)
	return diag.Diagnostics{}
}
