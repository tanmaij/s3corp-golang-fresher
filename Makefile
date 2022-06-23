path="./data/migrations"

migrate_create:
	migrate create -ext sql -dir data/migrations -seq research_document
migrate_up:
	migrate -path $(path) -database "postgresql://mai:1@localhost:5432/researchdocument?sslmode=disable" -verbose up
migrate_down:
	migrate -path $(path) -database "postgresql://mai:1@localhost:5432/researchdocument?sslmode=disable" -verbose down
build_image:
	docker build -t app .
migrate_go:
	migrate -path  $(path) -database "postgresql://mai:1@localhost:5432/researchdocument?sslmode=disable" goto ${v}
migrate_force:
	migrate -path  $(path) -database "postgresql://mai:1@localhost:5432/researchdocument?sslmode=disable" force ${v}
run:
	docker-compose up