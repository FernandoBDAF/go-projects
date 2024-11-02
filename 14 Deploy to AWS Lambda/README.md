# AWS Lambda Deployment Guide
AWS Lambda is a serverless function that can be used to run code in response to events.
This guide explains how to deploy a Go function to AWS Lambda using the AWS CLI.

## Prerequisites
1. AWS CLI installed and configured
2. Go installed on your machine
3. Basic understanding of AWS IAM roles and policies

## Step by Step Guide

### 1. Write the Go code
Create your Lambda function in `main.go`. Ensure it follows AWS Lambda function signature requirements.

### 2. Set up IAM Permissions
First, you need proper permissions to create and manage AWS resources.

#### Common Issue #1: Insufficient Permissions
If you see permission errors like:
```
is not authorized to perform: lambda:CreateFunction
```
Fix by attaching necessary policies to your IAM user:
```bash
aws iam attach-user-policy --user-name YOUR_USERNAME --policy-arn arn:aws:iam::aws:policy/AdministratorAccess
```

### 3. Create IAM Role for Lambda
Lambda needs a role to execute and access other AWS services.

#### Create trust policy
Create a file named `trust-policy.json`:
```json
{
    "Version": "2012-10-17",
    "Statement": [{ 
        "Effect": "Allow", 
        "Principal": {"Service": "lambda.amazonaws.com"}, 
        "Action": "sts:AssumeRole"
    }]
}
```

#### Create the role
```bash
aws iam create-role --role-name lambda-ex --assume-role-policy-document file://trust-policy.json
```

#### Common Issue #2: Role Already Exists
If you see:
```
An error occurred (EntityAlreadyExists) when calling the CreateRole operation: Role with name lambda-ex already exists
```
Either use the existing role or create a new one with a different name.

#### Attach execution policy
```bash
aws iam attach-role-policy \
    --role-name lambda-ex \
    --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole
```

### 4. Build and Package the Function
For `provided.al2` runtime, we need to create a bootstrap file:

```bash
# Build the Go binary
GOOS=linux GOARCH=amd64 go build -o bootstrap main.go

# Create deployment package
zip function.zip bootstrap
```

Important Notes:
1. The binary MUST be named `bootstrap` (not `main`)
2. Build for Linux as Lambda runs on Linux
3. Use GOARCH=amd64 for Lambda architecture
4. Make sure the binary is executable

#### Common Issue #4: Invalid Entrypoint
If you see:
```
{"errorType":"Runtime.InvalidEntrypoint","errorMessage":"Couldn't find valid bootstrap(s)"}
```
This means:
- The binary wasn't named `bootstrap`, or
- The binary wasn't built for Linux, or
- The binary wasn't included in the zip file correctly

### Updating an Existing Function

#### The Bootstrap Issue: A Common Gotcha
During our initial deployment, we encountered this error when testing the function:
```
{"errorType":"Runtime.InvalidEntrypoint","errorMessage":"RequestId: b1ae01d9-ba36-46ba-9b17-02c7337281ee Error: Couldn't find valid bootstrap(s): [/var/task/bootstrap /opt/bootstrap]"}
```

This error occurred because our initial deployment used a binary named `main` instead of `bootstrap`. While the function was deployed successfully, it failed at runtime because the `provided.al2` runtime specifically looks for an executable named `bootstrap` in specific locations.

To fix this, we needed to update our existing function rather than create a new one. Here's the process:

1. Rebuild the binary with the correct name:
```bash
GOOS=linux GOARCH=amd64 go build -o bootstrap main.go
```

2. Create a new deployment package:
```bash
zip function.zip bootstrap
```

3. Update the function code:
```bash
aws lambda update-function-code \
    --function-name go-aws-lambda \
    --zip-file fileb://function.zip
```

This update was necessary because:
- AWS Lambda's `provided.al2` runtime specifically looks for an executable named `bootstrap`
- The executable must be built for Linux (GOOS=linux) as Lambda runs on Linux
- The architecture must match Lambda's environment (GOARCH=amd64)
- The original deployment with a differently named binary won't work, even if the code is correct

The error message gives us two paths where Lambda looks for the bootstrap file:
- `/var/task/bootstrap`: This is where our uploaded code goes
- `/opt/bootstrap`: An alternative location for custom runtimes

By renaming our binary to `bootstrap` and updating the function, we ensure it's placed in the correct location that Lambda expects.

After updating, verify the function works by invoking it again:
```bash
aws lambda invoke \
    --function-name go-aws-lambda \
    --payload '{}' \
    response.json
```

### 5. Deploy the Lambda Function
```bash
aws lambda create-function \
    --function-name go-aws-lambda \
    --zip-file fileb://function.zip \
    --handler main \
    --runtime provided.al2 \
    --role arn:aws:iam::YOUR_ACCOUNT_ID:role/lambda-ex
```

#### Common Issue #3: Runtime Not Supported
If you see:
```
The runtime parameter of go1.x is no longer supported
```
Use `provided.al2` instead of `go1.x` as shown in the command above.

### 6. Verify Deployment
Check if your function was created successfully:
```bash
aws lambda get-function --function-name go-aws-lambda
```

### 7. Test the Function
```bash
aws lambda invoke \
    --function-name go-aws-lambda \
    --payload '{}' \
    response.json
```

```
aws lambda invoke \
    --function-name go-aws-lambda \
    --cli-binary-format raw-in-base64-out \
    --payload '{"What is your name?": "Fernando", "How old are you?": 33}' \
    response.json
```

```
cat response.json
```

## Troubleshooting Tips
1. Always check IAM permissions first if you get access denied errors
2. Ensure your Go function is properly built for AWS Lambda
3. Make sure the role has proper permissions attached
4. Use `provided.al2` runtime for Go functions
5. Wait a few seconds after creating roles before using them (AWS propagation delay)

## Additional Configuration
You can update function configuration like timeout and memory using:
```bash
aws lambda update-function-configuration \
    --function-name go-aws-lambda \
    --timeout 10 \
    --memory-size 256
```