# Slack + React Bot

## Introduction

A modern Slack bot implementation using React for the frontend and Go for the backend. This project showcases the integration of Slack's bot functionality using the Slacker framework, which provides an elegant way to handle Slack commands and events.

### Implementation Approaches

We chose the Slacker framework approach over other common methods:

### 1. Direct API Integration
- Raw implementation using Slack's Web API
- Direct HTTP requests to Slack endpoints
- Maximum flexibility but requires more boilerplate code
- Best for simple, specific integrations

### 2. Slack API Library/Wrapper (slack-go/slack)
- Uses `slack-go/slack` package
- Provides convenient Go functions for Slack API operations
- Handles authentication and request formatting
- Good for standard bot functionality

### 3. Slacker Framework (Our Choice)
- Built on top of the `slack-go/slack` package
- Provides an intuitive command-based interface
- Simplifies event handling and command parsing
- Includes built-in command definition system
- Perfect for interactive command-driven bots
- Reduces boilerplate code significantly

Our implementation leverages the Slacker framework's strengths:
- Easy command definitions
- Built-in parameter parsing
- Automatic help message generation
- Event-driven architecture
- WebSocket-based real-time messaging
- Simplified error handling

This project combines React's responsive frontend with Slacker's powerful backend capabilities to create an interactive and efficient bot experience.

# Setting Up Your Slack Bot

## Detailed Setup Process

### 1. Create a Slack App
1. Go to [Slack API](https://api.slack.com/apps)
2. Click "Create New App"
3. Choose "From Scratch"
4. Enter your app name and select your workspace
5. Click "Create App"

### 2. Enable Socket Mode
1. Navigate to "Socket Mode" in the left sidebar
2. Toggle "Enable Socket Mode" to On
3. You'll be prompted to create an App-Level Token
4. Name your token (e.g., "socket-token")
5. Add the `connections:write` scope
6. Click "Generate"
7. Save the generated token (starts with `xapp-`) - this is your `SLACK_APP_TOKEN`

### 3. Configure Bot Events
1. Go to "Event Subscriptions" in the sidebar
2. Toggle "Enable Events" to On
3. Under "Subscribe to bot events", add:
   - `app_mention`
   - `message.channels`
   - `message.groups`
   - `message.im`
   - `message.mpim`

### 4. Set Bot Token Scopes
1. Navigate to "OAuth & Permissions"
2. Under "Scopes" > "Bot Token Scopes", add:
   - `app_mentions:read`
   - `chat:write`
   - `channels:read`
   - `groups:read`
   - `im:read`
   - `mpim:read`

### 5. Install App to Workspace
1. Click "Install to Workspace" at the top of the OAuth page
2. Review and allow the permissions
3. Copy the "Bot User OAuth Token" (starts with `xoxb-`) - this is your `SLACK_BOT_TOKEN`

### 6. Final Configuration
1. Create a `.env` file in your project root
2. Add your tokens:

```env
SLACK_BOT_TOKEN=xoxb-your-bot-token
SLACK_APP_TOKEN=xapp-your-app-token
```

### 7. Test Your Bot
1. Invite your bot to a channel: `/invite @your-bot-name`
2. Mention your bot: `@your-bot-name hello`
3. Try the implemented commands

### Important Notes
- Always reinstall your app after making permission changes
- Keep your tokens secure and never commit them to version control
- Use Socket Mode for development and testing without exposing endpoints
- Monitor your bot's interactions in the Slack API dashboard

For specific implementation examples, see:
- File Upload Bot: 
```markdown:07 Slack BOT - File Upload/README.md
startLine: 15
endLine: 22
```
- Age Calculator Bot:
```markdown:08 Slack BOT - Calculate Age/README.md
startLine: 51
endLine: 55
```