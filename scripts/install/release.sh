#!/usr/bin/env bash

# Copyright 2024 slw <150657601@qq.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file.

# Build an IAM release.  This will build the binaries, create the Docker
# images and other build artifacts.

set -o errexit
set -o nounset
set -o pipefail

IAM_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
source "${IAM_ROOT}/scripts/common.sh"
source "${IAM_ROOT}/scripts/lib/release.sh"

IAM_RELEASE_RUN_TESTS=${IAM_RELEASE_RUN_TESTS-y}

golang::setup_env
build::verify_prereqs
release::verify_prereqs
#build::build_image
build::build_command
release::package_tarballs
release::updload_tarballs
release::github_release
release::generate_changelog
