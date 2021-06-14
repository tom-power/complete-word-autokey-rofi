#!/bin/bash
root="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
word=$1
if [[ ! -z "$word" ]]; then
  addFile=$root/.local/added.txt
  grep -x $word $addFile || echo $word >> $addFile
fi
