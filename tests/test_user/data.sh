#!/bin/bash

source config.sh

RESOURCE="user"
RESOURCE_NAME=test-"$RESOURCE"
DEST=$PRE$RESOURCE$END
WAITING=180

export CP_USER_NAME="$RESOURCE_NAME"-"$RANDOM"
echo -n "$CP_USER_NAME" > ref
export CP_USER_EMAIL="testing@empty.com"
export CP_USER_ROLE="User"
export CP_USER_ORG="test-all-organization-attach"
