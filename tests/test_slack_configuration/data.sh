#!/bin/bash

source config.sh

RESOURCE="slack-configuration"
RESOURCE_NAME=test-"$RESOURCE"
DEST=$PRE$RESOURCE$END
WAITING=180

export CP_SC_NAME="$RESOURCE_NAME"-"$RANDOM"
echo -n "$CP_SC_NAME" > ref
export CP_SC_ORG="test-all-organization"
export CP_SC_CHANNEL="random"
export CP_SC_TYPE="General"
export CP_SC_URL="https://test-crossplane.slack.com/archives/C043Z7TNESZ"
