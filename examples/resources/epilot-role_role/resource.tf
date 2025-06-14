terraform {
  required_providers {
    epilot-role = {
      source  = "epilot-dev/epilot-role"
      version = "0.20.8"
    }
  }
}

provider "epilot-role" {
  # Configuration options

  epilot_auth = "eyJraWQiOiJyd1wvSUdNeXJIU1wvV1wvMHlkOWoyT3Z6eURxTDEwM1RCUUNaUkVBejVMSUpBPSIsImFsZyI6IlJTMjU2In0.eyJzdWIiOiI3YmNjMzE2Ni03YmRhLTRmN2EtOTdjZi1mYjgxODAxYjUzNWUiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiaXNzIjoiaHR0cHM6XC9cL2NvZ25pdG8taWRwLmV1LWNlbnRyYWwtMS5hbWF6b25hd3MuY29tXC9ldS1jZW50cmFsLTFfNGh3VXlSWkVCIiwiY3VzdG9tOml2eV9vcmdfaWQiOiI3MzkyMjQiLCJjb2duaXRvOnVzZXJuYW1lIjoibi5nb2VsQGVwaWxvdC5jbG91ZCIsImN1c3RvbTppdnlfdXNlcl9pZCI6IjExMDAwMDM3IiwiYXVkIjoiNDdwcjdzdDdsNHVtYm1wZmJpanY2N21odWEiLCJldmVudF9pZCI6IjUwZjgxMGM0LTZlYzgtNDNjYS05ZmZlLTRlMTA3ZDllYWNiZiIsInRva2VuX3VzZSI6ImlkIiwiYXV0aF90aW1lIjoxNzQ5NTY5MDU5LCJleHAiOjE3NDk5MTI3MTAsImlhdCI6MTc0OTkwOTExMCwiZW1haWwiOiJuLmdvZWxAZXBpbG90LmNsb3VkIn0.X2ZQFHbIbxZ5s9Z3pseNBdmXnfFCQTQhsboNZI0Tb0-FRfLOhe80h1ej7RiYxmsCmTe5fYLHaVnBiPq2t4az4d7xPpQNRvZqRmQxbld9muvILzZAjcjfIn4PuxW3jNPN1SjoOPu-WozwjssxKCb4GOyyRkZa5m_r7CIN9tU9fnsHiIM9k1ZA_vNUACHnhcm-g0Z4y5Z0S_ClK8RQSzSCvx2GgF_u1NmQLKmKvOxStcb-B_Eu43W1ZgDq2YOz7tzO_o3RrOhdqzCZh7zwPF8gOka9hgWpdPge3_e7hv3-4a6b84ni09A0uON6jryKYxWwsMaf0raGayAKGPxoKQ4-bQ"
  server_url  = "https://permissions.dev.sls.epilot.io"
}


resource "epilot-role_role" "my_role" {
  grants = [
    {
      action   = "datalake:*"
      resource = "*"
    },
    {
      action   = "entity:*"
      resource = "*"
    },
    {
      action   = "entity:view"
      resource = "*"
      conditions = [
        {
          operation = "equals"
          attribute = "_acl.view"
          values    = [
            "org_739224"
          ]
        }
      ]
    },
  ]
  name = "Terraform Created Role"
  slug = "terraform-created-role"
  type = "user_role"
}
