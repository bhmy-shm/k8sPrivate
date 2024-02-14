# 使用官方的 Golang 镜像作为构建环境
FROM golang:1.22-alpine3.19 as builder

# Label 设置标签
LABEL stage = apiBuilder

# 设置环境变量
ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# 更新 apk 的索引并安装 tzdata 包，这个包是用来设置时区的
RUN apk update --no-cache && apk add --no-cache tzdata

# 设置打包的工作目录
WORKDIR /build

# 将 go.mod 和 go.sum 复制到容器中
ADD go.mod .
ADD go.sum .

# 下载依赖项（使用 go mod）
RUN go mod download

# 将源代码复制到容器中
COPY . .

# 构建应用程序。
RUN go build -ldflags="-s -w" -installsuffix cgo -o myapp .
RUN pwd
RUN ls -la


# 使用 scratch 作为最小运行时环境
FROM scratch

# 从 builder 阶段复制 SSL 证书，以便容器可以安全地与其他服务通信
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# 设置容器的时区为 Asia/Shanghai
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai

# 设置环境变量 TZ 以指定容器时区
ENV TZ Asia/Shanghai

# 设置工作目录为 /app
WORKDIR /app

# 从构建器镜像中复制构建好的应用程序
COPY --from=builder /build/myapp /app/myapp
COPY --from=builder /build/application.yaml /app/application.yaml

# 运行应用程序
CMD ["./myapp"]