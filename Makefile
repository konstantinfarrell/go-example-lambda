LAMBDA_NAME = go-example-shovel


package:
	go build lambda/main.go
	zip packaged.zip main

publish:
	aws lambda publish-version --function-name $(LAMBDA_NAME)