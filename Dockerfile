FROM golang:1.17-alpine as builder
WORKDIR /usr/src
COPY go.mod .
COPY go.sum .
RUN GOPROXY=${PROXY} go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o pomodoro .

FROM alpine
WORKDIR /usr/app
COPY --from=builder /usr/src/pomodoro .
CMD ["./pomodoro"]