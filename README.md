# gomono-layout -基于kratos+iam+整洁架构的企业级go微服务脚手架

基于大仓模式的微服务框架，集成常用组件和规范化、自动化代码生成。支持直接部署及容器化k8s部署

## 功能特性

本项目用到了Go企业开发的大部分核心技能点，见下图：

![技术思维导图](./docs/images/技术思维导图.png)

更多请参考：[shiqinfeng1/gocollect](https://github.com/shiqinfeng1/gocollect)

## 软件架构

![gomono-layout架构](./docs/images/gomono-layout架构.png)

架构解析见：[gomono-layout 架构 & 能力说明](./docs/guide/zh-CN/installation/installation-architecture.md)

## 快速开始

### 依赖检查

1. 服务器能访问外网

2. 操作系统：CentOS Linux 8.x (64-bit) / ubuntu Linux 22.04 (64-bit)

### 快速部署

快速部署请参考：[gomono-layout 部署指南](docs/guide/zh-CN/installation/README.md#部署指南)

> gomono-layout 项目还提供了更详细的部署文档，请参考：[手把手教你部署gomono-layout系统](docs/guide/zh-CN/installation/installation-procedures.md)

### 克隆当前模版

1. 安装kratos工具

    ```bash
    go install github.com/shiqinfeng1/gomono-layout/cmd/gomonoctl@latest
    ```

2. 基于当前模版创建一个项目仓库

    创建项目名称为helloworld，服务名称为user的项目：

    ```bash
    $gomonoctl newservice helloworld --service user  
    ```

3. 在项目名称为helloworld仓库中添加一个新的服务

    ```bash
    $gomonoctl newservice helloworld --service order  
    ```

### 构建

如果你需要重新编译gomono-layout项目，可以执行以下 2 步：

1. 首次使用时，先安装工具
    
    ```bash
    cd gomono-layout
    make tools
    ```
    
1. 编译

    ```bash
    cd gomono-layout
    make all
    ```

构建后的二进制文件保存在 `_output/platforms/linux/amd64/` 目录下。

## 使用指南

[Documentation](docs/guide/zh-CN)

## 如何贡献

欢迎贡献代码，贡献流程可以参考 [developer's documentation](docs/devel/zh-CN/development.md)。

## 社区

You are encouraged to communicate most things via [GitHub issues](https://github.com/shiqinfeng1/gomono-layout/issues/new/choose) or pull requests.

## 关于作者

## 谁在用

## 许可证

gomono-layout is licensed under the MIT. See [LICENSE](LICENSE) for the full license text.
