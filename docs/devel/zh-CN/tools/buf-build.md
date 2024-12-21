# buf.build 使用指南

buf.build是专门用于构建protobuf API的工具

## 安装

```bash
# Substitute BIN for your bin directory.
# Substitute VERSION for the current released version.
BIN="/usr/local/bin" && \
VERSION="1.47.2" && \
curl -sSL \
"https://github.com/bufbuild/buf/releases/download/v${VERSION}/buf-$(uname -s)-$(uname -m)" \
-o "${BIN}/buf" && \
chmod +x "${BIN}/buf"
```

查看最新的[版本](https://github.com/bufbuild/buf/releases)

## 配置

以下面的文件目录为例：

```bash
workspace_root
├── buf.yaml
└── proto
    ├── foo
    │   └── foo.proto
    └── bar
        ├── a
        │   └── d.proto
        ├── b
        │   ├── e.proto
        │   └── f.proto
        └── c
            ├── g.proto
            └── h.proto
```

1. 在根目录下运行

    ```bash
    buf config init
    ```

    生成buf.yaml文件：

    ```yaml
    version: v2
    # v2 buf.yaml 文件指定了一个本地工作区，该工作区至少由一个模块组成。
    # buf.yaml 文件应放置在工作区的根目录下，通常也就是源代码控制库的根目录下。
    # 通常就是源代码控制仓库的根目录。
    modules:
    # 每个模块条目都定义了一个路径，该路径必须是相对于
    # buf.yaml 所在的目录。您还可以从模块中指定要排除的目录。   
    - path: proto/foo
        # 模块可以指定BSR(Buf Schema Repository).
        name: buf.build/acme/foo
    - path: proto/bar
      name: buf.build/acme/bar
      # 排除子目录和特定 .proto 文件。请注意，排除
      # 是相对于 buf.yaml 文件的。
      excludes:
        - proto/bar/a
        - proto/bar/b/e.proto
      # 一个模块可以有自己的校验和配置，该配置会覆盖该模块的默认校验和配置。所有来自
      # 默认配置中的所有值都会被覆盖，并且不会合并任何规则。
      lint:
        use:
          - STANDARD
        except:
          - IMPORT_USED
        ignore:
          - proto/bar/c
        ignore_only:
        ENUM_ZERO_VALUE_SUFFIX:
          - proto/bar/a
          - proto/bar/b/f.proto
      # v1 配置中有一个 allow_comment_ignores 选项，可选择忽略评论。
      # 在 v2 版中，我们默认允许忽略评论，并允许使用 disallow_comment_ignores 选项取消忽略评论。
      disallow_comment_ignores: false
      enum_zero_value_suffix: _UNSPECIFIED
      rpc_allow_same_request_response: false
      rpc_allow_google_protobuf_empty_requests: false
      rpc_allow_google_protobuf_empty_responses: false
      service_suffix: Service
      disable_builtin: false
      # Breaking configuration for this module only. Behaves the same as a module-level
      # lint configuration.
      breaking:
      use:
        - FILE
      except:
        - EXTENSION_MESSAGE_NO_DELETE
      ignore_unstable_packages: true
      disable_builtin: false
      # Multiple modules are allowed to have the same path, as long as they don't share '.proto' files.
    - path: proto/common
      module: buf.build/acme/weather
      includes:
        - proto/common/weather
    - path: proto/common
      module: buf.build/acme/location
      includes:
        - proto/common/location
      excludes:
        # Excludes and includes can be specified at the same time, but if they are, each directory
        # in excludes must be contained in a directory in includes.
        - proto/common/location/test
    - path: proto/common
      module: buf.build/acme/other
      excludes:
        - proto/common/location
        - proto/common/weather
    
    # 工作区中所有模块共享的依赖关系。必须是 Buf 模式注册中心BSR托管的模块。
    # 这些依赖关系的解析存储在 buf.lock 文件中。
    deps:
      - buf.build/acme/paymentapis # The latest accepted commit.
      - buf.build/acme/pkg:47b927cbb41c4fdea1292bafadb8976f # The '47b927cbb41c4fdea1292bafadb8976f' commit.
      - buf.build/googleapis/googleapis:v1beta1.1.0 # The 'v1beta1.1.0' label.
    
    # The default lint configuration for any modules that don't have a specific lint configuration.
    #
    # If this section isn't present, AND a module doesn't have a specific lint configuration, the default
    # lint configuration is used for the module.
    lint:
      use:
        - STANDARD
        - TIMESTAMP_SUFFIX # This rule comes from the plugin example below.
    # Default breaking configuration. It behaves the same as the default lint configuration.
    breaking:
      use:
        - FILE
    
    # 可选的 Buf 插件。在本地安装的插件中指定的自定义校验或更改规则。每个 Buf 插件都单独列出，如果插件允许，还可以包含选项
    # 如果某个规则的 `default` 值被设为 true，那么即使'lint' 和 'breaking' 字段未设置。
    #
    # 请参阅 https://github.com/bufbuild/bufplugin-go/blob/main/check/internal/example/cmd/buf-plugin-timestamp-suffix/main.go 上的示例
    # 有关下面示例的更多细节。
    plugins:
    - plugin: plugin-timestamp-suffix # Specifies the installed plugin to use
        options:
        # The TIMESTAMP_SUFFIX rule specified above allows the user to change the suffix by providing a
        # new value here.
        timestamp_suffix: _time
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
     - 'buf.build/gnostic/gnostic'
    ```

4. 下载依赖

    ```bash
    buf dep update
    ```

5. 安装protoc插件

    见 ：`scripts/make-rules/tools.mk` 中`install.protoc*`的命令
    同时在`scripts/make-rules/common.mk`中的`CRITICAL_TOOLS`变量中添加安装的工具，有：
    `protoc-gen-go protoc-gen-go-grpc protoc-gen-go-http protoc-gen-go-errors protoc-gen-validate protoc-gen-openapi`

    在执行`make tools`后，会自动安装这些工具

6. 在`buf.gen.yaml`中配置插件

    插件包括本地已安装的插件和远端插件，根据需要配置

7. 生成go文件

    在根目录下执行`buf genegrate`，在`api/gen`会自动生成go文件

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
