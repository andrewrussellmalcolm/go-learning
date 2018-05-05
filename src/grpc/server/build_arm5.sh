#!/bin/bash
GOARCH=amd64 GOARM=5 go build
sshpass -p PTdcI69z scp server pi@86.26.15.118:.
