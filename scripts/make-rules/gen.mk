# Copyright 2024 slw <150657601@qq.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file.

# ==============================================================================
# Makefile helper functions for generate necessary files
#

# generate 
ifeq ($(GOOS), darwin)
	wireCmd=xargs -I F sh -c 'cd F && echo && wire'
else
	wireCmd=xargs -i sh -c 'cd {} && echo && wire'
endif


.PHONY: gen.run
#gen.run: gen.errcode gen.docgo
gen.run:  gen.wire gen.pb gen.clean gen.errcode gen.docgo.doc

.PHONY: gen.pb
gen.pb: tools.verify.buf
	@echo "===========> Generating pb files *.go from proto file through buf.build"
	@${ROOT_DIR}/scripts/buf.sh

.PHONY: gen.wire
gen.wire: tools.verify.wire
	@echo "===========> Generating wire_gen.go from wire.go file through wire"
	@find internal  -mindepth 2 -maxdepth 2 | grep server | $(wireCmd)

.PHONY: gen.errcode
gen.errcode: gen.errcode.code gen.errcode.doc

.PHONY: gen.errcode.code
gen.errcode.code: tools.verify.codegen
	@echo "===========> Generating iam error code go source files"
	@codegen -type=int ${ROOT_DIR}/pkg/code

.PHONY: gen.errcode.doc
gen.errcode.doc: tools.verify.codegen
	@echo "===========> Generating error code markdown documentation"
	@codegen -type=int -doc \
		-output ${ROOT_DIR}/docs/guide/zh-CN/api/error_code_generated.md ${ROOT_DIR}/pkg/code

.PHONY: gen.docgo.doc
gen.docgo.doc:
	@echo "===========> Generating doc.go for go packages"
	@${ROOT_DIR}/scripts/gendoc.sh

.PHONY: gen.docgo.check
gen.docgo.check: gen.docgo.doc
	@n="$$(git ls-files --others '*/doc.go' | wc -l)"; \
	if test "$$n" -gt 0; then \
		git ls-files --others '*/doc.go' | sed -e 's/^/  /'; \
		echo "$@: untracked doc.go file(s) exist in working directory" >&2 ; \
		false ; \
	fi

.PHONY: gen.docgo.add
gen.docgo.add:
	@git ls-files --others '*/doc.go' | $(XARGS) -- git add

.PHONY: gen.defaultconfigs
gen.defaultconfigs:
	@${ROOT_DIR}/scripts/gen_default_config.sh

.PHONY: gen.clean
gen.clean:
	@rm -rf ./api/client/{clientset,informers,listers}
	@$(FIND) -type f -name '*_generated.go' -delete
