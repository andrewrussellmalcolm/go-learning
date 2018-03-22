GOOS=linux go build -o main

zip aws-lambda-test.zip main

aws lambda update-function-code \
--function-name aws-lambda-test \
--zip-file fileb://./aws-lambda-test.zip \

#aws lambda create-function \
#--region eu-west-2 \
#--function-name aws-lambda-test \
#--zip-file fileb://./aws-lambda-test.zip \
#--runtime go1.x \
#--role arn:aws:iam::407811233194:role/aws-lamba-test \
#--handler main

#--tracing-config Mode=Active \
