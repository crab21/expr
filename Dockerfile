# Build the manager binary
FROM golang:1.16.6-alpine as builder

WORKDIR /workspace

# Copy the go source
COPY . .

# Build
RUN  export GOPROXY="https://goproxy.io" && cd cal && go mod vendor&& cd ../ && cd ast && go mod vendor&& cd ../ && go mod vendor \
     && CGO_ENABLED=0  GOOS=linux GOARCH=amd64 GO111MODULE=on go build -ldflags "-s -w" -a -o expr main.go \
     && apk add --no-cache upx ca-certificates tzdata \
     && upx --best expr -o upx_server  \
     && mv -f upx_server expr

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM busybox
WORKDIR /
COPY --from=builder /workspace/expr .