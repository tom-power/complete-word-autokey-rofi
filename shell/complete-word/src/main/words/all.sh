#!/bin/bash
root="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

wordsPath=~/.config/complete-word-autokey-rofi/words
readarray -t wordFiles < <(find $wordsPath -maxdepth 1 -type f -printf '%P\n')
wordFilePaths=()
for wordFile in "${wordFiles[@]}"
  do
    wordFilePaths+=("$wordsPath/$wordFile")
done

cat "${wordFilePaths[@]}"