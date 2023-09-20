# Lab Homepage

## Build

```shell
$ docker compose build --build-arg GOPROXY=<proxy-url>
```

## Testing for devlopment

```shell
$ bash ./test.sh start_mysql
$ bash ./test.sh dev_backend
$ bash ./test.sh dev_frontend
```