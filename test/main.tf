variable "api_key" {
  type = string
}

variable "api_user" {
  type = string
}

variable "site_url" {
  type = string
}

provider "atlassian" {
  api_user = var.api_user
  api_key  = var.api_key
  site_url = var.site_url
  version  = "~> 0.0"
}

data "atlassian_jira_project" "help_project" {
  key = "HELP"
}


data "atlassian_jira_issue_type" "get_help" {
  id = 10011
}


# resource "atlassian_jira_project" "test_project" {
#   key                  = "TEST"
#   name                 = "Terraform Test Project"
#   project_type_key     = data.atlassian_jira_project.help_project.project_type_key
#   project_template_key = "com.atlassian.servicedesk:simplified-external-service-desk"
#   lead_account_id      = "557058:0e8f96f7-a27f-4701-aff1-4969c29259eb"
# }

