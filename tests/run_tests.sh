#!/bin/bash

cd tests

ls > out_all
grep "test_" out_all > .out_test
TESTS=`cat .out_test`

#for t in $TESTS; do
#    echo $t
#done

source run_requirements.sh

#cd test_organization
#./run.sh
#cd ..

#cd test_alerting_profile
#./run.sh
#cd ..


#for curr_test in "$TESTS"; do
#    echo "$curr_test"
#    cd "$curr_test"
#    ./run.sh
#    cd ..
#done

while read curr; do
    #echo "$curr"
    cd "$curr"
    #ls
    source ./run.sh
    cd ..
done < ".out_test"

rm secret*
rm required-test-cloud-credential.yaml
