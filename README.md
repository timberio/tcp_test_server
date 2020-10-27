# TCP Test Server

A simple TCP server useful for testing. Has the option to write all messages
to a file.

This is used in [Vector]'s [test harness] to test and benchmark TCP performance.

## Getting started

1. Run `go build`
2. Run `./tcp_test_server --address=0.0.0.0:9000`

### Docker image

The HTTP test server is also available as a Docker image:

```bash
docker pull timberiodev/tcp_test_server:latest
```

[test harness]: https://github.com/timberio/vector-test-harness
[Vector]: https://github.com/timberio/vector