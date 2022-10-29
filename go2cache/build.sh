#!/usr/bin/env bash
protoc --go_out=plugins=grpc,import_path=go2cachev2:. *.proto
