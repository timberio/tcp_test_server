# TCP Test Server

A simple TCP server useful for testing. Has the option to write all messages
to a file.

This is used in [Vector]'s [test harness] to test and benchmark TCP performance.

## Getting started

1. Run `dep ensure`
2. Run `go build`
3. Run `./tcp_test_server --address=0.0.0.0:9000`

[test harness]: https://github.com/timberio/vector-test-harness
[Vector]: https://github.com/timberio/vector