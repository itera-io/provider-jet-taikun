#!/bin/bash

source data.sh
cp example_"$RESOURCE".yaml $DEST

sed -i "s^POLICY_PROFILE^$CP_POLICY_PROFILE^g" $DEST
sed -i "s^ORGANIZATION_REF^$CP_POLICY_ORG^g" $DEST
sed -i "s^REPO^$CP_POLICY_REPOS^g" $DEST
sed -i "s^TAG^$CP_POLICY_TAGS^g" $DEST
sed -i "s^LOCK^$CP_POLICY_LOCK^g" $DEST
sed -i "s^HTTP^$CP_POLICY_HTTP^g" $DEST
sed -i "s^NODE^$CP_POLICY_NODE^g" $DEST
sed -i "s^PROBE^$CP_POLICY_PROBE^g" $DEST
sed -i "s^INGRESS^$CP_POLICY_INGRESS^g" $DEST
sed -i "s^SELECTOR^$CP_POLICY_SELECTOR^g" $DEST
