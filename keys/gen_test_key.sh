#!/bin/sh
openssl genrsa -out server.key 1024
openssl req -new -key server.key -out server.csr -subj "/C=/ST=/L=/O=/OU=/CN=127.0.0.1"
openssl x509 -req -days 366 -in server.csr -signkey server.key -out server.crt
