/**
 * provider/provider.go
 */

package main

import (
	// Provided as part of Terraform Core; used to ensure consistency between
	// providers and abstract away various implemention details.
        "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// A schema.Provider defines:
// - the configuration keys the provider accepts
// - the resources it supports
// - any callbacks to configure

// NB: The resources that are provided must have the provider name as a prefix.
// For example, a resource named "server" that is provided by the "example" provider
// must be named "example_server" in the resource below. The resource itself is
// defined in resource_server.go.

func Provider() *schema.Provider {
        return &schema.Provider{
                ResourcesMap: map[string]*schema.Resource{
			"telus_package_json": resourceServer(),
		},
        }
}
