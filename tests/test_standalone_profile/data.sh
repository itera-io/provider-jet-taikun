#!/bin/bash

source config.sh

RESOURCE="standalone-profile"
RESOURCE_NAME=test-"$RESOURCE"
DEST=$PRE$RESOURCE$END
WAITING=180

export CP_SA_NAME="ts-sp-$TAIKUN_USER-$RANDOM"
echo -n "$CP_SA_NAME" > ref
export CP_SA_PUBLIC_KEY=$PUBLIC_KEY
export CP_SA_ORG="ts-$TAIKUN_USER-org-attach"
