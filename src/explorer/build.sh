#!/bin/bash
GOARCH=arm GOARM=5 go build
#sshpass -p raspberry scp explorer pi@192.168.0.21:.
