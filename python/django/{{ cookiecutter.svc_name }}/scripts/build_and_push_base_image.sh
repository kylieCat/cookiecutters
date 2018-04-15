#!/usr/bin/env bash
set -e

APP={{ cookiecutter.svc_name }}
REPO={{ cookiecutter.docker_repo }}

if [[ $# -lt 1 ]]; then
    echo "Usage: build_and_push_base_image.sh -v IMAGE_VERSION [-p]" >&2
    echo "  -v: The version, used as the tag on the image."
    echo "  -p (optional): If passed image will be pushed to the remote repo."
    exit 1
fi

while getopts "v:p" OPTION
do
    case $OPTION in
        v | --version)
            new_tag=base-${OPTARG}
            echo "Using tag: ${new_tag}"
            docker build -t ${APP}:"${new_tag}" .
            docker tag ${APP}:"${new_tag}" ${REPO}/${APP}:"${new_tag}"
            echo "Built image: ${APP}:${new_tag}"
            ;;
        p | --push)
            echo "Pushing image to repo: ${APP}:${new_tag}"
            docker push ${REPO}/${APP}:${new_tag}
            ;;
    esac
done
