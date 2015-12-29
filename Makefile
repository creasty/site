HEROKU_APP_NAME ?= crst-site

.PHONY: heroku-app
heroku-app:
	@heroku create ${HEROKU_APP_NAME} -b https://github.com/heroku/heroku-buildpack-multi.git

.PHONY: build
build:
	@godep get
	@godep go build

.PHONY: run
run:
	@./site

.PHONY: dev
dev:
	@./site --dev
