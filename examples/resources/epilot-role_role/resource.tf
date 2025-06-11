resource "epilot-role_role" "my_role" {
  schemas = {
    expires_at = "2028-07-21T17:32:28Z"
    grants = [
      {
        action = "entity-read"
        conditions = [
          {
            equals_condition = {
              attribute = "workflows.primary.task_name"
              operation = "equals"
              values = [
                "Qualification"
              ]
            }
          }
        ]
        dependencies = [
          {
            action = "entity-read"
            conditions = [
              {
                equals_condition = {
                  attribute = "workflows.primary.task_name"
                  operation = "equals"
                  values = [
                    "Qualification"
                  ]
                }
              }
            ]
            effect   = "deny"
            resource = "entity:123:contact:f7c22299-ca72-4bca-8538-0a88eeefc947"
          }
        ]
        effect   = "allow"
        resource = "entity:123:contact:f7c22299-ca72-4bca-8538-0a88eeefc947"
      }
    ]
    id              = "123:owner"
    name            = "Owner"
    organization_id = "123"
    partner_org_id  = "123"
    pricing_tier    = "Professional"
    slug            = "owner"
    type            = "user_role"
  }
}