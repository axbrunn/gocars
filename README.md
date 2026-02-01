# gocars 🚗

gocars is een multi-tenant webapplicatie waarmee autobedrijven eenvoudig hun eigen website en voorraad kunnen beheren.
De applicatie is gebouwd in Go, gebruikt server-side templates en htmx voor een snelle SPA-achtige ervaring zonder zware frontend frameworks.

## Kernidee

- Er is één Go-applicatie
- Meerdere autobedrijven (tenants) gebruiken dezelfde app
- Elk autobedrijf heeft:
    - een eigen website
    - een eigen dashboard
    - volledig gescheiden data (auto’s, teksten, gebruikers)

## Multi-tenant structuur

Elke request hoort bij één tenant (autobedrijf).
De tenant wordt bepaald op basis van het domein of subdomein, bijvoorbeeld:

```
jans.gocars.nl
piet.gocars.nl
```

Middleware leest het domein en koppelt de request aan het juiste autobedrijf (dealer).

---

## Hosts file (voor lokale tenants)

Voeg tenant-subdomains toe aan je hosts-bestand.

### Windows

1. Open **Notepad als administrator**
2. Open:

   ```
   C:\Windows\System32\drivers\etc\hosts
   ```
3. Voeg toe:

   ```
   127.0.0.1 tenant1.localhost
   127.0.0.1 tenant2.localhost
   ```

### Linux / macOS

```bash
sudo nano /etc/hosts
```

Voeg toe:

```
127.0.0.1 tenant1.localhost
127.0.0.1 tenant2.localhost
```

Daarna:

* [http://tenant1.localhost:8080](http://tenant1.localhost:8080)
* [http://tenant2.localhost:8080](http://tenant2.localhost:8080)

---

## Running PostgreSQL in Docker Containers

This Docker setup includes PostgreSQL.

### Prerequisites

- Docker
- Docker Compose (comes with Docker)

### Setup Instructions

1. **Clone your repository** (if not already done)

2. **Create environment file** (optional - defaults are set in docker-compose.yml):
   
 ```bash
cp .env.example .env
```
   
 Edit `.env` if you want to customize the database credentials.

### Start all services

 ```bash
# Build and start containers (if needed)
docker compose up --build

# Or just start containers
docker compose up
```

This will start the PostgreSQL container.

**Run in detached mode**

 ```bash
docker compose up -d
```

Containers run in the background.

**View logs**

 ```bash
docker compose logs -f
```

Follow logs for all running services.

**Stop services**

 ```bash
docker compose down
```

**Stop and remove volumes (⚠️ deletes database data)**

 ```bash
docker compose down -v
```

---

## PostgreSQL Database Migrations

We use migrate to manage database schema changes.

**Create a new migration:**

```bash
migrate create -ext sql -dir ./migrations -seq add_example_table
```

**Run migrations:**

```bash
migrate -path=./migrations -database=$DB_DSN up
```

**Rollback last migration:**

```bash
migrate -path=./migrations -database=$DB_DSN down 1
```

- `.up.sql` → applies changes
- `.down.sql` → reverts changes
- `$DB_DSN` = your database connection string

---

## Using Makefile

### Prerequisite

Make sure make is installed on your system.

```bash
# On Linux (Ubuntu/Debian)
$ sudo apt install make

# On macOS (using Homebrew)
$ brew install make

# On Windows (using Chocolatey)
> choco install make
```

### Available Commands

**Show help (Linux only)**

```bash 
make help
```

Prints a usage overview of available Makefile targets.

**Run the Go API locally**

```bash 
make run/api
```

Runs the Go API using the DSN from .env.

**Open a psql session**

```bash 
make db/psql
```

Opens a PostgreSQL shell connected to your Docker container.

**Create a new migration**

```bash 
make db/migrations/new name=your_migration_name
```
Creates a new SQL migration file in the `./migrations` directory.
Replace `your_migration_name` with a descriptive name.

**Run all pending migrations**

```bash 
make db/migrations/up
```

Runs all migrations that haven’t been applied yet.
You will be prompted for confirmation before applying migrations.

---
