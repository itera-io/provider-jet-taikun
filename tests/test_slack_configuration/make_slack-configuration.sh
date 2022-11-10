#!/bin/bash

source data.sh
cp example_"$RESOURCE".yaml $DEST

sed -i "s/SLACK_CONFIGURATION/$CP_SC_NAME/g" $DEST
sed -i "s/ORGANIZATION_REF/$CP_SC_ORG/g" $DEST
sed -i "s/CHANNEL/$CP_SC_CHANNEL/g" $DEST
sed -i "s/TYPE/$CP_SC_TYPE/g" $DEST
sed -i "s^URL^$CP_SC_URL^g" $DEST
