
## variables
process=http-server
server=192.168.0.21
user=pi
password=PTdcI69z

## build step
echo building $process
GOARCH=arm GOARM=5 go build -o $process

## kill step
echo killing $process
sshpass -p $password ssh $user@$server pkill -f $process

## upload step
echo uploading $process
sshpass -p $password scp $process $user@$server:.

## run step
echo running $process
sshpass -p $password ssh $user@$server ./$process &
