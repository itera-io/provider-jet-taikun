#!/bin/bash

source config.sh

RESOURCE="organization"
RESOURCE_NAME=test-"$RESOURCE"
DEST=$PRE$RESOURCE$END
WAITING=55

export CP_ORG_NAME="$RESOURCE_NAME"-"$RANDOM"
echo -n "$CP_ORG_NAME" > ref
export CP_ORG_DESCRIPTION=test-organization-desc
export CP_ORG_DISCOUNT=42
export CP_ORG_CITY=test-organization-city
export CP_ORG_BEMAIL=test_organization@empty.com
export CP_ORG_EMAIL=test_organization@againempty.com
export CP_ORG_PHONE=+441989