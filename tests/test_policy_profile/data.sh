#!/bin/bash

source config.sh

RESOURCE="policy-profile"
RESOURCE_NAME=test-"$RESOURCE"
DEST=$PRE$RESOURCE$END
WAITING=180

export CP_POLICY_PROFILE="$RESOURCE_NAME"-"$RANDOM"
echo -n "$CP_POLICY_PROFILE" > ref
export CP_POLICY_ORG="test-all-organization-attach"
export CP_POLICY_REPOS="repo"
export CP_POLICY_TAGS="tag"
export CP_POLICY_LOCK="false"
export CP_POLICY_HTTP="true"
export CP_POLICY_NODE="true"
export CP_POLICY_PROBE="true"
export CP_POLICY_INGRESS="true"
export CP_POLICY_SELECTOR="true"
