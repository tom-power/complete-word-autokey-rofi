#!/bin/bash
root="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
word=$1
if [[ ! -z "$word" ]]; then
  addFile=~/.config/complete-word-autokey-rofi/words/added.txt
  if [[ ! -d "$addFile" ]]; then
    touch $addFile
  fi
  grep -x $word $addFile || echo $word >> $addFile
fi
