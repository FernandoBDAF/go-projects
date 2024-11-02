# AWS Lambda Function Client

This project demonstrates how to create a Go client that invokes an AWS Lambda function. It builds upon our previous Lambda function deployment project and shows how to securely call it from another application.

## Prerequisites
1. Go installed on your machine
2. AWS Account with access to Lambda
3. AWS Credentials (Access Key ID and Secret Access Key)

## Project Structure
```
getFromLambda/
├── .env                  # Environment variables (not in version control)
├── .gitignore           # Git ignore file
├── go.mod               # Go module file
├── go.sum               # Go dependencies checksum
├── getFromLambda.go     # Main application code
└── README.md            # This file
```

## Key Components

### 1. AWS Credentials Setup
The application uses AWS credentials stored in environment variables. These are loaded from a `.env` file:

```bash
AWS_ACCESS_KEY_ID=your_access_key_here        # Starts with 'AKIA'
AWS_SECRET_ACCESS_KEY=your_secret_key_here    # Your secret key
AWS_REGION=us-west-1                          # Your AWS region
AWS_LAMBDA_FUNCTION_NAME=go-aws-lambda        # Your Lambda function name
```

To get your AWS credentials:
1. Log into AWS Console
2. Click your username in top right
3. Select "Security credentials"
4. Under "Access keys", click "Create access key"
5. Save both the Access Key ID and Secret Access Key securely

IMPORTANT: Never commit your `.env` file to version control!

### 2. Code Structure

The main components of `getFromLambda.go`:

```go
// Event structure matching Lambda function input
type MyEvent struct {
    Name string `json:"What is your name?"`
    Age  int    `json:"How old are you?"`
}

// Response structure matching Lambda function output
type MyResponse struct {
    Message string `json:"Answer"`
}
```

Key functionality:
1. Loads environment variables using `godotenv`
2. Creates AWS configuration with credentials
3. Initializes Lambda client
4. Invokes Lambda function with test data
5. Processes and displays response

## Setup Instructions

1. Clone the repository
2. Create `.env` file with your AWS credentials:
```bash
AWS_ACCESS_KEY_ID=your_access_key_here
AWS_SECRET_ACCESS_KEY=your_secret_key_here
AWS_REGION=us-west-1
AWS_LAMBDA_FUNCTION_NAME=go-aws-lambda
```

3. Install dependencies:
```bash
go mod tidy
```

4. Run the application:
```bash
go run getFromLambda.go
```

## Common Issues and Solutions

### 1. Credential Issues
If you see authentication errors, check:
- Correct credentials in `.env` file
- Environment variables are loading properly
- AWS credentials have necessary permissions

### 2. Region Mismatch
Ensure the AWS_REGION in your `.env` matches where your Lambda function is deployed.

### 3. Function Name
The AWS_LAMBDA_FUNCTION_NAME should match exactly the name used when deploying the Lambda function.

## Security Best Practices

1. Environment Variables
   - Never hardcode credentials in source code
   - Use `.env` file for local development
   - Add `.env` to `.gitignore`

2. AWS Credentials
   - Create IAM users with minimal required permissions
   - Regularly rotate access keys
   - Never share or expose Secret Access Keys
   - Deactivate compromised keys immediately

3. Production Deployment
   - Consider using AWS Secrets Manager
   - Use IAM roles when possible
   - Implement proper error handling
   - Monitor AWS CloudWatch logs

## Dependencies

```go
require (
    github.com/aws/aws-sdk-go-v2
    github.com/aws/aws-sdk-go-v2/config
    github.com/aws/aws-sdk-go-v2/credentials
    github.com/aws/aws-sdk-go-v2/service/lambda
    github.com/joho/godotenv
)
```

## Related Projects
This client connects to the Lambda function created in our previous project. For Lambda function deployment details, see the main project README.