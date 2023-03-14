package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceProviderSCIM(t *testing.T) {
	rName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceProviderSCIM(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("authentik_provider_scim.name", "name", rName),
				),
			},
			{
				Config: testAccResourceProviderSCIM(rName + "test"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("authentik_provider_scim.name", "name", rName+"test"),
				),
			},
		},
	})
}

func testAccResourceProviderSCIM(name string) string {
	return fmt.Sprintf(`
resource "authentik_provider_scim" "name" {
  name  = "%[1]s"
  url   = "http://localhost"
  token = "foo"
}
`, name)
}
