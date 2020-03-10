NAME = kay/exchange-rate-app
TAG=v1
INSTANCE = api-exchange-rate-app

.PHONY: default build run buildrun

default: buildrun

build:
	docker-compose build 

notify:
	@echo "Name: "  $(NAME) " Tag: " $(TAG)  " Instance: " $(INSTANCE) 

run:
	docker-compose up

buildrun: build notify run

