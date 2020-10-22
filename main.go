/**
 * provider/main.go
 */

package main

import (
        "github.com/hashicorp/terraform-plugin-sdk/plugin"
        "github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Providers are distributed as standalone binaries. This is the entrypoint
// that is invoked by Terraform for this provider. During development you
// will need to sideload this provider; place the provider binary in the
// directory ~/.terraform.d/plugins, or symbolically link it, after which
// a `terraform init` should succeed. Note that the version should also be
// specified, e.g. terraform-provider-telus_v0.0.1.

func main() {
	// The plugin library handles the communication between Terraform
	// and the plugin.
        plugin.Serve(&plugin.ServeOpts{
                ProviderFunc: func() terraform.ResourceProvider {
                        return Provider()
                },
        })
}
