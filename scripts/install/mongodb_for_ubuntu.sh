#!/bin/bash

# Copyright 2024 slw <150657601@qq.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file.


# The root of the build/dist directory
IAM_ROOT=$(dirname "${BASH_SOURCE[0]}")/../..
[[ -z ${COMMON_SOURCED} ]] && source ${IAM_ROOT}/scripts/install/common.sh

# 安装后打印必要的信息
function mongodb::info() {
cat << EOF
MongoDB Login: mongo mongodb://${MONGO_USERNAME}:'${MONGO_PASSWORD}'@${MONGO_HOST}:${MONGO_PORT}/iam_analytics?authSource=iam_analytics
EOF
}

# 安装
function mongodb::install()
{
  # 1. 配置 MongoDB Apt 源
  echo ${LINUX_PASSWORD} | sudo -S apt-get install gnupg
  echo ${LINUX_PASSWORD} | sudo -S wget -qO - https://www.mongodb.org/static/pgp/server-4.4.asc | sudo apt-key add -
  echo ${LINUX_PASSWORD} | sudo -S echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu bionic/mongodb-org/4.4 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-4.4.list

  # 2. 安装 MongoDB 和 MongoDB 客户端
  common::sudo "apt-get update"
  common::sudo "apt -y install mongodb-org"

	# 3. 开启外网访问权限和登录验证
	echo ${LINUX_PASSWORD} | sudo -S sed -i '/bindIp/{s/127.0.0.1/0.0.0.0/}' /etc/mongod.conf
	echo ${LINUX_PASSWORD} | sudo -S sed -i '/^#security/a\security:\n  authorization: enabled' /etc/mongod.conf

  # 4. 启动 MongoDB，并设置开机启动
  common::sudo "systemctl enable mongod"
  common::sudo "systemctl start mongod"

  # 5. 创建管理员账号，设置管理员密码
	echo ${LINUX_PASSWORD} | sudo -S mongo --quiet "mongodb://${MONGO_HOST}:${MONGO_PORT}" <<EOF
use admin
db.createUser({user:"${MONGO_ADMIN_USERNAME}",pwd:"${MONGO_ADMIN_PASSWORD}",roles:["root"]})
db.auth("${MONGO_ADMIN_USERNAME}", "${MONGO_ADMIN_PASSWORD}")
EOF

	# 6. 创建 ${MONGO_USERNAME} 用户
	echo ${LINUX_PASSWORD} | sudo -S mongo --quiet mongodb://${MONGO_ADMIN_USERNAME}:${MONGO_ADMIN_PASSWORD}@${MONGO_HOST}:${MONGO_PORT}/iam_analytics?authSource=admin << EOF
use iam_analytics
db.createUser({user:"${MONGO_USERNAME}",pwd:"${MONGO_PASSWORD}",roles:["dbOwner"]})
db.auth("${MONGO_USERNAME}", "${MONGO_PASSWORD}")
EOF

  mongodb::status || return 1
  mongodb::info
  log::info "install MongoDB successfully"
}

# 卸载
function mongodb::uninstall()
{
  set +o errexit
  common::sudo "systemctl stop mongodb"
  common::sudo "systemctl disable mongodb"
  common::sudo "apt-get -y remove mongodb-org"
  common::sudo "rm -rf /var/lib/mongo"
  common::sudo "rm -f /etc/apt/sources.list.d/mongodb-org-4.4.list"
  common::sudo "rm -f /etc/mongod.conf"
  common::sudo "rm -f /lib/systemd/system/mongod.service"
  common::sudo "rm -f /tmp/mongodb-*.sock"
  set -o errexit
  log::info "uninstall MongoDB successfully"
}

# 状态检查
function mongodb::status()
{
  # 查看 mongodb 运行状态，如果输出中包含 active (running) 字样说明 mongodb 成功启动。
  systemctl status mongod |grep -q 'active' || {
    log::error "mongodb failed to start, maybe not installed properly"
    return 1
  }

	echo "show dbs" | mongo --quiet "mongodb://${MONGO_HOST}:${MONGO_PORT}" &>/dev/null || {
    log::error "cannot connect to mongodb, mongo maybe not installed properly"
    return 1
  }

	echo "show dbs" | \
		mongo --quiet mongodb://${MONGO_ADMIN_USERNAME}:${MONGO_ADMIN_PASSWORD}@${MONGO_HOST}:${MONGO_PORT}/iam_analytics?authSource=admin &>/dev/null || {
    log::error "can not login with ${MONGO_ADMIN_USERNAME}, mongo maybe not initialized properly"
    return 1
  }

	echo "show dbs" | \
		mongo --quiet mongodb://${MONGO_USERNAME}:${MONGO_PASSWORD}@${MONGO_HOST}:${MONGO_PORT}/iam_analytics?authSource=iam_analytics &>/dev/null|| {
    log::error "can not login with ${MONGO_USERNAME}, mongo maybe not initialized properly"
    return 1
  }
}

if [[ "$*" =~ mongodb:: ]];then
  eval $*
fi
