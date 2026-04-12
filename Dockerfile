# 使用官方 Go 镜像
FROM golang:1.25

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum
COPY go.mod go.sum ./
RUN go mod tidy

# 复制项目所有文件
COPY . .

# 编译
RUN go build -o app

# 暴露接口
EXPOSE 8080

# 启动
CMD ["./app"]