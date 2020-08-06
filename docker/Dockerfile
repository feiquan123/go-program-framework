# STEP 1 build executable binary
FROM golang:alpine as builder

# env
COPY ./internal/.ssh /root/.ssh
RUN echo "http://mirrors.aliyun.com/alpine/latest-stable/main/" > /etc/apk/repositories &&\
	echo "http://mirrors.aliyun.com/alpine/latest-stable/community/" >> /etc/apk/repositories &&\
	apk add make git openssh &&\
	sed -i "s/#PermitRootLogin.*/PermitRootLogin yes/g" /etc/ssh/sshd_config &&\
	chmod 600 /root/.ssh/gitlab/fengzhiquan &&\
	mkdir -p $GOPATH/src/github.com/feiquan123/go-program-framework &&\
	go env -w GOPROXY="https://goproxy.cn,direct" GO111MODULE=on GOSUMDB=off

COPY . $GOPATH/src/github.com/feiquan123/go-program-framework
WORKDIR $GOPATH/src/github.com/feiquan123/go-program-framework
# build the binary
RUN pwd &&\
	ls -lh . &&\
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 make build

# STEP 2 build a small image
# start from builder
FROM alpine:latest as production
MAINTAINER feiquanxx@gmail.com
RUN mkdir /lib64 &&\
	ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2 &&\
	echo "http://mirrors.aliyun.com/alpine/latest-stable/main/" > /etc/apk/repositories &&\
	echo "http://mirrors.aliyun.com/alpine/latest-stable/community/" >> /etc/apk/repositories &&\
	apk add tzdata &&\
	ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime &&\
	echo "Asia/Shanghai" > /etc/timezone &&\
	ifconfig &&\
	mkdir -p /app/ /config/ /applogs/your-appname/  /appdata/your-appname/

# Copy our startic executable
COPY --from=builder /go/src/github.com/feiquan123/go-program-framework/bin/* /app/
COPY --from=builder /go/src/github.com/feiquan123/go-program-framework/configs/* /config/
RUN pwd & ls -lh /app/ /config/
ENTRYPOINT ["/app/server","-c","/config/app.yaml"]
