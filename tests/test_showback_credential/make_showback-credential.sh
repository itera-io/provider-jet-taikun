#!/bin/bash

source data.sh

#MAKE SECRET YAML

cp example_secret.yaml secret.yaml
sed -i "s/PASSWORD/$CP_SHOWBACK_CRED_PASSWORD_B/g" secret.yaml
kubectl apply -f secret.yaml

#MAKE OS CREDS YAML

cp example_"$RESOURCE".yaml $DEST

sed -i "s^SHOWBACK_CREDENTIAL^$CP_SHOWBACK_CRED_NAME^g" $DEST
sed -i "s^URL^$CP_SHOWBACK_CRED_URL^g" $DEST
sed -i "s^USER^$CP_SHOWBACK_CRED_USER^g" $DEST
sed -i "s^ORGANIZATION_REF^$CP_SHOWBACK_CRED_ORG^g" $DEST
sed -i "s^LOCK^$CP_SHOWBACK_CRED_LOCK^g" $DEST
