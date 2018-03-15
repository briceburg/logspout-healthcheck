# logspout-healthcheck

adds a healtcheck endpoiunt to [gliderlabs/logspout](https://github.com/gliderlabs/logspout).

## usage

* add `logspout-healthcheck` to your modules.go. example;

  ```go
  package main

  import (
  	_ "github.com/gliderlabs/logspout/adapters/syslog"
  	_ "github.com/gliderlabs/logspout/transports/tcp"
  	_ "github.com/briceburg/logspout-healthcheck"
    ...
  )
  ```

* build logspout
* test `/healthcheck` endpoint. example with logspout listening on :8888;

```
$ docker run --rm -d \
  -p 8888:80 \
  -v /var/run/docker.sock:/var/run/docker.sock \
  logspout:custom

$ curl http://localhost:8888/healthcheck
OK
```

## usage as a Dockerfile HEALTCHECK

```
FROM gliderlabs/logspout

HEALTHCHECK --interval=10s --timeout=3s --retries=2 \
  CMD wget -q --spider http://localhost/healthcheck
```

> wget is included in alpine/busybox. curl is not.
