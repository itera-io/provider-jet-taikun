#!/bin/bash

source config.sh

RESOURCE="kubernetes-profile"
RESOURCE_NAME=test-"$RESOURCE"
DEST=$PRE$RESOURCE$END
WAITING=180

export CP_KP_NAME="$RESOURCE_NAME"-"$RANDOM"
echo -n "$CP_KP_NAME" > ref
export CP_KP_ORG="test-alexis-org-attach"
