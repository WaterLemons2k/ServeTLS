# 服务 TLS

使用 Go 简单提供 HTTPS/TLS 服务。

[English](README.md) | 简体中文

## 使用

```sh
$ ServeTLS
2009/11/10 23:00:00 Listening on :443

$ ServeTLS -h
Usage of ServeTLS:
  -addr string
        Address to listen on (default ":443")
```

## 生成密钥

```sh
# 生成私钥（.key）
# 算法 "ECDSA "的密钥注意事项（X25519 || ≥ secp384r1）
# https://safecurves.cr.yp.to/
# 列出 ECDSA 支持的曲线（openssl ecparam -list_curves）
openssl ecparam -genkey -name secp384r1 -out server.key

# 根据私钥（.key）生成自签名（x509）公钥（PEM 编码 .pem|.crt）
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```

## 参考链接
- https://github.com/denji/golang-tls
- https://pkg.go.dev/crypto/tls#example-X509KeyPair-HttpServer