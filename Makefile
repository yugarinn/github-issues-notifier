DC = docker compose -f docker-compose.yml

up:
	@${DC} up --build

up-prod:
	@${DC} up --build -d

down:
	@${DC} down

shell:
	@${DC} run --rm -it app bash -l

test:
	@${DC} run -e GIN_ENV=test app /go/bin/gotestsum --format testname ./tests/feature

build_for_production:
	GIN_ENV=production go build
