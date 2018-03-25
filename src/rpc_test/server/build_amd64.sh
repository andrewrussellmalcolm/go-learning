#!/bin/bash
GOARCH=amd64 GOARM=5 go build
sshpass -p PTdcI69z scp server andrew@10.5.1.126:.
