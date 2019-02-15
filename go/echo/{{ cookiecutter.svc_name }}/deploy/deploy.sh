#!/usr/bin/env bash

set -e

GREEN="\033[0;38;5;2m"
YELLOW="\033[0;38;5;11m"
RED="\033[0;38;5;9m"
COLOR_OFF="\033[0m"
NS="sre"

TARGET=$1
APP=${PWD##*/}

if [[ "${TARGET}" == "" ]]
then
    echo "ERROR: No target supplied, exiting"
    exit 1
fi

get_version() {
    current_version=$(cat .bumpversion.cfg | grep "current_version =" | sed 's/current_version = //')
    printf "${current_version}\n"
}

get_environment() {
    target=${1}
    printf "unity-ie-sre-isolated-${target}/sre-gke-main"
}
function get_message() {
    env=$(get_environment ${TARGET})
    svc_name=${APP}
    commit=$(git rev-parse HEAD)
    version=${1}
    echo "*Deployment into ${env}*\n*Service:* ${svc_name}\n*Revision:* ${commit}\n*Context:* ${env}\n*Chart Version:* ${version}"
}

get_slack_url(){
    vault read -format json -field data secret/sre/slackHook | jq -r '.slackURL'
}

get_payload() {
    message=${1}
    echo -n "payload={\"channel\": \"#sre-deploys\", \"username\": \"SRE Deployment\", \"text\": \"${message}\"}"
}

send_slack_message(){
    message=${1}
    url=$(get_slack_url)
    payload=$(get_payload "${message}")
    curl -X POST --data-urlencode "${payload}" ${url}
}

VERSION=$(get_version)
CONTEXT="sre-${TARGET}"
printf "Deploying ${GREEN}${APP}:${VERSION}${COLOR_OFF} to ${YELLOW}${CONTEXT}${COLOR_OFF}\n"
helm upgrade --install ${APP} ./deploy/${APP}-${VERSION}.tgz \
    --version ${VERSION} \
    --kube-context ${CONTEXT} \
    --values=deploy/${APP}/values.yaml \
    --values=deploy/${APP}/${TARGET}.yaml
send_slack_message "$(get_message ${VERSION})"