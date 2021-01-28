FROM golang:1
ADD . /greenlog
ENV GOPROXY "https://goproxy.cn"
WORKDIR /greenlog
RUN go get -u github.com/swaggo/swag/cmd/swag && \
    go mod tidy && \
    swag i && \
    go build
CMD ./greenlog
