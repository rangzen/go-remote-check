# Go Remote Check

Use this image to run command from a Docker Compose configuration.

My first goal is to check the network config from inside through a basic web app.

## Run Go directly

```shell
go run main.go
```

## Docker file to build in Docker Compose

```yaml
version: '3'
services:
  go-remote-check:
    build:
      context: ./go-remote-check
    # Additional service configuration
```
