migrate:
	cd sql/schema && goose postgres postgres://postgres:password@localhost:5432/gotest up
delmigrate:
	cd sql/schema && goose postgres postgres://postgres:password@localhost:5432/gotest down
sqlc:
	sqlc generate
refresh:
	go mod tidy && go mod vendor
start:
	go build
	./gotestproj

.PHONY: migrate delmigrate sqlc refresh start