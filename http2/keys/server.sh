// 服务端私钥
openssl genrsa -out server.key 2048
//服务端签名请求
openssl req -new -sha256 -out server.csr -key server.key -config server.conf
//用根证书签发服务端证书server.pem
openssl x509 -req -days 3650 -CA ca.crt -CAkey ca.key -CAcreateserial -in server.csr -out server.pem -extensions req_ext -extfile server.conf