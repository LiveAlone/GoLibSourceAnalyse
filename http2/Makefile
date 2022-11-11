certs:
	openssl req -newkey rsa:2048 -nodes -keyout data/server.key -x509 -days 365 -out data/server.crt

run-h1-server:
	go run main/h1Server.go

run-h1-client:
	go run main/Client.go