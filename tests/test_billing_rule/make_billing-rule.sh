#!/bin/bash

source data.sh
cp example_"$RESOURCE".yaml $DEST

sed -i "s^BILLING_RULE^$CP_BILL_RULE_NAME^g" $DEST
sed -i "s^BILLING_CRED_REF^$CP_BILL_RULE_CREDS^g" $DEST
sed -i "s^TYPE^$CP_BILL_RULE_TYPE^g" $DEST
sed -i "s^PRICE^$CP_BILL_RULE_PRICE^g" $DEST
