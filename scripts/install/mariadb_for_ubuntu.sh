#!/bin/bash

# Copyright 2024 slw <150657601@qq.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file.


# The root of the build/dist directory
IAM_ROOT=$(dirname "${BASH_SOURCE[0]}")/../..

[[ -z ${COMMON_SOURCED} ]] && source ${IAM_ROOT}/scripts/install/common.sh

# 安装后打印必要的信息
function mariadb::info() {
cat << EOF
MariaDB Login: mysql -h127.0.0.1 -u${MARIADB_ADMIN_USERNAME} -p'${MARIADB_ADMIN_PASSWORD}'
EOF
}

# 安装
function mariadb::install()
{
  # 1. 配置 MariaDB 10.5 apt 源
  common::sudo "apt-get install software-properties-common dirmngr apt-transport-https"
  echo ${LINUX_PASSWORD} | sudo -S apt-key adv --fetch-keys 'https://mariadb.org/mariadb_release_signing_key.asc'
  # add /etc/apt/sources.list
  echo ${LINUX_PASSWORD} | sudo -S add-apt-repository 'deb [arch=amd64,arm64,ppc64el,s390x] https://mirrors.aliyun.com/mariadb/repo/10.5/ubuntu focal main'

  # 2. 安装 MariaDB 和 MariaDB 客户端
  common::sudo "apt update"
  common::sudo "apt -y install mariadb-server"

  # 3. 启动 MariaDB，并设置开机启动
  common::sudo "systemctl enable mariadb"
  common::sudo "systemctl start mariadb"

  # 4. 设置 root 初始密码
  common::sudo "mysqladmin -u${MARIADB_ADMIN_USERNAME} password ${MARIADB_ADMIN_PASSWORD}"

  mariadb::status || return 1
  mariadb::info
  log::info "install MariaDB successfully"
}

# 卸载
function mariadb::uninstall()
{
  set +o errexit
  common::sudo "systemctl stop mariadb"
  common::sudo "systemctl disable mariadb"
  common::sudo "apt-get -y remove mariadb-server"
  common::sudo "rm -rf /var/lib/mysql"
  set -o errexit
  log::info "uninstall MariaDB successfully"
}

# 状态检查
function mariadb::status()
{
  # 查看 mariadb 运行状态，如果输出中包含 active (running) 字样说明 mariadb 成功启动。
  systemctl status mariadb |grep -q 'active' || {
    log::error "mariadb failed to start, maybe not installed properly"
    return 1
  }

  mysql -u${MARIADB_ADMIN_USERNAME} -p${MARIADB_ADMIN_PASSWORD} -e quit &>/dev/null || {
    log::error "can not login with root, mariadb maybe not initialized properly"
    return 1
  }
  log::info "MariaDB status active"
}

if [[ "$*" =~ mariadb:: ]];then
  eval $*
fi
