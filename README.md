site
====

Here comes v5 of http://creasty.com


Development
-----------

### Setup

Configure environmental variable:

```sh
$ cp .env{.sample,}
$ vi $_
```

Build application:

```sh
$ make build
```

### Run local server

Simple run:

```sh
$ make run
$ make dev  # with development mode enabled (auto compile & restart)
```


Heroku
------

```sh
$ HEROKU_APP_NAME=foo make heroku-app
```
