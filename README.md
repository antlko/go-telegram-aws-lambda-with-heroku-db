# Golang Telegram BOT with AWS lambda boilerplate and Heroku PostgresDB

Telegram BOT boilerplate is a bot with heroku PostgresDB and simple PostgresDB connection and deploying to AWS lambda possibility.

## Installation

Initialize serverless with the command below:
```bash
serverless
```

Download golang dependencies
```bash
go mod download
```

Create .env file with required attributes:
```bash
cp .env.dist .env
```
```bash
cp .env.local.dist .env.local
```

# Usage

## Main settings

API_TOKEN - bot token from created bot with BotFather in the Telegram

HEROKU_IS_ACTIVE - configuration for switching between heroku and simple PostgresDB 

HEROKU_API_KEY - secret key from Heroku account (on the site: Profile -> Account Settings -> API Key)

HEROKU_POSTGRES_ID - PostgresDB ID

In order to check all addons from heroku and hastily check Postgres ID run command below:

```bash
heroku addons
```

Other information here
[Heroku PostgreSQL](https://devcenter.heroku.com/articles/heroku-postgresql)

## Run

# Local usage

```bash
go run app_local/main.go
```

# Deploy to AWS lambda

At first set up credentials from Identity and Access Management (IAM) user.

```bash
export AWS_SECRET_ACCESS_KEY=<your_secret_key>
```
```bash
export AWS_ACCESS_KEY_ID=<your_access_key>
```

Build application

```bash
make build
```

or immediately deploy it

```bash
make deploy
```