[中文版](readme.zh.md)
  
#  go-program-framework
  
  
![go-log](/doc/assets/go-logo-black.png)
The `go-program-framework` is established by referring to the engineering architecture of `GO` source code and my development experience. You can use it as a structural reference for your development of `GO` projects, and I will constantly improve the toolkit and project:fist::fist:
:star::star::star::star::star:
  
  
- [go-program-framework](#go-program-framework )
  - [ProgramArch](#programarch )
  - [ProgramRename](#programrename )
  - [Makefile](#makefile )
    - [make godoc](#make-godoc )
  - [DockerFile](#dockerfile )
  - [pkg pakage](#pkg-pakage )
  
##  ProgramArch
  
  


```
├── bin	                    // compiled executable binary	
├── configs                // config file
├── doc                    // document for `godoc` view
├── docker                 // DockerFile
├── internal               // private file
├── logs                   // log file
├── scripts                // script files, such as py,js.sh,bat...
└── src                    // go source code
    ├── app                // app for assembly server
    │   └── ops            // diff app version 
    ├── cmd                // go build fold
    │   ├── client         // client
    │   └── server         // server
    ├── internal           // private package，can't improt and view by`godoc`
    └── pkg                // tool package
        ├── app            // define `application` interface
        │   └── testdata   // testdata
        ├── config         // read config file
        ├── log            // create log entity
        └── utils          // base utils
```

  
##  ProgramRename
  
  
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
  godoc          start godoc server, 127.0.0.1            6060
  wire           go wire
  di-clean       docker image clean
  d-clean        docker clean
  d-build        docker build
  md             rename readme

```

  
###  make godoc 
  
  
`make godoc` run local http server,`http://127.0.0.1:6060`. you can browse urls like.
  
```go
http://127.0.0.1:6060/pkg/   // 依赖的包名, ./src/internal 不可访问
http://127.0.0.1:6060/doc/   // 就是 ./doc/
```
  
##  DockerFile
  
  
In order to drop the dependency in the container, the local `.ssh` file can be copied to the container for construction. However, after the construction, the image capacity is too large and there is code leakage. We can use multi-stage deployment to solve these problems
  
##  pkg pakage
  
  
`pkg` package defines the `server` interface. As long as your expansion package implements these interfaces, you can register unlimited services, use `wire` to automate dependency injection, and smoothly exit the application after receiving the exit signal
  
```go
// Server define register server interface
type Server interface {
	Desc() string
	Start() error
	Stop() error
}
```
[Application_Example](./src/pkg/app/example_app_test.go )
  
==Note that you are required to provide the basic configuration file as follows==
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
  
__welcome to the issues :fist::fist::fist:__
  