#!/bin/bash

source config.sh

RESOURCE="project"
RESOURCE_NAME=test-"$RESOURCE"
DEST=$PRE$RESOURCE$END
WAITING=3689

export CP_PROJECT_NAME="$RESOURCE_NAME"-"$RANDOM"
echo -n "$CP_PROJECT_NAME" > ref
export CP_PROJECT_CLOUD_CREDS="test-all-cloud-credential"
export CP_PROJECT_ORG="test-all-organization-attach"
export CP_PROJECT_QDISK="38000"
export CP_PROJECT_QRAM="38000"
export CP_PROJECT_KP="test-kubernetes-profile-17643"
export CP_PROJECT_FLAVOR="m1.medium"
export CP_PROJECT_IMG="b2a2939d-3609-4f15-92b7-b5dc3dd5c9d3"
export CP_PROJECT_VM_NAME="delete-vm"
export CP_PROJECT_VM_FLAVOR="m1.medium"
export CP_PROJECT_VM_IMG_ID="b2a2939d-3609-4f15-92b7-b5dc3dd5c9d3"
export CP_PROJECT_VM_SA="$TESTS_STANDALONE_REF"
export CP_PROJECT_VM_VOL_SIZE="38"
export CP_PROJECT_VM_USER="alexis"
