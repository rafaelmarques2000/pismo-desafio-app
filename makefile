serve:
	@docker-compose -f ./docker-compose.yaml up
rebuild:
	@docker-compose -f ./docker-compose.yaml up --build
run-all-tests:
	@docker-compose -f ./docker-compose.yaml exec -T api sh -c "go test ./..."
shell:
	@docker-compose -f ./docker-compose.yaml exec api bash
