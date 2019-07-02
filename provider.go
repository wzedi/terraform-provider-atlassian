package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"site_url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions["site_url"],
			},
			"api_user": {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions["api_user"],
			},
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions["api_key"],
			},
			"api_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["api_version"],
				Default:     "3",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"atlassian_jira_project":    datasourceJiraProject(),
			"atlassian_jira_issue_type": datasourceJiraIssueType(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"atlassian_jira_project":    resourceJiraProject(),
			"atlassian_jira_issue_type": resourceJiraIssueType(),
		},
		ConfigureFunc: providerConfigure,
	}
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"api_key":  "The Atlassian API key.",
		"api_user": "The Atlassian API user email address.",
		"site_url": "The Atlassian API site URL, example: " +
			"For 'https://your-domain.atlassian.net' site-url is 'your-doman'",
		"api_version": "The Atlassian API version, defaults to 3",
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		SiteUrl:    d.Get("site_url").(string),
		ApiUser:    d.Get("api_user").(string),
		ApiKey:     d.Get("api_key").(string),
		ApiVersion: d.Get("api_version").(string),
	}

	return config.Client()
}
