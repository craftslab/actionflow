#!/bin/bash

# USAGE: https://github.com/swaggo/swag/blob/master/README.md
# WEB: http://ip:port/swagger/index.html

release=1.6.7

if [ -d swag ]; then
    rm -rf swag
fi

mkdir swag

curl -L https://github.com/swaggo/swag/releases/download/v${release}/swag_${release}_Linux_x86_64.tar.gz -o swag.tar.gz
tar zxvf swag.tar.gz -C swag/
rm -rf swag.tar.gz

./swag/swag init

rm -rf swag
