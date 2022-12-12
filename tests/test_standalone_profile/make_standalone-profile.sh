#!/bin/bash

source data.sh
cp example_"$RESOURCE".yaml $DEST

sed -i "s/PROVIDER/$PROVIDER_NAME/g" $DEST

sed -i "s/STANDALONE/$CP_SA_NAME/g" $DEST
sed -i "s^PUBLIC_KEY^$CP_SA_PUBLIC_KEY^g" $DEST
sed -i "s/ORGANIZATION_REF/$CP_SA_ORG/g" $DEST
