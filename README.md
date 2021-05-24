# A Go Lambda

## Introduction

Just a simple lambda that reads from AWS Kinesis.
This is the data processing component for the golang example microservices.

WRITE operations are sent to this lambda which are then processed and written to a database.

### Notes

Since lambda is charged by function call & execution duration, it's not advisable to perform any heavy processing inside the lambda.
A better architecture would be using the lambda as a shovel to pipe kinesis data into another REST API. This solution scales well and
is cheaper to run.

Another issue is that this lambda processes data for multiple operations. A better architecture would be to have one Kinesis stream per operation,
and have a second lambda set up for just that. This separates data which makes monitoring and debugging more straightforward. But also makes it 
easier to update and modify methods of handling data.

But since the goal of these services currently is to practice golang, there's no pressing need for any of this.



## Usage

To create an executable and bundle it in a .zip file, use

    make

OR

    make package

TODO: Add AWS CLI commands to makefile to automate lambda publishing based on template file