/**
 * provider/datasource_package_json.go
 */

// TODO make this a *data resource*

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"io/ioutil"
)

func hash(s string) string {
	sha := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sha[:])
}

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	// Typecasting to string is safe in this case because Terraform guarantees that
	// inputs match the types provided in the resource data schema.
	path := d.Get("path").(string)

	// TODO read the package.json and construct an ID from some mix of:
	// - file path
	// - package name
	// - package version
	// - hash
	// - list of dependencies
	// - list of dev dependencies

	// The existence of a non-blank ID is what tells Terraform that a resource was
	// created. If no ID is set then whether or not an error is returned, no state
	// is persisted. If an error is returned and an ID is set, the full state is
	// saved. In partial mode, only the explicitly enabled configuration keys are
	// persisted resulting in a partial state.
	d.SetId(path)

	return resourceServerRead(d, m)
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	// This method is used to sync the local state with the actual state (upstream).
	// It is called at various points by Terraform and should be a read-only operation.
	// If the ID is updated to blank, this tells Terraform that the resource no longer
	// exists.

	path := d.Get("path").(string)

	data, err := ioutil.ReadFile(path)

	if err != nil {
		d.SetId("")
		return nil
	}

	// Define data structure.
	type PackageJson struct {
		Name        string
		Version     string
		Description string
		Author      string
	}

	var obj PackageJson

	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Printf("error:", err)
	}

	fmt.Printf("name: %s", obj.Name)

	// TODO parse the file, extract the pieces of information we want, and set any
	// attributes.
	d.Set("path", path)

	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	/*
		// Enable partial state mode.
		d.Partial(true)

		if d.HasChange("example") {
			// Try updating the example attribute.
			if err := updateExample(d, m); err != nil {
				return err
			}

			d.SetPartial("example")
		}
		// If we were to return here, before disabling partial mode, then only the "example"
		// field would be saved.

	        // We succeeded, disable partial mode. This causes Terraform to save all fields again.
	        d.Partial(false)
	*/

	// If we return without error Terraform assumes that any requested changes were
	// applied without issue. The diff between current resource state and request state
	// is merged and the results persisted.
	return resourceServerRead(d, m)
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	// This is called automatically assuming no error is returned, but we add it here
	// to be explicit.
	d.SetId("")

	// You should always handle the case where the resource has already been destroyed.
	// Not error should be returned in that case.

	return nil
}

// The schema.Resource type defines the data schema and the CRUD operations required
// for the resource. These are the only things required to create a resource.

func resourceServer() *schema.Resource {
	return &schema.Resource{
		// The Create, Read, Update, and Delete operations are the only
		// required ones; others are available as well. Terraform determines
		// which function to call and with what data. Based on the schema and
		// current state of the resource, Terraform can determine if it needs
		// to create a new resource, update an existing one, or destroy one.
		// The create and update function should always return the read function
		// to ensure the state is reflected in the terraform.state file.

		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		// The schema defines the data attributes needed for the resource.
		// NB: terraform automatically performs validation and type casting.

		Schema: map[string]*schema.Schema{
			// The path to a package.json file.
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
