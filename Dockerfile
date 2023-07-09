FROM golang:1.19  
  
WORKDIR /app
RUN go install github.com/cweill/gotests/gotests@v1.6.0
RUN go install github.com/fatih/gomodifytags@v1.16.0
RUN go install github.com/josharian/impl@v1.1.0
RUN go install github.com/haya14busa/goplay/cmd/goplay@v1.0.0
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
RUN go install golang.org/x/tools/gopls@latest
RUN go install github.com/ramya-rao-a/go-outline@v0.0.0-20210608161538-9736a4bde949
COPY go.mod go.sum ./
RUN go mod download

# Run Hot Reload
CMD ["air", "-c", ".air.toml"]