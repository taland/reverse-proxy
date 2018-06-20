# reverse-proxy impl

## Build

```
$ docker build -t proxy .
```

or

```
$ docker build -t proxy -f Dockerfile.vgo .
```

## Run

```
$ docker run -p 5555:5555 --name proxy-app proxy
```
