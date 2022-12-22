#!/bin/bash

source config.sh

RESOURCE="backup-credential"
RESOURCE_NAME=test-"$RESOURCE"
DEST=$PRE$RESOURCE$END
WAITING=180

export CP_BACKUP_CRED_NAME="ts--bc-$TAIKUN_USER-$RANDOM"
echo -n "$CP_BACKUP_CRED_NAME" > ref
export CP_BACKUP_CRED_ORG="ts--$TAIKUN_USER-org-attach"
export CP_BACKUP_CRED_ACCESS_KEY=$S3_ACCESS_KEY
export CP_BACKUP_CRED_ENDPOINT="https://test-backup-region.taikun.dev"
export CP_BACKUP_CRED_REGION="taikun-east-1"
export CP_BACKUP_CRED_LOCK="false"

echo -n "$S3_SECRET_KEY" | base64 > passwordb
export CP_BACKUP_CRED_SECRET_KEY_B=`cat passwordb`
rm passwordb
