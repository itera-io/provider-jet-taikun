#!/bin/bash

source data.sh
cp example_"$RESOURCE".yaml $DEST

sed -i "s/PROVIDER/$PROVIDER_NAME/g" $DEST

sed -i "s/USER/$CP_USER_NAME/g" $DEST
sed -i "s/EMAIL/$CP_USER_EMAIL/g" $DEST
sed -i "s/ROLE/$CP_USER_ROLE/g" $DEST
sed -i "s/ORGANIZATION_REF/$CP_USER_ORG/g" $DEST
