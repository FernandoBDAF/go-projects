# Slack Bot - File Upload

A Slack bot that allows users to upload files to channels.

## Environment Variables Setup

To run this Slack bot, you'll need to configure the following environment variables in your `.env` file:

### 1. Create a Slack App
1. Go to [Slack API Apps page](https://api.slack.com/apps)
2. Click "Create New App"
3. Choose "From scratch"
4. Name your app and select your workspace

### 2. Configure Bot Token Scopes
1. In your app settings, go to "OAuth & Permissions"
2. Under "Scopes" > "Bot Token Scopes", add these permissions:
   - `files:read`
   - `files:write`
   - `chat:write`
   - `channels:read`
   - `channels:join`

### 3. Install App to Workspace
1. Go to "Install App" in the sidebar
2. Click "Install to Workspace"
3. Authorize the app

### 4. Get Required Tokens
After installation, you'll need these values for your `.env` file:
- `SLACK_BOT_TOKEN`: Find under "OAuth & Permissions" > "Bot User OAuth Token"
- `SLACK_APP_TOKEN`: Go to "Basic Information" > "App-Level Tokens" > Create "connections:write" scope token

### 5. Create .env File
Create a `.env` file in your project root with:
```env
SLACK_BOT_TOKEN=xoxb-your-bot-token
SLACK_APP_TOKEN=xapp-your-app-token
SLACK_CHANNEL=your-channel-name
```

### 6. Enable Socket Mode
1. Go to "Socket Mode" in your app settings
2. Enable Socket Mode
3. This is required for the app to work without exposing a public URL

Remember to never commit your `.env` file to version control. Make sure it's listed in your `.gitignore`.

## Usage

1. Make sure you have Go installed on your system
2. Clone this repository
3. Set up your `.env` file as described above
4. Run the bot:
```bash
go run main.go
```

## Features

- Upload files to specified Slack channels
- Supports various file types
- Real-time file upload notifications

## Dependencies

- Go 1.x
- slack-go/slack package