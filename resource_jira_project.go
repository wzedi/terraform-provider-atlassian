package main

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceJiraProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceJiraProjectCreate,
		Read:   resourceJiraProjectRead,
		Update: resourceJiraProjectUpdate,
		Delete: resourceJiraProjectDelete,

		Schema: map[string]*schema.Schema{
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Description: "Project keys must be unique and start with an uppercase letter followed by one or more uppercase alphanumeric characters. " +
					"The maximum length is 10 characters. Required when creating a project. Optional when updating a project.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the project. Required when creating a project. Optional when updating a project.",
			},
			"project_type_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Description: "The project type, which dictates the application-specific feature set. Required when creating a project. " +
					"Not applicable for the Update project resource." +
					"Valid values: ops, software, service_desk, business",
			},
			"project_template_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Description: "A prebuilt configuration for a project. The type of the projectTemplateKey must match with the type of the projectTypeKey. " +
					"Required when creating a project. Not applicable for the Update project resource." +
					"Valid values: com.pyxis.greenhopper.jira:gh-simplified-agility-kanban, com.pyxis.greenhopper.jira:gh-simplified-agility-scrum, " +
					"com.pyxis.greenhopper.jira:gh-simplified-basic, com.pyxis.greenhopper.jira:gh-simplified-kanban-classic, " +
					"com.pyxis.greenhopper.jira:gh-simplified-scrum-classic, com.atlassian.servicedesk:simplified-it-service-desk, " +
					"com.atlassian.servicedesk:simplified-internal-service-desk, com.atlassian.servicedesk:simplified-external-service-desk, " +
					"com.atlassian.jira-core-project-templates:jira-core-simplified-content-management, " +
					"com.atlassian.jira-core-project-templates:jira-core-simplified-document-approval, " +
					"com.atlassian.jira-core-project-templates:jira-core-simplified-lead-tracking, " +
					"com.atlassian.jira-core-project-templates:jira-core-simplified-process-control, " +
					"com.atlassian.jira-core-project-templates:jira-core-simplified-procurement, " +
					"com.atlassian.jira-core-project-templates:jira-core-simplified-project-management, " +
					"com.atlassian.jira-core-project-templates:jira-core-simplified-recruitment, " +
					"com.atlassian.jira-core-project-templates:jira-core-simplified-task-, " +
					"com.atlassian.jira.jira-incident-management-plugin:im-incident-management",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "A brief description of the project.",
			},
			"lead_account_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The account id of the project lead. Required when creating a project. Optional when updating a project.",
			},
			"url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "A link to information about this project, such as project documentation",
			},
			"assignee_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "UNASSIGNED",
				Description: "The default assignee when creating issues for this project." +
					"Valid values: PROJECT_LEAD, UNASSIGNED",
			},
			"avatar_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  int64(0),
				Description: "An integer value for the project's avatar. " +
					"Format: int64",
			},
			"issue_security_scheme": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
				Description: "The ID of the issue security scheme for the project, which enables you to control who can and cannot view issues. " +
					"Use the Get issue security schemes resource to get all issue security scheme IDs." +
					"Format: int64",
			},
			"permission_scheme": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
				Description: "The ID of the permission scheme for the project. Use the Get all permission schemes resource to see a list of all permission scheme IDs. " +
					"Format: int64",
			},
			"notification_scheme": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
				Description: "The ID of the notification scheme for the project. Use the Get notification schemes resource to get a list of notification scheme IDs. " +
					"Format: int64",
			},
			"category_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
				Description: "The ID of the project's category. A complete list of category IDs is found using the Get all project categories operation. " +
					"Format: int64",
			},
		},
	}
}

func createOrUpdateProject(d *schema.ResourceData, m interface{}, create bool) error {

	data := map[string]interface{}{
		"key":                d.Get("key").(string),
		"name":               d.Get("name").(string),
		"projectTypeKey":     d.Get("project_type_key").(string),
		"projectTemplateKey": d.Get("project_template_key").(string),
		"description":        d.Get("description").(string),
		"leadAccountId":      d.Get("lead_account_id").(string),
		"url":                d.Get("url").(string),
		"assigneeType":       d.Get("assignee_type").(string),
		// "avatarId":            d.Get("avatar_id").(int),
		// "issueSecurityScheme": d.Get("issue_security_scheme").(int),
		// "permissionScheme":    d.Get("permission_scheme").(int),
		// "notificationScheme":  d.Get("notification_scheme").(int),
		// "categoryId":          d.Get("category_id").(int),
	}

	client := m.(*ApiClient)

	reqBody, _ := json.Marshal(data)
	method := "POST"
	resource := "project"
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

func resourceJiraProjectCreate(d *schema.ResourceData, m interface{}) error {
	return createOrUpdateProject(d, m, true)
}

func resourceJiraProjectRead(d *schema.ResourceData, m interface{}) error {
	key := d.Get("key").(string)

	client := m.(*ApiClient)
	resp, err := client.request("GET", "project/"+key, "")
	if err != nil {
		log.Printf("[ERROR] Error sending API request: %s", err)
		return err
	} else {
		log.Printf("[INFO] Response: key: %s", resp["key"].(string))
		return nil
	}
}

func resourceJiraProjectUpdate(d *schema.ResourceData, m interface{}) error {
	return createOrUpdateProject(d, m, false)
}

func resourceJiraProjectDelete(d *schema.ResourceData, m interface{}) error {
	key := d.Get("key").(string)

	client := m.(*ApiClient)
	_, err := client.request("DELETE", "project/"+key, "")
	if err != nil {
		log.Printf("[ERROR] Error sending API request: %s", err)
		return err
	} else {
		log.Printf("[INFO] Project deleted")
		return nil
	}
}
