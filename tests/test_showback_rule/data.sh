#!/bin/bash

source config.sh

RESOURCE="showback-rule"
RESOURCE_NAME=test-"$RESOURCE"
DEST=$PRE$RESOURCE$END
WAITING=180

export CP_SHOWBACK_RULE_NAME="$RESOURCE_NAME"-"$RANDOM"
echo -n "$CP_SHOWBACK_RULE_NAME" > ref
export CP_SHOWBACK_RULE_ORG="test-all-organization-attach"
export CP_SHOWBACK_RULE_PRICE=1938
export CP_SHOWBACK_RULE_KIND="External"
export CP_SHOWBACK_RULE_TYPE="Sum"
export CP_SHOWBACK_RULE_CREDS="test-all-showback-credential"
export CP_SHOWBACK_RULE_GLOBAL_LIMIT=38
export CP_SHOWBACK_RULE_PROJECT_LIMIT=38
