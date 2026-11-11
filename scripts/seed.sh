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
run addfeed "Boot.dev Blog" https://blog.boot.dev/index.xml
run addfeed "Alberto.dev Blog" https://blog.alberto.dev/index.xml

run register lane
run addfeed "TechCrunch" https://techcrunch.com/feed/
run follow https://news.ycombinator.com/rss
run follow https://blog.alberto.dev/index.xml

run login alberto

echo "seeded: alberto follows 3 feed, lane follows 3 (current user: alberto)"
