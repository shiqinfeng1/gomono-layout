# buf.build 使用指南

buf.build是专门用于构建protobuf API的工具

## 安装

```bash
# Substitute BIN for your bin directory.
# Substitute VERSION for the current released version.
BIN="/usr/local/bin" && \
VERSION="1.45.0" && \
curl -sSL \
"https://github.com/bufbuild/buf/releases/download/v${VERSION}/buf-$(uname -s)-$(uname -m)" \
-o "${BIN}/buf" && \
chmod +x "${BIN}/buf"
```

查看最新的[版本](https://github.com/bufbuild/buf/releases)

## 配置

1. 在根目录下运行

    ```bash
    buf config init
    ```

    会生成buf.yaml文件：

    ```yaml
    version: v2

    lint:
    use:
    - STANDARD
    breaking:
    use:
    - FILE
    ```

2. 设置工作区

    把包含proto的目录配置到buf.yaml中：

    ```yaml
    modules:
    - path: api/protos
    ```

3. 设置项目依赖

    添加依赖文件到buf.yaml：

    ```yaml
    deps:
     - 'buf.build/googleapis/googleapis'
     - 'buf.build/envoyproxy/protoc-gen-validate'
     - 'buf.build/kratos/apis'
     - 'buf.build/gnostic/gnostic'
     - 'buf.build/gogo/protobuf'
     - 'buf.build/tx7do/pagination'
    ```

4. 下载依赖

    ```bash
    buf dep update
    ```

## 格式化

可以使用如下命令对 .proto 文件进行代码格式化：

```bash
# 查看格式化后的变更内容
buf format -d
```

```bash
# 将格式化后的变更写入文件
buf format -w
```
