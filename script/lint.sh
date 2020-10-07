#!/bin/bash

gofmt -s -w .
golangci-lint run

go mod tidy
go mod verify
