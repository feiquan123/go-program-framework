[Englist](readme.md)
  
#  go-program-framework
  
  
![go-log](/doc/assets/go-logo-black.png ) 
:star::star::star::star::star:
`go-program-framework` 是参考 `go` 源码的工程架构以及本人的开发经验建立的,你可以作为你开发`GO`项目的一个结构参考,我也会不断完善工具包和项目的:fist::fist:
  
  
- [go-program-framework](#go-program-framework )
  - [工程结构](#工程结构 )
  - [重命名项目名](#重命名项目名 )
  - [Makefile](#makefile )
    - [make godoc](#make-godoc )
  - [DockerFile](#dockerfile )
  - [pkg 工具包的扩展](#pkg-工具包的扩展 )
  
##  工程结构
  
  


```
├── bin	                    //编译后的可执行文件		
├── configs                // 配置文件
├── doc                    // 文档,可用于`godoc`浏览
├── docker                 // DockerFile
├── internal               // 私人文件
├── logs                   // 日志
├── scripts                // 脚本 如:py,js.sh,bat...
└── src                    // go 的源码
    ├── app                // app 是组装服务
    │   └── ops            // 不同的app版本 
    ├── cmd                // go build  的时的直接目录
    │   ├── client         // 客户端
    │   └── server         // 服务端
    ├── internal           // go 的私有包，外部不可导入,并且`godoc`不可查看
    └── pkg                // 工具包
        ├── app            // 定义`application`的接口
        │   └── testdata   // 测试数据
        ├── config         // 读取配置文件
        ├── log            // 创建日志对象
        └── utils          // 基础工具包
```

  
##  重命名项目名
  
  
```sh
./scripts/rename.sh your-appname
```
  
##  Makefile
  
  
```sh{code_chunk_offset=1,
make help
```

```

 Choose a command run in go-program-framework:

  compile        Compile the binary of server and client
  build-server   build server
  build-client   build client
  build          build client and server
  clean-server   clean server
  clean-client   clean client
  clean          clean server and client
  start-server   start server
  start-client   start client
  godoc          start godoc server
  wire           go wire
  di-clean       docker image clean
  d-clean        docker clean
  d-build        docker build
  md             rename readme

```

  
###  make godoc 
  
  
`make godoc` 将会在本地运行http server,`http://127.0.0.1:6060`. 注意此时你可以访问到
  
```go
http://127.0.0.1:6060/pkg/   // 依赖的包名, ./src/internal 不可访问
http://127.0.0.1:6060/doc/   // 就是 ./doc/
```
  
##  DockerFile
  
  
为了可以在容器中下拉依赖,可以将本地的`.ssh`文件复制到容器中进行构建.但是这样构建完后镜像容量过大且存在代码泄露问题,我们可以使用多阶段部署来解决这些问题
  
##  pkg 工具包的扩展
  
  
pkg 包中定义了`Server`接口,只要你的扩展包实现了这些接口，那你就可以无限的注册服务，使用`wire`来自动化进行依赖注入，并在收到退出信号后平滑退出应用
  
```go
// Server define register server interface
type Server interface {
	Desc() string
	Start() error
	Stop() error
}
```
[Application_Example](./src/pkg/app/example_app_test.go )
  
注意,需要你提供基本的配置文件内容如下
```yaml
app:
  name: your-app-name
  
log: 
  filename: /tmp/your-app-name.log
  maxSize: 500
  maxBackups: 3
  maxAge: 3
  level: debug
  stdout: true
```  
  
__限于本人能力有限,欢迎各位的issues :fist:__
  