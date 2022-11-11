#!/bin/bash

export JUMP_TIME=2

echo -n "test-all-organization-attach" > ref

cp example_secret_cloud.yaml secret_cloud.yaml
echo -n "$CLOUD_PASSWORD" | base64 > passwordb
CLOUD_PASSWORD_B=`cat passwordb`
rm passwordb
sed -i "s/PASSWORD/$CLOUD_PASSWORD_B/g" secret_cloud.yaml
kubectl apply -f secret_cloud.yaml

cp example_secret_showback.yaml secret_showback.yaml
echo -n "$PROMETHEUS_PASSWORD" | base64 > passwordb
PROMETHEUS_PASSWORD_B=`cat passwordb`
rm passwordb
sed -i "s/PASSWORD/$PROMETHEUS_PASSWORD_B/g" secret_showback.yaml
kubectl apply -f secret_showback.yaml


kubectl apply -f required-test-organization.yaml
kubectl apply -f required-test-user.yaml

kubectl apply -f required-test-billing-credential.yaml
kubectl apply -f required-test-billing-rule.yaml
kubectl apply -f required-test-showback-credential.yaml

#source get_org_ref.sh

echo REF

cp example-cloud-credential.yaml required-test-cloud-credential.yaml
sed -i "s/ORG_REF/$ORG_REF/g" required-test-cloud-credential.yaml
kubectl apply -f required-test-cloud-credential.yaml

echo ENDING

#sleep 35
