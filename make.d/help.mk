#
# Self documented makefile
#
# Usage:
#	## Some help message
#	target:
#
# @author: Alexandru Guzinschi <alex@gentle.ro>
# @author: Crifan Li <admin@crifan.com>

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)

TARGET_MAX_CHAR_NUM=15

## Show this help message
help:
	@echo '$(NAME) $(VERSION)'
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}$(MAKE)${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-_0-9]+:/ { \
	helpMessage = match(lastLine, /^## (.*)/); \
	if (helpMessage) { \
		helpCommand = substr($$1, 0, index($$1, ":")); \
		sub(/:/, "", helpCommand); \
		helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
		printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
	} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST) | sort
