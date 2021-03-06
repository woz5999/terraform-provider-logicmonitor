package logicmonitor

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccLogicMonitorDeviceGroup_import(t *testing.T) {
	resourceName := "logicmonitor_device_group.testgroup"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testDeviceGroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckLogicMonitorDeviceGroupImport,
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"properties",
				},
			},
		},
	})
}

const testAccCheckLogicMonitorDeviceGroupImport = `
resource "logicmonitor_device_group" "testgroup" {
    name = "TestGroup123"
    disable_alerting = true
    description = "testing group123"
    properties {
     "group" = "test"
     "system.categories" = "a,b,c,d"
    }
}
`
