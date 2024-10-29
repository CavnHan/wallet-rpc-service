<!--
parent:
  order: false
-->

<div align="center">
  <h1> wallet-rpc-service </h1>
</div>

<div align="center">
  <a href="https://github.com/CavnHan/wallet-rpc-service/releases/latest">
    <img alt="Version" src="https://img.shields.io/github/tag/CavnHan/wallet-rpc-service.svg" />
  </a>
  <a href="https://github.com/CavnHan/wallet-rpc-service/blob/main/LICENSE">
    <img alt="License: Apache-2.0" src="https://img.shields.io/github/license/the-web3/rpc-service.svg" />
  </a>
  <a href="https://pkg.go.dev/github.com/CavnHan/wallet-rpc-service/">
    <img alt="GoDoc" src="https://godoc.org/github.com/the-web3/rpc-service?status.svg" />
  </a>
</div>


**Tips**: need [Go 1.22+](https://golang.org/dl/)

## Install

### Install dependencies
```bash
go mod tidy
```
### build
```bash
make 
```

### start
```bash
./rpc-service -c ./config.yml
```

### Start the RPC interface test interface

```bash
grpcui -plaintext 127.0.0.1:8089
```