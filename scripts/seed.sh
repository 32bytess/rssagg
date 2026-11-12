#!/usr/bin/env bash
# Seeds the database with test data through the CLI.
# WARNING: starts with `reset`, which deletes every user, feed and follow in
# the database named in ~/.gatorconfig.json.
set -euo pipefail

cd "$(dirname "$0")/.."

go build -o rss-agg .
run() { ./rss-agg "$@"; }

run reset

run register alberto
run addfeed "Hacker News" https://news.ycombinator.com/rss

run register lane
run addfeed "TechCrunch" https://techcrunch.com/feed/

run login alberto
run follow https://techcrunch.com/feed/

echo "seeded: alberto follows 2 feed, lane follows 1 (current user: alberto)"
