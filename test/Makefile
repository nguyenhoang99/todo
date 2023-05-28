generate-sqlboiler:
	sqlboiler psql -c sqlboiler.toml --wipe --no-tests

run:
	go run main.go

run-docker:
	docker compose up

down-docker:
	docker compose down

gen-mocks:
	mockery --dir api/controller --all --recursive --inpackage
	mockery --dir api/repository --all --recursive --inpackage

stop-postgres:
	systemctl stop postgresqldocker