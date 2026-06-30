# Gator
Gator is a multi-user RSS feed aggregator CLI. It continuously fetches posts from registered feeds in the background and stores them in a PostgreSQL database for browsing.

---

## Features

- Add RSS feeds from across the web to be aggregated
- Store aggregated posts in a PostgreSQL database
- Follow and unfollow RSS feeds added by other users
- View summaries of aggregated posts in the terminal

---

## Requirements

- Postgres: https://www.postgresql.org/download/
- Go: https://go.dev/doc/install

---

## How to Run

### 1. Install the program

```bash
go install github.com/tomkalva/gator@latest
```

### 2. Create a config file
In your home directory create a file named: ".gatorconfig.json"
with this structure:
```json
{
  "db_url": "postgres://username:@localhost:5432/database?sslmode=disable"
}
```


### 3. Run the program
```bash
gator <command>
```

### 4. Usage Examples

#### Register a user
```bash
gator register <name>
```

#### Login
```bash
gator login <name>
```

#### List all users and show current one
```bash
gator users
```

#### Add a feed
```bash
gator addfeed <name> <url>
```

#### List all feeds
```bash
gator feeds
```

#### Aggregate over feeds
```bash
gator agg <timeBetweenRequests>
gator agg 30s
gator agg 2s
```

#### Following other users feeds
```bash
gator follow <url>
gator unfollow <url>
gator following
```

#### Browse feeds
```bash
gator browse <limit(default=2)>
```
