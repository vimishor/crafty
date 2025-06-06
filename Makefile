.POSIX:

ROOT_DIR 	:= $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
SHELL 		:= /bin/bash -o pipefail

NAME		:= crafty
VERSION 	?= 0.1.0
GIT_COMMIT	:= $(shell git rev-parse --short HEAD)
BUILD_DATE	:= $(shell date +'%y%m%d%H%M')

GOLDFLAGS	+= -X github.com/vimishor/crafty/internal/version.Version=${VERSION}
GOLDFLAGS	+= -X github.com/vimishor/crafty/internal/version.BuildDate=${BUILD_DATE}
GOLDFLAGS	+= -X github.com/vimishor/crafty/internal/version.GitCommit=${GIT_COMMIT}
LDFLAGS 	= -ldflags "$(GOLDFLAGS)"

# Set out default programs
GO 		= go
RM 		= rm -f
CP 		= cp
MV 		= mv
TAR 		= tar
FIND 		= find
INSTALL 	= install
MKDIR 		= mkdir -p
CRAFTY		= $(GO) run ./cmd/crafty

distdir 	= $(ROOT_DIR)/dist
builddir 	= $(ROOT_DIR)/build

prefix 		?= /usr/local
bindir 		= $(prefix)/bin
datadir 	= $(prefix)/share
mandir 		= $(prefix)/man

test_opts 	?=
os_arch 	= $(word 4, $(shell go version))
os 		?= $(word 1,$(subst /, ,$(os_arch)))
arch 		?= $(word 2,$(subst /, ,$(os_arch)))

LUA_TEST_FILES	= $(wildcard $(shell find internal/modules -type f -name '*_test.lua'))

.PRECIOUS: Makefile

ifeq (true, $(FIX))
	FIX = --fix
endif

-include make.d/*.mk

.DEFAULT_GOAL := help

## Compile program(s)
build:
	$(GO) mod tidy -diff
	GOARCH=$(arch) GOOS=$(os) $(GO) build $(LDFLAGS) -o $(builddir)/$(NAME)-$(os)-$(arch) cmd/$(NAME)/main.go

.PHONY: clean
## Remove files created by build target
clean:
	$(RM) $(builddir)/$(NAME)-$(os)-$(arch)

.PHONY: dist
## Create a distribution tar file
dist: build
	$(MKDIR) $(distdir)/$(NAME)-$(VERSION)/{,bin,man,meta,lib}
	$(CP) $(builddir)/$(NAME)-$(os)-$(arch) $(distdir)/$(NAME)-$(VERSION)/bin/$(NAME)
	$(FIND) internal/modules -type f -name '*.lua' ! -name '*_test.lua' -exec cp {} $(distdir)/$(NAME)-$(VERSION)/meta \;
	$(FIND) contrib/meta -type f -name '*.lua' -exec cp {} $(distdir)/$(NAME)-$(VERSION)/meta \;
	$(CP) -rf contrib/lib/* $(distdir)/$(NAME)-$(VERSION)/lib
	$(CP) $(ROOT_DIR)/LICENSE $(ROOT_DIR)/Readme.md $(distdir)/$(NAME)-$(VERSION)/
	$(TAR) -czvf $(NAME)-$(VERSION)-$(os)-$(arch).tar.gz -C $(distdir)/ $(NAME)-$(VERSION)

.PHONY: distclean
## Remove files created by building this program
distclean:
	$(RM) -r $(distdir)/$(NAME)-$(VERSION)
	$(RM) $(NAME)-$(VERSION)-$(os)-$(arch).tar.gz

.PHONY: install
## Compile and install the program
install: build
	$(INSTALL) -D $(builddir)/$(NAME)-$(os)-$(arch) $(DESTDIR)$(bindir)/$(NAME)
	$(INSTALL) -d $(DESTDIR)$(datadir)/$(NAME)/meta $(DESTDIR)$(datadir)/$(NAME)/lib
	$(FIND) internal/modules -type f -name '*.lua' ! -name '*_test.lua' -exec cp {} $(DESTDIR)$(datadir)/$(NAME)/meta \;
	$(FIND) contrib/meta -type f -name '*.lua' -exec cp {} $(DESTDIR)$(datadir)/$(NAME)/meta \;
	$(CP) -rf contrib/lib/* $(DESTDIR)$(datadir)/$(NAME)/lib

.PHONY: install-strip
## Strip executable(s) while installing them. For GO programs this will also recompile with LDGLAGS+="-w"
install-strip:
	$(MAKE) install GOLDFLAGS="$(GOLDFLAGS) -w"
	$(INSTALL) --strip -D $(builddir)/$(NAME)-$(os)-$(arch) $(DESTDIR)$(bindir)/$(NAME)

.PHONY: uninstall
## Remove all the installed files.
uninstall:
	$(RM) $(DESTDIR)$(bindir)/$(NAME)
	$(RM) -r $(DESTDIR)$(datadir)/$(NAME)

.PHONY: test
## Run the entire test suite
test: test-go test-lua

.PHONY: test-go
test-go:
	$(GO) test $(test_opts) ./...

.PHONY: test-lua
test-lua: $(LUA_TEST_FILES)

.PHONY: $(LUA_TEST_FILES)
$(LUA_TEST_FILES):
	LUA_PATH="${ROOT_DIR}/contrib/lib/?/?.lua" $(CRAFTY) run $@

.PHONY: fmt
## Code formating
fmt:
	gofumpt -l -w .
	$(SHELL) -c 'test -z "$$(gofmt -l .)"'

.PHONY: lint
## Code linting
lint:
ifndef CI
	@# disabled because in a CI `golangci/golangci-lint-action` will be used
	golangci-lint run $(FIX)
endif
	$(SHELL) -c 'test -z "$$(gofmt -l .)"'
ifndef WITHOUT_STATICCHECK
	$(GO) run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-ST1003,-U1000 ./...
endif

.PHONY: tidy
## Code tidy
tidy:
	$(GO) mod tidy -v

.PHONY: audit
## Quality checks
audit:
	$(GO) mod verify
	$(GO) vet ./...
	$(GO) run golang.org/x/vuln/cmd/govulncheck@latest ./...

## Build manual(s)
man:
	@echo "TODO"
