#!/bin/bash

source config.sh

RESOURCE="access-profile"
RESOURCE_NAME=test-"$RESOURCE"
DEST=$PRE$RESOURCE$END
WAITING=180

export CP_ACCESS_NAME="$RESOURCE_NAME"-"$RANDOM"
echo -n "$CP_ACCESS_NAME" > ref
export CP_ACCESS_ORG="test-$USER_TEST-org-attach"
