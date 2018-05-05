#!/bin/bash
GOARCH=arm GOARM=5 go build
sshpass -p PTdcI69z scp server pi@86.26.15.118:.
