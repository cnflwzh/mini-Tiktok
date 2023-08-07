# 使用官方的Go image作为基础image进行构建
FROM golang:1.20.7 AS builder

# 在image内部设置工作目录
WORKDIR /app

# 把当前目录的内容复制到image的工作目录中
COPY . .

# 设置Go Modules的proxy
ENV GOPROXY=https://goproxy.cn,direct

# 下载依赖项并编译项目，生成可执行文件
# 开启CGO_ENABLED=0来做完全静态编译
RUN CGO_ENABLED=0 go mod download && \
    CGO_ENABLED=0 go build -ldflags '-w -extldflags "-static"' -o /myapp .

################################################

# 使用轻量级的基础镜像
FROM debian:buster-slim

# 创建非root用户
RUN useradd -m myuser

# 切换到非root用户
USER myuser

# 将工作目录设置为该用户的home目录
WORKDIR /home/myuser

# 从builder阶段复制构建的可执行文件
COPY --from=builder /myapp /myapp

# 编译应用之后，复制配置文件
COPY config/config.toml config/config.toml


CMD ["/myapp"]

# 暴露端口
EXPOSE 8888

# 容器启动时运行的命令
CMD ["/myapp"]