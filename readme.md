# Slack File Bot in Golang

A simple Go application to upload files to a channel within a Slack workspace. This bot makes it easy to share files with your team members without leaving Slack.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
  - [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)

## Features

- Upload files to a specific Slack channel.
- Supports various file types such as documents, images, and more.
- Easily configurable through environment variables.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- You have a Slack workspace where you have the necessary permissions to add a bot and upload files.
- You have Go installed on your local machine.

## Getting Started

To get started with this Slack File Bot, follow these steps:

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/zohaibsoomro/slack-file-bot.git
   ```

2. Change into the project directory:

   ```bash
   cd slack-file-bot
   ```

3. Install the necessary dependencies:

   ```bash
   go mod tidy
   ```

## Usage

To run the Slack File Bot, use the following command:

```bash
 go run main.go
```

## Contributing

Contributions are welcome! If you'd like to contribute to this project, please follow these guidelines:

Fork the repository.
Create a branch: git checkout -b feature/your-feature-name.
Commit your changes: git commit -m 'Add some feature'.
Push to the branch: git push origin feature/your-feature-name.
Create a pull request.
