#!/bin/bash

protoc protos/calculator.proto --go_out=plugins=grpc:./calculatorpb/