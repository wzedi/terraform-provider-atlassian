# terraform-provider-atlassian
Terraform Atlassian provider

## Testing

1. Build the provider executable: `go build -o terraform-provider-atlassian`
1. Initialise Terraform: `terraform init`
1. Run a Terraform plan: `terraform plan -v api_key <Atlassian API key> -v api_user <Atlassian login email address> -v site_url <Atlassian site URL>`
1. Apply the plan: `terraform apply -v api_key <Atlassian API key> -v api_user <Atlassian login email address> -v site_url <Atlassian site URL>`
