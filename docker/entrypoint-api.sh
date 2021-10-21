#!/bin/bash
set -e

migrate () {
  echo "migrations..."
}

runserver () {
  echo "run server"
}

case "$1" in
  migrate)
    shift
    migrate
    ;;
  runserver)
    shift
    runserver
    ;;
  *)
    exec "$@"
    ;;
esac