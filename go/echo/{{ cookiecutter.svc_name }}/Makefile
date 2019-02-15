# COLORS
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)
DIRTY := $(shell git diff --stat)
TARGETS := test prd

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

## Compile and build on the host machine and then run the resulting binary.
quick:
	@go build -o ./bin/{{cookiecutter.svc_name}}-server ./cmd/server/main.go
	@./bin/{{cookiecutter.svc_name}}-server -f local -d ./src/config

## Prune docker images
prune:
	yes | docker system prune

## Deploy a new release to the target environment that increments the patch version
## Usage:
##     make patch target=[test|prod]
patch: validate
	@./deploy/build.sh patch $(target)
	@./deploy/deploy.sh $(target)

## Deploy a new release to the target environment that increments the minor version
## Usage:
##     make minor target=[test|prod]
minor: validate
	@./deploy/build.sh minor $(target)
	@./deploy/deploy.sh $(target)

## Deploy a new release to the target environment that increments the major version
## Usage:
##     make major target=[test|prd]
major: validate
	@./deploy/build.sh major $(target)
	@./deploy/deploy.sh $(target)

## Deploy a release to the target environment that uses the current version
## Usage:
##     make rebuild target=[test|prd]
rebuild: validate
	@./deploy/build.sh rebuild $(target)
	@./deploy/deploy.sh $(target)

validate: installbump
ifndef target
	$(error 'target' is undefined. Use 'test' or 'prod' as the target argument)
endif
ifeq ($(filter $(target),$(TARGETS)),)
	$(info $(target) not in valid targets; choices are: $(TARGETS))
endif
ifneq ($(DIRTY), )
	$(error git working tree is dirty; commit or stash changes and try again)
endif

## Create a virtualenv and install package for managing versions
installbump:
ifeq ($(wildcard ./.venv/.*),)
	@python3 -m venv ./.venv
	@./.venv/bin/pip install -U pip git+https://github.com/kylie-a/bumpversion.git
endif
