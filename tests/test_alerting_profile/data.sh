#!/bin/bash

source config.sh

RESOURCE="alerting-profile"
RESOURCE_NAME=test-"$RESOURCE"
DEST=$PRE$RESOURCE$END
WAITING=55

export CP_ALERT_NAME="$RESOURCE_NAME"-"$RANDOM"
echo -n "$CP_ALERT_NAME" > ref
export CP_ALERT_EMAIL="alert@empty.com"
export CP_ALERT_ORG="test-all-organization-attach"
export CP_ALERT_REMINDER="Daily"
