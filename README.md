##Discord Go Bot
This is a small Discord Bot written in go for learning.

###Installation
First you need to clone the Repository
```
git clone https://github.com/UniqueEckart/GoBot
```
After that you need to configure the Config file of the Bot which is located in the source Directory.

Now you are able to start the Bot in 2 Ways.

The first would be to simply run
```go
go run bot.go
```
The Second is to First build the Bot and then run it.
```go
go build
```
Then give it Permissions to execute
```bash
chmod +x bot
```
And Finally run
```bash
./bot
```

##Features

The Bot has a small build-in webserver which is currently only for Healthcheck. But maybe there will come more the only endpoint is.
```url
http://locahost:8080/healthcheck
```