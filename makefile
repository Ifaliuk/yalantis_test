dc_up:
	docker-compose up --force-recreate -d

run_tests:
	docker run --rm -it -v ${shell PWD}:/app -w /app golang:1.15.13 sh -c "cd internal && go test"

dc_down:
	docker-compose down

dc_restart: dc-down dc-up