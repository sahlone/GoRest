#!/usr/bin/env bash
set -ufe -o pipefail

usage() {
  echo "Usage: $0  <action up|kill|rm|down>" 1>&2
  exit 1
}

SCRIPT_DIR=`dirname ${BASH_SOURCE[0]}`
echo $SCRIPT_DIR

if [ "$#" -gt "0" ]; then
    ACTION=${1}; else echo "Incorrect operation"; usage;
fi

echo $ACTION
if    !([ "${ACTION}" = "up" ]  || [ "${ACTION}" = "down" ] \
         || [ "${ACTION}" = "rm" ] ); then
         ( echo "Incorrect operation: ${ACTION}"; usage )

fi

shift $((OPTIND-1))

DOCKER_COMPOSE_FILE="$SCRIPT_DIR/../docker-compose.yml"

if [ "${ACTION}" = "up" ]; then
  ACTION="${ACTION} -d "
elif [ "${ACTION}" = "rm" ]; then
  ACTION="${ACTION} -f"
elif [ "${ACTION}" = "down" ]; then
    ACTION="${ACTION} -v --remove-orphans"
fi


# By default, docker-compose start network/container with prefix name to be a folder containing the yaml (i.e. docker)
# use --project-name to override with some random name to avoid conflict when running build at the same time
PROJECT_NAME=recipes
echo "Executing ${ACTION} with Docker-Compose with --project-name=${PROJECT_NAME}"
docker-compose --project-name ${PROJECT_NAME} -f ${DOCKER_COMPOSE_FILE} ${ACTION}

if [[ $ACTION == *"up"* ]]; then
    echo "Forwarding logs ..."
    docker logs -f "${PROJECT_NAME}_mongodb_1" > ./../docker/mongo.log &
    docker logs -f "${PROJECT_NAME}_go_1" > ./../docker/go-container.log &
    COMPOSE_HTTP_TIMEOUT=20 docker-compose -f ${DOCKER_COMPOSE_FILE} logs -f > ./../docker/docker-compose.log &
    docker exec -it "${PROJECT_NAME}_go_1" tail -f application.log
fi
