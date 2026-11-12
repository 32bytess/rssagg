# rssagg

A command-line RSS aggregator in Go. Follow feeds, save their posts to PostgreSQL, and read them in your terminal.

## Setup

You need Go 1.22+, PostgreSQL 15+ and [goose](https://github.com/pressly/goose).

1. Install the CLI (make sure `~/go/bin` is on your `PATH`):
   ```sh
   go install github.com/32bytess/rssagg@latest
   ```
2. Create the database:
   ```sh
   psql -U postgres -c "CREATE DATABASE gator;"
   ```
3. Create `~/.gatorconfig.json`:
   ```json
   { "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable" }
   ```
4. Run the migrations from a clone of this repo:
   ```sh
   goose -dir sql/schema postgres "<your db_url>" up
   ```

## Commands

| Command | Description |
| --- | --- |
| `register <name>` | Create a user and log in |
| `login <name>` | Switch user |
| `users` | List users |
| `addfeed <name> <url>` | Add a feed and follow it |
| `feeds` | List all feeds |
| `follow <url>` / `unfollow <url>` | Follow or unfollow a feed |
| `following` | List feeds you follow |
| `agg <duration>` | Fetch feeds every `<duration>` (e.g. `1m`) |
| `browse [limit]` | Show newest posts (default 2) |
| `reset` | **Wipe the database** (dev only) |

## Quick start

```sh
rssagg register alice
rssagg addfeed "Hacker News" https://news.ycombinator.com/rss
rssagg agg 1m        # leave running in another terminal
rssagg browse 10
```

## Development

Run with `go run . <command>`. `./scripts/seed.sh` loads test data (wipes the DB first).
Queries live in `sql/queries`; regenerate Go code with `sqlc generate`.
