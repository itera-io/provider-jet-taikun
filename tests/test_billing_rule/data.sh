#!/bin/bash

source config.sh

RESOURCE="billing-rule"
RESOURCE_NAME=test-"$RESOURCE"
DEST=$PRE$RESOURCE$END
WAITING=180

export CP_BILL_RULE_NAME="$RESOURCE_NAME"-"$RANDOM"
echo -n "$CP_BILL_RULE_NAME" > ref
export CP_BILL_RULE_CREDS="test-all-billing-credential"
export CP_BILL_RULE_TYPE="Sum"
export CP_BILL_RULE_PRICE=1938
