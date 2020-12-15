#!/bin/sh
protoc --proto_path=. --go_out=. msg.proto
