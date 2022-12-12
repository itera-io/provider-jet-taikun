#!/bin/bash

source data.sh

#MAKE SECRET YAML

cp example_secret.yaml secret.yaml
sed -i "s/PASSWORD/$CP_BILL_CRED_PASSWORD_B/g" secret.yaml
kubectl apply -f secret.yaml

#MAKE OS CREDS YAML

cp example_"$RESOURCE".yaml $DEST

sed -i "s/PROVIDER/$PROVIDER_NAME/g" $DEST

sed -i "s/BILLING_CREDENTIAL/$CP_BILL_CRED_NAME/g" $DEST
sed -i "s^URL^$CP_BILL_CRED_URL^g" $DEST
sed -i "s/USERNAME/$CP_BILL_CRED_USER/g" $DEST
sed -i "s/ORGANIZATION_REF/$CP_BILL_CRED_ORG/g" $DEST
sed -i "s/LOCK/$CP_BILL_CRED_LOCK/g" $DEST
