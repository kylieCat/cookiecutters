default: deps

# COLORS
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)


TARGET_MAX_CHAR_NUM=20
## Show help
help:
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

## Run the application using 'docker-compose up --build', no need to prebuild image
run:
	docker-compose up --build

## Same as 'run' except uses the '-d' argument to up so the application runs as a daemon
daemon:
	docker-compose up -d

## Runs docker system prune to remove cruft
prune:
	yes | docker system prune

## Runs the build script for a minor release. Requires the waypoint package
minor:
	./manifests/build.sh minor

## Runs the build script for a major release. Requires the waypoint package
major:
	./manifests/build.sh major

## Runs the build script for a patch release. Requires the waypoint package
patch:
	./manifests/build.sh patch

## Deploys the latest version to the staging clusters
stage:
	./manifests/deploy.sh stg

## Deploys the latest version to the prod clusters
deploy:
	./manifests/deploy.sh prd

## Run dep ensure to install all deps
deps:
	dep ensure

## Run dep ensure to update the sre provided packages.
internal:
	dep ensure -update gitlab.internal.unity3d.com/sre/pipboy gitlab.internal.unity3d.com/sre/pipeline