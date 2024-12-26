#!/bin/bash

function create_database() {
    local database_name=$1
    echo "CREATE DATABASE IF NOT EXISTS $database_name;" | mysql -u root -h 127.0.0.1 -P 3306
}

function migrate() {
    local database_name=$1
    docker run --rm --network host -v ./db/migrations:/migrations migrate/migrate \
        -path /migrations \
        -database "mysql://root@tcp(127.0.0.1:3306)/$database_name" \
        -verbose up
}

# create_database と create_table を使ってデータベースとテーブルを作成する
create_database hakushi
create_database hakushi_test

migrate hakushi
migrate hakushi_test
