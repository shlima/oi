PG_DSN?="postgres://localhost:5432?sslmode=disable"
REDIS_DSN?="redis://localhost:6379/12"

include db.mk

help: .db-help

test:
	PG_DSN=$(PG_DSN) \
	REDIS_DSN=$(REDIS_DSN) go test -timeout=1s -p 4 -count 1 -race ./...

.PHONY: proto
proto:
	protoc --proto_path=codec --go_out=. test.proto
