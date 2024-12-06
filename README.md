# Movie2.0

#Backend Directory structure
/backend/
│
├── /shared/                    # Shared directories for common utils across microservices
│   ├── /utils/                 # Shared utilities (JWT, logging, error handling, etc.)
│   ├── /config/                # Shared configuration (environment variables, common configs)
│   ├── /constants/             # Shared constants (status codes, roles, etc.)
│   └── /logger/                # Shared logger for all services
│
├── /internals/                 # Internal packages (not to be reused outside this repo)
│   ├── /auth/                  # JWT token validation logic, internal auth utils
│   ├── /payment/               # Stripe payment integration logic (internal use only)
│   └── /domain/                # Core domain logic (business rules for tickets, movies, etc.)
│
├── /pkg/                       # Reusable packages for the project (can be used across services)
│   ├── /auth/                  # Authentication helpers (JWT signing, token verification)
│   ├── /utils/                 # Utility functions (e.g., string manipulation, time)
│   ├── /logger/                # Logger package, used across all services
│   └── /config/                # Config manager to load and manage environment variables
│
├── /services/                  # Directory for individual microservices
│   ├── /users-service/         # User Management Microservice
│   │   ├── /cmd/               # Main entry point for this service
│   │   ├── /controllers/       # Handlers for user-related requests
│   │   ├── /models/            # User-related data models
│   │   ├── /routes/            # Routes for user-related APIs
│   │   ├── /services/          # Business logic for user-related operations
│   │   ├── /repository/        # Data access layer (DB operations)
│   │   └── main.go             # Main entry point for user service
│   │
│   ├── /theaters-service/      # Theater Management Microservice
│   │   ├── /cmd/               # Main entry point for this service
│   │   ├── /controllers/       # Handlers for theater-related requests
│   │   ├── /models/            # Theater-related data models
│   │   ├── /routes/            # Routes for theater-related APIs
│   │   ├── /services/          # Business logic for theater-related operations
│   │   ├── /repository/        # Data access layer (DB operations)
│   │   └── main.go             # Main entry point for theater service
│   │
│   ├── /movies-service/        # Movie Management Microservice
│   │   ├── /cmd/               # Main entry point for this service
│   │   ├── /controllers/       # Handlers for movie-related requests
│   │   ├── /models/            # Movie-related data models
│   │   ├── /routes/            # Routes for movie-related APIs
│   │   ├── /services/          # Business logic for movie-related operations
│   │   ├── /repository/        # Data access layer (DB operations)
│   │   └── main.go             # Main entry point for movie service
│   │
│   ├── /seats-service/         # Seat Reservation Microservice
│   │   ├── /cmd/               # Main entry point for this service
│   │   ├── /controllers/       # Handlers for seat reservation requests
│   │   ├── /models/            # Seat-related data models
│   │   ├── /routes/            # Routes for seat reservation APIs
│   │   ├── /services/          # Business logic for seat reservation operations
│   │   ├── /repository/        # Data access layer (DB operations)
│   │   └── main.go             # Main entry point for seat reservation service
│   │
│   ├── /tickets-service/       # Ticket Management Microservice
│   │   ├── /cmd/               # Main entry point for this service
│   │   ├── /controllers/       # Handlers for ticket-related requests
│   │   ├── /models/            # Ticket-related data models
│   │   ├── /routes/            # Routes for ticket-related APIs
│   │   ├── /services/          # Business logic for ticket-related operations
│   │   ├── /repository/        # Data access layer (DB operations)
│   │   └── main.go             # Main entry point for ticket service
│
├── /docker/                    # Docker files and configurations
│   ├── docker-compose.yml      # Docker Compose file to run all services together
│   └── /postgres/              # PostgreSQL Docker setup (docker-compose for DB)
│
├── /migrations/                # Database migrations (for PostgreSQL schema updates)
│   └── migration_file.sql      # SQL files to update DB schema
│
├── .env                        # Environment variables for configuration
├── go.mod                      # Go modules (dependencies)
├── go.sum                      # Go modules checksum
└── README.md                   # Project documentation
