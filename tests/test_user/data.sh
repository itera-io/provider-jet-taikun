#!/bin/bash

source config.sh

RESOURCE="user"
RESOURCE_NAME=test-"$RESOURCE"
DEST=$PRE$RESOURCE$END
WAITING=180

export CP_USER_NAME="ts-user-$TAIKUN_USER-$RANDOM"
echo -n "$CP_USER_NAME" > ref
export CP_USER_EMAIL="user-$TAIKUN_USER-$RANDOM@empty.com"
export CP_USER_ROLE="User"
export CP_USER_ORG="ts-$TAIKUN_USER-org-attach"
