#!/usr/bin/env bash
# Clears a "Dirty database version" state left behind by a failed migration.
# Mirrors what `migrate force <version>` does: sets schema_migrations to the
# given version and clears the dirty flag, without running any SQL from the
# migration files themselves.
set -euo pipefail

cd "$(dirname "$0")/.."

if [ -f .env ]; then
  set -a
  source .env
  set +a
fi

: "${MYSQL_USER:?MYSQL_USER not set (check your .env)}"
: "${MYSQL_PASSWORD:?MYSQL_PASSWORD not set (check your .env)}"
: "${MYSQL_DATABASE:?MYSQL_DATABASE not set (check your .env)}"

VERSION="${1:-}"

mysql_exec() {
  docker compose exec -T mysql mysql -u"$MYSQL_USER" -p"$MYSQL_PASSWORD" "$MYSQL_DATABASE" -e "$1"
}

usage() {
  echo "Usage: $0 <version>|DEL" >&2
  echo "  <version>  Sets schema_migrations to <version> and clears the dirty flag." >&2
  echo "  DEL        Drops the schema_migrations table entirely." >&2
  echo >&2
  echo "Current state:" >&2
  mysql_exec "SELECT * FROM schema_migrations;" >&2
  exit 1
}

if [ -z "$VERSION" ]; then
  usage
fi

if [ "$VERSION" == "DEL" ]; then
  mysql_exec "DROP TABLE schema_migrations;"
  echo "schema_migrations table dropped"
  exit 0
fi

if [[ ! "$VERSION" =~ ^[0-9]+$ ]]; then
  usage
fi

mysql_exec "UPDATE schema_migrations SET version=${VERSION}, dirty=0;"

echo "schema_migrations set to version=${VERSION}, dirty=0"
mysql_exec "SELECT * FROM schema_migrations;"