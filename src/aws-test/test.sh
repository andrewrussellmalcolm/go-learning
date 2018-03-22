aws lambda invoke --invocation-type RequestResponse --function-name aws-lambda-test --payload '{"message": "Hello Andrew"}' outfile.json
echo ------------------------------------
cat outfile.json
echo
echo ------------------------------------
echo
