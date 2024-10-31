package main

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
		return
	}

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	// Test the API connection first
	_, err = api.AuthTest()
	if err != nil {
		fmt.Printf("Auth Error: %v\n", err)
		return
	}

	channelID := "C07U5CH6MFF"
	fileArr := []string{"sherlock.txt"}

	for i := 0; i < len(fileArr); i++ {
		file, err := api.UploadFileV2(slack.UploadFileV2Parameters{
			Channel: channelID,
			File:    fileArr[i],
			Filename:       fileArr[i],
			FileSize:       100,
		})
		if err != nil {
			// More detailed error handling
			if rateLimitErr, ok := err.(*slack.RateLimitedError); ok {
				fmt.Printf("Rate limit error: %s\n", rateLimitErr)
			} else {
				fmt.Printf("Upload Error: %v\n", err)
			}
			return
		}
		fmt.Printf("Name: %s, URL: %s\n", file.ID, file.Title)
		fmt.Printf("file: %v\n", file)
	}
}
