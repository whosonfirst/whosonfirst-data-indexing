#!/bin/sh

# See also: ../update/cmd/update/main.go which does the same thing and can be run as a Lambda function.

# https://github.com/whosonfirst/docker-whosonfirst-data-indexing/issues/4

WHOAMI=`realpath $0`
FNAME=`basename $WHOAMI`
ROOT=`dirname $WHOAMI`

OS=`uname -s | tr '[:upper:]' '[:lower:]'`

BIN="/usr/bin"

# Pull in defaults from .env file
if [ -f ${ROOT}/${FNAME}.env ]
then
    source ${ROOT}/${FNAME}.env
fi

echo ${BIN}/wof-list-repos -org whosonfirst-data -prefix whosonfirst-data -updated-since ${UPDATED_SINCE}
REPOS=`${BIN}/wof-list-repos -org whosonfirst-data -prefix whosonfirst-data -updated-since ${UPDATED_SINCE}`

for REPO in ${REPOS}
do

    echo "Invoke ${ECS_TASK} for ${REPO}"

    # echo ${BIN}/ecs-launch-task -dsn 'credentials=iam: region=us-east-1' -task ${ECS_TASK} -container ${ECS_CONTAINER} -cluster ${ECS_CLUSTER} -launch-type FARGATE -public-ip ENABLED -security-group ${ECS_SECURITY_GROUP} -subnet ${ECS_SUBNET} ${ECS_COMMAND} ${REPO}
    
    ${BIN}/ecs-launch-task -dsn 'credentials=iam: region=us-east-1' -task ${ECS_TASK} -container ${ECS_CONTAINER} -cluster ${ECS_CLUSTER} -launch-type FARGATE -public-ip ENABLED -security-group ${ECS_SECURITY_GROUP} -subnet ${ECS_SUBNET} ${ECS_COMMAND} ${REPO}
    
done
