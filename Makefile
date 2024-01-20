.PHONY: run_restapi stop_restapi

run_restapi:
	docker build -t go-rest-api ./app && \
		docker run --rm --name go-rest-api -p 8091:8091 -e SERVICE__ENVIRONMENT=development -d go-rest-api

stop_restapi:
	docker stop go-rest-api