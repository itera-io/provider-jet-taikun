#!/bin/bash

source data.sh

#MAKE SECRET YAML

cp example_secret.yaml secret.yaml
sed -i "s^PASSWORD^$CP_BACKUP_CRED_SECRET_KEY_B^g" secret.yaml
kubectl apply -f secret.yaml

#MAKE BACKUP YAML

cp example_"$RESOURCE".yaml $DEST

sed -i "s^BACKUP_CREDENTIAL^$CP_BACKUP_CRED_NAME^g" $DEST
sed -i "s^ACCESS_KEY_ID^$CP_BACKUP_CRED_ACCESS_KEY^g" $DEST
sed -i "s^ENDPOINT^$CP_BACKUP_CRED_ENDPOINT^g" $DEST
sed -i "s^REGION^$CP_BACKUP_CRED_REGION^g" $DEST
sed -i "s^ORGANIZATION_REF^$CP_BACKUP_CRED_ORG^g" $DEST
sed -i "s^LOCK^$CP_BACKUP_CRED_LOCK^g" $DEST
