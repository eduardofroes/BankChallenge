BANK_VERSION=${shell cat ./VERSION}
IMAGE_TAG=bankchallenge:${BANK_VERSION}
CURRENT_DIR = $(shell pwd)

build: binary build-image

run-all: build create-net run-psql run-bank

stop-all: stop-bank stop-psql delete-net

.PHONY: binary
binary:
	docker run --rm -v ${CURRENT_DIR}/src:/go/src/bankchallenge\
				   -w /go/src/bankchallenge\
				   golang:1.15-alpine\
				   go build .

.PHONY: build-image
build-image: binary
	docker build -t $(IMAGE_TAG) .

.PHONY: run-bank
run-bank:
	docker run --name bankchallenge --rm \
			   -d \
			   -h bankchallenge \
			   -v ${CURRENT_DIR}:${CURRENT_DIR} \
			   -w ${CURRENT_DIR} \
			   -e DATABASE_HOST=postgres \
			   -e DATABASE_PORT=5432 \
			   -e DATABASE_USER=admin \
			   -e DATABASE_PASS=admin \
			   -e DATABASE_NAME=bank \
			   -e DATABASE_DRIVE=postgres \
			   --net=bank-network \
			   -p 8080:8080 \
			   bankchallenge:${BANK_VERSION}

.PHONY: stop-bank
stop-bank:
	docker stop bankchallenge

.PHONY: run-psql
run-psql:
	docker run --name postgres --rm \
	 -d \
	 -h postgres \
	 -v ${CURRENT_DIR}/sql:/docker-entrypoint-initdb.d/ \
	 -e POSTGRES_PASSWORD=admin \
	 -e POSTGRES_USER=admin \
	 -e POSTGRES_DB=bank \
	 --net=bank-network \
	 -p 5432:5432 \
	 postgres

.PHONY: stop-psql
stop-psql:
	docker stop postgres

.PHONY: create-net
create-net:
	docker network create bank-network

.PHONY: delete-net
delete-net:
	docker network rm bank-network