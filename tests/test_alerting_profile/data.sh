#!/bin/bash

source config.sh

RESOURCE="alerting-profile"
RESOURCE_NAME=test-"$RESOURCE"
DEST=$PRE$RESOURCE$END
WAITING=180

export CP_ALERT_NAME="ts--alert-$TAIKUN_USER-$RANDOM"
echo -n "$CP_ALERT_NAME" > ref
export CP_ALERT_EMAIL="alert.$TAIKUN_USER@empty.com"
export CP_ALERT_ORG="ts--$TAIKUN_USER-org-attach"
export CP_ALERT_REMINDER="Daily"
