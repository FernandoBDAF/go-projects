package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "os"

    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/credentials"
    "github.com/aws/aws-sdk-go-v2/service/lambda"
    "github.com/joho/godotenv"
)

type MyEvent struct {
    Name string `json:"What is your name?"`
    Age  int    `json:"How old are you?"`
}

type MyResponse struct {
    Message string `json:"Answer"`
}

func init() {
    // Load .env file
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file:", err)
    }
}

func main() {
    // Get credentials from environment variables
    accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
    secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
    region := os.Getenv("AWS_REGION")
    functionName := os.Getenv("AWS_LAMBDA_FUNCTION_NAME")

    // Validate required environment variables
    if accessKeyID == "" || secretAccessKey == "" || region == "" || functionName == "" {
        log.Fatal("Missing required AWS credentials in environment variables")
    }

    // Create credentials provider
    creds := credentials.NewStaticCredentialsProvider(
        accessKeyID,
        secretAccessKey,
        "", // Session token (optional)
    )

    // Load AWS configuration
    cfg, err := config.LoadDefaultConfig(context.TODO(),
        config.WithRegion(region),
        config.WithCredentialsProvider(creds),
    )
    if err != nil {
        log.Fatal("Unable to load SDK config:", err)
    }

    // Create Lambda client
    client := lambda.NewFromConfig(cfg)

    // Prepare the input payload
    event := MyEvent{
        Name: "Fernando",
        Age:  33,
    }
    
    payload, err := json.Marshal(event)
    if err != nil {
        log.Fatal("Error marshaling event:", err)
    }

    // Invoke Lambda function
    result, err := client.Invoke(context.TODO(), &lambda.InvokeInput{
        FunctionName: &functionName,
        Payload:      payload,
    })
    if err != nil {
        log.Fatal("Error invoking function:", err)
    }

    // Parse the response
    var response MyResponse
    err = json.Unmarshal(result.Payload, &response)
    if err != nil {
        log.Fatal("Error unmarshaling response:", err)
    }

    fmt.Println("Lambda Response:", response.Message)
}