#!/bin/bash

cd tests

ls > out_all
grep "test_" out_all > .out_test
TESTS=`cat .out_test`

source run_requirements.sh


WHITE='\033[0m'
RED='\033[0;31m'
GREEN='\033[0;32m'

while read curr; do
    cd "$curr"
    source ./run.sh

    RET=$?
    if [ "$RET" -eq "1" ]
    then
        cd ..
        rm .out*
        kubectl delete -f required-test-billing-rule.yaml
        kubectl delete -f required-test-billing-credential.yaml
        kubectl delete -f required-test-showback-credential.yaml
        kubectl delete -f required-test-cloud-credential.yaml
        kubectl delete -f secret_cloud.yaml
        kubectl delete -f required-test-kubernetes-profile.yaml
        kubectl delete -f required-test-user.yaml
        kubectl delete -f required-test-organization.yaml
        rm get*
        rm required*
        rm secret*

        printf "\n${RED}TESTSUITE FAILED !${WHITE}\n\n"
        exit 1
    fi

    cd ..
done < ".out_test"

rm .out*

#kubectl delete -f required-test-billing-rule.yaml
#kubectl delete -f required-test-billing-credential.yaml
#kubectl delete -f required-test-showback-credential.yaml

#kubectl delete -f required-test-cloud-credential.yaml
#kubectl delete -f secret_cloud.yaml

#kubectl delete -f required-test-kubernetes-profile.yaml
#kubectl delete -f required-test-user.yaml
#kubectl delete -f required-test-organization.yaml

kubectl delete -f secret_cloud.yaml

rm get*
rm required*
rm secret*

printf "\n${GREEN}TESTSUITE PASSED SUCCESSFULLY !!!${WHITE}\n\n"
exit 0
