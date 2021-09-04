THIS_FILE := $(lastword $(MAKEFILE_LIST))
PWD ?= pwd_unknown

# PROJECT_NAME defaults to name of the current directory.
# should not to be changed if you follow GitOps operating procedures.
PROJECT_NAME = $(notdir $(PWD))

.PHONY: help rebuild build up start clean down destroy stop restart logs logs-web ps

help:
	@echo ''
	@echo 'Usage: make [TARGET] [EXTRA_ARGUMENTS]'
	@echo 'Targets:'
	@echo '  build    	build docker --image-- for current user: $(HOST_USER)'
	@echo '  rebuild  	rebuild docker --image-- for current user: $(HOST_USER)'
	@echo '  test     	test docker --container-- for current user: $(HOST_USER)'
	@echo '  clean    	remove docker --image-- for current user: $(HOST_USER)'
	@echo '  prune    	shortcut for docker system prune -af. Cleanup inactive containers and cache.'
	@echo '  shell		run docker --container-- for current user: $(HOST_USER)(uid=$(HOST_UID))'
	@echo '  stop		stop the service'
	@echo '  start		start the service'
	@echo '  restart	stop and restart the service'
	@echo '  up		start the containers in the background'
	@echo '  ps		list containers'
	@echo '  push		push to repository'
	@echo '  pull		pull the image'
	@echo '  destroy	remove volumes'
	@echo '  logs		Displays log output from services'

rebuild:
	# force a rebuild by passing --no-cache
	docker-compose build --no-cache 	
build:
	docker-compose -f docker-compose.yml up -d --build $(c)
up:
	docker-compose -f docker-compose.yml up -d $(c)
start:
	docker-compose -f docker-compose.yml start $(c)
down:
	docker-compose -f docker-compose.yml down $(c)
destroy:
	docker-compose -f docker-compose.yml down -v $(c)
stop:
	docker-compose -f docker-compose.yml stop $(c)
restart:
	docker-compose -f docker-compose.yml stop $(c)
	docker-compose -f docker-compose.yml up -d $(c)
logs:
	docker-compose -f docker-compose.yml logs --tail=100 -f $(c)
logs-web:
	docker-compose -f docker-compose.yml logs --tail=100 -f web
ps:
	docker-compose -f docker-compose.yml ps
push:
	docker-compose -f docker-compose.yml push
pull:
	docker-compose -f docker-compose.yml push
prune:
	# clean all images that are not actively used
	docker system prune -af
clean:
	# remove created images orphans deleted images
	docker-compose -p down --remove-orphans --rmi all 2>/dev/null 
test:
	# here it is useful to add your own customised tests
	
	echo "I am `whoami`. My uid is `id -u`." && echo "Docker runs!"
	
	