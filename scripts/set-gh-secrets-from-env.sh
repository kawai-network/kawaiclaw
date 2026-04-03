#!/usr/bin/env bash
set -euo pipefail

REPO="getkawai/kawaiclaw"
ENV_FILE=".env"

usage() {
  cat <<'USAGE'
Usage: scripts/set-gh-secrets-from-env.sh

Reads KEY=VALUE pairs from .env and sets them as GitHub Actions secrets.
Blank lines and lines starting with # are ignored.
USAGE
}

if [[ $# -gt 0 ]]; then
  if [[ "$1" == "-h" || "$1" == "--help" ]]; then
    usage
    exit 0
  fi
  echo "No arguments are supported." >&2
  usage
  exit 1
fi

if [[ ! -f "$ENV_FILE" ]]; then
  echo "env file not found: $ENV_FILE" >&2
  exit 1
fi

if ! command -v gh >/dev/null 2>&1; then
  echo "gh CLI not found. Install gh and run: gh auth login" >&2
  exit 1
fi

# Skip strict auth check; gh secret set will fail if not authenticated.

failures=0

while IFS= read -r line || [[ -n "$line" ]]; do
  trimmed="${line#"${line%%[![:space:]]*}"}"
  trimmed="${trimmed%"${trimmed##*[![:space:]]}"}"

  [[ -z "$trimmed" ]] && continue
  [[ "$trimmed" == \#* ]] && continue

  if [[ "$trimmed" == export\ * ]]; then
    trimmed="${trimmed#export }"
    trimmed="${trimmed#"${trimmed%%[![:space:]]*}"}"
  fi

  if [[ "$trimmed" != *"="* ]]; then
    echo "Skipping invalid line (no '='): $line" >&2
    continue
  fi

  key="${trimmed%%=*}"
  value="${trimmed#*=}"

  key="${key%"${key##*[![:space:]]}"}"
  key="${key#"${key%%[![:space:]]*}"}"

  value="${value#"${value%%[![:space:]]*}"}"
  value="${value%"${value##*[![:space:]]}"}"

  if [[ "$value" == \"*\" ]]; then
    value="${value:1:${#value}-2}"
  elif [[ "$value" == \'*\' ]]; then
    value="${value:1:${#value}-2}"
  fi

  if [[ -z "$key" ]]; then
    echo "Skipping invalid line (empty key): $line" >&2
    continue
  fi

  echo "Setting secret: $key"
  if ! printf '%s' "$value" | gh secret set "$key" -R "$REPO"; then
    echo "Failed to set secret: $key" >&2
    failures=$((failures + 1))
  fi
done < "$ENV_FILE"

if [[ "$failures" -gt 0 ]]; then
  exit 1
fi
