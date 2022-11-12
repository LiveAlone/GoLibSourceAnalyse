#!/usr/bin/env bash

openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes

#openssl req -newkey rsa:2048 -nodes -keyout data/server.key -x509 -days 365 -out data/server.crt