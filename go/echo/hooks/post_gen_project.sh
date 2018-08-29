#!/usr/bin/env bash

pushd ../{{ cookiecutter.svc_name }}
make
popd
