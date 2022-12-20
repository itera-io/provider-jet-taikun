#!/bin/bash

source config.sh

RESOURCE="billing-rule"
RESOURCE_NAME=test-"$RESOURCE"
DEST=$PRE$RESOURCE$END
WAITING=180

export CP_BILL_RULE_NAME="billrule-$TAIKUN_USER-$RANDOM"
echo -n "$CP_BILL_RULE_NAME" > ref
export CP_BILL_RULE_CREDS="test-$TAIKUN_USER-billi-c"
export CP_BILL_RULE_TYPE="Sum"
export CP_BILL_RULE_PRICE=1938
