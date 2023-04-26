FROM golang:1.19-alpine
 
# 创建文件夹
RUN mkdir /blog
 
# 设置工作目录
WORKDIR /blog


ENV GO111MODULE=on
ENV CGO_ENABLED=0

ENV GOPROXY="https://goproxy.cn"
 
#将Dockerfile 中的文件存储到/app下
ADD . /blog
 
# 因为已经是在 /app下了，所以使用  ./

RUN go build -o main /goblog/main.go

# 暴露的端口
EXPOSE 3000
 
#设置容器的启动命令，CMD是设置容器的启动指令
CMD /blog/main
 