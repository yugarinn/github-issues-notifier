run:
	docker build -t github-issues-notifier . && docker run -d -v .:/github-issues-notifier github-issues-notifier

run-dev:
	docker build --build-arg GISN_ENV=development -t github-issues-notifier . && docker run -v .:/github-issues-notifier github-issues-notifier

stop:
	docker ps -q --filter "ancestor=github-issues-notifier" | xargs docker stop

shell:
	docker exec -it $$(docker ps -q --filter "ancestor=github-issues-notifier" | head -n 1) /bin/bash

logs:
	docker logs -f $$(docker ps -q --filter "ancestor=github-issues-notifier" | head -n 1)
