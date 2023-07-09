FROM golang:1.19  
  
WORKDIR /app  
RUN go install github.com/cosmtrek/air@latest  
RUN go install github.com/go-delve/delve/cmd/dlv@latest  
RUN go install golang.org/x/tools/gopls@latest
RUN go get -u github.com/go-chi/chi/v5

# Run Hot Reload
CMD ["air", "-c", ".air.toml"]