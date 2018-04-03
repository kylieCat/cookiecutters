#!/usr/bin/env bash

set -xe

VERSION=$1
TARGET=$2
APP={{ cookiecutter.svc_name }}

if [ "${VERSION}" == "" ] 
then
	echo "No version supplied, exiting"
	exit 1
fi

if [ "${TARGET}" == "" ] 
then
	echo "No target supplied, exiting"
	exit 1
fi

CLUSTERS="ause1a ause1b"

for CLUSTER in ${CLUSTERS}
do
	kubectl config set current-context ${TARGET}-${CLUSTER}
	kubectl config set-context $(kubectl config current-context) --namespace=sre
	helm upgrade ${APP} stg/${APP} --set image.tag=${VERSION} --values=manifests/${APP}/${TARGET}.yaml
done
