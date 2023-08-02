# xk6-harbor

This is a Harbor client library for [k6](https://github.com/loadimpact/k6).

## Build

To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. Clone this repostiroy
  ```shell
  git clone https://github.com/goharbor/xk6-harbor
  cd xk6-harbor
  ```

2. Build the binary
  ```shell
  make build
  ```

## Get the binrary
You can also download the latest pre-build binary file.

```shell
curl -sL $(curl -s https://api.github.com/repos/goharbor/xk6-harbor/releases/latest | grep 'http.*linux-amd64.tar.gz"' | awk '{print $2}' | sed 's|[\"\,]*||g') | tar -zx
```
