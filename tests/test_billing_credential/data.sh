#!/bin/bash

source config.sh

RESOURCE="billing-credential"
RESOURCE_NAME=test-"$RESOURCE"
DEST=$PRE$RESOURCE$END
WAITING=180

export CP_BILL_CRED_NAME="ts-billc-$TAIKUN_USER-$RANDOM"
echo -n "$CP_BILL_CRED_NAME" > ref
export CP_BILL_CRED_ORG="ts-$TAIKUN_USER-org-attach"
export CP_BILL_CRED_URL="https://prometheus.taikun.dev/"
export CP_BILL_CRED_USER=$PROMETHEUS_USER
export CP_BILL_CRED_LOCK="false"

echo -n "$PROMETHEUS_PASSWORD" | base64 > passwordb
export CP_BILL_CRED_PASSWORD_B=`cat passwordb`
rm passwordb
