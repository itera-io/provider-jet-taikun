#!/bin/bash

source config.sh

RESOURCE="slack-configuration"
RESOURCE_NAME=test-"$RESOURCE"
DEST=$PRE$RESOURCE$END
WAITING=180

export CP_SC_NAME="ts-slack-$TAIKUN_USER-$RANDOM"
echo -n "$CP_SC_NAME" > ref
export CP_SC_ORG="ts-$TAIKUN_USER-org-attach"
export CP_SC_CHANNEL="random"
export CP_SC_TYPE="General"
export CP_SC_URL="https://test-crossplane.slack.com/archives/C043Z7TNESZ"
