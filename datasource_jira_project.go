package main

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func datasourceJiraProject() *schema.Resource {
	return &schema.Resource{
		Read: datasourceJiraProjectRead,

		Schema: map[string]*schema.Schema{
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Description: "Project keys must be unique and start with an uppercase letter followed by one or more uppercase alphanumeric characters. " +
					"The maximum length is 10 characters. Required when creating a project. Optional when updating a project.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the project. Required when creating a project. Optional when updating a project.",
			},
			"project_type_key": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Description: "The project type, which dictates the application-specific feature set. Required when creating a project. " +
					"Not applicable for the Update project resource." +
					"Valid values: ops, software, service_desk, business",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A brief description of the project.",
			},
		},
	}
}

func datasourceJiraProjectRead(d *schema.ResourceData, m interface{}) error {
	key := d.Get("key").(string)

	client := m.(*ApiClient)
	resp, err := client.request("GET", "project/"+key, "")
	if err != nil {
		log.Printf("[ERROR] Error sending API request: %s", err)
		return err
	} else {
		//log.Println("[INFO] Response: ", resp)

		d.SetId(resp["id"].(string))
		d.Set("key", resp["key"].(string))
		d.Set("project_type_key", resp["projectTypeKey"].(string))
		d.Set("description", resp["description"].(string))

		return nil
	}
}

