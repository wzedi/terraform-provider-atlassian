package main

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceJiraIssueType() *schema.Resource {
	return &schema.Resource{
		Create: resourceJiraIssueTypeCreate,
		Read:   resourceJiraIssueTypeRead,
		Update: resourceJiraIssueTypeUpdate,
		Delete: resourceJiraIssueTypeDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique name for the issue type. The maximum length is 60 characters.",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the issue type.",
			},
			"type": &schema.Schema{
				Type:        schema.TypeBool,
				Required:    true,
				Description: "The issue type. Either subtask or standard.",
			},
		},
	}
}

func createOrUpdateIssueType(d *schema.ResourceData, m interface{}, create bool) error {

	data := map[string]interface{}{
		"name":        d.Get("name").(string),
		"description": d.Get("description").(string),
		"type":        d.Get("type").(string),
	}

	client := m.(*ApiClient)

	reqBody, _ := json.Marshal(data)
	method := "POST"
	resource := "issuetype"
	if !create {
		method = "PUT"
		resource += "/" + data["key"].(string)
	}
	resp, err := client.request(method, resource, string(reqBody))
	if err != nil {
		log.Printf("[ERROR] Error sending API request: %s", err)
		return err
	} else {
		if create {
			id := strconv.Itoa(int(resp["id"].(float64)))
			log.Printf("[INFO] Response: id: %s", id)
			d.SetId(id)
		}
		log.Printf("[INFO] Project created or updated")
		return nil
	}
}

func resourceJiraIssueTypeCreate(d *schema.ResourceData, m interface{}) error {
	return createOrUpdateIssueType(d, m, true)
}

func resourceJiraIssueTypeRead(d *schema.ResourceData, m interface{}) error {
	id := d.Get("id").(int)

	client := m.(*ApiClient)
	resp, err := client.request("GET", "issuetype/"+strconv.Itoa(id), "")
	if err != nil {
		log.Printf("[ERROR] Error sending API request: %s", err)
		return err
	} else {
		log.Printf("[INFO] Response: key: %s", resp["key"].(string))
		return nil
	}
}

func resourceJiraIssueTypeUpdate(d *schema.ResourceData, m interface{}) error {
	return createOrUpdateIssueType(d, m, false)
}

func resourceJiraIssueTypeDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Get("id").(int)

	client := m.(*ApiClient)
	_, err := client.request("DELETE", "issuetype/"+strconv.Itoa(id), "")
	if err != nil {
		log.Printf("[ERROR] Error sending API request: %s", err)
		return err
	} else {
		log.Printf("[INFO] Project deleted")
		return nil
	}
}
