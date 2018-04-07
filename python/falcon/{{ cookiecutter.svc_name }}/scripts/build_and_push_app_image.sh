#!/usr/bin/env bash
set -e

if [[ $# -lt 1 ]]; then
    echo "Usage: build_and_app_base_image.sh -v IMAGE_VERSION [-p]" >&2
    exit 1
fi

while getopts "v:p" OPTION
do
    case $OPTION in
        v | --version)
            new_tag=app-${OPTARG}
            echo "Using tag: $new_tag"
            docker build -t api:"$new_tag" .
            docker tag api:"$new_tag" imauld/api:"$new_tag"
            echo "Built image: api:$new_tag"
            ;;
        p | --push)
            echo "Pushing image to DockerHub: api:$new_tag"
            docker push imauld/api:$new_tag
            ;;
    esac
done
