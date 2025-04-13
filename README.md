# go_docker_api

This is a personal project to study REST APIs with Go and Docker.
For study purposes, some files may have explanatory comments.

references:
    https://github.com/arthur404dev/gopportunities - Repository used as a reference for this project.

commands:
    go:
        go mod <project_name> init                                                  --> create the go.mod file
        go get <github.com/package_name>                                            --> packge download
    docker:
        docker compose up -d <countainer_name>                                      --> build docker container

tools:
    migrations: github.com/pressly/goose

        make sure to create the .env variable
        .env variables: 
            GOOSE_DRIVER=postgres                                                   --> the database driver to use
            GOOSE_DBSTRING="postgresql://postgres:0000@localhost:5433/postgres"     --> the database connection string
            GOOSE_MIGRATION_DIR=./migrations                                        --> the directory containing the migration files (default: .)

        commands:
            goose create <migration_name> sql                                       --> create a migration file .sql
            goose up                                                                --> up
            goose down                                                              --> down


libs:
    gin(REST API): github.com/gin-gonic/gin

Go project file structurte:
/my-project 
├── /cmd                                    # Entry points for the application (main.go files)
│   └── /migrate                            # Command to run database migrations
│       └── main.go                         # Runs the migrations on the database
│   └── /server                             # Command to start the API server
│       └── main.go                         # Entry point for starting the server (e.g., using net/http)
├── /pkg                                    # Main package with reusable business logic
│   └── /database                           # Database management functions
│       └── migration.go                    # Logic to handle database migrations in Go
│       └── db.go                           # Database connection and configuration
│   └── /api                                # API logic
│       └── user.go                         # User model logic for interacting with the API
│   └── /services                           # Functions and implementations for services (e.g., data processing)
│       └── user_service.go                 # Functions that manage user data
├── /internal                               # Internal packages that are used only within your project
│   └── /utils                              # Helper or utility functions
│       └── logger.go                       # Logic for configuring logging in the project
│   └── /auth                               # Authentication and authorization logic
│       └── jwt.go                          # Logic for creating and validating JWT (if necessary)
├── /migrations                             # Directory containing database migration scripts
│   └── 001_create_users_table.up.sql       # Migration to create the users table
│   └── 001_create_users_table.down.sql     # Rollback migration for the users table creation
│   └── 002_add_email_column.up.sql         # Add an email column to the users table
│   └── 002_add_email_column.down.sql       # Rollback migration for adding the email column
├── /scripts                                # Scripts for maintenance, backups, or other external actions
│   └── backup.sh                           # Script for database backup
│   └── seed_users.sh                       # Script to insert initial data into the database
├── /configs                                # Configuration files for the project
│   └── config.yml                          # General configuration file (e.g., environment, database settings)
│   └── config.toml                         # Configuration file using TOML format
├── /api                                    # API logic: Controllers, Routing, Models
│   └── /controllers                        # REST endpoint controllers
│       └── user_controller.go              # Controller for user API endpoints
│   └── /routes                             # API route definitions
│       └── user_routes.go                  # Routes for user API (POST, GET, PUT, DELETE)
├── /docs                                   # Project documentation (README, manuals)
│   └── README.md                           # General project documentation
├── /tests                                  # Unit and integration tests
│   └── /unit                               # Unit tests for isolated functions
│       └── user_service_test.go            # Tests for user service functions
│   └── /integration                        # Tests for system integrations
│       └── user_integration_test.go        # Test to verify the user API integration
├── go.mod                                  # Go dependency manager file
├── go.sum                                  # Checksum file for Go dependencies
└── Makefile                                # Build and task automation file (optional)