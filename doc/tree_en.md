├── bin	 #// compiled executable binary	
├── configs #// config file
├── doc  #// document for `godoc` view
├── docker #// DockerFile
├── internal #// private file
├── logs #// log file
├── scripts #// script files, such as py,js.sh,bat...
└── src #// go source code
    ├── app #// app for assembly server
    │   └── ops #// diff app version 
    ├── cmd #// go build fold
    │   ├── client #// client
    │   └── server #// server
    ├── internal #// private package，can't improt and view by`godoc`
    └── pkg #// tool package
        ├── app #// define `application` interface
        │   └── testdata #// testdata
        ├── config #// read config file
        ├── log #// create log entity
        └── utils #// base utils
