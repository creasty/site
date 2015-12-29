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
$ go get
$ go build
```

### Run local server

Simple run:

```sh
$ ./site
$ ./site --dev  # with development mode enabled (auto compile & restart)
```


Heroku
------

```sh
$ heroku create crst-site -b https://github.com/heroku/heroku-buildpack-multi.git
```
