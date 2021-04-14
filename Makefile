BANK_VERSION=${shell cat ./VERSION}
IMAGE_TAG=bankchallenge:${BANK_VERSION}
CURRENT_DIR = $(shell pwd)

all: binary build-image save-image

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
	 -p 5432:5432 \
	 postgres

.PHONY: stop-psql
stop-psql:
	docker stop postgres