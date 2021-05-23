# A Go Lambda

Just a simple lambda that reads from AWS Kinesis.
This is the data processing component for the golang example microservices.

WRITE operations are sent to this lambda which are then processed and written to a database.

Since lambda is charged by function call & execution duration, it's not advisable to perform any heavy processing inside the lambda.
A better architecture would be using the lambda as a shovel to pipe kinesis data into another REST API. This solution scales well and
is cheaper to run.

But since the goal of these services currently is to practice golang, there's no pressing need to do this.

## Usage

To create an executable and bundle it in a .zip file, use

    make

OR

    make package

TODO: Add AWS CLI commands to makefile to automate lambda publishing based on template file