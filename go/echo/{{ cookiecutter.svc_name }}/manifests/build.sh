#!/usr/bin/env bash

VERSION=$1
NO_CLEANUP=$2

APP={{ cookiecutter.svc_name }}
CHARTS_MUSEUM=https://chartmuseum.stg.internal.unity3d.com/api/charts
DOCKER_REPO=artifactory.eu-cph-1.unityops.net:5010

if [ "${NO_CLEANUP}" == "" ]; then
	OLD_VERSION=$(echo "${VERSION} - .1"|bc)
	OLD_DOCKER_TAG="${DOCKER_REPO}/${APP}:"${OLD_VERSION}
	docker images rm $( docker images |grep "${OLD_DOCKER_TAG}"| awk '{ print $3}' )
	curl  --request DELETE ${CHARTS_MUSEUM}/${APP}/${OLD_VERSION}
fi
curl  --request DELETE ${CHARTS_MUSEUM}/${APP}/${VERSION}

DOCKER_TAG="${DOCKER_REPO}/${APP}:"${VERSION}

docker build -t ${DOCKER_TAG} .
docker push ${DOCKER_TAG}

REPO_URL=$( helm repo list | grep stg |awk '{ print $2}' )
echo ${REPO_URL}
helm package --version ${VERSION} manifests/${APP}

curl  --data-binary "@${APP}-${VERSION}.tgz" ${REPO_URL}/api/charts
helm repo index manifests/${APP} --url ${REPO_URL}
helm repo update
