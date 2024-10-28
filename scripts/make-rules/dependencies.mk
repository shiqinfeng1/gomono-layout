# Copyright 2024 slw <150657601@qq.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file.

# ==============================================================================
# Makefile helper functions for dependencies
#

.PHONY: dependencies.run
dependencies.run: dependencies.packages dependencies.tools

.PHONY: dependencies.packages
dependencies.packages:
# 在tidy之前先安装goconvey，否则会报错，原因见： https://github.com/smarty/assertions/issues/56
	@$(GO) get -u github.com/smartystreets/goconvey  
	@$(GO) mod tidy
	@$(GO) mod vendor

.PHONY: dependencies.tools
dependencies.tools: dependencies.tools.blocker dependencies.tools.critical

.PHONY: dependencies.tools.blocker
dependencies.tools.blocker: go.build.verify $(addprefix tools.verify., $(BLOCKER_TOOLS))

.PHONY: dependencies.tools.critical
dependencies.tools.critical: $(addprefix tools.verify., $(CRITICAL_TOOLS))

.PHONY: dependencies.tools.trivial
dependencies.tools.trivial: $(addprefix tools.verify., $(TRIVIAL_TOOLS))
