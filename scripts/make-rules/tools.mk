# Copyright 2024 slw <150657601@qq.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file.

# ==============================================================================
# Makefile helper functions for tools
#

TOOLS ?=$(BLOCKER_TOOLS) $(CRITICAL_TOOLS) $(TRIVIAL_TOOLS)

.PHONY: tools.install
tools.install: $(addprefix tools.install., $(TOOLS))

.PHONY: tools.install.%
tools.install.%:
	@echo "===========> Installing $*"
	@$(MAKE) install.$*

.PHONY: tools.verify.%
tools.verify.%:
	@if ! which $* &>/dev/null; then $(MAKE) tools.install.$*; fi

.PHONY: install.swagger
install.swagger:
	@$(GO) install github.com/go-swagger/go-swagger/cmd/swagger@latest

.PHONY: install.golangci-lint
install.golangci-lint:
	@$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2
	@golangci-lint completion bash > $(HOME)/.golangci-lint.bash
	@if ! grep -q .golangci-lint.bash $(HOME)/.bashrc; then echo "source \$$HOME/.golangci-lint.bash" >> $(HOME)/.bashrc; fi

.PHONY: install.go-junit-report
install.go-junit-report:
	@$(GO) install github.com/jstemmer/go-junit-report@latest

.PHONY: install.gsemver
install.gsemver:
	@$(GO) install github.com/arnaud-deprez/gsemver@latest

.PHONY: install.git-chglog
install.git-chglog:
	@$(GO) install github.com/git-chglog/git-chglog/cmd/git-chglog@latest

.PHONY: install.github-release
install.github-release:
	@$(GO) install github.com/github-release/github-release@latest

.PHONY: install.coscmd
install.coscmd:
	@if which pip &>/dev/null; then pip install coscmd; else pip3 install coscmd; fi

.PHONY: install.golines
install.golines:
	@$(GO) install github.com/segmentio/golines@latest

.PHONY: install.go-mod-outdated
install.go-mod-outdated:
	@$(GO) install github.com/psampaz/go-mod-outdated@latest

.PHONY: install.mockgen
install.mockgen:
	@$(GO) install github.com/golang/mock/mockgen@latest

.PHONY: install.gotests
install.gotests:
	@$(GO) install github.com/cweill/gotests/gotests@latest

.PHONY: install.protoc-gen-go
install.protoc-gen-go:
	@$(GO) install github.com/golang/protobuf/protoc-gen-go@latest

.PHONY: install.protoc-gen-go-grpc
install.protoc-gen-go-grpc:
	@$(GO) install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: install.protoc-gen-go-http
install.protoc-gen-go-http:
	@$(GO) install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest

.PHONY: install.protoc-gen-go-errors
install.protoc-gen-go-errors:
	@$(GO) install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest

.PHONY: install.protoc-gen-validate
install.protoc-gen-validate:
	@$(GO) install github.com/envoyproxy/protoc-gen-validate@latest

.PHONY: install.protoc-gen-openapi
install.protoc-gen-openapi:
	@$(GO) install github.com/google/gnostic/cmd/protoc-gen-openapi@latest

.PHONY: install.wire
install.wire:
	@$(GO) install github.com/google/wire/cmd/wire@latest

.PHONY: install.addlicense
install.addlicense:
	@$(GO) install github.com/marmotedu/addlicense@latest

.PHONY: install.goimports
install.goimports:
	@$(GO) install golang.org/x/tools/cmd/goimports@latest

.PHONY: install.depth
install.depth:
	@$(GO) install github.com/KyleBanks/depth/cmd/depth@latest

.PHONY: install.go-callvis
install.go-callvis:
	@$(GO) install github.com/ofabry/go-callvis@latest

.PHONY: install.gothanks
install.gothanks:
	@$(GO) install github.com/psampaz/gothanks@latest

.PHONY: install.richgo
install.richgo:
	@$(GO) install github.com/kyoh86/richgo@latest

.PHONY: install.rts
install.rts:
	@$(GO) install github.com/galeone/rts/cmd/rts@latest

.PHONY: install.codegen
install.codegen:
	@$(GO) install ${ROOT_DIR}/tools/codegen/codegen.go

.PHONY: install.kube-score
install.kube-score:
	@$(GO) install github.com/zegl/kube-score/cmd/kube-score@latest

.PHONY: install.go-gitlint
install.go-gitlint:
	@$(GO) install github.com/marmotedu/go-gitlint/cmd/go-gitlint@latest
