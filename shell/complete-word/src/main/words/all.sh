#!/bin/bash
root="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

readarray -t wordFiles < <(find "$root/.local" -maxdepth 1 -type f -printf '%P\n')
wordFilePaths=()
for wordFile in "${wordFiles[@]}"
  do
    wordFilePaths+=("$root/.local/$wordFile")
done

cat "${wordFilePaths[@]}"