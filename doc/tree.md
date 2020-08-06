├── bin	 #//编译后的可执行文件		
├── configs #// 配置文件
├── doc  #// 文档,可用于`godoc`浏览
├── docker #// DockerFile
├── internal #// 私人文件
├── logs #// 日志
├── scripts #// 脚本 如:py,js.sh,bat...
└── src #// go 的源码
    ├── app #// app 是组装服务
    │   └── ops #// 不同的app版本 
    ├── cmd #// go build  的时的直接目录
    │   ├── client #// 客户端
    │   └── server #// 服务端
    ├── internal #// go 的私有包，外部不可导入,并且`godoc`不可查看
    └── pkg #// 工具包
        ├── app #// 定义`application`的接口
        │   └── testdata #// 测试数据
        ├── config #// 读取配置文件
        ├── log #// 创建日志对象
        └── utils #// 基础工具包
