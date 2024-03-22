# Telegram Bot Stickers

## About it
This bot is able to create sticker sets, add to existing ones, receive sticker sets, and delete stickers.
> [!WARNING]
> The bot only works with those sticker sets that were created using it.

### Technologies

>redis https://redis.io/docs/install/install-redis/
>docker https://docs.docker.com/get-docker/
>docker-compose https://docs.docker.com/compose/install/
>golang https://go.dev/doc/install

## Using

### Clone the repositiry
```
git clone https://github.com/Lopsrc/tg-bot-sticker-go
```

### Preparation

Edit the local.env file. Specify your bot's token:
```
vim config/local.env
```

Edit the local.env file. Specify the host for redis. If you want to run locally, then comment out hostname = "localhost", if in a docker container, then comment out the string hostname = "redis":
```
vim config/local.yaml
```
> hostname: "localhost" 
> hostname: "redis"   # For docker

### Running

Running on a local host:
```
# start redis.
redis-server
# installing packages and running the bot.
go mod tidy && go run cmd/bot/main.go
```
Running in a docker container:
```
# build & start.
docker-compose up -d --build
# stop.
docker stop tg-bot-sticker-go redis
```
