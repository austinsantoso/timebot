# timebot

Telegram bot.

[t.me/austin_time_bot](http://t.me/austin_time_bot)

## Introduction

A telegram bot will usefull functions to tell the time.

## Prerequisites

- GO version `1.17.0`
- [mockgen](https://github.com/golang/mock) `v1.6.0`

## Deployment

Deployed on [Netlify](https://www.netlify.com/).
Using Netlify's Functions.

[Entrypoint](./netlify/functions/bot/main.go) for Netlify function.
Has been set as a webhook for bot.

## Development

### Install dependencies

```
go mod install
```

### Local Development

You can run the [main.go](./main.go) locally.

1. Get a bot token from [bot Father](t.me/botfather)
2. `go run main.go -token=<InsertBotToken>`

This will start a local server that will continually poll for updates from the bot.

## References

- [Netlify setup](https://travishorn.com/building-a-telegram-bot-with-netlify)
