# Build the manager binary
FROM golang:1.17-alpine as builder

WORKDIR /workspace

# Copy the go source
COPY . .

# Build
RUN  sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
     && export GOPROXY="https://goproxy.io" && cd cal && go mod vendor&& cd ../ && cd ast && go mod vendor&& cd ../ && go mod vendor \
     && CGO_ENABLED=0  GOOS=linux GOARCH=amd64 GO111MODULE=on go build -ldflags "-s -w" -a -o nop main.go \
     && apk add --no-cache upx ca-certificates tzdata \
     && upx --best nop -o upx_server  \
     && mv -f upx_server nop \
     && chmod +x nop

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM busybox
WORKDIR /ko-app/
COPY --from=builder /workspace/nop .
USER root

ENTRYPOINT ["/ko-app/nop"]