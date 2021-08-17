# Build the manager binary
FROM golang:1.16.6 as builder

WORKDIR /workspace

# Copy the go source
COPY . .
# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o expr expr.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM alpine:3.11
WORKDIR /tencent/
COPY --from=builder /workspace/expr .

ENTRYPOINT ["/tencent/expr"]