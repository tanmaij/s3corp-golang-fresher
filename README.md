## ABOUT THIS APP

    This app is used to study golang. The feature is very simple, manage research documents. The research documents have a lots of different subjects, and ech document have a lots of sub-documents.

## The project structure
    .
    ├── data/
    │   ├── migrations/
    │   │   ├── 000001_research_document.up.sql
    │   │   └── 000001_research_document.down.sql
    │   └── data.go
    ├── internal/
    │   ├── service/
    │   │   ├── subdocument.go
    │   │   └── document.go
    │   ├── reponsitory/
    │   │   ├── subdocument.go
    │   │   ├── new.go
    │   │   ├── mock.go
    │   │   └── document.go
    │   ├── models/
    │   │   ├── subdocument.go
    │   │   ├── psql_upsert.go
    │   │   ├── document.go
    │   │   ├── boil_view_names.go
    │   │   ├── boil_types.go
    │   │   ├── boil_table_names.go
    │   │   └── boil_queries.go
    │   └── handler/
    │       ├── subdocument.go
    │       ├── response.go
    │       └── document.go
    ├── main.go
    ├── go.sum
    ├── go.mod
    ├── Dockerfile
    ├── docker-compose.yml
    ├── sqlboiler.toml
    ├── README.md
    └── Makefile

# HOW TO RUN?
    - You need to install docker in your device
    - Open Terminal in the project path
    - Run "make build_image"
    - Run "make run"

# Technical stack
    - Golang 1.18.3
    - Docker
    - sqlBoiler
    - golang-migration
    - Postgresql
    - Go-chi
