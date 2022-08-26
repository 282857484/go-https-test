https://tonybai.com/2015/04/30/go-and-https/
https://blog.csdn.net/cuichenghd/article/details/109230584


##最新有效版
https://blog.csdn.net/Matthew__M/article/details/125163793

go run main.go

#这个在go1.15后为失效的证书方式
### 生成CA私钥
openssl genrsa -out ca.key 2048
### 生成CA根证书，-day指定证书有效期
openssl req -x509 -new -nodes -key ca.key -subj "/CN=localhost" -days 5000 -out ca.crt
### 生成服务端私钥
openssl genrsa -out server.key 2048
### 生成证书请求文件
openssl req -new -key server.key -subj "/CN=localhost" -out server.csr
### 生成最终证书文件，-day指定证书有效期
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 5000


#openssl v3正确答案
openssl genrsa -out ca.key 2048

openssl req -new -x509 -days 36500 -key ca.key -subj "/C=CN/ST=GD/L=SZ/O=Acme, Inc./CN=Acme Root CA" -out ca.crt

openssl req -newkey rsa:2048 -nodes -keyout server.key -subj "/C=CN/ST=GD/L=SZ/O=Acme, Inc./CN=adm-rds-svc.<nameSpaceName>.svc" -out server.csr

### //generate subjectAltName is different from normal ssl certificate
openssl x509 -req -extfile <(printf "subjectAltName=DNS:adm-rds-svc.<nameSpaceName>.svc") -days 36500 -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt

#这里的DNS后面跟域名，可以添加多个DNS
openssl x509 -req -extfile <(printf "subjectAltName=DNS:localhost") -days 36500 -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt