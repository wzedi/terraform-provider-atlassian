package main

import (
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
)

func datasourceJiraIssueType() *schema.Resource {
	return &schema.Resource{
		Read: datasourceJiraIssueTypeRead,

		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The ID of the issue type.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique name for the issue type. The maximum length is 60 characters.",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The description of the issue type.",
			},
			"icon_url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "",
			},
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The issue type. Either subtask or standard.",
			},
			"avatar_id": &schema.Schema{
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "The ID of an issue type avatar.",
			},
		},
	}
}

func datasourceJiraIssueTypeRead(d *schema.ResourceData, m interface{}) error {
	id := d.Get("id").(int)

	client := m.(*ApiClient)
	log.Printf("[INFO] ID is %d", id)
	resp, err := client.request("GET", "issuetype/"+strconv.Itoa(id), "")
	if err != nil {
		log.Printf("[ERROR] Error sending API request: %s", err)
		return err
	} else {
		//log.Println("[INFO] Response: ", resp)

		respId, _ := strconv.Atoi(resp["id"].(string))
		d.SetId(resp["id"].(string))
		d.Set("id", respId)
		d.Set("description", resp["description"].(string))
		d.Set("icon_url", resp["iconUrl"].(string))
		d.Set("avatar_id", resp["avatarId"].(float64))
		if resp["subtask"].(bool) {
			d.Set("type", "subtask")
		} else {
			d.Set("type", "standard")
		}

		return nil
	}
}
