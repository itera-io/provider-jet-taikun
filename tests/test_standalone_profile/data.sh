#!/bin/bash

source config.sh

RESOURCE="standalone-profile"
RESOURCE_NAME=test-"$RESOURCE"
DEST=$PRE$RESOURCE$END
WAITING=180

export CP_SA_NAME="$RESOURCE_NAME"-"$RANDOM"
echo -n "$CP_SA_NAME" > ref
export CP_SA_PUBLIC_KEY="ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDCt2lLq1PhVRkoUAvD7zzpedTwGg+0H9A90QPINxvFCpM30ggQ4AdqOXaDh/lyEI6ajygQmNwy6DZGPEgsYQzeOoV6Y2ITjw7RCidf29RrJ+VtWlRAXDWgaX8SY+o3fy6thbdzlwKh1//wIqx+mTB9XV7JKxJsRU94rBkBSI1HrWV00Ac1x0/nZxnMvcXAaO8YXNz01Fsd+572P/oufhGIOJszNoVNAfREbjtZ/gTkIar5hWAs3+brBYXCM1aG9+NWMoPnQt2z5My+zaw1hDMghpj+khRaECURgwX+DwabxexxUa3eELkxKXBXq1qM+Y2RPtuZ/sT3ndx793OkfAZCY1RP2U9hcoj/IBR9DlCc20+GGimTMGpKkhYK9OZ3FwGK0nf3GQXIxiKjTH2Q2cQW6tbMiwfLsB+9kOkyzZwRvnfO/wOhjIW2lR79A1mw78JamiPZDGWqnWpvCRPmXWjDWVYYj7owf7w+kijLYOsxblgC1fXAZ90YwjKIFYtOZnU= alexis@alexis"
export CP_SA_ORG="test-all-organization-attach"
