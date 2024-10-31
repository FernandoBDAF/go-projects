# Slack Age Calculator Bot

A simple Slack bot that calculates age based on the year of birth (YOB) input. Built with Go using the Slacker framework.

## Table of Contents
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [Development](#development)
- [Security](#security)
- [Contributing](#contributing)
- [License](#license)

## Features
- 🔢 Calculates age from year of birth
- 💬 Responds to Slack commands
- 📝 Real-time command event logging
- ⚙️ Environment variable configuration

## Prerequisites
- Go 1.x
- Slack Workspace
- Slack App with Bot Token and App-Level Token

## Installation

1. Clone the repository

git clone <repository-url>

2. Install dependencies

go mod download

3. Set up environment variables

### Environment Variables Setup

To set up your Slack bot environment variables, follow these steps:

1. **Create a Slack App**
   - Visit [Slack API Apps page](https://api.slack.com/apps)
   - Click "Create New App"
   - Choose "From scratch"
   - Name your app (e.g., "Age Calculator")
   - Select your workspace

2. **Configure Bot Token Scopes**
   - Navigate to "OAuth & Permissions" in sidebar
   - Under "Scopes" > "Bot Token Scopes", add:
     - `app_mentions:read`
     - `chat:write`
     - `commands`

3. **Enable Socket Mode**
   - Go to "Socket Mode" in sidebar
   - Toggle "Enable Socket Mode" to On
   - Create an App-Level Token with `connections:write` scope
   - Save the generated `xapp-` token

4. **Install App to Workspace**
   - Go to "Install App" in sidebar
   - Click "Install to Workspace"
   - Authorize the app

5. **Collect Tokens**
   - Bot Token: Find under "OAuth & Permissions" > "Bot User OAuth Token" (starts with `xoxb-`)
   - App Token: The previously generated token (starts with `xapp-`)

6. **Create .env File**
   Create a `.env` file in your project root with:
   ```env
   SLACK_BOT_TOKEN=xoxb-your-bot-token
   SLACK_APP_TOKEN=xapp-your-app-token
   ```

7. **Verify Setup**
   - Ensure your bot is invited to the desired channels
   - Test the connection by running the bot
   - Try the age calculation command

## Configuration

### Slack App Setup
1. Create a new Slack App at [api.slack.com/apps](https://api.slack.com/apps)
2. Enable Socket Mode
3. Add Bot Token Scopes:
   - `chat:write`
   - `commands`
4. Install the app to your workspace
5. Copy the tokens:
   - Bot User OAuth Token (starts with `xoxb-`)
   - App-Level Token (starts with `xapp-`)

### Environment Variables
| Variable | Description |
|----------|-------------|
| `SLACK_BOT_TOKEN` | Slack Bot User OAuth Token |
| `SLACK_APP_TOKEN` | Slack App-Level Token |

## Usage

1. Start the bot:
```bash
go run main.go
```

2. In Slack, use the command:
```
my yob is <year>
```

Example: `my yob is 2000`

The bot will respond with your calculated age.

### Command Events
The bot logs the following events:
- Timestamp
- Command used
- Parameters
- Event details

## Development

### Project Structure
```
├── main.go
├── .env
├── .env.example
├── README.md
└── .gitignore
```

### Error Handling
The bot includes error handling for:
- Invalid year inputs
- Missing environment variables
- Connection issues

## Security
- Never commit `.env` file
- Add `.env` to `.gitignore`
- Keep Slack tokens secure
- Rotate tokens if exposed
- Use environment variables for sensitive data
