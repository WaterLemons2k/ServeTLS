# Serve TLS

Simple serve HTTPS/TLS via Go.

English | [简体中文](README.zh-CN.md)

## Usage

```sh
$ ServeTLS
2009/11/10 23:00:00 Listening on :443

$ ServeTLS -h
Usage of ServeTLS:
  -addr string
        Address to listen on (default ":443")
```

## Generate key

```sh
# Generate private key (.key)
# Key considerations for algorithm "ECDSA" (X25519 || ≥ secp384r1)
# https://safecurves.cr.yp.to/
# List ECDSA the supported curves (openssl ecparam -list_curves)
openssl ecparam -genkey -name secp384r1 -out server.key

# Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```

## Reference Link
- https://github.com/denji/golang-tls
- https://pkg.go.dev/crypto/tls#example-X509KeyPair-HttpServer