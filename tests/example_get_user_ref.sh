#!/bin/bash

WAITING=185
COUNT=0
COUNT_LIMIT=$(($WAITING / $JUMP_TIME))
REF=""
while [ "$REF" = "" ];
do
    COUNT=$(($COUNT + 1))
    if [ $COUNT -gt $COUNT_LIMIT ];
    then
        echo "TIMEOUT!"
        kubectl delete -f "$DEST"
        rm out*
        rm ref
        rm "$DEST"
        exit 1
    fi

    #REF=`cat ref`
    kubectl get managed > out_raw
    grep "test-$TAIKUN_USER-user" "out_raw" > out
    CONTENT=`cat out`

    #echo "$CONTENT"

    READY=`grep -o True out | wc -l`

    sed -i "s/.* \([a-z0-9-]\+\) .*/\1/g" "out"

    CONTENT_AFTER=`cat out`

    if [ $READY -eq 2 -a "$CONTENT" != "$CONTENT_AFTER" ];
    then
        RESOURCES=`cat out | tr " " "\n"`
        for id in $RESOURCES; do
            REF="$id"
        done
    else
        REF=""
    fi

    sleep $JUMP_TIME
done

#echo $REF

export TESTS_USER_REF=$REF

rm out_raw
rm out
#rm ref
