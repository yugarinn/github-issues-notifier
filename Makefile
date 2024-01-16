build:
	docker build -t github-issues-notificator .

build-dev:
	docker build --build-arg GISN_ENV=development -t github-issues-notificator .

run:
	docker run -d -v .:/github-issues-notificator github-issues-notificator

stop:
	docker ps -q --filter "ancestor=github-issues-notificator" | xargs docker stop

shell:
	docker exec -it $$(docker ps -q --filter "ancestor=github-issues-notificator" | head -n 1) /bin/bash

logs:
	docker logs $$(docker ps -q --filter "ancestor=github-issues-notificator" | head -n 1)
