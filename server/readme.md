https://tonybai.com/2015/04/30/go-and-https/

cd /Users/houbin/go/src/https-test/server
openssl genrsa -out server.key 2048
openssl req -new -x509 -key server.key -out server.crt -days 365
openssl rsa -in server.key -out server.key.public
go run main.go