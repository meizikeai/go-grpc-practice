# go-grpc-practice

### 项目构架

基于[grpc](google.golang.org/grpc)构建，支持[mysql](https://github.com/go-sql-driver/mysql)、[redis](https://github.com/go-redis/redis)查询，依赖[golang](https://golang.google.cn/dl)环境。

#### 项目地址

https://gitlab-vywrajy.micoworld.net/ai-lab/blue-flame/go-grpc-practice.git

#### 项目结构

| 路径        | 描述          | 详情 |
|-------------|---------------|------|
| config      | config        | --   |
| controllers | controller    | --   |
| libs        | lib           | --   |
| models      | mysql / redis | --   |
| protos      | protobuf      | --   |
| routers     | router        | --   |
| test        | test          | --   |
| go.mod      | go modules    | --   |

#### 开发环境

  + 克隆项目 - `$ git clone https://gitlab-vywrajy.micoworld.net/ai-lab/blue-flame/go-grpc-practice.git`
  + 安装依赖 - `$ cd go-grpc-practice && go mod tidy`
  + 启动项目 - `$ go run .`

推荐[Visual Studio Code](https://code.visualstudio.com)编辑器，开发依赖[Tools](https://github.com/golang/vscode-go/blob/master/docs/tools.md)，请正确安装好后再开始。

#### 环境变量

在测试、线上环境需要设置 `GGP_MODE` 的值来区分相应环境。

需要在 `～/.zshrc` 里给定 `export GGP_MODE=release/test` 后，程序中通过 `mode := os.Getenv("GGP_MODE")` 获取。

如果 `GGP_MODE` 未设定，那么运行的是 `test` 环境。

为确保正确写日志，本地开发添加 `export GO_ENV=debug` 环境变量。

#### 访问地址

  + http://127.0.0.1:9527
