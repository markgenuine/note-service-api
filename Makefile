LOCAL_MIGRATION_DIR=./migrations
LOCAL_MIGRATION_DSN="host=localhost port=54321 dbname=service-db user=service-user password=service-password sslmode=disable"

PHONY: generate
generate:
	mkdir -p pkg/note_v1
	protoc --proto_path api/note_v1 \
		   --go_out=pkg/note_v1 --go_opt=paths=import \
		   --go-grpc_out=pkg/note_v1 --go-grpc_opt=paths=import \
		   api/note_v1/note.proto
	mv pkg/note_v1/github.com/markgenuine/note-service-api/pkg/note_v1/* pkg/note_v1/
	rm -rf pkg/note_v1/github.com

.PHONE: install-goose
.install-goose:
	go install github.com/pressly/goose/v3/cmd/goose@latest

.PHONE: local-migration-status
local-migration-status:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v


.PHONE: local-migration-up
local-migration-up:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

.PHONE: local-migration-down
local-migration-down:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v