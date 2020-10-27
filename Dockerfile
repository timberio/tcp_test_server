FROM golang:1.13 AS builder

WORKDIR /home
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -a -installsuffix nocgo -o /tcp_test_server .

FROM scratch
COPY --from=builder /tcp_test_server ./
ENTRYPOINT ["./tcp_test_server"]
