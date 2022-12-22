#!/bin/bash

source config.sh

RESOURCE="policy-profile"
RESOURCE_NAME=test-"$RESOURCE"
DEST=$PRE$RESOURCE$END
WAITING=180

export CP_POLICY_PROFILE="ts--ppr-$TAIKUN_USER-$RANDOM"
echo -n "$CP_POLICY_PROFILE" > ref
export CP_POLICY_ORG="ts--$TAIKUN_USER-org-attach"
export CP_POLICY_REPOS="repo"
export CP_POLICY_TAGS="tag"
export CP_POLICY_LOCK="false"
export CP_POLICY_HTTP="true"
export CP_POLICY_NODE="true"
export CP_POLICY_PROBE="true"
export CP_POLICY_INGRESS="true"
export CP_POLICY_SELECTOR="true"
