# password-manager
Creating a password manager in Golang

Start up local webserver with `go run main.go`

Navigate to https://localhost:8080/repo/

Note: cert.pem and key.pem are self generated certifactes to work with TLS

Run `go run $GOROOT/src/crypto/tls/generate_cert.go -host localhost` to create cert.pem and key.pem
