#!/bin/bash

source config.sh

RESOURCE="showback-credential"
RESOURCE_NAME=test-"$RESOURCE"
DEST=$PRE$RESOURCE$END
WAITING=180

export CP_SHOWBACK_CRED_NAME="test-sc-$TAIKUN_USER-$RANDOM"
echo -n "$CP_SHOWBACK_CRED_NAME" > ref
export CP_SHOWBACK_CRED_URL="https://prometheus.taikun.dev/"
export CP_SHOWBACK_CRED_USER=$PROMETHEUS_USER

echo -n "$PROMETHEUS_PASSWORD" | base64 > passwordb
export CP_SHOWBACK_CRED_PASSWORD_B=`cat passwordb`
rm passwordb

export CP_SHOWBACK_CRED_ORG="test-$TAIKUN_USER-org-attach"
export CP_SHOWBACK_CRED_LOCK="false"
