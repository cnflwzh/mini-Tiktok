# 使用官方的Go image作为基础image
FROM golang:1.21rc4 AS build

# 在image内部设置工作目录
WORKDIR /app

# 把当前目录的内容复制到image的工作目录中
COPY . .

# 设置Go Modules的proxy
ENV GOPROXY=https://goproxy.cn,direct

# 下载依赖项
RUN go mod download

# 编译项目，生成可执行文件
RUN go build -o main .

# 暴露端口，这个端口应该和你的应用程序实际使用的端口一致
EXPOSE 8888

# 容器启动时运行的命令
CMD ["./main"]
