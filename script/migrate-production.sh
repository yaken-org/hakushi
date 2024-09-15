#!/bin/bash
read -sp "Password: " PASSWORD
tty -s && echo

docker run --rm --network host -v ./db/migrations:/migrations migrate/migrate -path /migrations -database mysql://hakushi:$PASSWORD@tcp'(ishikari.node.yaken.org:3306)'/hakushi -verbose up
