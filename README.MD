# GoFrame Template For SingleRepo

Project Makefile Commands: 
- `make cli`: Install or Update to the latest GoFrame CLI tool.
- `make dao`: Generate go files for `Entity/DAO/DO` according to the configuration file from `hack` folder.
- `make service`: Parse `logic` folder to generate interface go files into `service` folder.
- `make image TAG=xxx`: Run `docker build` to build image according `manifest/docker`.
- `make image.push TAG=xxx`: Run `docker build` and `docker push` to build and push image according `manifest/docker`.
- `make deploy TAG=xxx`: Run `kustomize build` to build and deploy deployment to kubernetes server group according `manifest/deploy`.


FROM golang:alpine
# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \ CGO_ENABLED=0 \ GOOS=linux \ GOARCH=amd64
# 移动到工作目录：/buildWORKDIR /build
# 将代码复制到容器中COPY . .
# 将我们的代码编译成二进制可执行文件appRUN go build -o app .# 移动到用于存放生成的二进制文件的 /dist 目录WORKDIR /dist# 将二进制文件从 /build 目录复制到这里RUN cp /build/app .# 声明服务端口EXPOSE 8888# 启动容器时运行的命令CMD ["/dist/app"]